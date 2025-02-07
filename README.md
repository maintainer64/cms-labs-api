# CMS Labs Urfu

Монорепозиторий для модулей:

1. [CMS Labs Core](backend/README.md)
2. [CMS Labs Front](nextui-dashboard/README.md)
3. [PNET Lab Addon](pnetlabaddon/README.md)
4. [Go-Lang Shared](shared/README.md)
5. [Gen Model](gen/README.md)

## ⚡️ Быстрый старт

1. Для установки всех зависимостей GoLang выполните:

```bash
make dev
```

2. Для тестирования всего пакета достаточно вызывать:

```bash
make test
```

2.1. Прогоняет линтеры

2.2. форматеры Go и Frontend'а

2.3. Тестирует код

2.4. Пишет coverage

2.5. (В CI выполняется такая-же команда)

3. Для автоматической генерации документации везде:

```bash
make generate
```
