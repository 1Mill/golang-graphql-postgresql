FROM golang:1-alpine

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go build -o main ./cmd

CMD ["/app/main"]
