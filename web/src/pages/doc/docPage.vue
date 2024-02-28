<template>
<!-- 文章内容界面 -->
<div style="width: 85%">
  <div style="padding-left: 25px;padding-right: 16px;">
    <div class="prose w-full">
      <h1> {{ title }} </h1>
      <div v-if="date" style="margin-bottom: 1rem;"> {{ date }} </div>
    </div>
    <!--  使用markdown编辑器  -->
    <MarkDown class="w-full mt-4 mb-8" :data="data" toc></MarkDown>
  </div>
</div>
</template>
<script setup>
import MarkDown from '../../components/MarkDown.vue'
import axios from 'axios'
import { ref, onMounted, watch } from 'vue'
import fm from 'front-matter'
// 定义props
const props = defineProps({
  config: Object
})

const title = ref(props.config.title)
const date = ref(props.config.date)
const data = ref('')

// 定义一个名为 parse 的函数，用于解析 Markdown 内容。
const parse = (content) => {
  if (!content) {
    return
  }
  // 使用 front-matter 库解析 Markdown 内容，将结果存储在 page 变量中
  const page = fm(content)
  /* 将 page是.body 中的 Markdown 内容存储到 data 变量中，并通过正则表达式替换掉所有以 # 开头的行。
 这是为了移除 Markdown 内容中的第一个标题。目的为了避免在页面中重复显示标题。 */
  data.value = page.body.replace(/#.+\n/, '')
  // 如果 title 变量为空，将元数据中的标题存储到 title 中
  if (!title.value) {
    title.value = page.attributes.title
  }
  // 如果 date 变量为空且元数据中存在日期，将元数据中的日期存储到 date 变量中。
  if (!date.value && page.attributes.date) {
    date.value = page.attributes.date
  }
  // 如果 date 变量存在，将其转换为本地日期格式。
  if (date.value) {
    date.value = (new Date(date.value)).toLocaleDateString('en', { year: 'numeric', month: 'short', day: 'numeric' })
  }
}

// 使用 watch 监听 props.config 的变化，并在变化时执行回调函数
watch(
  () => props.config,
  (val) => {
    getData(val)
  }
)

// 定义一个名为 getData 的函数，用于获取文章数据。
const getData = (config) => {
  title.value = config.title
  date.value = config.date
  // 取config中的数据
  // 如果有，用上面定义的解释器解释数据
  // 如果没有 从config这个object类型(title,url)的url参数进行后端调用
  if (config.content) {
    parse(config.content)
  } else if (config.url) {
    axios.get(config.url).then((response) => {
      parse(response.data)
    })
  }
}

// 使用 onMounted 钩子，在组件挂载时调用 getData 函数。
onMounted(
  () => getData(props.config)
)
</script>

<style scoped>
</style>
