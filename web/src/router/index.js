import { createRouter, createWebHistory } from 'vue-router'
// 引入路由各页面路径配置
import routes from './routes'

const router = createRouter({
  routes,
  history: createWebHistory()
})

export default router
