# Этап 1: Сборка приложения
FROM golang:1.22-alpine3.18 AS builder
# Установка необходимых утилит
RUN apk update && \
    apk add --no-cache git

# Создание рабочей директории
WORKDIR /app

# Копирование файлов проекта и загрузка зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копирование остального кода приложения в рабочую директорию
COPY . .

# Проверка наличия cmd/main.go и других исходных файлов
RUN ls -la ./cmd/main.go

# Компиляция приложения
RUN go build -o /app/main ./cmd/main.go

# Проверка, что файл main был скомпилирован успешно
RUN if [ ! -f /app/main ]; then echo "Compilation failed"; exit 1; fi

# Этап 2: Запуск приложения
FROM alpine:3.18

# Установка необходимых утилит
RUN apk update && \
    apk add --no-cache libgcc libc6-compat \
    # Установка curl для проверки состояния
    && apk add --no-cache curl \
    # Установка tar для распаковки
    && apk add --no-cache tar \
    # Установка netcat для проверки состояния
    && apk add --no-cache netcat-openbsd \
    # Установка оболочки ash для проверки состояния
    && apk add --no-cache bash

# Создание рабочей директории
WORKDIR /app

# Копирование скомпилированного приложения из предыдущего этапа
COPY --from=builder /app/main /app/main

# Установка прав на выполнение файла main
RUN chmod +x /app/main

# Копирование статических файлов и конфигураций, если они существуют
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/.env /app/.env

# Определение переменных окружения
ENV DATABASE_URL=${DATABASE_URL}
ENV FRONTEND_PATH=${FRONTEND_PATH}
ENV JWT_SECRET=${JWT_SECRET}

# Открываем порт приложения
EXPOSE 8080

# Скрипт ожидания готовности базы данных и запуск приложения
CMD ["sh", "-c", "until nc -z -v -w30 db 5432; do echo 'Waiting for database connection...'; sleep 5; done; /app/main"]
