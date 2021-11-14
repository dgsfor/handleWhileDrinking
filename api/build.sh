#!/bin/bash
# 交叉编译，编译linux版本
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hwd-api
