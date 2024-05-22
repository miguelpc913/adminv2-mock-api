# Choose whatever you want, version >= 1.16
FROM cosmtrek/air

WORKDIR /app

COPY admin/ ./

RUN go mod download
# refresh go.sum hash
RUN go mod tidy  

# CMD ["air"]
# if we want to specify .air.toml config
CMD ["air", "-c", ".air.toml"]
