FROM golang:1.26-alpine

WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["CompileDaemon", "-polling", "-build=go build -o ./tmp/app .", "-command=./tmp/app"]
