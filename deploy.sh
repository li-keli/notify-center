#!/bin/bash

echo "[*] 开始构造comet包"
go build -mod=vendor -ldflags="-s -w" -installsuffix cgo -o ./comet/app ./comet
echo "[*] 开始构造server包"
go build -mod=vendor -ldflags="-s -w" -installsuffix cgo -o ./server/app ./server

echo "[*] 构造完成"