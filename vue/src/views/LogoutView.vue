<script>
import router from "@/router";
import { inject } from "vue";
import axios from "axios";
import Ws from "@/connection";
export default {
  setup() {
    const store = inject("store");
    axios
      .get("http://localhost:8009/logout", {
        withCredentials: true,
      })
      .then(() => {
        store.bobLogOut();
        store.bobDeleteUsername();
        Ws.disconnect();
        router.push("/");
        console.log("login v2lja");
      })
      .catch(() => {
        console.log("logout failed");
        router.push("/");
      });
  },
};
</script>
