FROM golang:1.16.3

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

ENV GO_VERSION 1.16.3
ENV NODE_VERSION 16.0.0

ENV PROTOBUF_VERSION 3.15.8
ENV PROTOTOOL_VERSION 1.10.0
ENV GRPC_WEB_VERSION 1.2.1
ENV GRPC_GATEWAY_VERSION 2.3.0

RUN apt-get update
RUN apt-get install -y --no-install-recommends \
  autoconf \
  automake \
  curl \
  git \
  libtool \
  tar
RUN rm -rf /var/lib/apt/lists/*

WORKDIR /tmp

# Install Node
RUN curl -fSLO "https://nodejs.org/dist/v${NODE_VERSION}/node-v{$NODE_VERSION}-linux-x64.tar.gz"
RUN curl -fSLO "https://nodejs.org/dist/v${NODE_VERSION}/SHASUMS256.txt"
RUN NODE_CHECKSUM=`cat SHASUMS256.txt | grep "node-v${NODE_VERSION}-linux-x64.tar.gz"` \
  && echo ${NODE_CHECKSUM} | sha256sum -c - \
  && tar -xvzf "node-v${NODE_VERSION}-linux-x64.tar.gz" -C . \
  && mv node-v${NODE_VERSION}-linux-x64/bin/* /usr/local/bin/ \
  && mv node-v${NODE_VERSION}-linux-x64/lib/node_modules/ /usr/local/lib/ \
  && rm -f "node-v${NODE_VERSION}-linux-x64.tar.gz" SHASUMS256.txt

# Protocol Buffers
RUN curl -fSLO "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protobuf-cpp-${PROTOBUF_VERSION}.tar.gz"
RUN tar -xvzf "protobuf-cpp-${PROTOBUF_VERSION}.tar.gz" -C . \
  && cd ./protobuf-${PROTOBUF_VERSION} \
  && ./autogen.sh \
  && ./configure \
  && make -j 4 \
  && make install

# # gRPC Web
# RUN curl -fSLO "https://github.com/grpc/grpc-web/releases/download/${GRPC_WEB_VERSION}/protoc-gen-grpc-web-${GRPC_WEB_VERSION}-$(uname -s)-$(uname -m)"
# RUN mv ./protoc-gen-grpc-web-${GRPC_WEB_VERSION}-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web \
#   && chmod +x /usr/local/bin/protoc-gen-grpc-web

# # gRPC Gateway
# WORKDIR /go/src/github.com/grpc-ecosystem
# RUN git clone https://github.com/grpc-ecosystem/grpc-gateway
#
# WORKDIR /go/src/github.com/grpc-ecosystem/grpc-gateway -b "v${GRPC_GATEWAY_VERSION}"

# Install Prototool
RUN curl -fSLO "https://github.com/uber/prototool/releases/download/v${PROTOTOOL_VERSION}/prototool-$(uname -s)-$(uname -m)"
RUN mv ./prototool-$(uname -s)-$(uname -m) /usr/local/bin/prototool \
  && chmod +x /usr/local/bin/prototool

# Install gRPC tools for Node.js
RUN npm install -g grpc grpc-tools grpc_tools_node_protoc_ts

RUN ldconfig
