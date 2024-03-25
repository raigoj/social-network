import { reactive, readonly } from "vue";

const state = reactive({
  loggedIn: false,
  loggedInUsername: "",
});

const bobLogIn = () => {
  state.loggedIn = true;
  localStorage.loggedIn = true;
};

const bobLogOut = () => {
  state.loggedIn = false;
  localStorage.clear();
};

const bobSetUsername = (username) => {
  state.loggedInUsername = username;
  localStorage.loggedInUsername = username;
};

const bobDeleteUsername = () => {
  state.loggedInUsername = "";
};

export default {
  state: readonly(state),
  bobLogIn,
  bobLogOut,
  bobDeleteUsername,
  bobSetUsername,
};

/*  mutations: {
    bobLogIn(state) {
      state.loggedIn = true;
    },
    bobLogOut(state) {
      state.loggedIn = false;
    },
  },
  actions: {},
  modules: {},
  getters: {}, */
