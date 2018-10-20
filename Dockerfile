FROM golang:1.11.1-alpine3.8 as builder
RUN apk --no-cache add make
WORKDIR /go/src/github.com/amnay-mo/fizzbuzz-api
COPY . /go/src/github.com/amnay-mo/fizzbuzz-api
RUN make build


FROM alpine:3.8
LABEL maintainer="amnay.m@gmail.com"
WORKDIR /app
ENTRYPOINT [ "/app/fizzbuzz-api" ]
COPY --from=builder /go/src/github.com/amnay-mo/fizzbuzz-api/build/* /app
