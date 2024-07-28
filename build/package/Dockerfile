FROM golang:1.22 AS builder

WORKDIR /usr/src/app

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go mod tidy

RUN sh ./scripts/gen-swag.sh

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main cmd/server/main.go

FROM scratch

COPY --from=builder /usr/src/app/main /

ENTRYPOINT [ "/main" ]