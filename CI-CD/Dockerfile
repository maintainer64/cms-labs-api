FROM golang:1.24.4
WORKDIR /src

# Установка базовых зависимостей
RUN apt-get update -y && \
    apt-get install -y --no-install-recommends \
        make \
        git \
        zip \
        unzip \
        pre-commit \
        dpkg \
        debconf \
        debhelper \
        jq \
        ca-certificates \
        curl \
        gnupg && \
    # Установка Node.js
    curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs && \
    # Установка Helm
    curl -fsSL https://baltocdn.com/helm/signing.asc | gpg --dearmor -o /usr/share/keyrings/helm.gpg && \
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main" | tee /etc/apt/sources.list.d/helm-stable-debian.list && \
    apt-get update && \
    apt-get install -y helm && \
    # Очистка кеша
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Установка Node.js инструментов
RUN npm install -g npm@latest && \
    npm install -g yarn semantic-release @semantic-release/git @semantic-release/gitlab @semantic-release/commit-analyzer @semantic-release/release-notes-generator

# Настройка Go окружения
RUN export GOPATH=$(go env GOPATH) && \
    export PATH=$GOPATH/bin:$PATH

# Копирование и запуск dev.sh (отдельные инструкции)
COPY dev.sh /dev.sh
RUN chmod +x /dev.sh && /dev.sh
