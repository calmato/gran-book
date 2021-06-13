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
