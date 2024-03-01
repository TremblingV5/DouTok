#!/bin/bash

# 设置环境变量
export DOUTOK_USER_SERVER_NAME=DouTokUserServer
export DOUTOK_USER_SERVER_PORT=8084
export DOUTOK_USER_ETCD_ADDRESS=localhost
export DOUTOK_USER_ETCD_PORT=2379

# 运行 Go 项目
go run main.go
