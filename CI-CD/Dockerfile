FROM golang:1.24.4
WORKDIR /src

# Установка базовых зависимостей
RUN apt update -y && \
    apt install -y make git zip unzip pre-commit dpkg debconf debhelper jq && \
    curl -sL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs && \
    curl https://baltocdn.com/helm/signing.asc | gpg --dearmor | sudo tee /usr/share/keyrings/helm.gpg > /dev/null && \
    apt-get install apt-transport-https --yes && \
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main" | tee /etc/apt/sources.list.d/helm-stable-debian.list && \
    apt-get install -y helm

# Установка Node.js инструментов
RUN npm install -g npm@latest && \
    npm install -g yarn semantic-release @semantic-release/git @semantic-release/gitlab @semantic-release/commit-analyzer @semantic-release/release-notes-generator

# Настройка Go окружения
RUN export GOPATH=$(go env GOPATH) && \
    export PATH=$GOPATH/bin:$PATH

# Копирование и запуск dev.sh (отдельные инструкции)
COPY dev.sh /dev.sh
RUN chmod +x /dev.sh && /dev.sh
