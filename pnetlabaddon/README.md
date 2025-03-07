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
                ProxyAddHeaders On
                ProxyPreserveHost On
                RequestHeader set "Host" expr=%{HTTP_HOST}
                RequestHeader set "X-Scheme" expr=%{REQUEST_SCHEME}
                RequestHeader set "X-Original-URI" expr=%{REQUEST_URI}
                RequestHeader set "X-Forwarded-For" expr=%{REMOTE_ADDR}
                RequestHeader set "X-Real-IP" expr=%{REMOTE_ADDR}
                RequestHeader set "X-Forwarded-Proto" expr=%{REQUEST_SCHEME}
                ProxyPass http://127.0.0.1:5885/pnet-lab-addon
                ProxyPassReverse http://127.0.0.1:5885/pnet-lab-addon
        </Location>
```

Перезагрузить apache2 и добавить модуль headers

```shell
sudo a2enmod headers
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

Для обновления пакета достаточно запускать

```shell
./updater.sh --package-version <master | pre | stage>
```

## ⚠️ License

Apache 2.0 &copy; [Vic Shóstak](https://shostak.dev/) & [True web artisans](https://1wa.co/).
