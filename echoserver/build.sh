#!/usr/bin/env bash

VERSION=`cat VERSION`
GOOS=linux GOARCH=amd64 go build -o echoserver
docker build --no-cache -t imroc.tencentcloudcr.com/test/echoserver:$VERSION .
rm echoserver