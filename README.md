DouTok目前已转战V2：https://github.com/cloudzenith/DouTok，本仓库不再高强度维护

**为什么要开一个V2版本的仓库？**

DouTok最初是字节跳动青训营的参赛项目（并取得了比较靠前的名次），在参加青训营的过程中，由于种种原因，导致项目中存在诸多不合理的地方。

例如：DouTok现版本的微服务划分不够合理，拆的过于零碎，也许看起来很“微服务”，但与实际工作生产环境上的服务划分却背道而驰，微服务的划分不应过分追求“微”，而是适应项目发展，在完善基本设计的前提下进行拆分。除此之外，过多的“微服务”导致维护过程中的困难，想要调试就需要启动非常多的服务，对维护过程非常折磨。

在青训营之后，我们开始考虑继续维护DouTok。让DouTok继续扩张的一个卡点是其本身没有前端，只能依赖青训营中提供的“抖声”APP。为了让DouTok顺利扩张，所以我们决定开发一个全新的V2版本。在V2版本中，DouTok减少了服务的划分，增加了前端项目，虽然现阶段依然不够完整，但是已经具备了继续扩张的土壤。

对参与过DouTok维护的所有同学表示感谢！

![image-20230223111012814](./documents/imgs/banner.jpeg)

DouTok is a backend for TikTok client based on Kitex and Hertz.

