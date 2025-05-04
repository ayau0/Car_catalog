# car-catalog-backend/Dockerfile
FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


# Устанавливаем зависимости для dockerize
RUN apt-get update && \
    apt-get install -y curl && \
    curl -sSL https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz | tar -xzv -C /usr/local/bin



COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
