## FavAni_FE

<p> 欢迎来到FavAni--一个基于Bangumi条目搜索的推荐算法工具！这个网站使用 Bangumi 的 API 实现数据检索。</p>

<a href="http://fa.viogami.me/" rel="nofollow">演示地址Demo</a>

前端使用 Vue3 和 Element Plus 通过vite进行构建，让你在使用该工具的同时享受良好的用户体验。
后端使用golang的GIN框架，使用gRPC的方式和算法后端通信(python)，此为**web文件夹👉前端代码**

你可以访问它的 GitHub 仓库地址：<a href="https://github.com/viogami/FavAni" target="_blank">FavAni</a>

## 首页面预览

![preview](https://github.com/viogami/FavAni/raw/master/web/public/preview.png)

## 前言
使用**vue3**，目前我希望可以通过自助调用bangumi的api，通过添加收藏，执行推荐算法，得到用户更倾向的anime。

后端依赖[bangumi API][1]调用实现数据检索以及算法设计。

该项目完全个人做全栈开发，不讲究界面设计了，css太困难，后面改用了`tailwindCSS`。我写了大量注释，每个文件和功能我也会具体介绍，也可以作为简单的vue3上手材料。

## src文件夹详解
### /api
- 定义axios接口文件,用于创建axios实例，定义请求拦截器和响应拦截器
- vite.config中定义代理，axios中创建不同的服务，用`axios.yourService`的方式使用多个url
- 新建user.js等文件，用于定义调用api的函数

### /assets
- 存放图片等静态资源
- 用于存放css文件

### /components
- 存放组件
- 用于存放页面组件，如`document.vue`，`login.vue`等

### /pages
- 页头导航，用props实现不同页面时，导航的显示位置
- 侧边栏收放和页面响应式设计(用`el-col`，自带响应式)
- 文档显示功能
    - 使用markdown-it，更多依赖见`package.json`文件
    - 文档使用动态路由，详见`router/index.js`文件
    - `document.vue`是文章主组件，负责显示页面，文章列表，文章点击路由，该页面定义文章列表的菜单样式(目录层次)。
      `docPage.vue`是负责调用markdown工具进行编译，返回编译后的数据。`docCat.vue`是文章的目录文件也是路由进行匹配的文件，负责导入文章并实现文章路由。
      最后再定义一个`markdown.vue`组件执行编译markdown文本，这个组件是仿用了别人的项目。
- 时间线，关于页面的简单路由，以及404页面的路由
- ### 用户页
- 有两种登陆方式，注册后用用户名和密码，或者直接使用bangumi的用户名登陆
- 显示用户的基本信息，创建一个卡片，显示用户的收藏，收藏请求页数异步拉取。
- ### 搜索
- 用户查询,可以基于用户名查询bangumi的用户信息，并返回用户的收藏
- 条目搜索，返回条目搜索列表，可以点击添加到收藏。
- ### 登录注册
- 连接后端数据库，存储用户名和密码
- 实现jwt验证，进入用户界面等需要通过鉴权，保证用户信息安全。
- 每次进入网站登录检查jwtToken，实现自动登录


### /router
- 存放路由文件，定义路由

### /store
- 存放pinia的store文件，定义store的增删查改界面
- 定义搜索的条目仓库，搜索的用户信息仓库，用户信息仓库

### /utils
- 存放可复用组件函数，包含jwt的认证请求，axios实例以及element-plus的请求封装



[1]: https://bangumi.github.io/api/#/