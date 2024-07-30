# Этап 1: Сборка фронтенда
FROM node:18-alpine AS frontend-builder

# Создание рабочей директории
WORKDIR /frontend

# Копирование файлов проекта и установка зависимостей
COPY webcms-vue/package.json webcms-vue/package-lock.json ./
RUN npm install

# Копирование остального кода фронтенда в рабочую директорию
COPY webcms-vue ./

# Сборка фронтенд-приложения и сохранение логов
RUN npm run build > /frontend/build.log 2>&1 || (cat /frontend/build.log && exit 1)

# Этап 2: Запуск сервера для статики
FROM node:18-alpine

# Установка http-server для обслуживания статических файлов
RUN npm install -g http-server

# Установка curl
RUN apk add --no-cache curl

# Создание рабочей директории
WORKDIR /frontend

# Копирование собранных файлов из предыдущего этапа
COPY --from=frontend-builder /frontend/dist /frontend/dist

# Определение порта для сервера
EXPOSE 80

# Запуск http-server для обслуживания файлов из директории dist
CMD ["http-server", "dist", "-p", "80"]
