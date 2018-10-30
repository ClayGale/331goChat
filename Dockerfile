FROM golang:1.11.1


COPY ./go /go/src/github.com/ClayGale/331goChat
WORKDIR /go/src/github.com/ClayGale/331goChat

RUN go get ./

RUN go build -o main .
EXPOSE 9090
CMD ["/go/main"]
