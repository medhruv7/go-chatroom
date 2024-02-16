# Docker container for chat room application built using websockets in go

FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go-chatroom

EXPOSE 8080

CMD ["/go-chatroom"]