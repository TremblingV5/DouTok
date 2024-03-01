## 技术栈

### 服务端

- go1.18 及以上
- HTTP 框架：[Hertz](https://www.cloudwego.io/docs/hertz/)
- RPC 框架：[Kitex](https://www.cloudwego.io/docs/kitex/)
- 中间件：kafka、redis
- 存储：mysql、hbase、minio

#### 项目启动

🌟 下面的内容需要配合 `DouTok/guidlines.md` 食用

0. 安卓手机 + pc，处在一个局域网内。
1. 克隆 reborn 分支代码（该分支仅用于验证项目启动，只能用于测试用户注册登陆流程，其他功能可能有问题，预计两天内合入main，开放所有功能）
2. 修改 hosts 文件，添加 `127.0.0.1 hb-master`。
3. 将 `./env/dependencies.yml` 中 `KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://192.168.1.119:9092` 的ip 替换成自己的局域网 IP。
4. 启动依赖环境：docker-compose -f ./env/dependencies.yml up -d（reborn 版相比 main 少了一些依赖，会影响部分功能，但是不影响开发者验证启动流程正确性）。
5. 登陆 mysql，在 DouTok 数据库中执行 scripts/DouTok.sql。
6. 通过容器访问 kafka，为需要 kafka 的服务创建 topic。

```shell
# 以交互式模式进入容器
docker exec -it 容器id /bin/bash
# 进入Kafka目录
cd /opt/kafka_2.13-2.8.1/
# 创建名为message_1的主题，1个分区，每个分区1个副本
bin/kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic message_1
# 创建名为relation_1的主题，1个分区，每个分区1个副本
bin/kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic relation_1
```

7. 启动服务（因为项目在迭代中，所以不同服务启动方式有些不同，详情参考 guidelines.md，将启动过程中涉及到的 kafka 的 ip 设置成局域网的 ip，对应配置文件路径为 DouTok/config/xxx，xxx 对应某一个服务。

8. 打开客户端，右下角“我”长按。

![image-20240227000742899](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240227000742899.png)

9. 跳转至输入连接的服务端口，这里输入连接的后端网关服务地址（最后一个 "/" 不要忘记）：

![image-20240227000823094](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240227000823094.png)

10. 体验注册登陆等功能。

### 前端

- 技术栈：React Hooks、TypeScript、Redux Saga、Vite

- 项目目录

![image-20240226234007434](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240226234007434.png)

#### 项目启动

```shell
npm install
npm run dev
```

