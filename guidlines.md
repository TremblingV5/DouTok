# Guidelines

> This file is used to tell everyone how to build, deploy and use DouTok.

## How to deploy the related dependencies

### Modify `./env/dependencies.yml`

`./env/dependencies.yml` is a configuration for `docker-compose`. In this file, now we import some dependencies which we need used in DouTok. In another words, we must deploy these dependencies so we can run DouTok.

In this file, now we imported:

- Redis
- MySQL
- etcd
- zookeeper
- HBase
- Kafka

Always we don't need to modify it but we must modify one line for Kafka.

In `services.kafka.environment`, there's an entry named `KAFKA_ADVERTISED_LISTENERS`. We should use the IP address of your local network. For example, I get my IP address by using `ifconfig` is `192.168.1.119` and I will configure it as `KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://192.168.1.119:9092`.

Don't forget modify config files for each modules, you can use a global serach for `YOUR_OWN_IP` to find all where you should replace with your own IP address.

### Modify `hosts` file

Something more, modification of `hosts` file is needed. This will help `HBase` run as we expected. We should add this in our `hosts` file:

```shell
127.0.0.1 hb-master
```

### Fire

Run `docker-compose -f ./env/dependencies.yml` to deploy all dependencies.

## How to build backend

For `User`, `UserDomain`, `Comment`, `CommentDomain`, `Favorite`:

1. Copy `./config/vscode_launch.jsonc` to `./vscode`
2. Use vscode to run these modules.

For others:

Run each modules by using `go run ./applications/xxx/`

## How to build frontend

## How to use DouTok
