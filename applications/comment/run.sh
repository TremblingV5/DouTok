#!/bin/bash

# 设置环境变量
export DOUTOK_COMMENT_SERVER_NAME=DouTokCommentServer
export DOUTOK_COMMENT_SERVER_PORT=8081
export DOUTOK_COMMENT_ETCD_ADDRESS=localhost
export DOUTOK_COMMENT_ETCD_PORT=2379

# 运行 Go 项目
go run main.go
