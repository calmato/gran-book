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

### Prometheus
apt install -y prometheus prometheus-node-exporter

### Grafana
wget -q -O - https://packages.grafana.com/gpg.key | sudo apt-key add -
echo "deb https://packages.grafana.com/oss/deb stable main" | sudo tee -a /etc/apt/sources.list.d/grafana.list
apt update -y
apt install -y grafana

### Service
systemctl disable --now ufw.service
systemctl restart rsyslog.service
systemctl enable --now nginx.service
systemctl enable --now grafana-server.service
systemctl enable --now prometheus prometheus-node-exporter
