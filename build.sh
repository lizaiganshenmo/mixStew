#!/bin/bash

# 检查参数的数量
if [ $# -ne 1 ]; then
    echo "Usage: $0 <command>"
    exit 1
fi

command=$1

case $command in
"stop")
    kill -9 $(lsof -ti:8001)
    kill -9 $(lsof -ti:8002)
    kill -9 $(lsof -ti:8003)
    kill -9 $(lsof -ti:8004)
    kill -9 $(lsof -ti:8888)
    ;;
"start")
    cd cmd

    # run user
    cd user
    go build -o ../output/main_user && nohup ../output/main_user &

    cd ..
    # run follow
    cd follow
    go build -o ../output/main_folllow && nohup ../output/main_follow &

    cd ..
    # run interaction
    cd interaction
    go build -o ../output/main_interaction && nohup ../output/main_interaction &

    cd ..
    # run article
    cd article
    go build -o ../output/main_article && nohup ../output/main_article &

    cd ..
    # run gateway
    cd api
    go build -o ../output/main_api && nohup ../output/main_api &
    ;;
"restart")
    kill -9 $(lsof -ti:8001)
    kill -9 $(lsof -ti:8002)
    kill -9 $(lsof -ti:8003)
    kill -9 $(lsof -ti:8004)
    kill -9 $(lsof -ti:8888)
    cd cmd

    # run user
    cd user
    go build -o ../output/main_user && nohup ../output/main_user &

    cd ..
    # run follow
    cd follow
    go build -o ../output/main_folllow && nohup ../output/main_follow &

    cd ..
    # run interaction
    cd interaction
    go build -o ../output/main_interaction && nohup ../output/main_interaction &

    cd ..
    # run article
    cd article
    go build -o ../output/main_article && nohup ../output/main_article &

    cd ..
    # run gateway
    cd api
    go build -o ../output/main_api && nohup ../output/main_api &

    ;;
*)
    echo "Invalid command."
    exit 1
    ;;
esac
