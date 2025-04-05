# Используем официальный образ Golang в качестве базового для сборки
FROM golang:alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем все файлы в контейнер
COPY . .

# Загружаем зависимости
RUN go mod download

# Собираем приложение
RUN go build -o app

# Используем образ Alpine для финальной сборки
FROM alpine

# Устанавливаем рабочую директорию
WORKDIR /

# Копируем скомпилированное приложение из стадии сборки
COPY --from=builder /app/app /app

COPY --from=builder /app/.env .env


# Устанавливаем точку входа для контейнера
ENTRYPOINT ["/app"]
