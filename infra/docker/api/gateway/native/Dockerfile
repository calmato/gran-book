# ビルド用コンテナ
FROM node:16.1.0-alpine as builder

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /var/gateway

RUN apk add --update \
  g++ \
  make \
  python

ADD ./api/gateway/native ./
RUN yarn install --production=false

RUN yarn build

# 実行用コンテナ
FROM node:16.1.0-alpine

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo
ENV NODE_ENV production

WORKDIR /var/gateway

RUN apk add \
  g++ \
  make \
  python

COPY --from=builder /var/gateway/package.json ./package.json
RUN yarn install --production=true

COPY --from=builder /var/gateway/dist ./dist

EXPOSE 3000
CMD ["yarn", "start"]
