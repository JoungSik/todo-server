FROM golang:1.18.2

RUN mkdir -p /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

RUN go install github.com/cosmtrek/air@latest

EXPOSE 1323
CMD ["/bin/sh", "air"]