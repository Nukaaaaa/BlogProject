# Используем официальный образ Go
FROM golang:1.23.6

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum из user-service
COPY ./user-service/go.mod ./go.mod
COPY ./user-service/go.sum ./go.sum
RUN go mod download

# Копируем остальные файлы user-service
COPY ./user-service/. .

# Собираем бинарный файл
RUN go build -o main .

# Указываем порт, на котором работает сервис
EXPOSE 8081

# Запуск приложения
CMD ["./main"]
