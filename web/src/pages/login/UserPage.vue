<template>
  <HeaderPage/>
<!-- 主体内容区 -->
<el-main class="UserPageMain">
  <!-- 主体内容 -->
  <el-avatar :size="60" :src="userAvatar" @error="errorHandler" >
    <img src="https://cube.elemecdn.com/e/fd/0fc7d20532fdaf769a25683617711png.png" alt="用户头像"/>
  </el-avatar>

  <!-- 用户昵称 -->
  <p>{{ nickname }}</p>

  <!-- 用户签名 -->
  <div>
    <h3>用户签名: {{ userSign }}</h3>
  </div>
  
  <!-- 用户收藏信息 -->
  <el-col style="display: flex;justify-content: center" >
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>用户收藏（当前第{{ page }}/ {{ totalpages }}页 ）</span>
        <el-button text :disabled="page === 1" @click="prevPage()">上一页</el-button>
        <el-button text :disabled="page === totalpages" @click="nextPage()">下一页</el-button>
      </div>
    </template>
    <div v-for="f_item in userProfile.getFavorList(page-1,8)" :key="f_item" class="text item">{{ f_item }}</div>
  </el-card>
  </el-col>
    
  <el-col style="display: flex;justify-content: center">
    <el-card >
      <template #header>
        <div class="card-header">
          <h3 style="color: pink"> 你可能也喜欢 </h3>
        </div>
      </template>
      <div> display soon .... </div>
    </el-card>
  </el-col>
  
</el-main>
</template>

<script setup>
import { useUserStore } from '../../store/userProfile.js'
import HeaderPage from '../../pages/Home/Header.vue'
import {userFavorite, userFavorite_Bangumi} from '../../api/user.js'
import { eleNotice } from '../../utils/notice.js'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const userProfile = useUserStore()
const username = userProfile.username
const nickname = userProfile.nickname
const userAvatar = userProfile.avatarUrl
const userSign = userProfile.sign === '' ? '该用户没有设置签名' : useUserStore().sign

const router = useRouter()

const subject_type = 2 // 1书籍 2动画 3音乐 4游戏 6三次元
const type = '' // 1: 想看 2: 看过 3: 在看 4: 搁置 5: 抛弃

// 头像加载失败
const errorHandler = () => true
// 定义页数，用于返回下一页用户收藏
const page = ref(1)
const totalpages = ref()
// 收藏拉取次数
const fetchcount = ref(1)

// 获取用户收藏,首次获取5页
function getfav_init() {
  if (userProfile.bangumiLogin) {
    userFavorite_Bangumi(username, subject_type, type, 40)
        .then(res => {
          userProfile.favorList = [] // 先清空，防止旧数据冗余
          userProfile.favorList_max = res.data.total
          // 循环res中收藏列表插入到store的favorlist中
          for (let item = 0; item < res.data.data.length; item++) {
            const name_cn = res.data.data[item].subject.name_cn
            name_cn === '' ? userProfile.addFavorList(res.data.data[item].subject.name) : userProfile.addFavorList(name_cn)
          }
          eleNotice('success', '该用户共有' + res.data.total + '个收藏条目!')
          // 设置总页数
          totalpages.value = Math.ceil(res.data.total / 8)
        })
        .catch(err => {
          eleNotice('error', '用户收藏请求失败~\n' + err.response.data.description)
        })
  } else {
    userFavorite(username)
        .then(res => {
          userProfile.favorList = [] // 先清空，防止旧数据冗余
          userProfile.favorList_max = res.data.data.length
          // 循环res中收藏列表插入到store的favorlist中
          for (let item = 0; item < res.data.data.length; item++) {
            const name_cn = res.data.data[item].data_anime.anime_cn
            name_cn === '' ? userProfile.addFavorList(res.data.data[item].data_anime.anime) : userProfile.addFavorList(name_cn)
          }
          eleNotice('success', '该用户共有' + res.data.data.length + '个收藏条目!')
          // 设置总页数
          totalpages.value = Math.ceil(res.data.data.length / 8)
        })
        .catch(err => {
          eleNotice('error', '非法进入用户界面,自动跳转主页~，error:\n' + err.message)
          router.push('/')
        })
  }
}
getfav_init()

// 获取用户收藏
const fetchData = () => {
  if (userProfile.bangumiLogin) {
    userFavorite_Bangumi(username, subject_type, type, 40, 40 * fetchcount.value)
        .then(res => {
          // 循环res中收藏列表插入到store的favorlist中
          for (let item = 0; item < res.data.data.length; item++) {
            const name_cn = res.data.data[item].subject.name_cn
            name_cn === '' ? userProfile.addFavorList(res.data.data[item].subject.name) : userProfile.addFavorList(name_cn)
          }
        })
        .catch(err => {
          eleNotice('error', '用户收藏请求失败~\n' + err.response.data.description)
        })
    fetchcount.value++
  } else {
    userFavorite(username)
        .then(res => {
          // 循环res中收藏列表插入到store的favorlist中
          for (let item = 0; item < res.data.data.length; item++) {
            const name_cn = res.data.data[item].data_anime.anime_cn
            name_cn === '' ? userProfile.addFavorList(res.data.data[item].data_anime.anime) : userProfile.addFavorList(name_cn)
          }
        })
        .catch(err => {
          eleNotice('error', '用户收藏请求失败~\n' + err.response.data.error)
        })
    fetchcount.value++
  }
}
// 翻页函数 上一页
const prevPage = () =>{
  if(page.value>1){
     page.value--
  }
}
// 翻页函数 下一页
const nextPage = () => {
  if(page.value<totalpages.value){
    page.value++
  }
  if (40*fetchcount.value < userProfile.favorList_max){
    fetchData()
  }
}
</script>

<style scoped>
.UserPageMain {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin: 0;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.box-card {
  width: 100%;
}
</style>
