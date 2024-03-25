// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			fmt.Println("wsconnection success")
			h.clients[client] = true
			var msg WSComm
			var userIDs []int
			for c := range h.clients {
				userIDs = append(userIDs, c.clientId)
			}
			msg.Type = "onlineUsers"
			msg.Body = userIDs
			out, _ := json.Marshal(msg)
			go func() {
				h.broadcast <- out
			}()

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				fmt.Println("wsconnection closed")
				delete(h.clients, client)
				close(client.send)
			}
			var msg WSComm
			var userIDs []int
			for c := range h.clients {
				userIDs = append(userIDs, c.clientId)
			}
			msg.Type = "onlineUsers"
			msg.Body = userIDs
			out, err := json.Marshal(msg)
			if err != nil {
				fmt.Println("tekkis error byte'deks tegemisel", err)
			}
			go func() {
				h.broadcast <- out
			}()
		case message := <-h.broadcast:
			fmt.Println("Broadcastib:", string(message))
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
