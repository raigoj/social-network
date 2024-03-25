<template>
  <div id="chats-container">
    <div id="chats"></div>
    <div>
      <div id="chat-messages">
        <p id="chat-messages-placeholder">Select chat to start messaging</p>
      </div>
      <form
        @submit.prevent="sendMessage()"
        id="message-form"
        style="display: block"
      >
        <input
          required
          v-model="sentMessage"
          type="text"
          id="message-input"
          size="64"
          placeholder="Send message"
          autocomplete="off"
          autofocus="false"
          style="display: none"
        />
      </form>
    </div>
  </div>
</template>
<script>
let vueUserId;
let vueReceiverID;
function formatDate(input) {
  const date = new Date(input);
  const year = date.getFullYear();
  const month = `${date.getMonth() + 1}`.padStart(2, "0");
  const day = `${date.getDate()}`.padStart(2, "0");
  const hours = `${date.getHours()}`.padStart(2, "0");
  const minutes = `${date.getMinutes()}`.padStart(2, "0");
  return `${year}.${month}.${day} ${hours}:${minutes}`;
}
function getCurrentDate() {
  const currentDate = new Date();
  const year = currentDate.getFullYear();
  const month = String(currentDate.getMonth() + 1).padStart(2, "0");
  const day = String(currentDate.getDate()).padStart(2, "0");
  const hour = String(currentDate.getHours()).padStart(2, "0");
  const minute = String(currentDate.getMinutes()).padStart(2, "0");
  const dateString = `${year}.${month}.${day} ${hour}:${minute}`;
  return dateString;
}

