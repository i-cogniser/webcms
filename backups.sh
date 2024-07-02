#!/bin/bash

# Переменные
DB_NAME="webcmsdb"
DB_USER="userwebcms"
DB_PASSWORD="adminwebcms"
DB_HOST="localhost"
BACKUP_DIR="/backups"
DATE=$(date +'%Y-%m-%d_%H-%M-%S')
BACKUP_FILE="${BACKUP_DIR}/backup_${DATE}.sql"

# Выполнение резервного копирования
PGPASSWORD=$DB_PASSWORD pg_dump -h $DB_HOST -U $DB_USER $DB_NAME > $BACKUP_FILE

# Проверка результата резервного копирования
if [ $? -eq 0 ]; then
  echo "Резервное копирование успешно завершено: ${BACKUP_FILE}"
else
  echo "Ошибка резервного копирования"
  exit 1
fi
