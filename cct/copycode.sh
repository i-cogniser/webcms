#!/bin/bash

# Определяем директорию, в которой находится скрипт
script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_root="$script_dir/.."
project_name="webcms"  # Имя проекта

# Путь к файлам input.txt и out.txt в директории скрипта
input_file="$script_dir/input.txt"
output_file="$script_dir/out.txt"

# Функция для логирования
log() {
    echo "$(date '+%Y-%m-%d %H:%M:%S') $1"
}

# Логирование: создание и очистка файлов input.txt и out.txt
log "Шаг 1: Создаем или очищаем файлы input.txt и out.txt"
if [ ! -f "$input_file" ]; then
    log "   Создаем файл input.txt"
    touch "$input_file"
fi
log "   Очищаем файл out.txt"
> "$output_file"

# Проверяем, пуст ли файл input.txt
if [ ! -s "$input_file" ]; then
    log "Файл input.txt пуст. Заполните его необходимыми именами файлов и запустите скрипт снова."
    exit 1
fi

# Логирование: выводим содержимое файла input.txt для проверки
log "Содержимое файла input.txt:"
cat "$input_file"

# Логирование: заходим в файл input.txt и начинаем чтение
log "Шаг 2: Заходим в файл input.txt и начинаем чтение"
while IFS= read -r file_name; do
    if [ -z "$file_name" ]; then
        continue  # Пропускаем пустые строки
    fi
    log "   Ищем файл по имени: $file_name"
    # Ищем файл по имени в проекте (включая поддиректории)
    found_files=$(find "$project_root" -name "$file_name")

    if [ -z "$found_files" ]; then
        log "   Не найден файл по имени: $file_name"
    else
        log "   Найдены следующие файлы:"
        echo "$found_files" | while IFS= read -r file; do
            relative_path="${file#"$project_root/"}"  # Получаем относительный путь от корня проекта
            log "      Имя Файла: $project_name/$relative_path"
            log "      Содержимое:"
            cat "$file"
            echo ""
            # Записываем информацию о файле в out.txt
            echo "Имя Файла: $project_name/$relative_path" >> "$output_file"
            echo "Содержимое:" >> "$output_file"
            cat "$file" >> "$output_file"
            echo "" >> "$output_file"
        done
    fi
done < <(grep -v '^\s*$' "$input_file")  # Используем grep для исключения пустых строк из input.txt

# Функция для вывода структуры проекта в виде дерева
print_tree() {
    local directory=$1
    local prefix=$2
    local exclude_dirs=(".git" ".idea")

    # Проверяем, исключен ли текущий каталог
    for exclude in "${exclude_dirs[@]}"; do
        if [[ "$(basename "$directory")" == "$exclude" ]]; then
            return
        fi
    done

    # Получаем список подкаталогов и файлов
    local items=("$directory"/*)
    local total=${#items[@]}
    local count=0

    for item in "${items[@]}"; do
        ((count++))
        local branch_prefix="$prefix"
        local item_name=$(basename "$item")

        if [ "$count" -eq "$total" ]; then
            branch_prefix+="└── "
            new_prefix="$prefix    "
        else
            branch_prefix+="├── "
            new_prefix="$prefix│   "
        fi

        if [ -d "$item" ]; then
            echo "${branch_prefix}${item_name}" >> "$output_file"
            print_tree "$item" "$new_prefix"
        else
            echo "${branch_prefix}${item_name}" >> "$output_file"
        fi
    done
}

# Выводим структуру проекта в out.txt
log "Шаг 3: Вывод структуры проекта"
echo "Структура проекта:" >> "$output_file"
echo "$project_name" >> "$output_file"
print_tree "$project_root" ""

log "Структура проекта добавлена в файл $output_file"

# Проверяем, была ли найдена информация
if [ -s "$output_file" ]; then
    log "Информация успешно записана в файл $output_file"
else
    log "Нет информации для записи в файл $output_file, так как файл input.txt не содержит указанных файлов."
fi
