#!/usr/bin/env bash

# cf. https://github.com/audacity/audacity/blob/master/linux/ubuntu-focal

conan --version

if [ ! -d "audacity" ]
then
    git clone https://github.com/audacity/audacity
fi

cd /audacity/build

cmake_options=(
    -G "Unix Makefiles"
    -DCMAKE_BUILD_TYPE=Release
)

cmake "${cmake_options[@]}" ../audacity

exit_status=$?

if [ $exit_status -ne 0 ]; then
    exit $exit_status
fi

make -j`nproc`

cd bin/Release
mkdir -p "Portable Settings"

ls -la 