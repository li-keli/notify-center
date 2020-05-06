#!/bin/bash

echo "[$(date +%T)] 开始编译"
CGO_ENABLED=0 go build -mod=vendor -ldflags="-s -w" -installsuffix cgo -o ./comet/app ./comet && echo "[$(date +%T)] comet编译成功" &
CGO_ENABLED=0 go build -mod=vendor -ldflags="-s -w" -installsuffix cgo -o ./server/app ./server && echo "[$(date +%T)] server编译成功" &
wait
echo "[$(date +%T)] 任务完成"