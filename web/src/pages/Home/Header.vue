<template>
  <el-container>
    <el-header class="fixed-header">
      <el-row>
        <el-col :xs="8" :sm="6" :md="6" :lg="6" :xl="6">
          <router-link to="/" >
            <h2 style="color: pink;">FavAni</h2>
          </router-link>
        </el-col>

        <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
          <el-menu mode="horizontal"  text-color="black" :default-active="menuIndex" >
            <el-menu-item index="1">
              <router-link to="/" >Home</router-link>
            </el-menu-item>
            <el-menu-item index="2">
              <router-link to="/docs" >Document</router-link>
            </el-menu-item>
            <el-menu-item index="3">
              <router-link to="/timeline">Timeline</router-link>
            </el-menu-item>
            <el-menu-item index="4">
              <router-link to="/about">About</router-link>
            </el-menu-item>
          </el-menu>
        </el-col>
        
        <el-col :xs="4" :sm="6" :md="6" :lg="6" :xl="6"  style="margin-top: 20px;">
            <el-button>
              <a href="https://github.com/viogami/FavAni" target="_blank">
                <img :src="github_icon" alt="GitHub Icon" width="18" height="18" />
              </a>
            </el-button>
          <span style="margin-right: 20px"></span>
            <el-dropdown>
              <el-button>
                <el-icon style="width: 18px"><UserFilled /></el-icon>
              </el-button>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item @click="toUserInfo">
                    <el-icon><User/></el-icon>
                    UserInfo
                  </el-dropdown-item>
                </el-dropdown-menu>
                <el-dropdown-menu>
                  <el-dropdown-item @click="notImplement('Setting')">
                    <el-icon><Setting /></el-icon>
                    Setting
                  </el-dropdown-item>
                </el-dropdown-menu>
                <el-dropdown-menu>
                  <el-dropdown-item v-if="userProfile.username === ''"  @click="router.push('/login')">
                    <el-icon><Promotion /></el-icon>
                    LogIn
                  </el-dropdown-item >
                  <el-dropdown-item v-else @click="logout()">
                    <el-icon><SwitchButton /></el-icon>
                    Logout
                  </el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
        </el-col>
      </el-row>
    </el-header>
    <!--分割线-->
    <el-divider style="margin-bottom: 0;" />
  </el-container>
</template>


<script setup>
import github_icon from '../../assets/github-mark.svg'
import {User,UserFilled,Setting,SwitchButton,Promotion} from "@element-plus/icons-vue";
import {ElMessage, ElNotification} from "element-plus";
import router from "../../router/index.js";
// 引入store
import { useUserStore } from '../../store/userProfile.js'
import { Logout } from "../../api/user.js";

defineProps(['menuIndex']) // 定义props实现页面跳转 导航标签也跳转

// 实例用户仓库
const userProfile = useUserStore()
const username = userProfile.username

// 判断是否登陆

//退出登陆
const token = localStorage.getItem('jwtToken');
function logout() {
  // 如果令牌存在，则添加到请求头中
  if (token) {
    const headers = {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
    };
  Logout(headers).then(() => {
    userProfile.$reset()
    localStorage.removeItem('jwtToken')
    router.push('/')
    ElNotification({
      message: '成功退出登陆！',
      type: 'success',
      duration: 2000
    })
  }).catch((error) => {
    console.log(error)
  })
}
  else {
    userProfile.$reset()
    router.push('/')
    ElNotification({
      message: 'bangumi用户，无token成功注销！',
      type: 'success',
      duration: 2000
    })
  }
}
// 跳转用户信息
function toUserInfo() {
  if (username){
    router.push({
      name: 'UserPage',
      params: { username: username }
    })
  }
  else
    ElMessage({
      message: 'Please login first',
      type: 'warning',
      duration: 2000,
    })
}
//显示未实现的提示
function notImplement(name) {
  ElMessage({
    message: name + ' Coming Soon',
    type: 'warning',
    duration: 1500,
  })
}
</script>


<style>

</style>
