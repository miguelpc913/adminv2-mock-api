# Choose whatever you want, version >= 1.16
FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY admin/go.mod admin/go.sum admin/.air.toml ./

RUN go mod download
# refresh go.sum hash
RUN go mod tidy  

# CMD ["air"]
# if we want to specify .air.toml config
CMD ["air", "-c", ".air.toml"]