const addMessageToChat = (message) => {
  const messageDate = document.createElement("p");
  messageDate.classList.add("message-date");
  //console.log(message.Sentat);
  if (message.Sentat.length === 0) {
    messageDate.innerText = getCurrentDate();
  } else {
    messageDate.innerText = formatDate(message.Sentat);
  }
  let messages = document.getElementById("chat-messages");
  let messag = document.createElement("div");
  //lisa id ja class ja p s6numi sisuga
  messag.id = "message-" + message.Id;
  let test;
  if (message.Senderid === vueUserId) {
    test = "sended";
  } else {
    test = "received";
  }

  messag.classList.add("message");
  messag.classList.add(test + "-message");
  let sisu = document.createElement("p");
  sisu.classList.add("message-text");
  sisu.textContent = message.Text;
  //let msgInfo = document.createElement("div");
  messag.appendChild(sisu);
  messag.appendChild(messageDate);
  if (messages != null) {
    //messages.appendChild(messag);
    return messag;
  }
};
const addValuesToChat = (user) => {
  if (user.User.Username === localStorage.loggedInUsername) {
    vueUserId = user.User.Id;
    return;
  }
  //console.log(user);
  let chats = document.getElementById("chats");
  let chat = document.createElement("div");
  let chatInfo = document.createElement("div");
  let unreadMessagesCount = document.createElement("div");
  unreadMessagesCount.classList.add(`unread-messages-count`);
  unreadMessagesCount.id = `chat-${user.User.id}-unread-messages-count`;
  let lastMessage = document.createElement("p");
  let lastMessageDate = document.createElement("p");
  let userName = document.createElement("p");
  let pic = document.createElement("div");
  //pic.src = "http://localhost:8081/images/default-male-avatar.jpg";
  userName.textContent = user.User.Firstname + " " + user.User.Lastname;
  chat.classList.add("chat");
  //chat.classList.add("active");
  chat.setAttribute("id", "chat-" + user.User.Id);
  chat.style.cursor = "pointer";
  chatInfo.classList.add("chat-info");
  lastMessage.setAttribute("id", "chat-" + user.User.Id + "-lastMessage");
  if (user.UnreadMessagesCount != 0) {
    unreadMessagesCount.textContent = user.UnreadMessagesCount;
    unreadMessagesCount.style.opacity = 100;
  } else {
    unreadMessagesCount.textContent = 0;
    unreadMessagesCount.style.opacity = 0;
  }
  lastMessageDate.setAttribute(
    "id",
    "chat-" + user.User.Id + "-lastMessageDate"
  );

  chat.addEventListener("click", () => {
    //const test = document.getElementById("message-input");
    document.getElementById("message-input").style.display = "block";
    Array.from(document.getElementsByClassName("active")).forEach((el) => {
      el.classList.remove("active");
    });
    chat.classList.add("active");
    let test = chat.id.split("-");
    vueReceiverID = test[1];
    var msg1 = {
      type: "messagesRequest",
      body: {
        userID: test[1],
        messagesCount: 0,
      },
    };
    Ws.send(msg1);
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
  });
  if (chats != null) {
    chats.appendChild(chat);
  }

  chat.appendChild(pic);
  chat.appendChild(chatInfo);
  chat.appendChild(unreadMessagesCount);
  chatInfo.appendChild(userName);
  chatInfo.appendChild(lastMessage);
  chatInfo.appendChild(lastMessageDate);
};
export async function addAllChats(users) {
  //console.log("HAKKAME EHITAMA");
  let chatid;
  Array.from(document.getElementsByClassName("active")).forEach((el) => {
    let test = el.id.split("-");
    chatid = parseInt(test[1]);
  });
  let chats = document.getElementById("chats");
  if (chats === null) {
    return;
  }
  chats.innerHTML = "";
  users.forEach((element) => {
    // console.log("elementid = ", element.User.Id);
    //console.log("chatid = ", chatid);
    console.log("unreadmessagecount :", element.UnreadMessagesCount);
    addValuesToChat(element);
    if (element.User.Id === chatid) {
      let el = document.getElementById(`chat-${chatid}`);
      el.classList.add("active");
    }
  });
}
export async function appendMessagesToChat(messages) {
  var element = document.getElementById("chat-messages");
  let chatHeight = element.scrollHeight;
  messages.forEach((message) => {
    element.prepend(addMessageToChat(message));
  });
  let chatHeight2 = element.scrollHeight;
  element.scrollTop = chatHeight2 - chatHeight;
}
export async function addMessagesToChat(messages) {
  let messages2 = document.getElementById("chat-messages");
  messages2.innerHTML = "";
  messages.forEach((message) => {
    //console.log(message);
    messages2.appendChild(addMessageToChat(message));
  });
  messages2.scrollTop = messages2.scrollHeight;
}
export async function addOnlineUsers(users) {
  console.log("praegused online userid:", users);
  Array.from(document.getElementsByClassName("chat")).forEach((el) => {
    el.classList.remove("online");
  });
  if (users != null) {
    users.forEach((u) => {
      var chat = document.getElementById(`chat-${u}`);
      if (chat) {
        //console.log("lisan online chat-i l6ppu");
        chat.classList.add("online");
      }
    });
  }
}
export async function addNewMessageToEndOrGiveNotification(message) {
  if (vueReceiverID === message.Senderid.toString()) {
    console.log("j6uab ilust");
    var elem = document.getElementById("chat-messages");
    if (elem === null) {
      return;
    }
    elem.appendChild(addMessageToChat(message));
    elem.scrollTop = elem.scrollHeight;
    //enne getallchatsi peaks saatma s6numi et uued s6numid on loetud
    var msg2 = {
      type: "getAllChats2",
      body: {
        username: localStorage.loggedInUsername,
        receiverid: vueReceiverID,
      },
    };
    Ws.send(msg2);
    var msg3 = {
      type: "getOnlineUsers",
      body: localStorage.loggedInUsername,
    };
    Ws.send(msg3);
  }
}
</script>
<script setup>
import { onMounted, ref } from "vue";
import Ws from "@/connection";
const sentMessage = ref();
const sendMessage = () => {
  //console.log("Saadan teele s6numi: ", sentMessage.value);
  var msg = {
    type: "sendMessage",
    body: {
      Senderid: vueUserId,
      Receiverid: vueReceiverID,
      Text: sentMessage.value,
    },
  };
  sentMessage.value = "";
  Ws.send(msg);
  var msg2 = {
    type: "getAllChats",
    body: localStorage.loggedInUsername,
  };
  Ws.send(msg2);
  var msg1 = {
    type: "messagesRequest",
    body: {
      userID: vueReceiverID,
      messagesCount: 0,
    },
  };
  Ws.send(msg1);
  var msg3 = {
    type: "getOnlineUsers",
    body: localStorage.loggedInUsername,
  };
  Ws.send(msg3);
};
/*
    onMounted()
    onUpdated()
    onUnmounted()
    onBeforeMount()
    onBeforeUpdate()
    onBeforeUnmount()
    onErrorCaptured()
    onRenderTracked()
    onRenderTriggered()
    onActivated()
    onDeactivated()
    onServerPrefetch()
*/

