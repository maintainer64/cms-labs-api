# Аддон для PNETLab реализующий OpenID аутентификацию

## 1. Первичная установка

1. Подключитесь к консоли по SSH до PNETLab и скопируйте содержимое файла [updater.sh](updater.sh)

2. Выдайте разрешение на запуск командой:

```shell
chmod +x updater.sh
```

3. Установите пакет jq и unzip. Следующим образом

```shell
sudo apt install jq unzip -y
```

Или

```shell
snap install jq unzip
```

4. Выполните команду

Скачивание и обновление пакета для канала stage:

```shell
./updater.sh --package-version stage
```

Скачивание и обновление пакета для канала pre:

```shell
./updater.sh --package-version pre
```

Скачивание и обновление пакета для канала master:

```shell
./updater.sh --package-version master
```

5. После успешной установки ваши логи будут выглядеть примерно так

```shell
Removed /etc/systemd/system/multi-user.target.wants/pnetlabaddon.service.
Unpacking pnetlabaddon (1.0) over (1.0) ...
Setting up pnetlabaddon (1.0) ...
DB_TABLE_PREFIX = addon_ | DB_NAME = pnetlab_db | DB_USER = pnetlab | DB_HOST = 127.0.0.1 | DB_PORT = 3306
Замена шаблонов завершена. Измененные файлы сохранены в директории processed_migrations.
⚡️ (Sql-Migrate) up
Applied 0 migrations
/etc/pnetlabaddon
/
Archive:  /etc/pnetlabaddon/pnetlab-inject.zip
replace /etc/pnetlabaddon/resources/views/pnetlab-inject/pnetlabapijs.html
  inflating: /etc/pnetlabaddon/resources/views/pnetlab-inject/pnetlabapijs.html
  inflating: /etc/pnetlabaddon/resources/views/pnetlab-inject/main.blade.php
 extracting: /etc/pnetlabaddon/resources/views/pnetlab-inject/pnetlabapijs/main.js
  inflating: /etc/pnetlabaddon/resources/views/pnetlab-inject/pnetlabapijs/lab.js
  inflating: /etc/pnetlabaddon/resources/views/pnetlab-inject/pnetlabapijs/assets/index-4e5c0698.css
  inflating: /etc/pnetlabaddon/resources/views/pnetlab-inject/pnetlabapijs/assets/index-940f062f.js
Created symlink /etc/systemd/system/multi-user.target.wants/pnetlabaddon.service → /lib/systemd/system/pnetlabaddon.service.
root@pnet3
```

6. Проверьте работоспособность пакета используя команду

```shell
root@pnet3:~# service pnetlabaddon status
● pnetlabaddon.service - PNET lab addon Service
   Loaded: loaded (/lib/systemd/system/pnetlabaddon.service; enabled; vendor preset: enabled)
   Active: active (running) since Fri 2025-03-07 06:44:55 UTC; 1min 28s ago
 Main PID: 17349 (pnetlabaddon)
    Tasks: 6 (limit: 2307)
   CGroup: /system.slice/pnetlabaddon.service
           └─17349 /usr/local/bin/pnetlabaddon

Mar 07 06:44:55 pnet3 systemd[1]: Started PNET lab addon Service.
Mar 07 06:44:55 pnet3 pnetlabaddon[17349]:  ┌───────────────────────────────────────────────────┐
Mar 07 06:44:55 pnet3 pnetlabaddon[17349]:  │                   Fiber v2.52.6                   │
Mar 07 06:44:55 pnet3 pnetlabaddon[17349]:  │               http://127.0.0.1:5885               │
Mar 07 06:44:55 pnet3 pnetlabaddon[17349]:  │                                                   │
Mar 07 06:44:55 pnet3 pnetlabaddon[17349]:  │ Handlers ............ 12  Processes ........... 1 │
Mar 07 06:44:55 pnet3 pnetlabaddon[17349]:  │ Prefork ....... Disabled  PID ............. 17349 │
Mar 07 06:44:55 pnet3 pnetlabaddon[17349]:  └───────────────────────────────────────────────────┘
```

```shell
root@pnet3:~# curl -v 'http://127.0.0.1:5885/pnet-lab-addon/api/docs'
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 5885 (#0)
> GET /pnet-lab-addon/api/docs HTTP/1.1
> Host: 127.0.0.1:5885
> User-Agent: curl/7.58.0
> Accept: */*
>
< HTTP/1.1 301 Moved Permanently
< Date: Fri, 07 Mar 2025 06:53:36 GMT
< Content-Length: 0
< X-Request-Id: 375ed821-12e8-407e-b025-2f4d21dd3c5a
< Location: /pnet-lab-addon/api/docs/index.html
<
* Connection #0 to host 127.0.0.1 left intact
```

7. Изменить параметры Apache2

Перейти в настройку apache2

```shell
cd /etc/apache2/sites-enabled
nano pnetlab.conf
nano pnetlabs.conf
```

Добавить следующие дерективы

