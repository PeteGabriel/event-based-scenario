FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .

EXPOSE 8200

CMD make run