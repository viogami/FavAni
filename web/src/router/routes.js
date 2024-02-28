import about from '../pages/about.vue'
import NotFound from '../pages/404.vue'
import home from '../pages/Home/Home.vue'
import timeline from '../pages/Timeline.vue'
import UserPage from '../pages/login/UserPage.vue'
import docs from '../pages/doc/documents.vue'
import docCat from '../pages/doc/docCat.vue'
import login from "../pages/login/login.vue"; 
import posttest from "../pages/PostTest.vue"
import SubjectSearch from "../pages/search/SearchSubject.vue";
import SearchUser from "../pages/search/SearchUser.vue";

const routes = [
  { path: '/', component: home }, //  首页
  { path: '/login', component: login }, //  登陆
  { path: '/about', component: about }, //  关于页面
  { path: '/user/:username', name: 'UserPage', component: UserPage }, // 使用动态路由创建用户界面
  { path: '/timeline', component: timeline }, // 时间线界面
  { path: '/docs',
    component: docs,
    redirect: '/docs/introduce', // 重定向到介绍界面
    children: [
      {
        path: ':page',
        name: 'DocPage',
        component: docCat // 必须定位至目录，不然无法config值无法获取
      }
    ]
  }, // 文章路由
  { path: '/search/:keywords', name: 'SubjectSearch', component: SubjectSearch}, // 使用动态路由创建搜索用户结果界面
  { path: '/search/user/:user', name: 'SearchUser', component: SearchUser}, // 使用动态路由创建条目搜索结果界面
  { path: '/:patchMatch(.*)', name: NotFound, component: NotFound }, // 404 页面
  { path: '/posttest',component: posttest } // post 测试页面
]

export default routes
