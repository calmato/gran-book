FROM alpine:3.12.3

ENV TERRAFORM_VERSION 0.14.4

RUN apk add --no-cache \
      ca-certificates \
      make \
      unzip \
      wget

RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip
RUN unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip

RUN mv terraform /usr/local/bin
RUN chmod +x /usr/local/bin/terraform
