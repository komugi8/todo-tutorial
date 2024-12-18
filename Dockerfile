FROM golang:1.23-alpine3.21

WORKDIR /app

RUN apk update && apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]
