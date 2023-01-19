# DouTok

- 编写业务参考
    - 参考代码结构/目录结构：https://github.com/cloudwego/kitex-examples/tree/main/bizdemo/easy_note
    - 参考业务逻辑：https://github.com/a76yyyy/tiktok

- 目录说明

```shell
.
├── README.md
├── applications       # 各模块业务 
│   ├── api            # 网关
│   │   ├── pack       # 需要包装的结构（dto）
│   │   ├── handlers   # 接口回调函数
│   │   ├── main.go    # hertz 主函数
│   │   ├── middleware # 中间件
│   │   └── rpc        # rpc 客户端操作
│   ├── comment  # 评论
│   │   ├── Makefile
│   │   ├── build.sh
│   │   ├── dal        # 数据库操作
│   │   ├── handler.go # rpc 服务端的实现
│   │   ├── main.go    # rpc 服务主函数
│   │   ├── pack       # 需要包装的结构（dto）
│   │   ├── rpc        # rpc 客户端操作
│   │   ├── script     
│   │   └── service    # 业务逻辑（MVC）
│   ├── favorite # 点赞
│   ├── feed 	 # 视频流
│   ├── message  # 即时通讯
│   ├── publish  # 发布视频
│   ├── relation # 关注
│   └── user     # 用户/登陆/注册
├── config 			 # 配置文件
├── deploy           # 项目环境部署文件
├── documents        # 文档相关
├── go.mod
├── go.sum
├── kitex_gen        # kitex generate
├── pkg~~~~
│   ├── configurator # 配置初始化器
│   ├── constants    # 常量库/错误库
│   ├── mysqIniter   # MySQL 初始化器
│   └── redisHandle  # Redis 初始化器
├── proto        # IDL 文件
└── scripts      # 脚本文件
```

- test