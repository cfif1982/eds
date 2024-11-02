# Указываем базовый образ для сборки проекта
FROM golang:1.22.1 AS build

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы проекта в контейнер
COPY . .

# Загружаем зависимости и собираем проект
RUN go mod tidy
RUN go build -o eds  ./cmd/eds

# Используем минимальный образ для выполнения собранного приложения
FROM debian:bookworm-slim

# Копируем собранный бинарник из предыдущего контейнера
COPY --from=build /app/eds /usr/local/bin/eds

# Копируем файл конфигурации
COPY --from=build /app/config/prod.yaml /usr/local/bin/prod.yaml

# Копируем папку с миграциями
COPY --from=build /app/migrations /usr/local/bin/migrations

# Указываем рабочую директорию
WORKDIR /usr/local/bin

# Определяем команду, которая будет выполнена при запуске контейнера
# CMD ["eds", "--config ./srv.yaml"] - такой синтаксис неправильно принимает аргументы командной строки
CMD eds --config ./prod.yaml

