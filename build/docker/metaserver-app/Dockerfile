FROM golang:1.22.3-alpine

ARG WEB_APPLICATION_PORT=8081

RUN apk add inkscape

WORKDIR /app
COPY . /app
RUN go build -o bin/metaserver ./cmd/metadata-server/main.go

EXPOSE ${WEB_APPLICATION_PORT}

CMD ["./bin/metaserver"]
