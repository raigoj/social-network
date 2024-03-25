import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import PostView from "../views/PostView.vue";
import RegisterView from "@/views/RegisterView.vue";
import LoginView from "@/views/LoginView.vue";
import LogoutView from "@/views/LogoutView.vue";
import CreatePostView from "@/views/CreatePostView.vue";
import ErrorView from "@/views/ErrorView.vue";
import ChatView from "@/views/ChatView.vue";
import PostsView from "@/views/PostsView.vue";
import store from "../store.js";

const routes = [
  {
    path: "/",
    name: "Home",
    component: HomeView,
  },
  {
    path: "/posts",
    name: "Posts",
    component: PostsView,
    meta: { requiresAuth: true },
  },
  {
    path: "/posts/:id",
    name: "Postdetails",
    component: PostView,
    props: true,
  },
  {
    path: "/register",
    name: "Register",
    component: RegisterView,
    meta: { guest: true },
  },
  {
    path: "/login",
    name: "Login",
    component: LoginView,
    meta: { guest: true },
  },
  {
    path: "/logout",
    name: "Logout",
    component: LogoutView,
  },
  {
    path: "/createpost",
    name: "Createpost",
    component: CreatePostView,
    meta: { requiresAuth: true },
  },
  {
    path: "/chat",
    name: "Chat",
    component: ChatView,
    meta: { requiresAuth: true },
    //props: (route) => ({ foo: route.query.foo }),
  },
  {
    path: "/:catchAll(.*)",
    name: "notFound",
    component: ErrorView,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});
router.beforeResolve((to) => {
  //console.log("routerist tule 3 consolelogi");
  //console.log(store.state.loggedIn);
  //console.log(to.meta.guest);
  //console.log(store.state.loggedInUsername);
  if (to.meta.requiresAuth && !store.state.loggedIn) {
    return { name: "Login" };
  } else if (to.meta.guest && store.state.loggedIn) {
    //console.log("vanamiis");
    return { name: "Home" };
  } else {
    return true;
  }
});

export default router;
