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
log "Шаг 1: Создаем или очищаем файл out.txt"
> "$output_file"

# Проверяем, пуст ли файл input.txt
if [ ! -s "$input_file" ]; then
    log "Файл input.txt пуст. Заполните его необходимыми именами файлов и запустите скрипт снова."
    exit 1
fi

# Функция для поиска файла во всем проекте, исключая node_modules
find_file() {
    local file_name="$1"
    find "$project_root" -type f -iname "$file_name" ! -path "*/node_modules/*" -print 2>/dev/null
}

# Читаем файл input.txt и ищем файлы
log "Шаг 2: Заходим в файл input.txt и начинаем чтение"
first_file=true
not_found_files=()
ignore_paths=()

while IFS= read -r line || [ -n "$line" ]; do
    if [[ "$line" == \** ]]; then
        file_name="${line#\*}"  # Убираем знак * в начале строки
        file_name=$(echo "$file_name" | xargs)  # Убираем пробелы в начале и конце строки
        if [ -z "$file_name" ]; then
            continue  # Пропускаем пустые строки
        fi
        # Под этим комментарием часть кода изменять нельзя. Ищем строки, начинающиеся с *
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
    elif [[ "$line" == +* ]]; then
        path="${line#+}"  # Убираем + в начале строки
        path=$(echo "$path" | xargs)  # Убираем пробелы в начале и конце строки
        if [ -n "$path" ]; then
            ignore_paths+=("$path")
        fi
    fi
done < "$input_file"

# Функция для проверки, игнорируется ли путь
is_ignored() {
    local path="$1"
    for ignore in "${ignore_paths[@]}"; do
        if [[ "$path" == "$ignore" || "$path" == "$ignore/"* ]]; then
            return 0
        fi
    done
    return 1
}

# Функция для вывода структуры проекта
print_tree() {
    local directory="$1"
    local prefix="$2"

    # Пропускаем директории и файлы, указанные для игнорирования
    if is_ignored "$directory" || is_ignored "$(basename "$directory")"; then
        return
    fi

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
            # Перед выводом папки проверяем, является ли она игнорируемой
            if ! is_ignored "$item" && ! is_ignored "$(basename "$item")"; then
                echo "${branch_prefix}${item_name}/" >> "$output_file"
                print_tree "$item" "$new_prefix"
            fi
        else
            echo "${branch_prefix}${item_name}" >> "$output_file"
        fi
    done
}

# Выводим структуру проекта
log "Шаг 3: Создание структуры проекта"
echo "=============================================================" >> "$output_file"
echo "Структура проекта:" >> "$output_file"
echo "${project_name}/" >> "$output_file"
print_tree "${project_root}" ""

# Выводим список ненайденных файлов
if [ ${#not_found_files[@]} -gt 0 ]; then
    echo "=============================================================" >> "$output_file"
    echo "Список не найденных файлов:" >> "$output_file"
    for file in "${not_found_files[@]}"; do
        echo "$file" >> "$output_file"
    done
fi

log "Шаг 4: Информация успешно записана в файл out.txt"
