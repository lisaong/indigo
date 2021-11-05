#!/usr/bin/env bash

# https://github.com/audacity/audacity/tree/master/linux/ubuntu-focal

export TZ=Europe/London
ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

apt_packages_basic=(
    build-essential
    python3-minimal
    python3-pip
    python3-setuptools
    python3-wheel
    g++-8
    libstdc++-8-dev
    git
    wget
)

apt_packages_minimal_deps=(
    libgtk2.0-dev
    libasound2-dev
    libavformat-dev
    libjack-jackd2-dev
)

apt_packages_full_deps=(
    zlib1g-dev
    libexpat1-dev
    libmp3lame-dev
    libsndfile-dev
    libsoxr-dev
    portaudio19-dev
    libsqlite3-dev
    libavcodec-dev
    libavformat-dev
    libavutil-dev
    libid3tag0-dev
    libmad0-dev
    libvamp-hostsdk3v5
    libogg-dev
    libvorbis-dev
    libflac-dev
    libflac++-dev
    lv2-dev
    liblilv-dev
    libserd-dev
    libsord-dev
    libsratom-dev
    libsuil-dev
    libportmidi-dev
    libportsmf-dev
    libsbsms-dev
    libsoundtouch-dev
    libtwolame-dev
    libssl-dev
    libcurl4-openssl-dev
    libpng-dev
    libjpeg-turbo8-dev
    libopus-dev
)

apt-get install -y --no-install-recommends \
  "${apt_packages_basic[@]}" \
  "${apt_packages_minimal_deps[@]}" \
  "${apt_packages_full_deps[@]}"


pip3 install wheel conan

wget -O /tmp/cmake.sh https://github.com/Kitware/CMake/releases/download/v3.21.4/cmake-3.21.4-linux-x86_64.sh
sh /tmp/cmake.sh --skip-license
ln -s /audacity/bin/cmake /usr/bin/cmake