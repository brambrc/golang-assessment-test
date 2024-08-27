FROM golang:1.20-alpine

RUN apk add --no-cache curl bash

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz -o migrate.tar.gz \
    && tar -xvzf migrate.tar.gz \
    && mv migrate /usr/local/bin/migrate \
    && rm migrate.tar.gz

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .

EXPOSE 8080

CMD ["migrate", "-path", "db/migrations", "-database", "postgres://postgres:yourpassword@db:5432/mezink?sslmode=disable", "up"]

CMD ["go", "run", "main.go"]
