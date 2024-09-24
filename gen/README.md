# Fiber golang generation Model

## Использование

Автоматическое создание модулей для проекта по модели в БД

Для создания связных файлов используйте

1. Перейдите в директорию с проектом

```shell
cd /work/your/backend/directory
```

2. Проверьте что ваша директория примерно такого типа

```shell
ls -ll
---
|-app
  |-controllers
  |-di
  |-models
  |-queris
  |-usecases
|-pkg
  |-routes
|-platform
```

3. Выполните команду указывая свои CLI параметры

```shell
go run ../gen -name $(MODEL) -root $(ROOT) -directory $(APP_DIRECTORY)
---
go run ../gen -name "Users" -root "gitlab.com/a10869/api-modules/backend" -directory "~/go/labs/backend"
```

4. В монорепе используется Make для быстрого вызова команды используй

```shell
MODEL="Users" make crud
```

## Разработка

Нужно использовать некоторые дополнительные компоненты разработки
Для установки может понадобиться официальные сайты с указанием твоей ОС

1. make
2. gocritic
3. gosec
4. golangci-lint
5. pre-commit

Для тестирования всей интеграции как в пайпланах вызывай

```shell
make test
```
