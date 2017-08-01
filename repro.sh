#!/bin/bash

# use Minikube Docker daemon
eval $(minikube docker-env)

# build image and mount the `mounted` directory
docker build --rm -t repro:latest .

# Run image in container and attach
docker run --name repro-fs-reloading -v "$PWD/mounted":"/app/mounted" repro
