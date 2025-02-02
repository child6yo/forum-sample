FROM golang:1.23

WORKDIR /forum-service

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o forum-service ./cmd/main.go

CMD ["./forum-service"]

