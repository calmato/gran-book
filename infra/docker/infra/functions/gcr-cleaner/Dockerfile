FROM golang:1.15.6-alpine AS builder

WORKDIR /go/src/github.com/calmato/gran-book/infra/functions/gcr-cleaner

RUN apk add --update --no-cache \
      git \
      make

ADD ./infra/functions/gcr-cleaner ./
RUN go build -o app ./cmd/server/main.go

FROM alpine

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /var/functions

RUN apk add --update --no-cache \
      ca-certificates

COPY --from=builder /go/src/github.com/calmato/gran-book/infra/functions/gcr-cleaner/app ./app

EXPOSE 8080
CMD ["./app"]
