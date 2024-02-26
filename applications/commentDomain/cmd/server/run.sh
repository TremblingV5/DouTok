#!/bin/bash

# 设置环境变量
export DOUTOK_COMMENT_DOMAIN_SERVER_NAME=DouTokCommentDomainServer
export DOUTOK_COMMENT_DOMAIN_SERVER_PORT=8083
export DOUTOK_COMMENT_DOMAIN_ETCD_ADDRESS=localhost
export DOUTOK_COMMENT_DOMAIN_ETCD_PORT=2379
export DOUTOK_COMMENT_DOMAIN_MYSQL_USERNAME=admin
export DOUTOK_COMMENT_DOMAIN_MYSQL_PASSWORD=root
export DOUTOK_COMMENT_DOMAIN_MYSQL_DATABASE=DouTok

# 运行 Go 项目
go run main.go
