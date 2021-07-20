# FROM alpine:3.13.5 AS builder

# ENV NGINX_VERSION 1.21.0
# ENV GEOIP_VERSION 3.3

# WORKDIR /tmp

# RUN apk add --update \
#   autoconf \
#   automake \
#   curl \
#   g++ \
#   gcc \
#   geoip-dev \
#   git \
#   libc-dev \
#   libmaxminddb-dev \
#   libtool \
#   linux-headers \
#   make \
#   openssl-dev \
#   pcre-dev \
#   tar \
#   zlib-dev

# RUN curl -fSLO "https://github.com/leev/ngx_http_geoip2_module/archive/refs/tags/${GEOIP_VERSION}.tar.gz"
# RUN tar -xvzf "${GEOIP_VERSION}.tar.gz" -C . \
#   && mv ./ngx_http_geoip2_module-${GEOIP_VERSION} ./ngx_http_geoip2_module

# RUN curl -fSLO "http://nginx.org/download/nginx-${NGINX_VERSION}.tar.gz"
# RUN tar -xvzf "nginx-${NGINX_VERSION}.tar.gz" -C . \
#   && cd ./nginx-${NGINX_VERSION} \
#   && ./configure --add-dynamic-module=/tmp/ngx_http_geoip2_module --with-stream \
#   && make \
#   && make install

# 実行用コンテナ
FROM nginx:1.21.0-alpine

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# COPY --from=builder /usr/local/nginx/modules/ngx_http_geoip2_module.so /etc/nginx/modules/ngx_http_geoip2_module.so
# COPY --from=builder /usr/local/nginx/modules/ngx_stream_geoip2_module.so /etc/nginx/modules/ngx_stream_geoip2_module.so

EXPOSE 80 443
CMD ["nginx", "-g", "daemon off;"]
