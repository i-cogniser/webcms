#!/bin/bash

# Указываем корневую папку проекта как текущую директорию
projectRoot="."

# Создаем подкаталоги в корневой папке проекта
directories=(
    "cmd"
    "config"
    "controllers"
    "middlewares"
    "models"
    "repositories"
    "services"
    "templates/admin"
    "utils"
)

for dir in "${directories[@]}"; do
    mkdir -p "$projectRoot/$dir"
done

# Создаем основные файлы в корневой папке проекта
files=(
    "cmd/main.go"
    "config/config.go"
    "controllers/auth_controller.go"
    "controllers/content_controller.go"
    "controllers/user_controller.go"
    "middlewares/auth_middleware.go"
    "models/user.go"
    "models/post.go"
    "models/page.go"
    "repositories/user_repository.go"
    "repositories/post_repository.go"
    "repositories/page_repository.go"
    "services/auth_service.go"
    "services/content_service.go"
    "services/user_service.go"
    "templates/base.html"
    "templates/index.html"
    "templates/admin/dashboard.html"
    "templates/admin/edit_post.html"
    "utils/logger.go"
    ".env"
    "go.mod"
    "go.sum"
)

for file in "${files[@]}"; do
    touch "$projectRoot/$file"
done

# Заполняем файлы базовым содержимым

# main.go
cat <<EOF > "$projectRoot/cmd/main.go"
package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("CMS is running...")
    http.ListenAndServe(":8080", nil)
}
EOF

# config.go
cat <<EOF > "$projectRoot/config/config.go"
package config

import (
    "os"
)

func GetConfig(key string) string {
    return os.Getenv(key)
}
EOF

# base.html
cat <<EOF > "$projectRoot/templates/base.html"
<!DOCTYPE html>
<html>
<head>
    <title>{{ .Title }}</title>
</head>
<body>
    {{ block "content" . }}{{ end }}
</body>
</html>
EOF

# index.html
cat <<EOF > "$projectRoot/templates/index.html"
{{ define "content" }}
<h1>Welcome to the CMS</h1>
{{ end }}
EOF

# .env
cat <<EOF > "$projectRoot/.env"
# Environment variables
PORT=8080
DATABASE_URL=postgres://user:password@localhost:5432/webcms
EOF

# Создаем содержимое для go.mod с объявлением модуля и с последней версией Go
goModContent=$(cat <<EOF
module example.com/webcms

go 1.18
EOF
)

# Записываем содержимое go.mod в файл без BOM
echo "$goModContent" > "$projectRoot/go.mod"

# Переходим в рабочую директорию проекта
cd "$projectRoot" || exit

# Выполняем go mod tidy для управления зависимостями
echo "Running go mod tidy..."
go mod tidy

echo "Project structure for CMS has been created successfully."
