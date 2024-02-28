<template>
<!-- 文章目录页面，负责导入文章并实现文章路由(路由匹配) -->
  <docPage :config="item"> </docPage>
</template>

<script setup>
import docPage from './docPage.vue'
import { useRoute } from 'vue-router'
import { ref } from 'vue'

const route = useRoute()
const page = route.params.page

// 定义一个pages对象，用于路由匹配，其元素就是config的具体
const pages = {
  introduce: {
    title: 'Introduce',
    url: 'https://raw.githubusercontent.com/viogami/FavAni_FE/main/README.md'
  },
  frontend: {
    title: 'Front End',
    content: '这将介绍前端技术栈 '
  },
  backend: {
    title: 'Back End',
    content: '这将介绍后端技术栈'
  }
}

// 定义传入doc-page的props的objects类型参数
const item = ref(pages[page])

</script>
