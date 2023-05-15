FROM golang:1.19-alpine

RUN mkdir /app

ADD . /app

# static web site
# TODO

WORKDIR /app

RUN go build -o server main.go

CMD ["/app/server"]
