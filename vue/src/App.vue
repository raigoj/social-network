<template>
  <nav>
    <div class="grid-container">
      <router-link class="home neon-button" :to="{ name: 'Home' }"
        >Home</router-link
      >
      <!-- v-if="store.state.loggedIn" -->
      <router-link
        class="posts neon-button"
        v-if="store.state.loggedIn"
        :to="{ name: 'Posts' }"
        >Posts</router-link
      >
      <router-link
        class="messages neon-button"
        v-if="store.state.loggedIn"
        :to="{ name: 'Chat' }"
        >Messages
        <div
          id="displayUnreadMessage"
          style="
            display: none;
            width: 25px;
            height: 25px;
            border-radius: 50%;
            background-color: white;
          "
        ></div>
      </router-link>
      <router-link
        class="createpost neon-button"
        v-if="store.state.loggedIn"
        :to="{ name: 'Createpost' }"
        >Create post</router-link
      >
      <router-link
        class="logout neon-button"
        v-if="store.state.loggedIn"
        :to="{ name: 'Logout' }"
        >Logout</router-link
      >
      <!-- v-if="!store.state.loggedIn" -->
      <div v-if="!store.state.loggedIn"></div>
      <div v-if="!store.state.loggedIn"></div>
      <router-link
        class="login neon-button"
        v-if="!store.state.loggedIn"
        :to="{ name: 'Login' }"
        >Login</router-link
      >
      <router-link
        class="register neon-button"
        v-if="!store.state.loggedIn"
        :to="{ name: 'Register' }"
        >Register</router-link
      >
    </div>
  </nav>
  <router-view />
</template>
<script>
export async function unreadMessageDisplayOn() {
  console.log("lisan punkti");
  let elem = document.getElementById("displayUnreadMessage");
  elem.style.display = "block";
}
export async function unreadMessageDisplayOff() {
  console.log("yritan maha v5tta punkti");
  let elem = document.getElementById("displayUnreadMessage");
  elem.style.display = "none";
}
</script>

<script setup>
import { inject } from "vue";
import Ws from "@/connection";
const store = inject("store");
//const connection = inject("connection");
/* this.$root.$on("sendMessage", () => {
  // your code goes here
  connection.send(JSON.stringify({ type: "type", body: "msg" }));
}); */
//console.log(localStorage.loggedIn + "loggedin");
//console.log(localStorage.loggedInUsername + "loggedinusername");

if (localStorage.loggedIn && document.cookie.length > 0) {
  store.bobLogIn();
  Ws.connect();
} else {
  store.bobLogOut();
}
if (localStorage.loggedInUsername) {
  store.bobSetUsername(localStorage.loggedInUsername);
}

/* onMounted(() => {
  connection.onopen = function () {
    connection.send(JSON.stringify({ type: "token", body: "ws t66tas" }));
  };

  connection.onmessage = async function (evt) {
    console.log("Receiving transmission:");
    console.log(evt);
  };
}); */

/* const connection = new WebSocket("ws://localhost:8009/socket");
connection.onmessage = function (message) {
  const data = JSON.parse(message.data);
  console.log(data);
};
let msg = {
  type: "tere",
  body: "mina",
};
try {
  connection.send(JSON.stringify(msg));
  console.log("success saatmisel?");
} catch (err) {
  console.log(err);
} */

/*
#E8D7FF
#81968f
#fcd29f
#3185fc
#3a5683
*/
/*
dae0e6
0079d3
ffffff
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
body {
  background-color: #dae0e6;
}

nav {
  position: relative;
  padding: 5px;
  background-color: #ffffff;
  height: 20px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}

.posts {
  position: absolute;
  margin-left: 0px;
  color: #0079d3;
}

.createpost {
  position: absolute;
  right: 200px;
  color: #0079d3;
}

.register {
  position: absolute;
  right: 0px;
  color: #0079d3;
}

.login {
  position: absolute;
  right: 200px;
  color: #0079d3;
}

.logout {
  position: absolute;
  right: 0px;
  color: #0079d3;
}

.messages {
  position: absolute;
  right: 550px;
  color: #0079d3;
}
*/
</script>

<style>
.grid-container {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
  grid-auto-rows: 100px;
}
body {
  place-items: center;
  background: var(--clr-bg);
  font-family: "Balsamiq Sans", cursive;
  color: var(--clr-neon);
}
:root {
  --clr-neon: #dcf900;
  --clr-bg: rgb(68, 67, 67);
}
nav a.router-link-exact-active {
  color: #42b983;
}
nav a.router-link-exact-active:hover {
  box-shadow: inset 0 0 0.5em 0 var(--clr-neon), 0 0 0.5em 0 var(--clr-neon);
}

.neon-button {
  font-size: 2.5rem;

  display: inline-block;
  cursor: pointer;
  text-decoration: none;
  color: var(--clr-neon);
  border: var(--clr-neon) 0.125em solid;
  padding: 0.25em 1em;
  border-radius: 0.25em;

  text-shadow: 0 0 0.125em hsl(0 0% 100% / 0.3), 0 0 0.45em currentColor;

  box-shadow: inset 0 0 0.5em 0 var(--clr-neon), 0 0 0.5em 0 var(--clr-neon);

  position: relative;
}

.neon-button:hover {
  box-shadow: inset 0 0 0.5em 0 var(--clr-neon), 0 0 2em 0.5em var(--clr-neon);
}
</style>