![Hertz](https://img.shields.io/static/v1?label=Golang&message=1.18&color=brightgreen&style=plastic&logo=go) ![Hertz](https://img.shields.io/static/v1?label=Hertz&message=using&color=green&style=plastic&logo=go) ![Hertz](https://img.shields.io/static/v1?label=Kitex&message=using&color=yellowgreen&style=plastic&logo=go) ![Hertz](https://img.shields.io/static/v1?label=gorm/gen&message=using&color=yellow&style=plastic&logo=etcd) ![Hertz](https://img.shields.io/static/v1?label=etcd&message=3.4&color=orange&style=plastic&logo=etcd) ![Hertz](https://img.shields.io/static/v1?label=MySQL&message=8.0&color=red&style=plastic&logo=mysql) ![Hertz](https://img.shields.io/static/v1?label=Redis&message=7.0&color=blue&style=plastic&logo=redis) ![Hertz](https://img.shields.io/static/v1?label=HBase&message=2.1.3&color=blueviolet&style=plastic&logo=ApacheHadoop) ![Hertz](https://img.shields.io/static/v1?label=kafka&message=Tencent&color=ff69b4&style=plastic&logo=ApacheKafka)

## Documents

DouTok now have a documents site: [https://doutok.zhengfei.xin](https://doutok.zhengfei.xin)

## Quick Start

1. Deploy dependencies
   Deploy some dependencies such as MySQL, Redis etc. Run them by using `docker-compose` with `.yml` files in `./deploy` or deploy them by yourself.

   ```sh
   docker-compose -f ./deploy/env.yml up -d
   ```

2. Update config files

   There's an elegant way to run applications in this repo which is using `docker-compose`. So update config files in `./config_docker_compose` if you use `docker-comopse`. If you don't want to run them in this way, you must update config files in `./config` for ensuring them working.

3. Run applications by using `docker-compose`
   Everything is ready now, ensure that there's a `docker-compose.yml` in the root directory and run this command in terminal

   ```sh
   docker-compose up -d
   ```

### Download Client

We can download `.apk` file from `./ui` of this repo. Now the client is only support Android. After downloading and installing of this app. We can open it firstly. Then we can click `我` on the right bottom to enter configure page. After opening `高级配置`, we can input base url of the backend. An example is `http://localhost:8080/`.

## Architecture

### Technology Architecture

![image-20230223111253963](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20230223111253963.png)

### Tracing

> Visit `http://127.0.0.1:16686/` on browser.

![image-20230223111410359](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20230223111410359.png)

### Metric

> Visit `http://127.0.0.1:3000/` on browser.

![image-20230223111633341](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20230223111633341.png)

### API Introduce

| 接口                            | 功能       | 读/写 | 特点         | 备注               |
| ------------------------------- | ---------- | ----- | ------------ | ------------------ |
| /douyin/feed/                   | 视频流接口 | 读    | 根本功能接口 | 存储优化，缓存优化 |
| /douyin/user/register/          | 注册接口   | 写    |              |                    |
| /douyin/user/login/             | 登陆接口   | 读    |              |                    |
| /douyin/user/                   | 用户信息   | 读    |              |                    |
| /douyin/publish/action/         | 发布视频   | 写    |              | 存储优化           |
| /douyin/publish/list/           | 发布列表   | 读    |              | 缓存优化           |
| /douyin/favorite/action/        | 点赞接口   | 写    | 操作数大     | 延迟处理           |
| /douyin/favorite/list/          | 点赞列表   | 读    |              | 缓存优化           |
| /douyin/comment/action/         | 评论接口   | 写    | 操作数大     | 延迟处理           |
| /douyin/comment/list/           | 评论列表   | 读    |              | 缓存优化           |
| /douyin/relation/action/        | 关系操作   | 写    | 操作数大     | 延迟处理           |
| /douyin/relation/follow/list/   | 关注列表   | 读    |              | 缓存优化           |
| /douyin/relation/follower/list/ | 粉丝列表   | 读    |              | 缓存优化           |
| /douyin/relation/friend/list/   | 朋友列表   | 读    |              | 缓存优化           |
| /douyin/message/chat/           | 聊天记录   | 读    |              | 存储优化，缓存优化 |
| /douyin/message/action/         | 消息操作   | 写    |              | 延迟处理           |

## CI/CD

- Build docker images by using github actions
- Push docker images to private repo created with Aliyun by using github actions
- DOING: Try to run applications with k8s automatically

## Directories

```tree
.
├── applications    // 模块逻辑目录
│   ├── api    // api模块逻辑实现                        
│   │   ├── biz    // api模块主要逻辑实现
│   │   │   ├── handler
│   │   │   │   ├── api
│   │   │   │   ├── pack.go
│   │   │   │   └── ping.go
│   │   │   ├── model
│   │   │   │   └── api
│   │   │   └── router
│   │   │       ├── api
│   │   │       └── register.go
│   │   ├── chat
│   │   ├── initialize
│   │   │   ├── init_hertz.go
│   │   │   ├── jwt.go
│   │   │   ├── redis.go
│   │   │   ├── rpc
│   │   │   └── viper.go
│   │   ├── main.go
│   │   ├── Makefile
│   │   ├── router_gen.go
│   │   └── router.go
│   ├── comment    // comment模块逻辑实现
│   │   ├── build.sh
│   │   ├── dal    // 数据层
│   │   │   ├── gen.go    // gorm/gen生成代码
│   │   │   ├── migrate    // gorm自动迁移
│   │   │   │   └── main.go
│   │   │   ├── model    // model层
│   │   │   └── query    // gorm/gen生成结果
│   │   │       ├── comment_counts.gen.go
│   │   │       ├── comments.gen.go
│   │   │       ├── gen.go
│   │   │       └── var.go
│   │   ├── handler    // RPC服务的接口入口
│   │   ├── main.go    // 模块入口
│   │   ├── Makefile
│   │   ├── misc    // 模块所需的一些简单零散逻辑
│   │   ├── pack    // 将service层提供的数据查询结果包装成接口的返回消息
│   │   ├── rpc    // 初始化调用其他微服务
│   │   ├── script
│   │   ├── service // 数据查询
│   ├── favorite // favorite模块
│   ├── feed    // feed模块
│   ├── message // message模块
│   ├── publish // publish模块
│   ├── relation // relation模块
│   └── user // user模块
├── config    // 项目所需要的配置文件
├── deploy    // 项目所需的环境部署
├── documents    // 相关文档
├── go.mod
├── go.sum
├── kitex_gen    // Kitex生成的代码
│   ├── comment
│   ├── favorite
│   ├── feed
│   ├── message
│   ├── publish
│   ├── relation
│   └── user
├── pkg    // 项目所依赖的一些公共包
│   ├── constants    // 常量包
│   ├── dlog    // 日志包
│   ├── dtviper    // 配置包
│   ├── errno    // 错误码包
│   ├── hbaseHandle    // HBase操作包
│   ├── initHelper    // 初始化服务助手
│   ├── kafka    // Kafka操作包
│   ├── middleware    // 中间件
│   ├── misc    // 一些零散逻辑
│   ├── mock    // Mock测试包
│   │   ├── comment
│   │   ├── favorite
│   │   ├── feed
│   │   ├── message
│   │   ├── publish
│   │   ├── relation
│   │   └── user
│   ├── mysqlIniter    // MySQL操作包
│   ├── ossHandle    // OSS操作包
│   ├── redisHandle    // Redis操作包
│   ├── safeMap    // 线程安全的Map
│   └── utils    // 其他组件
├── proto    // 基于Protobuf3完成的IDL
├── README.md
└── scripts
```
