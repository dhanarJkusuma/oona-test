FROM golang:1.10

RUN mkdir -p /go/src/oona-test
WORKDIR /go/src/oona-test

ADD . /go/src/oona-test

RUN go get -v

RUN go build -o main . 

EXPOSE 8080
CMD ["/go/src/oona-test/main"]

