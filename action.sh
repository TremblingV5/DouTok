#!/bin/bash

services=(
    "api" "comment" "commentDomain" "favorite" "favoriteDomain"
    "feed" "message" "messageDomain" "publish" "relation" "relationDomain"
    "user" "userDomain" "videoDomain"
)

if [ $1 == "start" ]; then
    for i in 0 1 2 3 4 5 6 7 8 9 10 11 12 13
    do
        nohup go run ./applications/${services[i]}/ > ./log/${services[i]}.out 2>&1 &
    done
fi

if [ $1 == "stop" ]; then
    for i in 0 1 2 3 4 5 6 7 8 9 10 11 12 13
    do
        process=`ps -ef | grep "go run ./applications/${services[i]}/" | grep -v grep | awk '{print $2}'`
        for j in $process
        do
            echo "Kill: ${services[i]} $j"
            kill -9 $j
        done
    done

    for i in 8088 8087 8086 8085 8084 8083 8082 8081 8079 8078 8077 8076 8075 8074
    do
        process=`lsof -i:$i | grep LISTEN | awk '{print $2}'`
        for j in $process
        do
            echo "Kill: port $i $j"
            kill -9 $j
        done
    done
fi