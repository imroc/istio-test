#!/usr/bin/env bash


kubectl delete pod -l app=benchechoclient &
kubectl delete pod -l app=echoserver &
