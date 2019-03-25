FROM golang:alpine

ADD ./src /tmp/src/

#COPY /tmp/src/ /go/src/

WORKDIR /tmp/src/

RUN apk add --no-cache git mercurial \
    && go get -d \
    && apk del git mercurial

ENV PORT=3001

CMD ["go", "build", "main.go"]
