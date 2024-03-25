import axios from "axios";
import VueAxios from "vue-axios";
import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
//import connection from "./connection";

const app = createApp(App);
app.provide("store", store);
// dno kas seda vaja
//app.provide("connection", connection);

app.use(router).use(VueAxios, axios).mount("#app");
