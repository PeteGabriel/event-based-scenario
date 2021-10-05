FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

#CMD go test -v ./...

EXPOSE 8200

RUN go build -o bin/main main.go
CMD ./bin/main