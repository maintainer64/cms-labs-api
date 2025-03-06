#!/bin/bash

# Параметры по умолчанию
PACKAGE_VERSION="master"
PACKAGE_NAME="pnetlabaddon"
FULL_PATH="a10869"
BASE_URL="https://gitlab.com"
FIRST=1

# Обработка аргументов командной строки
while [[ $# -gt 0 ]]; do
  case $1 in
    --package-version)
      PACKAGE_VERSION="$2"
      shift
      shift
      ;;
    *)
      echo "Неизвестный параметр: $1"
      exit 1
      ;;
  esac
done

# Получение ID пакета
PACKAGES=$(curl --silent --location "$BASE_URL/api/graphql" \
  --header "Content-Type: application/json" \
  --data @- <<EOF
{
  "operationName": "getPackages",
  "variables": {
    "fullPath": "$FULL_PATH",
    "packageName": "$PACKAGE_NAME",
    "packageVersion": "$PACKAGE_VERSION",
    "first": $FIRST
  },
  "query": "query getPackages(\$fullPath: ID!, \$packageName: String, \$packageVersion: String, \$first: Int) { group(fullPath: \$fullPath) { id packages(packageName: \$packageName, packageVersion: \$packageVersion, first: \$first) { nodes { id } } } }"
}
EOF
)

# Извлечение ID пакета с помощью jq
PACKAGE_ID=$(echo "$PACKAGES" | jq -r '.data.group.packages.nodes[0].id')
if [ -z "$PACKAGE_ID" ] || [ "$PACKAGE_ID" == "null" ]; then
  echo "Пакет не найден."
  exit 1
fi
echo "ID пакета: $PACKAGE_ID"

# Получение информации о файлах пакета
FILES=$(curl --silent --location "$BASE_URL/api/graphql" \
  --header "Content-Type: application/json" \
  --data @- <<EOF
{
  "operationName": "getPackageFiles",
  "variables": {
    "id": "$PACKAGE_ID",
    "first": $FIRST
  },
  "query": "query getPackageFiles(\$id: PackagesPackageID!, \$first: Int) { package(id: \$id) { id packageFiles(first: \$first) { nodes { id fileName downloadPath } } } }"
}
EOF
)

# Извлечение пути для скачивания файла с помощью jq
DOWNLOAD_PATH=$(echo "$FILES" | jq -r '.data.package.packageFiles.nodes[0].downloadPath')
if [ -z "$DOWNLOAD_PATH" ] || [ "$DOWNLOAD_PATH" == "null" ]; then
  echo "Файл пакета не найден."
  exit 1
fi
echo "Путь для скачивания: $DOWNLOAD_PATH"

# Скачивание файла
curl --fail --location --output "${PACKAGE_NAME}-${PACKAGE_VERSION}.deb" \
  "$BASE_URL$DOWNLOAD_PATH"

echo "Файл успешно скачан: ${PACKAGE_NAME}-${PACKAGE_VERSION}.deb"

sudo dpkg -i ${PACKAGE_NAME}-${PACKAGE_VERSION}.deb
