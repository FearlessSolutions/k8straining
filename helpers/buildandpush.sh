#!/bin/bash

docker-compose build

docker tag k8straining_add us.gcr.io/fearless-sandbox-2019/k8straining_add:latest
docker tag k8straining_multiply us.gcr.io/fearless-sandbox-2019/k8straining_multiply:latest
docker tag k8straining_nest1 us.gcr.io/fearless-sandbox-2019/k8straining_nest1:latest
docker tag k8straining_nest2 us.gcr.io/fearless-sandbox-2019/k8straining_nest2:latest

docker push us.gcr.io/fearless-sandbox-2019/k8straining_add:latest
docker push us.gcr.io/fearless-sandbox-2019/k8straining_multiply:latest
docker push us.gcr.io/fearless-sandbox-2019/k8straining_nest1:latest
docker push us.gcr.io/fearless-sandbox-2019/k8straining_nest2:latest