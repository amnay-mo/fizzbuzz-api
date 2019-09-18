FROM golang:1.12-stretch as builder
ENV GO111MODULE=auto
WORKDIR /go/src/github.com/amnay-mo/fizzbuzz-api
COPY . /go/src/github.com/amnay-mo/fizzbuzz-api
RUN make get build


FROM debian:stretch-slim
LABEL maintainer="amnay.m@gmail.com"
WORKDIR /app
ENTRYPOINT [ "/app/fizzbuzz-api" ]
COPY --from=builder /go/src/github.com/amnay-mo/fizzbuzz-api/build/* /app
