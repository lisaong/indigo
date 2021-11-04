#!/usr/bin/env bash

docker build -t audacity_linux_env .

if [ ! -d "audacity" ]
then
    git clone https://github.com/audacity/audacity
fi

cd audacity
mkdir build/linux-system

docker run -v ${pwd}:/audacity/audacity/ -v ${pwd}/build/linux-system:/audacity/build -it audacity_linux_env
