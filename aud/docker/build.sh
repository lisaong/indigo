#!/usr/bin/env bash

# docker builder prune
docker build -t audacity_linux_env .

if [ ! -d "audacity" ]
then
    git clone https://github.com/audacity/audacity
fi

docker run --rm -v `pwd`/audacity:/audacity/audacity/ -v `pwd`/build:/audacity/build -it audacity_linux_env
