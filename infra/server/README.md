# サーバ構築関連のメモ

## ファイル配置先

* bin
  * certbot-auth-hook.sh -> `/var/scripts/certbot-auth-hook.sh`
  * certbot-cleanup-hook.sh -> `/var/scripts/certbot-cleanup-hook.sh`
  * ubuntu-setup.sh -> `/var/scripts/ubuntu-setup.sh`
  * ubuntu-test.sh -> `/var/scripts/ubuntu-test.sh`
  * update-certificate.sh -> `/var/scripts/update-certificate.sh`
* config
  * crontab -> `/etc/crontab`
  * loki.service -> `/etc/systemd/system/loki.service`
  * loki.yaml -> `/etc/loki/config.yaml`
  * nginx.conf -> `/etc/nginx/conf.d/default.conf`
* github.com
  * calmato/gran-book -> `/opt/gran-book`
  * swagger-api/swagger-ui -> `/opt/swagger-ui`

## 実行手順メモ

1. 各種ファイル配置
    > $ mkdir -p /var/scripts

    > $ scp -r ./gran-book/infra/server/bin/*.sh [user]@[remote]:/var/scripts

2. セットアップの実行
    > $ /var/scripts/ubuntu-setup.sh

3. Let's Enctyptのテスト実行
    > $ certbot certonly \  
    >   --dry-run \  
    >   --manual \  
    >   --agree-tos \  
    >   --no-eff-email \  
    >   --manual-public-ip-logging-ok \  
    >   --preferred-challenges dns-01 \  
    >   --server https://acme-v02.api.letsencrypt.org/directory \  
    >   --manual-auth-hook /var/scripts/certbot-auth-hook.sh \  
    >   --manual-cleanup-hook /var/scripts/certbot-cleanup-hook.sh \  
    >   -m "[email]" \  
    >   -d "[domain]"

4. 証明書の発行
    > $ certbot certonly \  
    >   --manual \  
    >   --agree-tos \  
    >   --no-eff-email \  
    >   --manual-public-ip-logging-ok \  
    >   --manual-public-ip-logging-ok \  
    >   --preferred-challenges dns-01 \  
    >   --server https://acme-v02.api.letsencrypt.org/directory \  
    >   --manual-auth-hook /var/scripts/certbot-auth-hook.sh \  
    >   --manual-cleanup-hook /var/scripts/certbot-cleanup-hook.sh \  
    >   -m "[email]" \  
    >   -d "[domain]>"

5. 証明書更新のテスト実行
    > $ certbot renew --force-renewal --dry-run

6. Cron設定ファイルの編集
    > $ vi /etc/crontab

7. Cronの再起動
    > $ systemctl restart cron.service

8. Nginx設定ファイルをscpで転送
    > $ scp infra/server/config/nginx.conf [account]@[ip-addr]:/etc/nginx/conf.d/default.conf

9. Nginxサービスの再起動
    > $ systemctl restart nginx.service
