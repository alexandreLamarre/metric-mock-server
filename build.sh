#!/bin/bash

docker build -t alex7285/instrumentation:latest .;

echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin;

docker push alex7285/instrumentation:latest;