# ビルド用コンテナ
FROM golang:1.16.7-alpine AS builder

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /go/src/github.com/calmato/gran-book/api/gateway/admin

RUN apk add --update --no-cache \
  git \
  make

ADD ./api/gateway/admin ./
RUN make build

# 実行用コンテナ
FROM alpine

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /var/gateway

RUN apk add --update \
  ca-certificates

COPY --from=builder /go/src/github.com/calmato/gran-book/api/gateway/admin/app ./app

EXPOSE 8080
CMD ["./app"]
