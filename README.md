# CMS Labs Urfu

[![pipeline status](https://gitlab.com/a10869/api-modules/badges/master/pipeline.svg)](https://gitlab.com/a10869/api-modules/-/commits/master)

[![coverage report](https://gitlab.com/a10869/api-modules/badges/master/coverage.svg)](https://gitlab.com/a10869/api-modules/-/commits/master)

[![Latest Release](https://gitlab.com/a10869/api-modules/-/badges/release.svg)](https://gitlab.com/a10869/api-modules/-/releases)

## CI-CD:

1. [Pipelines](CI-CD/README.md)
2. [Docker files](CI-CD/docker.md)

## Монорепозиторий для модулей:

1. [CMS Labs Core](backend/README.md)
2. [CMS Labs Front](nextui-dashboard/README.md)
3. [PNET Lab Addon](pnetlabaddon/README.md)
4. [Go-Lang Shared](shared)
5. [Gen Model](gen/README.md)

## Документация методов Swagger

1. [CMS Labs Core](backend/docs/swagger.json)
2. [PNET Lab Addon](pnetlabaddon/docs/swagger.json)

## ⚡️ Быстрый старт

1. Для установки всех зависимостей GoLang выполните:

```bash
make dev
```

2. Для тестирования всего пакета достаточно вызывать:

```bash
# Тестирует код
# Пишет coverage
make test

# Форматеры Go и Frontend'а
make pre_commit

# Линтер и инъекции кода
make security
```

3. Для автоматической генерации документации везде:

```bash
make generate
```
