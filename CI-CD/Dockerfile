FROM golang:1.23.5
WORKDIR /src

RUN apt update -y
RUN apt install make git zip unzip pre-commit -y
RUN curl -sL https://deb.nodesource.com/setup_20.x | bash
RUN apt-get install nodejs -yq
RUN npm install -g yarn
RUN export GOPATH=$(go env GOPATH)
RUN export PATH=$GOPATH/bin:$PATH
# Dev mode local
RUN go install gotest.tools/gotestsum@latest
RUN go install github.com/go-critic/go-critic/cmd/gocritic@latest
RUN go install github.com/securego/gosec/v2/cmd/gosec@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
RUN go install github.com/jstemmer/go-junit-report/v2@latest
RUN go install github.com/rubenv/sql-migrate/...@latest
RUN go install golang.org/x/tools/cmd/cover@latest
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install github.com/t-yuki/gocover-cobertura@latest
