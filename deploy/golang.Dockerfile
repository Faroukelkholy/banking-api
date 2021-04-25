FROM golang:1.15

WORKDIR /go/src/bank

COPY . .

RUN go mod tidy

EXPOSE 3000

CMD ["go","run","cmd/main.go"]