FROM golang:1.23.5
WORKDIR /src

RUN apt update -y
RUN apt install make git zip unzip pre-commit dpkg debconf debhelper lintian -y
RUN curl -sL https://deb.nodesource.com/setup_20.x | bash
RUN apt-get install nodejs -yq
RUN npm install -g yarn
RUN export GOPATH=$(go env GOPATH)
RUN export PATH=$GOPATH/bin:$PATH
# Dev mode local
COPY dev.sh ./dev.sh
RUN ./dev.sh
