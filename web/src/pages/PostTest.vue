<template>
  <div id="app">
    <input v-model="addr" placeholder="post地址" />
    <input v-model="username" placeholder="post内容" />
    <button @click="postData">Post Data</button>
  </div>
</template>

<script setup>
import {DelUser} from "../api/user.js";
import {ref} from "vue";

const addr = ref()
const username = ref()

const postData =() => {
  switch (addr){
    case '/deluser' :
      // 调用函数，post给后端
      DelUser(username.value)
        .then(res => {
          console.log(res)
        })
        .catch(err => {
          console.log(err)
        })
  }
}
</script>
<style scoped>

</style>
