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

4. API Docs страница (Swagger): [https://cms-lab.gubanov.site/api/docs](https://cms-lab.gubanov.site/api/docs)

![Screenshot](https://user-images.githubusercontent.com/11155743/112715187-07dab100-8ef0-11eb-97ea-68d34f2178f6.png)

## 🗄 Структура шаблона

### ./app

**Папка с бизнес-логикой**. Этот каталог не зависит от _того, какой драйвер базы данных вы используете_ или _какое
решение для кэширования вы выбрали_, или любых других сторонних вещей.

- `./app/controllers` папка для функциональных контроллеров (используются в маршрутах)
- `./app/models` папка для описания бизнес-моделей и методов вашего проекта
- `./app/queries` папка для описания запросов к моделям вашего проекта

### ./docs

**Папка с документацией API**. Этот каталог содержит конфигурационные файлы для автоматически генерируемой документации
API с помощью Swagger.

### ./pkg

**Папка с функциональностью, специфичной для проекта**. Этот каталог содержит весь код, специфичный для вашего
бизнес-кейса, такой как _конфигурации_, _middleware_, _маршруты_ или _утилиты_.

- `./pkg/configs` папка для функций конфигурации
- `./pkg/middleware` папка для добавления middleware (встроенного в Fiber и вашего)
- `./pkg/repository` папка для описания `const` вашего проекта
- `./pkg/routes` папка для описания маршрутов вашего проекта
- `./pkg/utils` папка с утилитарными функциями (запуск сервера, проверка ошибок и т.д.)

### ./platform

**Папка с логикой уровня платформы**. Этот каталог содержит всю логику уровня платформы, которая будет использоваться
для построения фактического проекта, такая как _настройка базы данных_ или _экземпляра кэш-сервера_ и _хранение
миграций_.

- `./platform/cache` папка с функциями настройки in-memory кэша (по умолчанию, Redis)
- `./platform/database` папка с функциями настройки базы данных (по умолчанию, PostgreSQL)
- `./platform/migrations` папка с файлами миграций (используется с
  инструментом [golang-migrate/migrate](https://github.com/golang-migrate/migrate))

## ⚙️ Конфигурация

Смотри пример в файле [.env.test](.env.test)

## ⚙️ Тесты & Линтер

```bash
make test
make lint
pre-commit run --all-files
```

## ⚙️ DevOps

```bash
make test
make lint
pre-commit run --all-files
```

## ⚠️ Лицензия

Apache 2.0 &copy; [Vic Shóstak](https://shostak.dev/) & [True web artisans](https://1wa.co/).
