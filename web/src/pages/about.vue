<script setup>
import { ref } from 'vue'
import Header from './Home/Header.vue'
import FooterPage from './Home/footer.vue'

const count = ref(0)
</script>

<template>
<!-- 添加页头 -->
  <Header :menu-index="'4'"/>
  <div style="height: calc(100vh - 130px); overflow-y: auto;">
<!-- 关于页面 -->
  <div class="about-container">
    <h2>
      文档的 Introduce 页面调用的github的api实现，读取仓库的README.md文件。
      如果因为网络问题不显示，可以访问该项目的 GitHub 仓库地址：<a href="https://github.com/viogami/FavAni/tree/master/web" target="_blank">FavAni</a>
    </h2>

    <h1> 更新日志 </h1>
    <p>2024/3</p>
    <p>完善图算法中，使用pytorch编写算法，部署该接口后将使用gprc和后端通信。修改了后端收藏表的数据库模型，使用redis实现收藏请求的消息队列，可以执行高并发操作。</p>
    <p>2024/2</p>
    <p>寒假期间，做了中间件鉴权功能，给了鉴权路由分组。登录功能使用jwt。同时完成了收藏数据库的建设</p>
    <p>2024/1/21</p>
    <p>重命名的项目。完善了搜索界面，登陆和用户查询功能用两个store存放不同的信息，复用会产生一些权限问题。解决了前端分页设计的一些问题，接下来着重着手完善后端</p>
    <p>2024/1/2</p>
    <p>添加了登陆注册功能，连接后端数据库存放用户信息，也可以使用bangumi直接用用户名登陆</p>
    <p>2023/12/2</p>
    <p>更新了doc文档显示，更加符合视觉。添加了条件搜索，增加了默认搜索信息。注意到,bangumiAPI每次最多返回1000条查询信息。</p>
    <p>2023/11/29</p>
    <p>更新了用户收藏信息的调用算法，用缓存和阶段调用加速了翻页体验，并完善了响应式布局</p>
    <p>2023/11/26</p>
    <p>前端基本首页做完了(基本能看)，用户的功能也大差不大了，项目准备build发布一次，后续开发一下检索的功能，并开发后端和数据库，进行数据处理。</p>
    <p>2023/10/10</p>
    <p>
      个人作全栈开发，用vue这种较为轻量的前端开发框架。目前决定使用vue3+element-plus做前端。
      框架搭建：下载node.js，配置环境。创建第一个vue单页应用。根据elementUI的官方步骤，引入依赖。</p>

    <div class="card">
      <button type="button" @click="count++">功德 {{ count }}</button>
      <p>
        如果你觉得无聊，可以加加功德
      </p>
    </div>
  </div>

  <footer-page/>
  </div>
</template>

<style scoped>
.about-container {
  max-width: 800px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  text-align: center;
  flex: 1;
  padding-bottom: 60px; /* 底部内边距，确保内容不会被页脚遮挡 */
}

.card {
  background-color: #f5f5f5;
  padding: 15px;
  border-radius: 8px;
  margin-top: 20px;
}

button {
  background-color: #28a745;
  color: #fff;
  padding: 10px 15px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #218838;
}
</style>
