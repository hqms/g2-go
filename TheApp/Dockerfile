FROM golang:alpine as build

COPY . /app
WORKDIR /app
RUN apk update && apk add gcc libc-dev syslinux-dev
RUN go get && go build main.go

# Prepare final, minimal image
FROM alpine

COPY --from=build /app/main /app/TheApp
ENV HOME /app
WORKDIR /app
EXPOSE 8080
ENTRYPOINT /app/TheApp