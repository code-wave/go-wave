FROM golang:1.15.7 as builder

WORKDIR /go/src
COPY ./src /go/src
RUN go build -o start

FROM golang:1.15.7
COPY --from=builder /go/src /go/src
WORKDIR /go/src