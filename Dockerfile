FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build webinjection/cmd/webinjection/main.go

RUN go install github.com/pressly/goose/v3/cmd/goose@latest


CMD ["./main"]