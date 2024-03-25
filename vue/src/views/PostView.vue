<template>
  <div class="grid-post-container">
    <div class="grid-post-container2">
      <div class="post-title">{{ title }}</div>
      <div class="post-user">Posted by {{ username }}</div>
      <div class="post-content">{{ content }}</div>
      <div class="post-time">{{ formatDate(creationtime) }}</div>
    </div>
    <div class="grid-form">
      <form @submit.prevent="createcomment()">
        <div>
          <textarea v-model="commentcontent"></textarea>
          <button>Post</button>
        </div>
      </form>
    </div>
    <div>Comments:</div>
    <div v-for="comment in comments" :key="comment.Id" class="grid-comments">
      <div class="grid-comment">
        <div class="grid-comment-content">{{ comment.Content }}</div>
        <div>Commented by {{ comment.Username }}</div>
        <div>{{ formatDate(comment.Creationtime) }}</div>
      </div>
    </div>
  </div>
</template>

<script>
import router from "@/router";
import axios from "axios";
import { onUnmounted, ref } from "vue";
const commentcontent = ref();
function formatDate(input) {
  const date = new Date(input);
  const year = date.getFullYear();
  const month = `${date.getMonth() + 1}`.padStart(2, "0");
  const day = `${date.getDate()}`.padStart(2, "0");
  const hours = `${date.getHours()}`.padStart(2, "0");
  const minutes = `${date.getMinutes()}`.padStart(2, "0");
  return `${year}.${month}.${day} ${hours}:${minutes}`;
}
/*
https://www.reddit.com/r/vuejs/comments/s13biy/vue3_routerlink_cant_pass_value_to_params/
*/
export default {
  props: ["id", "username", "title", "content", "creationtime"],
  setup(props) {
    if (props.username === undefined) {
      router.push("/error404");
    }
    var comments = ref({});
    async function loadComments() {
      const url = "http://localhost:8009/loadcomments";
      const payload = {
        postid: props.id,
      };
      await axios
        .post(url, JSON.stringify(payload), {
          withCredentials: true,
        })
        .then((response) => {
          //console.log(response);
          comments.value = response.data;
          //console.log(comments);
        })
        .catch((error) => {
          console.log(error);
          //console.log("@@@@@@@@@@@@@@@@@@@@2");
        });
    }
    async function createcomment() {
      const url = "http://localhost:8009/createcomment";
      //console.log("this is postid" + props.id);
      const payload = {
        content: commentcontent.value,
        username: localStorage.loggedInUsername,
        postid: props.id,
      };
      await axios
        .post(url, JSON.stringify(payload), {
          withCredentials: true,
        })
        .then(() => {
          //console.log(response);
          loadComments();
          commentcontent.value = "";
        })
        .catch((error) => {
          console.log(error);
        });
    }
    loadComments();
    onUnmounted(() => {
      commentcontent.value = "";
    });
    /*onUpdated(() => {
      loadComments();
    });*/
    return {
      loadComments,
      commentcontent,
      createcomment,
      comments,
      formatDate,
    };
  },
  onBeforeMount(props) {
    console.log(props.Content);
    console.log("mounted");
  },
};
</script>

<style>
button {
  background: red;
  border: 0;
  padding: 10px 20px;
  color: white;
  border-radius: 15px;
  width: 170px;
  height: 55px;
  font-size: 2rem;
}
.grid-comment-content {
  margin-bottom: 25px;
}

.grid-post-container {
  margin-top: 30px;
  margin-left: 30px;
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: auto;
  grid-gap: 20px;
  font-size: 2rem;
  color: white;
}

.post-title {
  grid-area: title;
  margin-bottom: 10px;
}
.post-user {
  grid-area: user;
}
.post-content {
  grid-area: content2;
  margin-bottom: 10px;
}
.post-time {
  grid-area: time;
}

.grid-post-container2 {
  display: grid;
  grid-template-columns: 1fr 5fr;
  grid-template-rows: auto;
  grid-template-areas:
    "title user"
    "content2 content2"
    "time time";
}

.grid-form {
  display: grid;
}
textarea {
  width: 600px;
  height: 150px;
  font-size: 2rem;
  resize: none;
}

.grid-comments {
  display: grid;
  border: var(--clr-neon) 0.125em solid;
  border-radius: 25px;
  width: 1000px;
  padding: 15px;
  text-overflow: ellipsis;
  white-space: pre;
  overflow: clip;
  word-break: break-word;
  max-width: 100%;
}

.grid-comment {
  display: grid;
  grid-template-rows: auto;
  grid-template-columns: 1fr 5fr;
}

.grid-comment-content {
  width: 600px;
  white-space: initial;
}
</style>
