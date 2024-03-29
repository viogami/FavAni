# FavAni

<p> FavAni，基于神经网络推荐算法，找到你最喜欢的动画！</p>

<a href="http://fa.viogami.me/" rel="nofollow">演示地址Demo</a>

<p> 前端使用 Vue3 和 Element Plus 进行构建，让你在使用该工具的同时享受良好的用户体验。
    后端使用golang的GIN框架,简洁高效，性能卓越，打包main文件一键服务器端部署。图算法使用pytorch库编写，处理为独立接口，使用grpc的方式和此后端进行通信。</p>


后端地址:<a>http://faapi.viogami.me/</a>

使用【/routes】路由可查看后端全部接口 --> <a>http://faapi.viogami.me/routes</a>

<p> 你可以访问它的 GitHub 仓库地址：<a href="https://github.com/viogami/FavAni_FE" target="_blank">FavAni_FE</a> </p>

## 首页面预览

![preview](https://github.com/viogami/FavAni_FE/raw/main/public/preview.png)

## 前言
个人开发的基于 BangumiAPI  执行推荐算法的数据分析网站。
调用 Bangumi的 API  实现数据检索，并且构建知识图谱。
后端 RestfulAPI ，通过 Gin  实现。接入 gorm  使用 mysql ，存放知识图谱
和用户的收藏信息，每个用户的收藏用来构建子图，使用图卷积神经网络（GCN）并引入注意力机制实现推荐算法。

功能仍在完善中....

## 实现细节
 - 前端界面设计
 - 后端api接口设计
 - 登录注册功能
 - jwt鉴权
 - 数据库表设计
 - 基于bangumi构建知识图谱
 - 用户收藏构建子图，执行推荐算法
 - grpc通信
 - docker文件编写

 ## 快速上手
<a href="http://fa.viogami.me/" rel="nofollow">演示地址Demo</a>

使用【/routes】查看后端全部接口 --> <a>http://faapi.viogami.me/routes</a>

python的gcn模型接口部署在9999端口

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
        # example: return gcn_pb2.GCNResult(node_scores={"node1": 0.5, "node2": 0.8})

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


