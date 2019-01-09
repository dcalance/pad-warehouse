FROM golang:latest

RUN mkdir /app

RUN go get github.com/gorilla/mux
RUN go get github.com/clbanning/anyxml
RUN go get github.com/gocql/gocql

ADD . /app/

WORKDIR /app/

RUN ls -la

RUN go build -o main .

CMD ["/app/main"]
