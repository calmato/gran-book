##################################################
# OS: Ubuntu 20.04 LTS
##################################################

#!/bin/bash

##################################################
# Variables
##################################################

##################################################
# Main
##################################################
### Check User
if [ "${USER}"  != "root" ]; then
  echo 'You must be logged in as "root" to run this tool.'
  exit 1
fi

### Packages
apt update -y && apt upgrade -y
apt install -y \
  apt-transport-https \
  bash-completion \
  ca-certificates \
  chrony \
  curl \
  git \
  gnupg \
  gnupg2 \
  jq \
  keyboard-configuration \
  language-pack-ja \
  make \
  net-tools \
  openssl \
  software-properties-common \
  traceroute \
  tree \
  unzip \
  vim \
  wget

### OS Basic
localectl set-locale LANG=C.UTF-8
localectl set-x11-keymap jp

timedatectl set-timezone Asia/Tokyo

### Cloud SDK
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
apt update -y
apt install -y google-cloud-sdk

### Cloud Proxy
# wget https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O cloud_sql_proxy
# chmod +x cloud_sql_proxy
# mv ./cloud_sql_proxy /usr/local/bin/cloud_sql_proxy

### kubectl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee -a /etc/apt/sources.list.d/kubernetes.list
apt update -y
apt install -y kubectl

### Let's Encrypt (Certbot)
apt install -y certbot

### Nginx
apt install -y nginx

### Nodejs
curl -fsSL https://deb.nodesource.com/setup_16.x | sudo -E bash -
apt install -y nodejs
npm install -g yarn

### Java
apt install -y openjdk-11-jre-headless

### Prometheus
apt install -y prometheus prometheus-node-exporter

### Grafana
wget -q -O - https://packages.grafana.com/gpg.key | sudo apt-key add -
echo "deb https://packages.grafana.com/oss/deb stable main" | sudo tee -a /etc/apt/sources.list.d/grafana.list
apt update -y
apt install -y grafana

### Loki
mkdir -p /etc/loki
curl -sLO https://github.com/grafana/loki/releases/download/v2.2.1/loki-linux-amd64.zip
curl -sL -o /etc/loki/config.yaml https://raw.githubusercontent.com/grafana/loki/master/cmd/loki/loki-local-config.yaml
unzip loki-linux-amd64.zip
mv loki-linux-amd64 /usr/local/bin/loki
chmod +x /usr/local/bin/loki

cat <<EOF > /etc/systemd/system/loki.service
[Unit]
Description=Grafana loki, a log aggregation system
Documentation=https://grafana.com/docs/loki
Before=grafana-server.service
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/loki --config.file=/etc/loki/config.yaml
Restart=always

[Install]
WantedBy=multi-user.target
EOF

cat <<EOF > /etc/rsyslog.d/91-loki.conf
# Log kernel generated Grafana loki log messages to file
:syslogtag, contains, "loki" /var/log/loki.log

# comment out the following line to allow loki messages through.
# Doing so means you'll also get loki messages in /var/log/syslog
# & stop
EOF

### Service
systemctl daemon-reload
systemctl disable --now ufw.service
systemctl restart rsyslog.service
systemctl enable --now nginx.service
systemctl enable --now loki.service
systemctl enable --now grafana-server.service
systemctl enable --now prometheus prometheus-node-exporter
