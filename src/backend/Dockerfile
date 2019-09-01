FROM golang:1-alpine

RUN apk update \
	&& apk upgrade \
	&& apk add git

# * Install required packages
RUN go get -u \
	github.com/graph-gophers/graphql-go \
	github.com/graph-gophers/graphql-go/relay \
	github.com/jinzhu/gorm \
	github.com/jinzhu/gorm/dialects/postgres

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go build -o main ./cmd

CMD ["/app/main"]
