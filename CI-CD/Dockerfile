FROM golang:1.23.5
WORKDIR /src

RUN apt update -y && \
    apt install -y make git zip unzip pre-commit dpkg debconf debhelper jq && \
    curl -sL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g yarn semantic-release @semantic-release/git @semantic-release/gitlab @semantic-release/commit-analyzer @semantic-release/release-notes-generator && \
    export GOPATH=$(go env GOPATH) && \
    export PATH=$GOPATH/bin:$PATH \
# Dev mode local
COPY dev.sh ./dev.sh
RUN ./dev.sh
