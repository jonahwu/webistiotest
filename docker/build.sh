#!/bin/bash
#REPO=docker-reg.emotibot.com.cn:55688
REPO=172.16.155.136:5000
REPO=${REPO_JENKINS:-$REPO}

CONTAINER="cicdtestimage"
CONTAINER=${CONTAINER_JENKINS:-$CONTAINER}

#TAG=$(git rev-parse --short HEAD)
TAG=test


DOCKER_IMAGE=$REPO/$CONTAINER:$TAG

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo $DIR
BUILDROOT=$DIR/..

# Build dockerbas
#cmdsource="$DIR/appscript.sh"
#echo $cmdsource
#eval $cmdsource
#cmd="docker build -t $CONTAINER:$TAG -f $DIR/Dockerfile $BUILDROOT"
# $JOB_NAME is provided by Jenkins
cmd="docker build --no-cache -t $DOCKER_IMAGE -f $DIR/Dockerfile $BUILDROOT"
echo $cmd
eval $cmd
