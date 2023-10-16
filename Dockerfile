FROM golang:1.21.3-alpine3.18 as binary
LABEL authors="kalunik"
WORKDIR /go/src/github.com/kalunik/urShorty/
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o urShorty cmd/main.go

FROM alpine:3.18
WORKDIR /root/
COPY --from=binary /go/src/github.com/kalunik/urShorty/urShorty ./
CMD ["./urShorty"]