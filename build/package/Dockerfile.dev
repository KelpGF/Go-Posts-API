FROM golang:1.22

WORKDIR /usr/src/app

RUN apt update

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go mod tidy

CMD ["tail", "-f", "/dev/null"]
