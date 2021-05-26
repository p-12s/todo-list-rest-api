FROM golang:1.16.4-alpine3.13 AS builder

RUN go version
RUN apk add git

COPY ./ /github.com/p-12s/todo-list-rest-api
WORKDIR /github.com/p-12s/todo-list-rest-api

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

#lightweight docker container with binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /github.com/p-12s/todo-list-rest-api/.bin/app .
COPY --from=0 /github.com/p-12s/todo-list-rest-api/pkg/configs/ ./configs/

EXPOSE 8000

CMD [ "./app"]

# в деплой пока углубляться не буду
# продолжить выполнение отсюда: https://youtu.be/ceRTzKuCLnw?list=PLbTTxxr-hMmxZMXsvaE-PozXxktdJ5zLR&t=143
