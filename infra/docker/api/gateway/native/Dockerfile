# ビルド用コンテナ
FROM golang:1.16.7-alpine AS builder

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /go/src/github.com/calmato/gran-book/api/gateway/native

RUN apk add --update --no-cache \
  git \
  make

ADD ./api/gateway/native ./
RUN make build

# 実行用コンテナ
FROM alpine

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /var/gateway

RUN apk add --update \
  ca-certificates

COPY --from=builder /go/src/github.com/calmato/gran-book/api/gateway/native/app ./app

EXPOSE 8080
CMD ["./app"]