onMounted(() => {
  //console.log("Chatvue onmounted", localStorage.loggedInUsername);
  var msg1 = {
    type: "getAllChats",
    body: localStorage.loggedInUsername,
  };
  Ws.send(msg1);
  var msg2 = {
    type: "getOnlineUsers",
    body: localStorage.loggedInUsername,
  };
  Ws.send(msg2);
  var element = document.getElementById("chat-messages");
  // add a scroll event listener to the element
  element.addEventListener("scroll", function () {
    // check if the user has scrolled to the top of the element
    //console.log("scrolltop value = ", element.scrollTop);
    if (this.scrollTop === 0) {
      // do something
      let loadedMsgs = Array.from(document.getElementsByClassName("message"));
      var msg1 = {
        type: "messagesRequest",
        body: {
          userID: vueReceiverID,
          messagesCount: loadedMsgs.length,
        },
      };
      Ws.send(msg1);
    }
  });
});
</script>
<style>
.body {
  word-wrap: break-word;
  overflow-wrap: break-word;
}

#chats-container {
  padding: 10px;
  display: grid;
  grid-template-columns: 300px auto;
  background-color: #292a2d;
  border-radius: 5px;
}

#chats {
  height: 580px;
  overflow-y: scroll;
}

.chats-title {
  font-size: 18px;
  font-weight: 600;
  padding: 10px;
  margin: 0px;
}

.chat {
  display: grid;
  grid-template-columns: 50px 175px 25px;
  grid-gap: 10px;
  padding: 10px;
  margin-bottom: 10px;
  border-radius: 5px;
  background-color: #3b3d42;
  color: #c9c9d8;
}

.chat.active {
  background-color: #c9c9d8;
  color: #292a2d;
}

.chat-info p {
  word-wrap: normal;
  max-width: 188px;
  overflow: hidden;
  text-overflow: ellipsis;
}

.unread-messages-count {
  align-self: center;
  width: 25px;
  height: 25px;
  font-size: 13px;
  text-align: center;
  line-height: 25px;
  border-radius: 100%;
  vertical-align: middle;
  background-color: #c9c9d8;
  color: #292a2d;
}

.chat-unread-messages-count.active {
  background-color: #3b3d42;
  color: #c9c9d8;
}

.chat img {
  border-radius: 50%;
  max-width: 50px;
  margin: 0;
  padding: 0;
}

.chat.online {
  border: 3px solid greenyellow;
}

.chat p {
  padding: 0;
  margin: 0;
}

#chat-messages-placeholder {
  position: relative;
  top: 50%;
  text-align: center;
  font-size: 22px;
}

#chat-messages {
  padding: 10px;
  overflow: auto;
  height: 520px;
}

.message {
  width: max-content;
  min-width: 180px;
  max-width: 500px;
  margin: 10px 10px 10px 10px;
  padding: 10px;
  border-radius: 5px;
  background-color: #3b3d42;
  color: #c9c9d8;
}

.message-text {
  margin: 0;
  margin-bottom: 10px;
}

.message-info {
  display: grid;
  grid-template-columns: auto 40px;
}

.message-info p {
  margin: 0;
  font-size: 14px;
  margin-left: auto;
  margin-right: 0;
}

.message-status {
  letter-spacing: -5px;
}

.sended-message {
  margin-left: auto;
  margin-right: 0;
}

#message-form {
  padding-left: 10px;
}

#message-input {
  width: 100%;
  height: 40px;
  padding: 10px;
  background-color: #3b3d42;
  color: #c9c9d8;
}
input,
select,
textarea {
  padding: 5px;
  border: none;
  border-radius: 5px;
  box-sizing: border-box;
  color: #a9a9b3;
  font-size: 16px;
  background-color: #292a2d;
  outline: none;
}

.typing-indicator {
  display: none;
  will-change: transform;
  padding: 5px 0;
  animation: 2s infinite ease-out;
}

.typing-indicator p {
  /* display: inline; */
  margin: 0;
  padding: 0;
}

.typing-indicator::before,
.typing-indicator::after {
  height: 15px;
  width: 15px;
  border-radius: 50%;
}

.typing-indicator span {
  height: 9px;
  width: 9px;
  float: left;
  margin: 0 1px;
  background-color: #c9c9d8;
  border-radius: 50%;
  opacity: 0.4;
}

.chat.active .typing-indicator span {
  background-color: #3b3d42;
}

.typing-indicator span:nth-of-type(1) {
  animation: 1s blink infinite 0.3333s;
}

.typing-indicator span:nth-of-type(2) {
  animation: 1s blink infinite 0.6666s;
}

.typing-indicator span:nth-of-type(3) {
  animation: 1s blink infinite 0.9999s;
}

@keyframes blink {
  50% {
    opacity: 1;
  }
}
</style>
