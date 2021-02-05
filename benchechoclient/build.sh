#!/usr/bin/env bash

VERSION=`cat VERSION`
GOOS=linux GOARCH=amd64 go build -o benchechoclient
docker build --no-cache -t imroc.tencentcloudcr.com/test/benchechoclient:$VERSION .
rm benchechoclient