<template>
  <div class="grid-createpost">
    <form @submit.prevent="createpost()">
      <div class="wrapper">
        <h4 class="info">Create a post</h4>

        <select class="categories" v-model="category">
          <option value="buy">Buy</option>
          <option value="sell">Sell</option>
          <option value="exchange">Exchange</option>
        </select>

        <div class="smallwrapper">
          <div>
            <textarea
              class="title"
              placeholder="Title"
              required
              v-model="title"
            ></textarea>
          </div>

          <div>
            <textarea
              class="content"
              placeholder="Text"
              required
              v-model="content"
            ></textarea>
          </div>

          <button class="postpost">Post</button>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import router from "@/router";
import axios from "axios";
import { onUnmounted, ref } from "vue";
const category = ref();
const title = ref();
const content = ref();
export default {
  setup() {
    async function createpost() {
      const url = "http://localhost:8009/createpost";
      const payload = {
        category: category.value,
        title: title.value,
        content: content.value,
        username: localStorage.loggedInUsername,
      };
      await axios
        .post(url, JSON.stringify(payload), {
          withCredentials: true,
        })
        .then(() => {
          router.push("/");
        })
        .catch((error) => {
          console.log(error);
        });
    }
    onUnmounted(() => {
      category.value = "";
      title.value = "";
      content.value = "";
    });
    return { category, title, content, createpost };
  },
};
</script>

<style>
.grid-createpost {
  display: grid;
  margin-left: 30px;
  margin-top: 30px;
  font-size: 2rem;
}

.title {
  height: 50px;
  border-radius: 10px;
}

.content {
  border-radius: 10px;
}

.info {
  margin: 0;
}
.categories {
  font-size: 2rem;
}
</style>
