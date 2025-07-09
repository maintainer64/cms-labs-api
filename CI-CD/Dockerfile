FROM golang:1.24.4
WORKDIR /src

# Установка базовых зависимостей
RUN apt update -y && \
    apt install -y make git zip unzip pre-commit dpkg debconf debhelper jq && \
    curl -sL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs

# Установка Node.js инструментов
RUN npm install -g npm@latest && \
    npm install -g yarn semantic-release @semantic-release/git @semantic-release/gitlab @semantic-release/commit-analyzer @semantic-release/release-notes-generator

# Настройка Go окружения
RUN export GOPATH=$(go env GOPATH) && \
    export PATH=$GOPATH/bin:$PATH

# Копирование и запуск dev.sh (отдельные инструкции)
COPY dev.sh /dev.sh
RUN chmod +x /dev.sh && /dev.sh
