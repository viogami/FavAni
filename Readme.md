# FavAni

<p> FavAni，基于神经网络推荐算法，找到你最喜欢的动画！</p>

<a href="http://fa.masterkagami.com/" rel="nofollow">演示地址Demo</a>

<p> 前端使用 Vue3 和 Element Plus 进行构建，让你在使用该工具的同时享受良好的用户体验。
    后端使用golang的GIN框架,简洁高效，性能卓越，打包main文件一键服务器端部署。</p>


<p> 你可以访问它的 GitHub 仓库地址：<a href="https://github.com/viogami/FavAni_FE" target="_blank">FavAni_FE</a> </p>

## 首页面预览

![preview](https://github.com/viogami/FavAni_FE/raw/main/public/preview.png)

## 前言
个人开发的基于 BangumiAPI  执行推荐算法的数据分析网站。
调用 Bangumi的 API  实现数据检索，并且构建知识图谱。
后端 RestfulAPI ，通过 Gin  实现。接入 gorm  使用 mysql ，存放知识图谱
和用户的收藏信息，每个用户的收藏用来构建子图，使用图神经网络（GCN）实现推荐算法。

功能仍在完善中....

## 实现细节
 - 前端界面设计
 - 后端api接口设计
 - 登录注册功能
 - jwt鉴权
 - 数据库表设计
 - 基于bangumi构建知识图谱
 - 用户收藏构建子图，执行推荐算法
 - docker文件编写

 ## 快速上手
~~ coming soon ~~


