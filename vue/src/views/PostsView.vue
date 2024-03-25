<template>
  <div class="grid-posts">
    <div class="grid-buttons">
      <a class="grid-buttons2" href="#" v-on:click="loadPosts(0)">All</a>
      <a class="grid-buttons2" href="#" v-on:click="loadPosts(1)">Buy</a>
      <a class="grid-buttons2" href="#" v-on:click="loadPosts(2)">Sell</a>
      <a class="grid-buttons2" href="#" v-on:click="loadPosts(3)">Exchange</a>
    </div>
    <div class="grid-post">
      <div v-for="post in posts" :key="post.Id">
        <div class="grid-text">
          <router-link
            :to="{
              name: 'Postdetails',
              params: {
                id: post.Id,
                title: post.Title,
                content: post.Content,
                username: post.Username,
                creationtime: post.Creationtime,
              },
            }"
          >
            <div class="grid-info">
              <div class="title">{{ post.Title }}</div>
              <div>Posted by: {{ post.Username }}</div>
              <div>{{ formatDate(post.Creationtime) }}</div>
            </div>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
/*
https://www.reddit.com/r/vuejs/comments/s13biy/vue3_routerlink_cant_pass_value_to_params/
*/
import axios from "axios";
import { ref } from "vue";
var posts = ref();
function formatDate(input) {
  const date = new Date(input);
  const year = date.getFullYear();
  const month = `${date.getMonth() + 1}`.padStart(2, "0");
  const day = `${date.getDate()}`.padStart(2, "0");
  const hours = `${date.getHours()}`.padStart(2, "0");
  const minutes = `${date.getMinutes()}`.padStart(2, "0");
  return `${year}.${month}.${day} ${hours}:${minutes}`;
}
export default {
  setup() {
    async function loadPosts(category) {
      const url = "http://localhost:8009/home";
      const payload = {
        category: category,
      };
      await axios
        .post(url, JSON.stringify(payload), {
          withCredentials: true,
        })
        .then((response) => {
          posts.value = response.data;
        })
        .catch((error) => {
          console.log(error);
        });
    }
    return { loadPosts, posts, formatDate };
  },
};
</script>
<style>
.grid-posts {
  display: grid;
  grid-template-columns: 270px 1fr;
  grid-template-areas:
    "sidebar content"
    "sidebar content"
    "sidebar content"
    "sidebar content"
    "sidebar content";
}

.grid-posts {
  margin-top: 20px;
}
.grid-post {
  display: grid;
  grid-gap: 20px;
}
.grid-buttons {
  display: grid;
  grid-auto-rows: 100px;
  grid-area: sidebar;
}

.grid-buttons2 {
  font-size: 2.5rem;
  color: white;
  text-decoration: none !important;
  padding: 1rem 2rem;
}

.grid-text {
  display: grid;
  border: var(--clr-neon) 0.125em solid;
  border-radius: 15px;
  text-decoration: none;
}
.grid-info {
  display: grid;
  grid-template-columns: 1fr 200px;
  color: white;
  font-size: 1.25rem;
  padding: 10px;
  grid-gap: 25px;
}

.grid-buttons2:focus {
  color: red;
}

.grid-buttons2:hover {
  box-shadow: inset 0 0 0.5em 0 white, 0 0 2em 0.5em white;
  color: red;
}

a {
  text-decoration: none;
}
</style>
