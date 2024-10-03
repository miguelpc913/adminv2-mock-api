FROM golang:1.23-alpine

WORKDIR /app

#RUN go install github.com/cosmtrek/air@latest
# Issue reported, solution:
# https://github.com/air-verse/air/issues/532
##############################################
RUN go install github.com/air-verse/air@latest

COPY admin/ ./

RUN go mod download
# refresh go.sum hash
RUN go mod tidy

HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl -f http://localhost:8080/health || exit 1

# CMD ["air"]
# if we want to specify .air.toml config
CMD ["air", "-c", ".air.toml"]
