#!/usr/bin/env bash

kubectl exec -it -c benchechoclient `kubectl get pod -l app=benchechoclient | grep Running | awk -F ' ' '{print $1}'` bash