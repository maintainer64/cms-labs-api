#!/bin/sh

# Функция для загрузки переменных из .env файла
load_env_file() {
  ENV_FILE="$1"  # Первый аргумент функции — имя файла

  # Проверяем, существует ли файл
  if [ -f "$ENV_FILE" ]; then
    # Читаем файл, игнорируем комментарии и пустые строки
    while IFS= read -r line; do
      # Убираем строки, начинающиеся с комментариев или пустые строки
      if [ -n "$line" ] && [ "$(echo "$line" | cut -c1)" != "#" ]; then
        # Убираем комментарии в конце строки
        line=$(echo "$line" | sed 's/\s*#.*//')
        # Экспортируем переменную
        eval "export $line"
      fi
    done < "$ENV_FILE"
  else
    # Если файл не найден, выводим предупреждение
    echo "Warning: Файл $ENV_FILE не найден"
  fi
}

# Импортируем переменные окружения из .env.test, если файл существует
load_env_file ".env.test"
load_env_file ".env"

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
# Например, platform/migrations и конфигурационный файл
files_to_process="platform/migrations"

# Проходим по каждому файлу или директории
for item in $files_to_process; do
  if [ -d "$item" ]; then
    # Если это директория, обрабатываем все файлы в ней
    find "$item" -type f \( -name '*.sql' -o -name '*.yml' \) | while read -r file; do
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
