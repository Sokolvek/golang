FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 8080

# Передаем строку подключения к базе данных как аргумент
CMD ["/app/main", "postgres://postgres:3344@localhost:5432/go"]

