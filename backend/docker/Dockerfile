FROM golang:1.23.0-alpine

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPROXY=https://proxy.golang.org,direct

RUN apk add --no-cache bash git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN chmod +x ./docker/Entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["docker/Entrypoint.sh"]
