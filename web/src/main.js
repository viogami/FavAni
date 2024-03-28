import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue' // 全局引入图标
import { createPinia } from 'pinia' // pinia
import {jwtLogin} from "./utils/auth.js";
// 初始化app
const app = createApp(App)
// 引入组件
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

// use函数每次只接受一个参数
app.use(router)
app.use(ElementPlus)
app.use(createPinia())
app.use(jwtLogin)
app.mount('#app')


