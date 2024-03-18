<template>
  <!-- 添加页头 -->
  <Header />
    <el-col
        v-for="items in searchRes.getSearchResList(page-1,page_items)"
        :key="items.name"
        :offset="2"
        :span="20"
    >
    <el-card
        shadow="hover"
        style="margin-top: 10px;"
    >
      <el-row >
      <img
          :src="items.img"
          class="image"
          alt="no img"/>
      <div style="padding: 10px;">
        <span>条目类型：{{ getTypeName(items.type) }}</span>
        <br>
        <span>原名：{{ items.name }}</span>
        <br>
        <span>中文名：{{ items.name_cn === ''?'暂无':items.name_cn }}</span>
        <br>
        <span>bangumi评分：</span>
        <el-rate
            v-model="items.rate"
            :max="10"
            disabled
            show-score
            text-color="#ff9900"
            score-template="{value}分"
        />
      </div>
      </el-row>
    </el-card>
    <el-button  @click="addToFav(items.id,items.name,items.name_cn)">添加到收藏</el-button>
    </el-col>
    <el-pagination 
        v-model:current-page="page"
        @change="changePage()"
        background 
        layout="prev, pager, next"
        :total="searchRes.searchResList_max" 
        style="margin-top: 30px;margin-bottom:20px;display:flex;justify-content: center"
    />
</template>

<script setup>
import {GetSubjectById, SearchSubject} from "../../api/subject.js";
import {ElNotification,ElMessage} from "element-plus";
import {useSearchEntryStore} from "../../store/SearchEntry.js";
import {useUserStore} from "../../store/userProfile.js";
import {useRoute} from "vue-router";
import Header from "../Home/Header.vue";
import {ref} from "vue";
import router from "../../router/index.js";
import { AddFav } from "../../api/fav.js";
// 路由实例
const route = useRoute()
// 使用查询结果的store，读取查询类型
const searchRes = useSearchEntryStore() 

const searchType = Number(route.query.cat)
const keywords = route.params.keywords
const responseGroup = 'large'
const start = 0
const max_results = 25
// 当前页数
const page = ref(Number(route.query.page))
const totalpages = ref()
const page_items = 10

SearchSubject(keywords, searchType, responseGroup, start, max_results)
    .then(res => {
      eleNotice('success','成功查询' + keywords + '！共有' + res.data.results + '个条目！')
      // 先清空结果列表，防止旧数据冗余
      searchRes.searchResList = []
      // 设置查询结果store信息
      searchRes.searchResList_max = res.data.results
      // 设置总页数
      totalpages.value=Math.ceil(res.data.results / page_items)
      // 循环拉取元素
      if (searchRes.searchResList.length <= searchRes.searchResList_max){
        for (let i = 0;  i*max_results < res.data.results; i++) {
          fetchData(i)
        }
      }
    })
    .catch(err => {
      eleNotice('error','条目查询失败~\n' + err)
    })
// 拉取条目
const fetchData = (fetchCount) => {
  SearchSubject(keywords, searchType, responseGroup, fetchCount*max_results, max_results)
      .then(res => {
        // 循环插入元素
        for (let i = 0;  i < max_results; i++) {
          const currentItem = res.data.list[i] ?? '';
          const currentItemImg = currentItem.images ? currentItem.images.common : 'https://cube.elemecdn.com/e/fd/0fc7d20532fdaf769a25683617711png.png'
          const currentItemRate = currentItem.rating ? currentItem.rating.score : "暂无评"
          searchRes.addSearchRes(currentItem.type,currentItem.id,currentItemImg, currentItem.name, currentItem.name_cn,currentItemRate);
        }
      })
      .catch(err => {
        eleNotice('error','第' + fetchCount + '批次条目加载异常\n' + err)
      })
}

// 跳页函数
const changePage = () => {
  // 跳转到条目结果页面
  router.push({
    name: 'SubjectSearch',
    params: { keywords:keywords },
    query: { cat: searchRes.getSearchType(),page:page.value }
  })
}

// 使用用户信息的store
const userProfile = useUserStore() 
// 根据id获取标签
const fetchTags = (id) => {
  return GetSubjectById(id)       //为了取消异步，此行必须return
    .then(response => {
      return response.data.tags
    })
    .catch(error => {
      eleNotice('error', '根据ID获取条目失败:'+ error )
    })
}
// 添加到收藏
const addToFav = async (id,anime,anime_cn) => {
  if (userProfile.username === '') {
    eleNotice('warning', '添加收藏前，请先登录~')
    return
  }
  // 获取条目tags
  const tags = await fetchTags(id)
  // 构造收藏数据
  const data_anime = {anime_id: id, anime: anime,anime_cn:anime_cn,tags: tags}
  // 添加收藏
  AddFav(userProfile.username, id,data_anime)
    .then(res => {
      eleNotice('success', '添加收藏成功!')
    })
    .catch(err => {
      eleNotice('error', '添加收藏失败:'+ err.response.data.error )
    })
}

// 条目类型转换函数
function getTypeName(subject){
  switch (subject) {
    case 1:
      return '书籍'
    case 2 :
      return '动画'
    case 3:
      return '音乐'
    case 4:
      return '游戏'
    case 6:
      return '三次元'
  }
  return 'unknown'
}
// 通知显示函数
function eleNotice(type,msg){
  switch (type) {
    case 'success':
      ElNotification({
            message: msg,
            type: 'success',
            duration: 2000
          }
      )
      break
    case 'warning':
      ElMessage({
        message: msg,
        type: 'warning',
        duration: 2000,
      })
      break
    case 'error':
      ElNotification({
        title: 'ERROR',
        message: msg,
        type: 'error',
        duration: 2000
      })
      break
    default:
      ElNotification({
        title: 'ERROR',
        message: 'please input correct notice type ,it`s a string',
        type: 'error',
        duration: 2000
      })
  }
}
</script>


<style scoped>
</style>