#!/bin/bash

# 设置环境变量
export DOUTOK_FAVORITE_SERVER_NAME=DouTokFavoriteServer
export DOUTOK_FAVORITE_SERVER_PORT=8082
export DOUTOK_FAVORITE_ETCD_ADDRESS=localhost
export DOUTOK_FAVORITE_ETCD_PORT=2379

# 运行 Go 项目
go run main.go
