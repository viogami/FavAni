# FavAni

<p> FavAni，基于神经网络推荐算法，找到你最喜欢的动画！</p>

<a href="http://fa.viogami.me/" rel="nofollow">演示地址Demo</a>

前端使用 Vue3 和 Element Plus 进行构建，让你在使用该工具的同时享受良好的用户体验。
后端使用golang的GIN框架,简洁高效，性能卓越，打包main文件一键服务器端部署。
图算法依赖pytorch，处理为接口独立部署，通过grpc和此后端进行通信。


后端地址:<a>http://faapi.viogami.me/</a>

使用【/routes】路由可查看后端全部接口 --> <a>http://faapi.viogami.me/routes</a>

<p> 你可以访问它的 GitHub 仓库地址：<a href="https://github.com/viogami/FavAni" target="_blank">FavAni</a> </p>

## 首页面预览

![preview](https://github.com/viogami/FavAni/raw/master/web/public/preview.png)

## 前言
个人开发的基于 BangumiAPI  执行推荐算法的数据分析网站。
调用 Bangumi的 API  实现数据检索，并且构建知识图谱。使用jwt鉴权，并自动登录。

后端 RestfulAPI ，通过 Gin  实现。接入 gorm  使用 mysql ，接入鉴权中间件。
算法后端存放知识图谱和用户的收藏信息，每个用户的收藏用来构建子图，使用图卷积神经网络（GCN）并引入注意力机制实现推荐。


## 实现细节
 - 前端界面设计
 - 后端api接口设计
 - 登录注册功能
 - jwt鉴权，根据token有效与否自动登录
 - 数据库表设计
 - 基于bangumi构建知识图谱
 - 用户收藏构建子图，执行推荐算法
 - grpc通信
 - docker文件编写

 ## 快速上手
<a href="http://fa.viogami.me/" rel="nofollow">演示地址Demo</a>

使用【/routes】查看后端全部接口 --> <a>http://faapi.viogami.me/routes</a>

python的gcn模型接口部署在9999端口

## 文件夹解释
- auth : 鉴权相关文件，目前为jwt
- config：定义配置文件格式并存放配置文件，用于上线灵活修改项目
- database: 定义数据库模型，以及创建mysql和redis
- middleware：定义中间件，包含中间件函数
- proto：定义grpc的protoc文件
- repos：仓库的接口文件，每个数据库表对应一个仓库
- server：启动服务，包括mysql，redis，grpc等
- web：前端文件夹

## grpc说明
本后端做grpc客户端，gcn模型处理代码用py编写，做grpc的服务端。

协议文件编写在`proto/gcn.proto`中

### python服务端简要代码
```python
import grpc
from concurrent import futures
from proto import gcn_pb2_grpc, gcn_pb2

class GCNServicer(gcn_pb2_grpc.GCNServiceServicer):
    def ProcessGraph(self, request, context):
        # Process graph data using GCN and return result
        # example: 
        # return gcn_pb2.GCNResult(node_scores={"node1": 0.5, "node2": 0.8})

def server():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    gcn_pb2_grpc.add_GCNServiceServicer_to_server(GCNServicer(), server)
    server.add_insecure_port('[::]:9999')
    server.start()
    print('gRPC 服务端已开启,端口为9999...')
    server.wait_for_termination()

if __name__ == '__main__':
    server()

```


