FROM golang:alpine

RUN addgroup -g 1337 go && \
    adduser -D -u 1337 -G go go
USER go

WORKDIR /harvey

ADD ./harvey/ /harvey/

CMD go version
