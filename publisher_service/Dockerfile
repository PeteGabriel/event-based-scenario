FROM golang:1.17.1 as builder

ENV GO111MODULE=on

RUN mkdir /app
ADD . /app/
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

#CMD go test -v ./...

EXPOSE 8080

RUN go build -o bin/main main.go
CMD ./bin/main