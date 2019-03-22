FROM golang:alpine

ADD ./src   /go/src/

WORKDIR  /go/src/

RUN apk add --no-cache git mercurial \
    && go get -d \
    && apk del git mercurial

ENV PORT=3001

# CMD ["go", "build", "main.go"]
