# CMS Labs FRONTEND

Работает на Next-UI и Vite - сборщик

[NextUI](https://nextui.org/)

## Быстрый запуск

Установка зависимостей

```bash
yarn install
```

Запуск сервера

```bash
yarn run dev
```

После этого открываем https://localhost:3000 в браузере.

Для работы с бекендом нужно настроить прокси локального сервера.
Найди следующую строчку в проекте и измени содержимое

```js
function setupProxy(): CommonServerOptions["proxy"] {
```
