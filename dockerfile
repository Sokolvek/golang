FROM golang:apline as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV GOCACHE=/go/.cache
RUN go build -o cmd/main.go

FROM ubuntu:22.04
RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/app .
ENTRYPOINT ["./app"]