
FROM golang:1.19-alpine AS builder

# install git (required for fetching the dependencies)
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/mypackage/myapp/

COPY . .

# fetch dependencies
RUN go get -d -v

# build the binary.
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/server

# copy static content
RUN cp -r webjs /app
RUN cp .env /app

RUN ls -l /app

# STEP 2 build a small image

FROM scratch

COPY --from=builder /app /app

WORKDIR /app

ENTRYPOINT ["/app/server"]