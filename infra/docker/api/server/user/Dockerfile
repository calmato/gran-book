# ビルド用コンテナ
FROM golang:1.15.6-alpine as builder

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /go/src/github.com/calmato/gran-book/api/server/user

RUN apk add --no-cache \
      git \
      make

ADD ./api/user ./
RUN make build

# 実行用コンテナ
FROM alpine

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /var/api

RUN apk add --no-cache \
      ca-certificates

COPY --from=builder /go/src/github.com/calmato/gran-book/api/server/user/app ./app

EXPOSE 8080
CMD ["./app"]
