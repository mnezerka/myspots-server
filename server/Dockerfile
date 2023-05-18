# STEP 1 build executable binary

FROM golang:1.19-alpine AS builder

# install git (required for fetching the dependencies)
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN mkdir app

# fetch dependencies
RUN go get -d -v

# build the binary.
RUN go build -o /app/server

# copy static content
RUN cp -r webjs /app/
#ENTRYPOINT ["/bin/sh"]

# STEP 2 build a small image

FROM scratch

# copy our static executable.
COPY --from=builder /app/server /app/server

#COPY --from=builder /app/webjs /app/

#WORKDIR /app

ENTRYPOINT ["/app/server"]