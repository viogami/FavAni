<template>
  <Header :menu-index="'-1'"/>
  <el-col  style="margin-top: 15vh" align="middle">
  <el-card class="box-card"  >
    <el-tabs type="border-card" v-model="activeTab" @tabClick="TabChange" v-if="showLogin">
      <el-tab-pane label="使用账号密码登陆">
          <!--默认登陆表单-->
          <el-form
              ref="DefaultFormRef"
              :rules="loginRules"
              label-width="120px"
              :model="DefaultForm"
              label-position="top"
              style="height: 50px">
            <el-form-item label=" " prop="username" >
              <el-input v-model="DefaultForm.username"
                        placeholder="Pleas input your username">
                <template #prefix>
                  <el-icon :size="20"> <UserFilled /> </el-icon> <!--图标-->
                </template>
              </el-input>
            </el-form-item>
            <el-form-item label=" " prop="password" >
            <el-input v-model="DefaultForm.password"
                      type="password"
                      placeholder="Please input password"
                      show-password>
              <template #prefix>
                <el-icon :size="20"> <Lock />  </el-icon> <!--图标-->
              </template>
            </el-input>
            </el-form-item>
          </el-form>
      </el-tab-pane>

      <el-tab-pane label="使用Bangumi用户名登陆">
          <!--用户名表单-->
          <el-form
              ref="BangumiLoginRef"
              :rules="loginRules"
              :model="BangumiForm"
              label-width="120px"
              style="height: 50px"
          >
          <el-input v-model="BangumiForm.username"
                    placeholder="Pleas input your username"
                    style="margin-top: 30px">
            <template #prefix>
              <el-icon :size="20"> <UserFilled />  </el-icon> <!--图标-->
            </template>
          </el-input>
          </el-form>
      </el-tab-pane>
      <!--不同标签激活，按钮实行不同登陆-->
      <el-button type="primary" round @click="activeTab=='0' ? LoginDefault() : LoginBangumi() "
                 color="pink"
                 style="margin-top: 50px"
      >登陆</el-button>
    </el-tabs>
    <!-- 注册面板 -->
    <el-form
        ref="RegisterFormRef"
        :model="RegisterForm"
        :rules="registerRules"
        label-width="120px"
        label-position="top"
        v-else
    >
      <el-form-item label="Username" prop="username" >
        <el-input v-model="RegisterForm.username" placeholder="Input your username"/>
      </el-form-item>
      <el-form-item label="Password" prop="password" >
        <el-input v-model="RegisterForm.password" placeholder="Input your password"/>
      </el-form-item>
      <el-form-item label="E-mail" prop="email">
        <el-input v-model="RegisterForm.email" placeholder="Input your email"/>
      </el-form-item>
      
      <el-button type="primary" round @click="Register()"
                 color="pink"
                 style="margin-top: 10px"
      >注册</el-button>
    </el-form>
    
  </el-card>
    
  <div style="margin-top: 10px" v-if="showLogin">
    <el-button link @click="showLogin=false">SIGN UP</el-button>
  </div>
    
  <div style="margin-top: 10px" v-else>
    <el-button link @click="showLogin=true">SIGN IN</el-button>
  </div>
  </el-col>
</template>

<script lang="ts" setup >
import Header from '../Home/Header.vue'
import { BangumiLogin,DefaultLogin,UserRegister } from '../../api/user' // login方法
import { reactive, ref } from 'vue'
// 引入Vue Router
import { useRouter } from 'vue-router'
// ele组件
import {ElMessage, ElNotification} from 'element-plus'
import { UserFilled,Lock } from '@element-plus/icons-vue'
// 引入store
import { useUserStore } from '../../store/userProfile.js'
// 获取路由实例
const router = useRouter()

// 活跃标签,以及改变函数
let activeTab = ref('0')
const TabChange = (tab) => {
  activeTab = tab.paneName
}
// 是否显示登陆/注册
const showLogin = ref(true)

// 表单信息ref
const DefaultFormRef = ref()
const BangumiLoginRef = ref()
const RegisterFormRef = ref()
// 默认登陆的默认表单
const DefaultForm = reactive({
  // 默认输入guest
  username: 'admin',
  password: '123456'
})
// bangumi登陆的默认表单
const BangumiForm = reactive({
  // 默认输入violetmail
  username: 'violetmail'
})
// 注册信息的表单
const RegisterForm = reactive({
  username: '',
  email: '',
  password: ''
})

// 实例用户仓库
const userProfile = useUserStore()

// 用户名规则验证器
const UserFormValidator = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('Please input'))
  }
  callback()
}
// 登陆表单规则
const loginRules = {
  username: [
    {
      validator: UserFormValidator,
      message: '用户名还没有填写',
      trigger: 'blur'
    }
  ],
  password:[
    {
      message: '密码还没有填写',
      trigger: 'blur'
    }
  ]
}
// 注册表单规则
const registerRules = {
  username: [
    {
      required: true,
      validator: UserFormValidator,
      message: '用户名还没有填写',
      trigger: 'blur'
    }
  ],
  email:[
    {
      message: '邮箱还没有填写',
      trigger: 'blur'
    }
  ],
  password:[
    {
      required: true,
      message: '密码还没有填写',
      trigger: 'blur'
    }
  ]
}

// 默认登陆
const LoginDefault = () => {
  DefaultFormRef.value.validate(async (valid) => {
    if (valid){
      DefaultLogin(DefaultForm.username,DefaultForm.password)
          .then(res => {
            // 设置token
            const token = res.data.token;
            localStorage.setItem('jwtToken', token);
            // 设置用户信息
            userProfile.setUserInfo(res.data.user.username, res.data.user.username, '', '')
            // 显示登陆通知
            eleNotice('success','欢迎，' + res.data.user.username + '!')
            // 跳转到主页
            router.push('/')
          })
          .catch(err => {
            eleNotice('error',err.response.data)
          })
  }
    else eleNotice('warning','Please input the valid information')
})
}
// bangumi用户名登陆
const LoginBangumi = () => {
  // 用户表单验证
  BangumiLoginRef.value.validate(async (valid) => {
    if (valid) {
      // 根据响应中的信息判断是否是 Bangumi 用户
      BangumiLogin(BangumiForm.username)
          .then(res => {
            // 设置用户信息
            userProfile.setUserInfo(res.data.username, res.data.nickname, res.data.avatar.large, res.data.sign)
            // 显示登陆通知
            eleNotice('success','欢迎，' + res.data.nickname + '!')
            // 跳转到主页
            router.push('/')
          })
          .catch(err => {
            eleNotice('error',err.response.data)
          })
    }
    else eleNotice('warning','Please input the valid information')
  })
}
// 注册按钮
const Register =() => {
  RegisterFormRef.value.validate(async (valid) => {
  if (valid) {
    // 调用注册函数，post给后端
    UserRegister(RegisterForm.username,RegisterForm.email,RegisterForm.password)
        .then(res => {
          // 显示注册通知
          eleNotice('success','注册成功！')
          // 跳转到主页
          router.push('/')
        })
        .catch(err => {
          eleNotice('error',err.response.data)
        })
  }
  else eleNotice('warning','Please input the valid information')
})
}
// 通知显示函数
function eleNotice(type,msg){
  switch (type) {
    case 'success':
      ElNotification({
            message: msg,
            type: 'success',
            duration: 2000 
          })
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
        message: 'please input correct notice type ,it`s must a string',
        type: 'error',
        duration: 2000
      })
  }
}
</script>

<style scoped>
.box-card {
  max-height: 500px;
  max-width: 670px;
  min-width: 200px;
  text-align: center;
}
</style>
