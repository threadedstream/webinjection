FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build backend/cmd/backend/main.go && \
    go install github.com/pressly/goose/v3/cmd/goose@latest

COPY entrypoint.sh .

CMD ["bash", "entrypoint.sh"]