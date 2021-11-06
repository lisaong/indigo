# Credits: https://github.com/audacity/audacity/tree/master/linux/ubuntu-focal

FROM ubuntu:bionic

LABEL maintainer="d.vedenko@audacityteam.org"
LABEL description="A build environment to check the builds for the package maintainers"
LABEL version="3.0"

RUN apt-get update && apt-get install -y dos2unix

WORKDIR /audacity

COPY ["dependencies.sh", "/audacity/"]
RUN dos2unix dependencies.sh

RUN ["bash", "-ex", "dependencies.sh"]

COPY ["entrypoint.sh", "/audacity/"]
RUN dos2unix entrypoint.sh

CMD ["bash", "-ex", "./entrypoint.sh"]