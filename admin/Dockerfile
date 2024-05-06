FROM golang:alpine

WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .

CMD ["air"]