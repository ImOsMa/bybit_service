FROM golang:alpine

RUN go version

COPY . /github.com/ImOsMa/bybit_service/
WORKDIR /github.com/ImOsMa/bybit_service/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/ImOsMa/bybit_service/.bin/service .
COPY --from=0 /github.com/ImOsMa/bybit_service/configs configs/

CMD ["./service"]
