# Choose whatever you want, version >= 1.16
FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY admin/ ./

RUN go mod download
# refresh go.sum hash
RUN go mod tidy  

CMD ["air"]
# CMD ["air", "-c", ".air.toml"]
