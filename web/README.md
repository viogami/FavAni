## FavAni_FE

<p> 欢迎来到FavAni--一个基于Bangumi条目搜索的推荐算法工具！这个网站使用 Bangumi 的 API 实现数据检索。</p>

<a href="http://fa.masterkagami.com/" rel="nofollow">演示地址Demo</a>

<p> 前端使用 Vue3 和 Element Plus 进行构建，让你在使用该工具的同时享受良好的用户体验。
    后端使用golang的GIN框架，简单尝试中，本仓库只提交前端代码，后端在搭建中，完善后会全部移放到后端仓库中</p>

<p> 目前正在仍在开发中，兴趣项目<s>（精力有限）</s> </p>

<p> 你可以访问它的 GitHub 仓库地址：<a href="https://github.com/viogami/FavAni_FE" target="_blank">FavAni_FE</a> </p>

## 首页面预览

![preview](https://github.com/viogami/FavAni_FE/raw/main/public/preview.png)

## 前言
使用**vue3**，目前我希望可以通过自助调用bangumi的api，通过添加收藏，执行推荐算法，得到用户更倾向的anime。

后端通过[bangumi API][1]调用。

该项目完全个人做全栈开发，不讲究界面设计了，css太困难，后面改用了`tailwindCSS`。我写了大量注释，每个文件和功能我也会具体介绍，也可以作为简单的vue3上手材料。

## 功能详解
### 首页
- 页头导航，用props实现不同页面时，导航的显示位置
- 侧边栏收放和页面响应式设计(用`el-col`，自带响应式)
- 文档显示功能
  - 使用markdown-it，更多依赖见`package.json`文件
  - 文档使用动态路由，详见`router/index.js`文件
  - `document.vue`是文章主组件，负责显示页面，文章列表，文章点击路由，该页面定义文章列表的菜单样式(目录层次)。
 `docPage.vue`是负责调用markdown工具进行编译，返回编译后的数据。`docCat.vue`是文章的目录文件也是路由进行匹配的文件，负责导入文章并实现文章路由。
  最后再定义一个`markdown.vue`组件执行编译markdown文本，这个组件是仿用了别人的项目。
- 时间线，关于页面的简单路由，以及404页面的路由
- 登陆注册功能，连接后端数据库，存储用户名和密码

### 用户页
- 有两种登陆方式，注册后用用户名和密码，或者直接使用bangumi的用户名登陆
- 通过菜单栏可以进入用户页
- 显示用户的基本信息，创建一个卡片，显示用户的收藏
- 用pinia的一个store存放所有用户信息(userProfile.js),在其中定义增删查改。

### 搜索
- 用户查询,可以基于用户名查询bangumi的用户信息，并返回用户的收藏
- 条目搜索，返回条目搜索列表，可以点击添加到收藏。

### 调用api
- 创建一个axios.js文件
- vite.config中定义代理，axios中创建不同的服务，用`axios.yourService`的方式使用多个url
- 新建user.js等文件，用于定义调用api的函数

[1]: https://bangumi.github.io/api/#/