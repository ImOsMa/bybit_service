FROM golang:alpine

ENV GIN_MODE=release

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o bybit_service ./cmd/main.go

CMD ["./bybit_service"]
