FROM golang:1.18-alpine AS dev
COPY . /go/src/api
WORKDIR /go/src/api
RUN apk update && apk upgrade && apk add --no-cache git mysql-client \
    && go get -u \
    && apk del --purge git \
    && go install github.com/cosmtrek/air@latest

ENV SERVER_PORT=3000

FROM golang:1.18-alpine AS build
COPY . /go/src/api
WORKDIR /go/src/api
RUN apk update && apk upgrade && apk add --no-cache git \
    && go get -u \
    && GOOS=linux GOARCH=amd64 go build -o app .

FROM gcr.io/distroless/base AS deploy
WORKDIR /go/src/api
COPY --from=build /go/src/app/api .

EXPOSE 3000
ENV SERVER_PORT=3000
EXPOSE 3000
CMD ["./app"]
