// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"social-network/dbFunc"
	"social-network/structs"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send     chan []byte
	clientId int
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
type WSComm struct {
	Type string
	Body interface{}
}

func sendMsg(c *Client, msg WSComm) {
	out, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	//go func() {
	//	c.send <- out
	//}()
	//fmt.Println("Saadan s6numi: ", string(out))
	c.send <- out
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		/*_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		fmt.Println(string(message))*/
		//c.hub.broadcast <- message
		var msg WSComm
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("error reading JSON: %s\n", err.Error())
			break
		}
		fmt.Println("Loen s6numit:", msg)
		switch msg.Type {
		case "token":
			fmt.Println(msg.Body)
			session, err := dbFunc.GetSessionBySessionId(dbName, msg.Body.(string))
			if err != nil {
				fmt.Println(err)
				fmt.Println("sessionnit pole olemas")
				return
			}
			c.clientId = session.Userid
			var msg WSComm
			var userIDs []int
			for c := range c.hub.clients {
				userIDs = append(userIDs, c.clientId)
			}
			msg.Type = "onlineUsers"
			msg.Body = userIDs
			out, _ := json.Marshal(msg)
			go func() {
				c.hub.broadcast <- out
			}()
		case "getAllChats":
			users, err := dbFunc.GetUsersWithLastMessage(dbName, c.clientId)
			user, err := dbFunc.GetUserById(dbName, c.clientId)
			var test structs.Chat
			test.User = user
			users = append(users, test)
			if err != nil {
				fmt.Println(err)
			}
			msg.Type = "getAllChats"
			msg.Body = users
			sendMsg(c, msg)
		case "getAllChats2":
			var receiverID int
			var username string
			m, ok := msg.Body.(map[string]interface{})
			if !ok {
				fmt.Println("t6rge fix it")
			}
			for key, value := range m {
				switch key {
				case "receiverid":
					receiverID, _ = strconv.Atoi(value.(string))
				case "username":
					username, _ = value.(string)
				}
			}
			user2, _ := dbFunc.GetUserByUsername(dbName, username)
			dbFunc.UpdateMessageRead(dbName, user2.Id, receiverID)
			users, err := dbFunc.GetUsersWithLastMessage(dbName, c.clientId)
			user, err := dbFunc.GetUserById(dbName, c.clientId)
			var test structs.Chat
			test.User = user
			users = append(users, test)
			if err != nil {
				fmt.Println(err)
			}
			msg.Type = "getAllChats"
			msg.Body = users
			sendMsg(c, msg)
		case "messagesRequest":
			var receiverid int
			var limit int
			m, ok := msg.Body.(map[string]interface{})
			if !ok {
				fmt.Println("t6rge fix it")
			}
			for k, v := range m {
				switch k {
				case "userID":
					//not pretty what's in here, but it is what it is
					test := v.(string)
					test2, _ := strconv.Atoi(test)
					receiverid = test2
				case "messagesCount":
					test := v.(float64)
					test2 := int(test)
					limit = test2
				}
			}
			messages, err := dbFunc.GetMessages(dbName, c.clientId, receiverid, limit)
			err = dbFunc.UpdateMessageRead(dbName, c.clientId, receiverid)
			if err != nil {
				fmt.Println("getmessages ei tyyta :((")
			}
			if limit == 0 {
				msg.Type = "getMessagesFromChat"
				msg.Body = messages
				sendMsg(c, msg)
			} else {
				msg.Type = "appendMessagesFromChat"
				msg.Body = messages
				sendMsg(c, msg)
			}

		case "sendMessage":
			var receiverid int
			var senderid int
			var text string
			m, ok := msg.Body.(map[string]interface{})
			if !ok {
				fmt.Println("t6rge fix it")
			}
			for k, v := range m {
				switch k {
				case "Senderid":
					senderid = int(v.(float64))
				case "Receiverid":
					receiverid, _ = strconv.Atoi(v.(string))
				case "Text":
					text = v.(string)
				}
			}
			dbFunc.SetMessage(dbName, senderid, receiverid, text)
			//fmt.Println(c.hub.clients)
			for c := range c.hub.clients {
				//fmt.Println("clientid = ", c.clientId, " v = ", v)
				if receiverid == c.clientId {
					msg.Type = "newMessage"
					var out structs.Message
					out.Senderid = senderid
					out.Text = text
					msg.Body = out
					sendMsg(c, msg)
					users, err := dbFunc.GetUsersWithLastMessage(dbName, c.clientId)
					//fmt.Println("kontrollime unreadmessagecount", users)
					user, err := dbFunc.GetUserById(dbName, c.clientId)
					var test structs.Chat
					test.User = user
					users = append(users, test)
					if err != nil {
						fmt.Println(err)
					}
					msg.Type = "getAllChats"
					msg.Body = users
					sendMsg(c, msg)
					var userIDs []int
					for c := range c.hub.clients {
						userIDs = append(userIDs, c.clientId)
					}
					msg.Type = "onlineUsers"
					msg.Body = userIDs
					//c.conn.WriteJSON(msg)
					fmt.Println(msg)
					sendMsg(c, msg)
					break
				}
			}
		case "getOnlineUsers":
			var userIDs []int
			for c := range c.hub.clients {
				userIDs = append(userIDs, c.clientId)
			}
			msg.Type = "onlineUsers"
			msg.Body = userIDs
			//c.conn.WriteJSON(msg)
			fmt.Println(msg)
			sendMsg(c, msg)

		default:
			fmt.Println(c.clientId, msg, "eksinud s6num defaultis")
		}
		//fmt.Printf("Message Received, Type: %s\nBody: %s\nTo: ", msg.Type, msg.Body)

	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				fmt.Println("The hub closed the channel.")
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			/* n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}
			*/
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))

			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	if !websocket.IsWebSocketUpgrade(r) {
		// The connection is not a valid WebSocket upgrade request, so return an error
		http.Error(w, "Not a valid WebSocket upgrade request", http.StatusBadRequest)
		return
	}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		//check the http.Request
		// make sure it's ok to access
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

/* //https://gowebexamples.com/websockets/
//https://medium.com/@matryer/the-http-handler-wrapper-technique-in-golang-updated-bc7fbcffa702#.e4k81jxd3
//https://tutorialedge.net/projects/chat-system-in-go-and-react/part-4-handling-multiple-clients/

package main

import (
	"fmt"
	"net/http"

	"01.kood.tech/git/harlet/real-time-forum.git/dbFunc"
	"github.com/gorilla/websocket"
)

type Message struct {
	Type string
	Body string
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConn *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WSConnection established")
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		//check the http.Request
		// make sure it's ok to access
		return true
	}

	wsConn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer wsConn.Close()

	for {
		var msg Message
		err := wsConn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("error reading JSON: %s\n", err.Error())
			break
		}
		switch msg.Type {
		case "getAllChats":
			users, err := dbFunc.GetAllUsers(dbName)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(users)
			wsConn.WriteJSON(users)
		case "userOnline":
			dbFunc.ChangeUserOnlineStatus(dbName, msg.Body, 1)
		case "userOffline":
			dbFunc.ChangeUserOnlineStatus(dbName, msg.Body, 0)
			fmt.Println("kasutaja logib v2lja")
		default:
			fmt.Println(msg)
		}
		//fmt.Printf("Message Received, Type: %s\nBody: %s\nTo: ", msg.Type, msg.Body)

	}
	fmt.Println("SIIA EI TOHIKS J^UDA WSCONNECTION")
}

func SendMessage(msg string) {
	err := wsConn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		fmt.Printf("error sending message: %s\n", err.Error())
	}
} */
