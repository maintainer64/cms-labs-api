#!/bin/sh

# Список разрешённых переменных (через пробел)
ALLOWED_VARS="DB_TYPE DB_HOST DB_PORT DB_USER DB_PASSWORD DB_NAME DB_SSL_MODE DB_TABLE_PREFIX"

# Функция для загрузки переменных из .env файла с фильтрацией
load_env_file() {
  ENV_FILE="$1"
  if [ -f "$ENV_FILE" ]; then
    while IFS= read -r line; do
      # Пропускаем комментарии и пустые строки
      if [ -n "$line" ] && [ "$(echo "$line" | cut -c1)" != "#" ]; then
        # Убираем комментарии в конце строки
        line=$(echo "$line" | sed 's/\s*#.*//')
        # Извлекаем имя переменной (часть до =)
        var_name=$(echo "$line" | cut -d= -f1)
        # Проверяем, входит ли имя в список разрешённых
        if echo "$ALLOWED_VARS" | grep -qw "$var_name"; then
          # Безопасный экспорт: используем export с подстановкой, но без eval
          # Убираем возможные кавычки вокруг значения (опционально)
          value=$(echo "$line" | cut -d= -f2-)
          # Удаляем обрамляющие кавычки, если они есть (одинарные или двойные)
          value=$(echo "$value" | sed -e "s/^['\"]//" -e "s/['\"]$//")
          export "$var_name"="$value"
        fi
      fi
    done < "$ENV_FILE"
  else
    echo "Warning: Файл $ENV_FILE не найден"
  fi
}

# Загружаем переменные
load_env_file ".env.test"
load_env_file ".env"

# Указываем директорию для сохранения измененных файлов
output_dir="processed_migrations"
echo "DB_TABLE_PREFIX = $DB_TABLE_PREFIX | DB_NAME = $DB_NAME | DB_USER = $DB_USER | DB_HOST = $DB_HOST | DB_PORT = $DB_PORT"

mkdir -p "$output_dir"

# Функция для замены шаблона в файле
replace_template() {
  local input_file="$1"
  local output_file="$2"
  cp "$input_file" "$output_file"
  if [ "$(uname)" = "Darwin" ]; then
      sed -i '' -e "s/{{.DB_TABLE_PREFIX}}/$DB_TABLE_PREFIX/g" "$output_file"
  else
      sed -i -e "s/{{.DB_TABLE_PREFIX}}/$DB_TABLE_PREFIX/g" "$output_file"
  fi
}

# Функция для вызова sql-migrate независимо от окружения
sql_migrate_cmd() {
  if [ -f /usr/local/bin/sql-migrate ]; then
      /usr/local/bin/sql-migrate "$@"
  elif [ -f ./sql-migrate ]; then
      ./sql-migrate "$@"
  elif [ -f /app/sql-migrate ]; then
      /app/sql-migrate "$@"
  else
    sql-migrate "$@"
  fi
}

# Указываем директорию или файлы, в которых нужно произвести замену
files_to_process="platform/migrations"

for item in $files_to_process; do
  if [ -d "$item" ]; then
    find "$item" -type f \( -name '*.sql' -o -name '*.yml' \) | while read -r file; do
      relative_path="${file#$item/}"
      output_file="$output_dir/$relative_path"
      mkdir -p "$(dirname "$output_file")"
      replace_template "$file" "$output_file"
    done
  elif [ -f "$item" ]; then
    output_file="$output_dir/$(basename "$item")"
    replace_template "$item" "$output_file"
  else
    echo "Файл или директория $item не найдены."
  fi
done

echo "Замена шаблонов завершена. Измененные файлы сохранены в директории $output_dir."
cd "$output_dir" || exit
command="$1"
if [ "$command" = 'up' ]; then
    echo "⚡️ (Sql-Migrate) up"
    sql_migrate_cmd up sslmode=disable
elif [ "$command" = 'd' ]; then
    echo "⚡️ (Sql-Migrate) down"
    sql_migrate_cmd down sslmode=disable
elif [ "$command" = 's' ]; then
    echo "⚡️ (Sql-Migrate) status"
    sql_migrate_cmd status sslmode=disable
else
    echo "⚡️ (Sql-Migrate) Type: <up | d | s> "
fi
cd - || exit
rm -rf "$output_dir"
