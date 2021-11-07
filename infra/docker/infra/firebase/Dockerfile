FROM openjdk:17.0.1-jdk-slim

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

ENV NODE_VERSION 16.13.0
ENV FIREBASE_VERSION 9.22.0

RUN apt update
RUN apt install -y --no-install-recommends \
  curl \
  tar

WORKDIR /tmp

# Install Node
RUN curl -fSLO "https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-linux-x64.tar.gz"
RUN curl -fSLO "https://nodejs.org/dist/v${NODE_VERSION}/SHASUMS256.txt"
RUN cat SHASUMS256.txt | grep "node-v${NODE_VERSION}-linux-x64.tar.gz" | sha256sum -c - \
  && tar -xvzf "node-v${NODE_VERSION}-linux-x64.tar.gz" -C . \
  && mv node-v${NODE_VERSION}-linux-x64/bin/* /usr/local/bin/ \
  && mv node-v${NODE_VERSION}-linux-x64/lib/node_modules/ /usr/local/lib/ \
  && rm -f "node-v${NODE_VERSION}-linux-x64.tar.gz" SHASUMS256.txt

RUN npm install -g firebase-tools@${FIREBASE_VERSION}

WORKDIR /var/firebase

EXPOSE 4000
EXPOSE 8080
EXPOSE 9099

CMD ["firebase", "emulators:start"]
