#!/bin/bash

# Устанавливаем локаль UTF-8
export LC_ALL=en_US.UTF-8
export LANG=en_US.UTF-8

# Определяем директорию, в которой находится скрипт
script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_name="webcms"  # Имя проекта

# Путь к файлам input.txt и out.txt в директории скрипта
input_file="$script_dir/input.txt"
output_file="$script_dir/out.txt"

# Функция для логирования
log() {
    echo "$(date '+%Y-%m-%d %H:%M:%S') $1"
}

# Функция для поиска корня проекта
find_project_root() {
    local dir="$script_dir"
    while [ "$dir" != "/" ] && [ "$dir" != "" ]; do
        if [ "$(basename "$dir")" == "$project_name" ]; then
            echo "$dir"
            return 0
        fi
        dir="$(dirname "$dir")"
    done
    echo "$script_dir/.."  # Возвращаем значение по умолчанию, если ничего не найдено
    return 1
}

project_root=$(find_project_root)

# Создаем или очищаем файл out.txt
log "Шаг 1: Создаем или очищаем файлы input.txt и out.txt"
> "$output_file"

# Проверяем, пуст ли файл input.txt
if [ ! -s "$input_file" ]; then
    log "Файл input.txt пуст. Заполните его необходимыми именами файлов и запустите скрипт снова."
    exit 1
fi

# Функция для поиска файла во всем проекте
find_file() {
    local file_name="$1"
    find "$project_root" -type f -iname "$file_name" -print 2>/dev/null
}

# Читаем файл input.txt и ищем файлы
log "Шаг 2: Заходим в файл input.txt и начинаем чтение"
first_file=true
not_found_files=()
while IFS= read -r line || [ -n "$line" ]; do
    # Ищем строки, начинающиеся с *
    if [[ "$line" == \** ]]; then
        file_name="${line#\*}"  # Убираем знак * в начале строки
        file_name=$(echo "$file_name" | xargs)  # Убираем пробелы в начале и конце строки
        if [ -z "$file_name" ]; then
            continue  # Пропускаем пустые строки
        fi

        # Ищем файл во всем проекте
        found_files=$(find_file "$file_name")

        if [ -z "$found_files" ]; then
            not_found_files+=("$line")
        else
            if ! $first_file; then
                echo "=============================================================" >> "$output_file"
            fi
            echo "Имя Файла: $file_name" >> "$output_file"
            echo "$found_files" | while IFS= read -r file; do
                relative_path="${file#"$project_root/"}"
                echo "Путь: $relative_path" >> "$output_file"
                echo "Содержимое:" >> "$output_file"
                cat "$file" >> "$output_file" 2>/dev/null || echo "Не удалось прочитать содержимое файла" >> "$output_file"
                echo "" >> "$output_file"
            done
            first_file=false
        fi
    fi
done < "$input_file"

# Выводим список ненайденных файлов
if [ ${#not_found_files[@]} -gt 0 ]; then
    echo "=============================================================" >> "$output_file"
    echo "Список не найденных файлов:" >> "$output_file"
    for file in "${not_found_files[@]}"; do
        echo "$file" >> "$output_file"
    done
fi

# Функция для вывода структуры проекта
#print_tree() {
#    local directory="$1"
#    local prefix="$2"
#    local exclude_dirs=(".git" ".idea")
#
#    for exclude in "${exclude_dirs[@]}"; do
#        if [[ "$(basename "$directory")" == "$exclude" ]]; then
#            return
#        fi
#    done
#
#    local items=("$directory"/*)
#    local total=${#items[@]}
#    local count=0
#
#    for item in "${items[@]}"; do
#        ((count++))
#        local branch_prefix="$prefix"
#        local item_name=$(basename "$item")
#
#        if [ "$count" -eq "$total" ]; then
#            branch_prefix+="└── "
#            new_prefix="$prefix    "
#        else
#            branch_prefix+="├── "
#            new_prefix="$prefix│   "
#        fi
#
#        if [ -d "$item" ]; then
#            echo "${branch_prefix}${item_name}" >> "$output_file"
#            print_tree "$item" "$new_prefix"
#        else
#            echo "${branch_prefix}${item_name}" >> "$output_file"
#        fi
#    done
#}

# Выводим структуру проекта
#log "Шаг 3: Cоздание структуры проекта"
#echo "=============================================================" >> "$output_file"
#echo "Структура проекта:" >> "$output_file"
#echo "${project_name}" >> "$output_file"
#print_tree "${project_root}" ""

log "Шаг 4: Информация успешно записана в файл out.txt"