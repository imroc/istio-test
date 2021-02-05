#!/usr/bin/env bash

./build.sh
./push.sh

kubectl delete pod -l app=echoserver