import {
  addAllChats,
  addMessagesToChat,
  addOnlineUsers,
  addNewMessageToEndOrGiveNotification,
  appendMessagesToChat,
} from "./views/ChatView.vue";
import { unreadMessageDisplayOn, unreadMessageDisplayOff } from "./App.vue";
var connection;
const getConnection = () => {
  /*if (connection && connection.readyState < 2) {
    console.log("connection juba loodud");
    return Promise.resolve(connection);
  }*/
  if (connection && connection.readyState < 2) {
    return Promise.resolve(connection);
  }
  console.log("Websocketi esmane loomine");
  return new Promise((resolve) => {
    if (window["WebSocket"]) {
      let token = document.cookie.substring(8);
      if (token == undefined) {
        return;
      }

      const conn = new WebSocket(`ws://localhost:8009/socket`);

      conn.onopen = function () {
        conn.send(JSON.stringify({ type: "token", body: token }));
        var msg2 = {
          type: "getAllChats",
          body: localStorage.loggedInUsername,
        };
        Ws.send(msg2);
        var msg3 = {
          type: "getOnlineUsers",
          body: localStorage.loggedInUsername,
        };
        Ws.send(msg3);
      };

      conn.onmessage = async function (msg) {
        //console.log("Wbsocketisse tulev s6num: ");
        //let test = JSON.stringify(msg.data);
        let jsonMsg = JSON.parse(msg.data);
        console.log("Wbsocketisse tulev s6num: ", jsonMsg);
        switch (jsonMsg.Type) {
          case "getAllChats":
            if (
              jsonMsg.Body.some((element) => {
                if (element["UnreadMessagesCount"] > 0) {
                  return true;
                }
              })
            ) {
              unreadMessageDisplayOn();
            } else {
              unreadMessageDisplayOff();
            }
            addAllChats(jsonMsg.Body);
            break;
          case "getMessagesFromChat":
            addMessagesToChat(jsonMsg.Body);
            break;
          case "appendMessagesFromChat":
            appendMessagesToChat(jsonMsg.Body);
            break;
          case "newMessage":
            //siia tuleb kirjutada, et n2idata notificationit userile, et uus s6num on tulnud
            console.log("newmessage entry = ", jsonMsg.Body);
            addNewMessageToEndOrGiveNotification(jsonMsg.Body);
            break;
          case "onlineUsers":
            addOnlineUsers(jsonMsg.Body);
            break;
          default:
            console.log("backendi websocketist tule defaulti, mv");
            console.log(jsonMsg);
        }
      };
      resolve(conn);
    } else {
      alert("Your browser does not support WebSockets");
    }
  });
};

const Ws = {
  connect: async () => {
    connection = await getConnection();
  },

  send: async (e) => {
    connection = await getConnection();
    connection.send(JSON.stringify(e));
    console.log("sending: ", e);
  },

  disconnect: async () => {
    connection.close();
  },
};

export default Ws;

/* 
let connection = new WebSocket(`ws://localhost:8009/socket`);
function open() {
  connection = new WebSocket(`ws://localhost:8009/socket`);
}
function send(msg) {
  connection.send(JSON.stringify(msg));
}
connection.onopen = function () {
  connection.send(
    JSON.stringify({ type: "userOnline", body: localStorage.loggedInUsername })
  );
};
connection.onclose = function () {
  connection.send(
    JSON.stringify({ type: "userOffline", body: localStorage.loggedInUsername })
  );
};

function close() {
  connection.close();
}

connection.onmessage = async function (evt) {
  console.log("Receiving transmission:");
  console.log(evt);
};
export default {
  connection,
  send,
  close,
  open,
};
 */
