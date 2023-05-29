FROM golang:1.20.3 AS builder

ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1
ENV GOOS=linux

WORKDIR /go/src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o restuarant-service

ENTRYPOINT ["/go/src/restuarant-service"]
