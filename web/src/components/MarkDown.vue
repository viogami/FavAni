<template>
  <div class="container0" >
    <el-col :xs="24" :sm="24" :md="20" :lg="20" :xl="20" v-html="parse()"/> <!-- eslint-disable-line -->
    <el-col :xs="0" :sm="0" :md="4" :lg="4" :xl="4"  v-if="props.toc" >
<!-- 右侧文章导航 -->
    <div  class="container1">
      <nav class="no-scrollbar font-medium leading-loose border-solid border-l-4 pl-4 mr-8">
<!-- 通过Vue.js的v-for指令，遍历tocList.c数组中的每个元素，item表示数组中的每个项目，i表示当前项目的索引。-->
        <ul v-for="(item, i) in tocList.c" :key="item">
<!-- text-emerald-grey类表示在悬停时文本颜色变为灰色。-->
<!-- 动态绑定类名，如果tocIndex等于当前项的索引i，则添加text-emerald-black类，表示高亮当前选定的目录项。-->
          <li>
            <a class="text-emerald-grey" :class="{ 'text-emerald-pink': tocIndex === i }" @click="tocIndex = i"
               :href=getHref(item.n)> {{ item.n }} </a>
            <!-- 二级标题 -->
            <ul v-for="(t, j) in item.c" :key="t">
              <li><a class="text-emerald-grey" :class="{ 'text-emerald-pink': tocIndex === (i + 1) * 1000 + j }"
                     @click="tocIndex = (i + 1) * 1000 + j" :href=getHref(t.n)> {{ t.n }}
              </a></li>
            </ul>
          </li>
        </ul>
      </nav>
    </div>
    </el-col>
  </div>
</template>

<script setup>
import MarkDownIt from 'markdown-it'
import prism from 'prismjs'
import MarkDownDoneRight from 'markdown-it-toc-done-right'
import uslug from 'uslug'
import { ref } from 'vue'
import anchor from 'markdown-it-anchor'

const props = defineProps({
  data: { type: String },
  toc: { type: Boolean, default: false }
})

const tocList = ref([])
const tocIndex = ref(-1)

const uslugify = (s) => {
  return uslug(s)
}

const md = MarkDownIt({
  breaks: true,
  html: true,
  linkify: true,
  typographer: true,
  highlight: (code, lang) => { // 代码高亮
    if (lang === 'golang') { // golang简写为go
      lang = 'go'
    }

    if (prism.languages[lang]) {
      return prism.highlight(code, prism.languages[lang], lang)
    } else {
      return code
    }
  }
}).use(anchor, {
  permalink: anchor.permalink.ariaHidden({ // ARIA hidden
    placement: 'before'
  })
}).use(MarkDownDoneRight, {
  slugify: uslugify,
  listType: 'ul',
  callback: function (html, ast) {
    if (tocList.value.length === 0) {
      tocList.value = ast
    }
  }
})

const parse = () => {
  if (!props.data) {
    return ''
  }
  return md.render(props.data)
}

const getHref = (target) => {
  return '#' + uslug(target)
}

</script>

<style scoped>
.container0{
  display: flex;
  flex-direction: row;
  margin-left: 15px;
  overflow: hidden;
}
.container1{
  position: fixed;
  top: 110px;
  bottom: 0;
  right: 100px;
}

.text-emerald-grey{
  color: grey;
}

.text-emerald-pink{
  color: pink;
}

</style>
