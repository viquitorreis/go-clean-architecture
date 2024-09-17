FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o main

EXPOSE 6969

CMD ["./main"]
