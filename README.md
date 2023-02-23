![image-20230223111012814](./documents/imgs/banner.jpeg)

DouTok is a backend for TikTok client based on Kitex and Hertz.

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

## Quick Start

1. env

```shell
docker-compose -f pro.yml up -d
```

2. hostname

```shell
append [hbase_addr hb-master] in /etc/hosts
```

3. run each service

```shell
relation/message/feed/publish/favorite/user/comment/api
```

