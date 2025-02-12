# CMS Labs Urfu

[![pipeline status](https://gitlab.com/a10869/api-modules/badges/master/pipeline.svg)](https://gitlab.com/a10869/api-modules/-/commits/master)

[![coverage report](https://gitlab.com/a10869/api-modules/badges/master/coverage.svg)](https://gitlab.com/a10869/api-modules/-/commits/master)

[![Latest Release](https://gitlab.com/a10869/api-modules/-/badges/release.svg)](https://gitlab.com/a10869/api-modules/-/releases)

Монорепозиторий для модулей:

1. [CMS Labs Core](//gitlab.com/a10869/api-modules/backend)
2. [CMS Labs Front](//gitlab.com/a10869/api-modules/backend/nextui-dashboard)
3. [PNET Lab Addon](//gitlab.com/a10869/api-modules/pnetlabaddon)
4. [Go-Lang Shared](//gitlab.com/a10869/api-modules/shared)
5. [Gen Model](//gitlab.com/a10869/api-modules/gen)

## Документация методов Swagger

1. [CMS Labs Core](https://cms-lab.gubanov.site/api/docs)
2. [PNET Lab Addon](https://editor-next.swagger.io/?url=https://cms-lab-docs.gubanov.site/files/app/gitlab.com/a10869/api-modules/pnetlabaddon/docs/swagger.json)

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
