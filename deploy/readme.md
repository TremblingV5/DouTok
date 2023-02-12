# 部署

[TOC]

## 不同尺寸的环境配置

|类型|配置文件|备注|
|:---:|:---:|:---:|
|全部环境|env.yml||
|基础环境，MySQL Redis Etcd|basic.yml||
|HBase环境|hbase.yml|采用alone版本的HBase,自带zookeeper,故简单隔离。试用时需将访问机的hosts添加一个名为“hb-master”的远端记录|
|链路追踪所需环境|pro.yml|吃配置，慎重！|
|其他配置|others.yml|目前仅包含mongo|

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
- jaeger
- elasticsearch

在使用docker-compose启动后，可以使用`docker ps`命令查看正在运行的容器，若有容器运行不正常，可以使用在`docker ps`的列表中找到不正常的容器id,然后使用命令`docker logs xxxxx`来查看该容器的日志。

## 几个UI界面

kafka-manager: 9099端口
hbase: 16010端口
jaeger-UI: 16686端口
