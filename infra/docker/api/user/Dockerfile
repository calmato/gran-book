# ビルド用コンテナ
FROM golang:1.17.2-alpine as builder

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /go/src/github.com/calmato/gran-book/api

RUN apk add --update --no-cache \
      git \
      make

ADD ./api ./
RUN make build SERVICE=user

# 実行用コンテナ
FROM alpine

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /var/api

RUN apk add --update --no-cache \
      ca-certificates

COPY --from=builder /go/src/github.com/calmato/gran-book/api/app ./app

EXPOSE 8080
CMD ["./app"]
