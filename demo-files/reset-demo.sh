#!/bin/bash

# Demo reset.
# Reset quota file.
# Delete images from GCR.
kubectl delete deployment gif-maker
cp Dockerfile-local ../Dockerfile
cp cloudbuild.v001.yaml ../cloudbuild.yaml
kubectl delete deployments gif-maker --context dev
kubectl delete deployments gif-maker --context bretmcg
kubectl delete deployments gif-maker --context production
kubectl config use-context dev
# git tag -d 0.0.1
# git push --delete origin 0.0.1


# Just once per env.
# kubectl create secret generic gif-maker --from-file=/var/run/secret/cloud.google.com/service-account.json
# kubectl expose deployment gif-maker --type=LoadBalancer --name=gif-maker-service

