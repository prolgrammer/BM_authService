FROM golang:1.23.1-alpine as builder

WORKDIR /auth

COPY go.mod go.sum ./
RUN go mod download

COPY . .
#RUN go install github.com/swaggo/swag/cmd/swag@latest
#RUN swag init -g cmd/main.go -o ./docs
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR etc/auth

COPY --from=builder /auth/main .
COPY --from=builder /auth/config ./config

ENTRYPOINT ["./main"]

