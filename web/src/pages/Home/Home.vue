<template>
<!--页头-->
  <HeaderPage :menuIndex="'1'"/>

<!--侧边栏-->
<el-container>
<el-col :xs="6" :sm="6" :md="6" :lg="6" :xl="6">
<el-aside width="250px" class="sidebar">
  <h4 style="text-align:start;margin-left: 30px;">查询类型</h4>
  <el-menu
      default-active="全部类型"
      class="el-menu-vertical-demo"
      :collapse="isCollapse"
      @select="index => {
        searchRes.searchType=index
        ElNotification({
          message: '查询类型已经切换到：' + index,
          type: 'success',
          duration: 2000 // 持续两秒
        })}"
  >
    <el-menu-item index="全部类型">
      <el-icon><icon-menu /></el-icon>
      <span>全部类型</span>
    </el-menu-item>
    <el-menu-item index="书籍">
      <el-icon><reading /></el-icon>
      <span>书籍</span>
    </el-menu-item>
    <el-menu-item index="动画">
      <el-icon><video-play /></el-icon>
      <span>动画</span>
    </el-menu-item>
    <el-menu-item index="音乐" >
      <el-icon><headset /></el-icon>
      <span>音乐</span>
    </el-menu-item>
    <el-menu-item index="游戏">
      <el-icon><switch-filled /></el-icon>
      <span>游戏</span>
    </el-menu-item>
    <el-menu-item index="三次元">
      <el-icon><avatar /></el-icon>
      <span>三次元</span>
    </el-menu-item>
  </el-menu>
</el-aside>
</el-col>
<!--搜索框-->
<el-main class="main-container">
  <!--  走马灯-->
  <el-col :xs="18" :sm="18" :md="18" :lg="18" :xl="18" >
    <el-carousel :interval="4000" type="card" height="150px">
      <el-carousel-item  v-for="(item, index) in carouselItem" :key="index">
        <h3>{{ item }}</h3>
      </el-carousel-item>
    </el-carousel>
  </el-col>
  
  <el-col :xs="18" :sm="18" :md="18" :lg="18" :xl="18" style="margin-top: 30px" >
  <SearchPage/>
  </el-col>
</el-main>
</el-container>

<!--侧边栏放缩图标按钮-->
<div class="collapseIcon">
  <div v-if="isCollapse">
    <el-icon v-model="isCollapse" size="20" @click="collapseMenu">
      <fold/>
    </el-icon>
  </div>
  <div v-else>
    <el-icon v-model="isCollapse" size="20" @click="collapseMenu">
      <expand/>
    </el-icon>
  </div>
</div>

<!--页脚-->
<footer-page/>
</template>


<script lang="ts" setup>
import SearchPage from './SearchPage.vue'
import HeaderPage from './Header.vue'
import FooterPage from './footer.vue'
import { useSearchEntryStore } from '../../store/SearchSubject.js'
import { ref } from 'vue'
import {
  Menu as IconMenu,
  Headset,
  SwitchFilled,
  VideoPlay,
  Avatar,
  Reading,
  Expand,
  Fold
} from '@element-plus/icons-vue'
import { ElNotification } from "element-plus";

// 侧边栏折叠面板逻辑
const isCollapse = document.body.clientWidth < 950 ? ref(true) : ref(false)
const collapseMenu = () => {
  isCollapse.value = !isCollapse.value
}
const searchRes = useSearchEntryStore()

const carouselItem = ['First  --> search your favourite anime', 'Second --> add it to your fav list', 'Third --> Get the recommendation ']
</script>


<style scoped>
.sidebar {
  padding-bottom: 25px;
  margin-top: 40px;
}

.collapseIcon {
  margin-top: 1rem;
  text-align: left;
  padding-left: 1.5rem;
}
.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>
