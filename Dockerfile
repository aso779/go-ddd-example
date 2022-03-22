#Migrations
FROM golang:1.17.6 as migrations

WORKDIR /app
COPY .go/pkg/mod /go/pkg/mod
COPY go.mod go.sum ./
COPY ./resources/config.yml ./resources/config.yml
RUN go mod download
COPY src ./src


#Build
FROM golang:1.17.6 as builder

WORKDIR /app
COPY .go/pkg/mod /go/pkg/mod
COPY go.mod go.sum ./
COPY src ./src
COPY gqlgen.yml ./gqlgen.yml

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o executable ./src/application

#Run
FROM alpine:3.12
LABEL maintainer="RZS crew <andrushchak@rozetka.com.ua>"

RUN apk update && \
    apk add --no-cache tzdata ca-certificates && \
    rm -rf /var/cache/apk/*

ENV TZ=Europe/Kiev
WORKDIR /root/
COPY --from=builder /app/executable .
COPY ./resources ./resources
ENV RESOURCES_PATH=/root/resources

EXPOSE 8080
CMD ["./executable", "-c", "./resources/config.yml"]