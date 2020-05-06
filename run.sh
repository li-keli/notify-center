#!/bin/bash

go run comet/main.go &
go run server/main.go &

wait