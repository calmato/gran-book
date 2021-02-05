FROM golang:1.15.6-alpine

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

ENV NODE_VERSION 15.4.0
ENV PROTOBUF_VERSION 3.14.0
ENV GPRC_WEB_VERSION 1.2.1
ENV GRPC_GATEWAY_VERSION 2.0.1

RUN apk add --update \
  autoconf \
  automake \
  build-base \
  curl \
  git \
  libtool \
  nodejs \
  npm \
  tar \
  yarn

# Protocol Buffers
WORKDIR /tmp
RUN curl -L \
  -o /tmp/protobuf.tar.gz \
  "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protobuf-cpp-${PROTOBUF_VERSION}.tar.gz"
RUN tar xvzf protobuf.tar.gz

WORKDIR /tmp/protobuf-${PROTOBUF_VERSION}
RUN ./autogen.sh && ./configure && make -j 4 && make install

# gRPC Web
WORKDIR /tmp
RUN curl -L \
  -o /tmp/protoc-gen-grpc-web \
  "https://github.com/grpc/gprc-web/releases/download/${PROTOBUF_VERSION}/protoc-gen-grpc-web-${PROTOBUF_VERSION}-linux-x86_64"
RUN mv ./protoc-gen-grpc-web /usr/local/bin/protoc-gen-grpc-web
RUN chmod +x /usr/local/bin/protoc-gen-grpc-web

# ts-protoc-gen
RUN yarn global add ts-protoc-gen

# gRPC Gateway
# WORKDIR /go/src/github.com/grpc-ecosystem
# RUN git clone https://github.com/grpc-ecosystem/grpc-gateway
#
# WORKDIR /go/src/github.com/grpc-ecosystem/grpc-gateway -b "v${GRPC_GATEWAY_VERSION}"
