#Build
FROM golang:1.18.1 as builder

WORKDIR /app
COPY src ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o executable ./application

#Run
FROM alpine:3.15.4
LABEL maintainer="aso779"

RUN apk update && \
    apk add --no-cache tzdata ca-certificates && \
    rm -rf /var/cache/apk/*

ENV TZ=Europe/London
WORKDIR /app/
COPY --from=builder /app/executable .
COPY ./resources ./resources

EXPOSE 8080
CMD ["./executable", "-c", "./resources/config.yml"]