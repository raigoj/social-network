<template>
  <div class="grid-login">
    <form @submit.prevent="registerUser">
      <div class="form">
        <input
          id="username"
          type="text"
          required
          v-model="username"
          class="form__input"
          autocomplete="off"
          placeholder=" "
        />
        <label for="username" class="form__label">Username</label>
      </div>
      <div v-if="usernameError" class="error">{{ usernameError }}</div>

      <div class="form">
        <input
          id="age"
          class="form__input"
          type="date"
          required
          v-model="age"
          autocomplete="off"
          placeholder=" "
        />
        <label for="age" class="form__label">Age</label>
      </div>
      <div v-if="ageError" class="error">{{ ageError }}</div>

      <div class="form genderindput">
        <select v-model="gender" id="gender" class="form__input">
          <option value="male">Male</option>
          <option value="female">Female</option>
        </select>
        <label for="gender" class="form__label">Gender</label>
      </div>

      <div class="form">
        <input
          id="firstname"
          class="form__input"
          type="text"
          required
          v-model="firstname"
          autocomplete="off"
          placeholder=" "
        />
        <label for="firstname" class="form__label">First name:</label>
      </div>
      <div v-if="firstnameError" class="error">{{ firstnameError }}</div>

      <div class="form">
        <input
          id="lastname"
          class="form__input"
          type="text"
          required
          v-model="lastname"
        />
        <label for="lastname" class="form__label">Last name:</label>
      </div>
      <div v-if="lastnameError" class="error">{{ lastnameError }}</div>

      <div class="form">
        <input
          id="email"
          class="form__input"
          type="email"
          required
          v-model="email"
        />
        <label for="email" class="form__label">Email:</label>
      </div>
      <div v-if="emailError" class="error">{{ emailError }}</div>

      <div class="form">
        <input
          id="password"
          class="form__input"
          type="password"
          required
          v-model="password"
        />
        <label for="password" class="form__label">Password:</label>
      </div>

      <div class="form">
        <input
          id="password2"
          class="form__input"
          type="password"
          required
          v-model="password2"
        />
        <label for="password2" class="form__label">Retype password</label>
      </div>
      <div v-if="passwordError" class="error">{{ passwordError }}</div>

      <div class="submit">
        <button>Register</button>
      </div>
    </form>
  </div>
</template>

<script>
import router from "@/router";
import axios from "axios";
import { onUnmounted, ref } from "vue";
var ageError = ref();
var emailError = ref();
var usernameError = ref();
var firstnameError = ref();
var lastnameError = ref();
var passwordError = ref();
var username = ref();
var age = ref();
var gender = ref();
var firstname = ref();
var lastname = ref();
var email = ref();
var password = ref();
var password2 = ref();
export default {
  setup() {
    async function registerUser() {
      passwordError.value =
        password.value === password2.value ? "" : "Passwords don't match";
      ageError.value =
        new Date() > new Date(age.value) ? "" : "You from the future?";
      firstnameError.value = /^[A-Za-z]+$/.test(firstname.value)
        ? ""
        : "Letters only";
      lastnameError.value = /^[A-Za-z]+$/.test(lastname.value)
        ? ""
        : "Letters only";
      if (
        !passwordError.value &&
        !ageError.value &&
        !lastnameError.value &&
        !firstnameError.value
      ) {
        const payload = {
          username: username.value,
          age: age.value,
          gender: gender.value,
          firstname: firstname.value,
          lastname: lastname.value,
          email: email.value,
          password: password.value,
        };
        console.log(payload);
        console.log(passwordError.value);
        axios
          .post("http://localhost:8009/register", JSON.stringify(payload))
          .then((response) => {
            console.log(response);
            console.log("success");
            //user registred, redirect to login or home or login
            router.push("/login");
          })
          .catch((error) => {
            if (error.response.data.EmailTaken) {
              emailError.value = "This email already taken";
            }
            if (error.response.data.UsernameTaken) {
              usernameError.value = "This username already taken";
            }
          });
      }
    }
    onUnmounted(() => {
      username.value = "";
      age.value = "";
      gender.value = "";
      firstname.value = "";
      lastname.value = "";
      email.value = "";
      password.value = "";
      password2.value = "";
    });
    return {
      registerUser,
      ageError,
      emailError,
      usernameError,
      firstnameError,
      lastnameError,
      passwordError,
      username,
      age,
      gender,
      firstname,
      lastname,
      email,
      password,
      password2,
    };
  },
};

/*
  data() {
    return {
      username: "",
      age: "",
      gender: "",
      firstname: "",
      lastname: "",
      email: "",
      password: "",
      password2: "",
      usernameError: "",
      emailError: "",
      passwordError: "",
      ageError: "",
      firstnameError: "",
      lastnameError: "",
    };
  },
  methods: {
    registerUser() {
      this.ageError = "";
      this.emailError = "";
      this.usernameError = "";
      this.firstnameError = "";
      this.lastnameError = "";

      this.passwordError =
        this.password === this.password2 ? "" : "Passwords don't match";
      this.ageError =
        new Date() > new Date(this.age) ? "" : "You from the future?";
      this.firstnameError = /^[A-Za-z]+$/.test(this.firstname)
        ? ""
        : "Letters only";
      this.lastnameError = /^[A-Za-z]+$/.test(this.lastname)
        ? ""
        : "Letters only";

      if (
        !this.passwordError &&
        !this.ageError &&
        !this.lastnameError &&
        !this.firstnameError
      ) {
        const payload = {
          username: this.username,
          age: this.age,
          gender: this.gender,
          firstname: this.firstname,
          lastname: this.lastname,
          email: this.email,
          password: this.password,
        };
        axios
          .post("http://localhost:8009/register", JSON.stringify(payload))
          .then((response) => {
            console.log(response);
            console.log("success");
            //user registred, redirect to login or home or login
            router.push("/login");
          })
          .catch((error) => {
            if (error.response.data.EmailTaken) {
              this.emailError = "This email already taken";
            }
            if (error.response.data.UsernameTaken) {
              this.usernameError = "This username already taken";
            }
          });
      }
    },
  },*/
/*
#E8D7FF
#81968f
#fcd29f
#3185fc
#3a5683

form {
  max-width: 420px;
  margin: 30px auto;
  background: white;
  text-align: left;
  padding: 40px;
  border-radius: 10px;
}
label {
  color: #aaa;
  display: inline-block;
  margin: 25px 0 15px;
  font-size: 0.6em;
  text-transform: uppercase;
  letter-spacing: 1px;
  font-weight: bold;
}
input,
select {
  display: block;
  padding: 10px 6px;
  width: 100%;
  box-sizing: border-box;
  border: none;
  border-bottom: 1px solid #ddd;
  color: #555;
}
button {
  background: #0b6dff;
  border: 0;
  padding: 10px 20px;
  margin-top: 20px;
  color: white;
  border-radius: 20px;
}
.submit {
  text-align: center;
}
.error {
  color: #ff0062;
  margin-top: 10px;
  font-size: 0.8em;
  font-weight: bold;
}
*/
</script>

<style>
.genderindput {
  height: 5rem;
  margin-bottom: 30px;
}
</style>
