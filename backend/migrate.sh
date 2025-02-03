#!/bin/bash

# Import env variables from ./.env
set -a
[ -f .env.test ] && . .env.test || echo "Warning: .env.test file not found"
[ -f .env ] && . .env || echo "Warning: .env file not found"

# Указываем директорию для сохранения измененных файлов
output_dir="processed_migrations"
echo "DB_TABLE_PREFIX = $DB_TABLE_PREFIX | DB_NAME = $DB_NAME | DB_USER = $DB_USER | DB_HOST = $DB_HOST | DB_PORT = $DB_PORT"

# Создаем директорию, если она не существует
mkdir -p "$output_dir"

# Функция для замены шаблона в файле
replace_template() {
  local input_file="$1"
  local output_file="$2"
  # Копируем файл в новую директорию
  cp "$input_file" "$output_file"
  # Используем sed для замены шаблона на значение переменной
  # Проверяем операционную систему
  if [[ "$(uname)" == "Darwin" ]]; then
      sed -i '' -e "s/{{.DB_TABLE_PREFIX}}/$DB_TABLE_PREFIX/g" "$output_file"
  else
      sed -i -e "s/{{.DB_TABLE_PREFIX}}/$DB_TABLE_PREFIX/g" "$output_file"
  fi
}

# Указываем директорию или файлы, в которых нужно произвести замену
# Например, platform/migrations и конфигурационный файл
files_to_process=(
  "platform/migrations"
)

# Проходим по каждому файлу или директории
for item in "${files_to_process[@]}"; do
  if [ -d "$item" ]; then
    # Если это директория, обрабатываем все файлы в ней
    find "$item" -type f -name '*.sql' -or -name '*.yml' | while read -r file; do
      # Определяем путь для сохранения измененного файла
      relative_path="${file#$item/}"
      output_file="$output_dir/$relative_path"
      # Создаем необходимые директории
      mkdir -p "$(dirname "$output_file")"
      replace_template "$file" "$output_file"
    done
  elif [ -f "$item" ]; then
    # Если это файл, обрабатываем его
    output_file="$output_dir/$(basename "$item")"
    replace_template "$item" "$output_file"
  else
    echo "Файл или директория $item не найдены."
  fi
done

echo "Замена шаблонов завершена. Измененные файлы сохранены в директории $output_dir."
cd "$output_dir"
command="$1"
if [[ $command == 'up' ]]
then
    echo "⚡️ (Sql-Migrate) up"
    sql-migrate up sslmode=disable
elif [[ $command == 'd' ]]
then
    echo "⚡️ (Sql-Migrate) down"
    sql-migrate down sslmode=disable
elif [[ $command == 's' ]]
then
    echo "⚡️ (Sql-Migrate) status"
    sql-migrate status sslmode=disable
else
    echo "⚡️ (Sql-Migrate) Type: <up | d | s> "
fi
cd -
rm -rf "$output_dir"
