# CMS Labs

- Core уровень для интеграции между LMS системой (Moodle) по протоколам LTI
- Предоставление SSO OPENID формата
- Управление серверами и кластерами
- Распределение пользователей между сегментами сервером
(пока только PNETLab)

Шаблон [Create Go App CLI](https://github.com/create-go-app/cli)

<img src="https://img.shields.io/badge/Go-1.19+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
&nbsp;<a href="https://goreportcard.com/report/gitlab.com/a10869/api-modules/backend" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>
&nbsp;<img src="https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none" alt="license" />

## ⚡️ Быстрый старт

1. Использовать для быстрого создания моделей:

```bash
MODEL='book' make crud
```

2. Для миграции используй:

```bash
make migrate.up
```

3. Для запуска проекта:

```bash
make run
```

4. API Docs страница (Swagger): [127.0.0.1:5000/docs](http://127.0.0.1:5000/docs)

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

## ⚙️ Конфигурация

Смотри пример в файле [.env.test](.env.test)

## ⚙️ Тесты & Линтер

```bash
make test
make lint
pre-commit run --all-files
```

## ⚠️ Лицензия

Apache 2.0 &copy; [Vic Shóstak](https://shostak.dev/) & [True web artisans](https://1wa.co/).
