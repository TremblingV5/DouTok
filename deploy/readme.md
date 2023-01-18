# 部署

[TOC]

## 依赖环境

`env.yml`即使用`docker-compose`进行依赖环境创建的配置文件。

在`deploy`目录下使用命令`docker-compose -f env.yml up -d`来启动各个服务的容器，使用命令`docker-compose -f env.yml down`来关闭各个容器。

其中包含的服务有：

- redis
- mysql
- zookeeper
- kafka
- kafka-manager
- mongo
- hbase
- etcd
