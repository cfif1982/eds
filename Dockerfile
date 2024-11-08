# Указываем базовый образ для сборки проекта
FROM golang:1.22.1 AS build

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы проекта в контейнер
COPY . .

# Загружаем зависимости и собираем проект
RUN make tidy
RUN make build_linux
RUN make build_migrator_linux

# Используем минимальный образ для выполнения собранного приложения
FROM debian:bookworm-slim

# Копируем собранный бинарник из предыдущего контейнера
COPY --from=build /app/eds /usr/local/bin/eds
COPY --from=build /app/migrator /usr/local/bin/migrator

# Копируем файл конфигурации
COPY --from=build /app/config/prod.yaml /usr/local/bin/prod.yaml

# Копируем папку с миграциями
COPY --from=build /app/internal/migrator/migrations /usr/local/bin/migrations

# Указываем рабочую директорию
WORKDIR /usr/local/bin

RUN ./migrator --config ./prod.yaml

# Определяем команду, которая будет выполнена при запуске контейнера
# CMD ["eds", "--config ./srv.yaml"] - такой синтаксис неправильно принимает аргументы командной строки
CMD eds --config ./prod.yaml

