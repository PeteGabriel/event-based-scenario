FROM golang:1.17.1 as builder

ENV GO111MODULE=on
WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

EXPOSE 8208

CMD make run