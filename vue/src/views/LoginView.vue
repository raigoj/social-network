<template>
  <div class="grid-login">
    <form @submit.prevent="loginUser()" class="loginform">
      <div>
        <div class="form">
          <input
            required
            v-model="username"
            type="text"
            id="email"
            class="form__input"
            autocomplete="off"
            placeholder=" "
          />
          <label for="email" class="form__label">Username or email</label>
        </div>
      </div>
      <div>
        <div class="form">
          <input
            required
            v-model="password"
            type="password"
            id="password2"
            class="form__input"
            autocomplete="off"
            placeholder=" "
          />
          <label for="password2" class="form__label">Password</label>
        </div>
        <div v-if="inputError" class="error">{{ inputError }}</div>
      </div>

      <div class="submit">
        <button>Log in</button>
      </div>
    </form>
  </div>
</template>

<script>
import router from "@/router";
import axios from "axios";
import { onUnmounted, ref } from "vue";
import { inject } from "vue";
import Ws from "@/connection";
const username = ref();
const password = ref();
const inputError = ref();
export default {
  setup() {
    const store = inject("store");
    async function loginUser() {
      const url = "http://localhost:8009/login";
      const payload = {
        username: username.value,
        password: password.value,
      };
      await axios
        .post(url, JSON.stringify(payload), {
          withCredentials: true,
        })
        .then((response) => {
          if (response.status === 200) {
            console.log("You are logged in");
            store.bobLogIn();
            store.bobSetUsername(payload.username);
            Ws.connect();
            router.push("/");
          }
        })
        .catch(() => {
          console.log("Login failed");
          inputError.value = "Invalid username or password";
          console.log(inputError.value);
        });
    }
    onUnmounted(() => {
      username.value = "";
      password.value = "";
      inputError.value = "";
    });
    return { username, password, inputError, loginUser };
  },
};
</script>

<style>
.grid-login {
  display: grid;
  grid-template-columns: 1fr;
  grid-auto-rows: auto;
  color: red;
  margin-top: 50px;
  margin-left: 50px;
}
.form {
  font-size: 1rem;
  position: relative;
  width: 20rem;
  height: 4rem;
  margin-bottom: 65px;
}
.form__input {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border: 2px solid gray;
  border-radius: 0.5rem;
  font-family: inherit;
  font-size: inherit;
  color: white;
  outline: none;
  padding: 1.25rem;
  background: none;
  /* Change border when input focus*/
}
.form__input:hover {
  border-color: red;
}
.form__input:focus {
  border-color: red;
}
.form__label {
  position: absolute;
  left: 1rem;
  top: 0.8rem;
  padding: 0 0.5rem;
  color: black;
  cursor: text;
  background-color: #dcf900;
  top: -0.5rem;
  font-size: 0.8rem;
  left: 0.8rem;
  font-size: 1rem;
}
.error {
  font-size: 1.5rem;
  margin-bottom: 1rem;
}
</style>