```shell
        <Location /pnet-lab-addon>
                Order allow,deny
                Allow from all
                ProxyPass http://127.0.0.1:5885/pnet-lab-addon
                ProxyPassReverse http://127.0.0.1:5885/pnet-lab-addon
                ProxyAddHeaders On
                ProxyPreserveHost On
        </Location>
```

Перезагрузить apache2

```shell
sudo /etc/init.d/apache2 restart
```

Проверить подключение к pnetlabaddon перейдя по ссылке вида

```shell
https://pnet.local/pnet-lab-addon/api/docs
```

Должна отобразиться структура методов Swagger UI

8. Установка API ключа и изменение .env параметров

Для того, чтобы pnetlabaddon поддерживал openID и передавал статистику в **core**,
необходимо использовать токен доступа.

Прописать его можно здесь

```shell
nano /etc/pnetlabaddon/.env
# CMS client
CMS_MAX_TIMEOUT=30  # Кол-во секунд для ответа core
CMS_CLIENT_ID=pnet3  # Уникальное название сервера для SSO
CMS_TOKEN=fe0810391cbab933d15d836f344c02629586f0ddda67c1414e093be17e3a3bc3  # Токен
CMS_BASE_URL=https://cms-lab.gubanov.site
```

Перезапускаем pnetlabaddon

```shell
service pnetlabaddon restart
```

Примечания

<img src="https://img.shields.io/badge/Go-1.19+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
&nbsp;<a href="https://goreportcard.com/report/gitlab.com/a10869/api-modules/pnetlabaddon" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>
&nbsp;<img src="https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none" alt="license" />

[Fiber](https://gofiber.io/) is an Express.js inspired web framework build on top of Fasthttp, the fastest HTTP engine
for Go. Designed to ease things up for **fast** development with **zero memory allocation** and **performance** in mind.

## 2. Обновления

## ⚡️ Quick start

1. Create a new project with Fiber:

```bash
cgapp create

# Choose a backend framework:
#   net/http
# > fiber
#   chi
```

2. Rename `.env.example` to `.env` and fill it with your environment values.
3. Install [Docker](https://www.docker.com/get-started) and the following useful Go tools to your system:

- [golang-migrate/migrate](https://github.com/golang-migrate/migrate#cli-usage) for apply migrations
- [github.com/swaggo/swag](https://github.com/swaggo/swag) for auto-generating Swagger API docs
- [github.com/securego/gosec](https://github.com/securego/gosec) for checking Go security issues
- [github.com/go-critic/go-critic](https://github.com/go-critic/go-critic) for checking Go the best practice issues
- [github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint) for checking Go linter issues

4. Run project by this command:

```bash
make docker.run
```

5. Go to API Docs page (Swagger): [127.0.0.1:5000/docs](http://127.0.0.1:5000/docs)

![Screenshot](https://user-images.githubusercontent.com/11155743/112715187-07dab100-8ef0-11eb-97ea-68d34f2178f6.png)

## 🗄 Template structure

### ./app

**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which
caching solution your choose_ or any third-party things.

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/models` folder for describe business models and methods of your project
- `./app/queries` folder for describe queries for models of your project

### ./docs

**Folder with API Documentation**. This directory contains config files for auto-generated API Docs by Swagger.

### ./pkg

**Folder with project-specific functionality**. This directory contains all the project-specific code tailored only for
your business use case, like _configs_, _middleware_, _routes_ or _utils_.

- `./pkg/configs` folder for configuration functions
- `./pkg/middleware` folder for add middleware (Fiber built-in and yours)
- `./pkg/repository` folder for describe `const` of your project
- `./pkg/routes` folder for describe routes of your project
- `./pkg/utils` folder with utility functions (server starter, error checker, etc)

### ./platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual
project, like _setting up the database_ or _cache server instance_ and _storing migrations_.

- `./platform/cache` folder with in-memory cache setup functions (by default, Redis)
- `./platform/database` folder with database setup functions (by default, PostgreSQL)
- `./platform/migrations` folder with migration files (used
  with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)

## ⚙️ Configuration

```ini
# .env

# Stage status to start server:
#   - "dev", for start server without graceful shutdown
#   - "prod", for start server with graceful shutdown
STAGE_STATUS="dev"

# Server settings:
SERVER_HOST="0.0.0.0"
SERVER_PORT=5000
SERVER_READ_TIMEOUT=60

# JWT settings:
JWT_SECRET_KEY="secret"
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=15
JWT_REFRESH_KEY="refresh"
JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT=720

# Database settings:
DB_TYPE="pgx"   # pgx or mysql
DB_HOST="cgapp-postgres"
DB_PORT=5432
DB_USER="postgres"
DB_PASSWORD="password"
DB_NAME="postgres"
DB_SSL_MODE="disable"
DB_MAX_CONNECTIONS=100
DB_MAX_IDLE_CONNECTIONS=10
DB_MAX_LIFETIME_CONNECTIONS=2

# Redis settings:
REDIS_HOST="cgapp-redis"
REDIS_PORT=6379
REDIS_PASSWORD=""
REDIS_DB_NUMBER=0
```

## ⚠️ License

Apache 2.0 &copy; [Vic Shóstak](https://shostak.dev/) & [True web artisans](https://1wa.co/).
