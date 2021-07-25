# インフラストラクチャ - 仮想マシン

[root](./../../../README.md) 
/ [14_infrastructure](./../README.md) 
/ [33_virtual_machine](./README.md)

## 実装メモ

### GCE初期設定

1. [ubuntu-setup.sh](./../../../infra/server/bin/ubuntu-setup.sh) をscpで転送
2. root権限に昇格
3. スクリプトの実行
4. [Cloud SDKの設定](#cloud-sdkの設定)
5. [Let's Encryptの設定](#lets-encryptの設定)
6. [Cronの設定](#cronの設定)

### Cloud SDKの設定

1. 初期化
  > $ gcloud init

2. GKE用の証明書取得
  > $ gcloud container clusters get-credentials [gke-cluster-name] \  
  >   --zone [zone] \  
  >   --project [project-id]

### Let's Encryptの設定

1. テスト実行
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

2. 証明書の発行
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

3. 証明書更新のテスト実行
  > $ certbot renew --force-renewal --dry-run

### Cronの設定

1. 設定ファイルの編集
  > $ vi /etc/crontab

2. サービスの再起動
  > $ systemctl restart cron.service

### Nginxの設定

1. 設定ファイルをscpで転送
  > $ scp infra/server/config/nginx.conf [account]@[ip-addr]:/etc/nginx/conf.d/default.conf

2. サービスの再起動
  > $ systemctl restart nginx.service

### Swagger UIの設定

1. リポジトリのクローン
    > $ cd /opt

    > $ git clone https://github.com/swagger-api/swagger-ui.git

    > $ git clone https://github.com/calmato/gran-book.git

2. Swagger UIのビルドファイルをコピー
    > $ mkdir -p /var/www/html/swagger

    > $ cp -r ./swagger-ui/dist/ /var/www/html/swagger/native

    > $ cp -r ./swagger-ui/dist/ /var/www/html/swagger/admin

3. openapi.yamlの生成
    > $ cd /opt/gran-book/docs/12_backend/11_swagger

    > $ yarn; yarn generate

3. シンボリックリンクの作成
    > $ ln -s \
    > /opt/gran-book/tmp/data/swagger/native/openapi/openapi.yaml \
    > /var/www/html/swagger/native/openapi.yaml

    > $ ln -s \
    > /opt/gran-book/tmp/data/swagger/admin/openapi/openapi.yaml \
    > /var/www/html/swagger/admin/openapi.yaml

4. 設定ファイルの編集
    > $ vi /var/www/html/swagger/native/index.html

    > $ vi /var/www/html/swagger/native/index.html

```html
 38     <script>
 39     window.onload = function() {
 40       // Begin Swagger UI call region
 41       const ui = SwaggerUIBundle({
-42         url: "https://petstore.swagger.io/v2/swagger.json",
+42         url: "./openapi.yaml",
 43         dom_id: '#swagger-ui',
 44         deepLinking: true,
 45         presets: [
 46           SwaggerUIBundle.presets.apis,
 47           SwaggerUIStandalonePreset
 48         ],
 49         plugins: [
 50           SwaggerUIBundle.plugins.DownloadUrl
 51         ],
 52         layout: "StandaloneLayout"
 53       });
 54       // End Swagger UI call region
 55
 56       window.ui = ui;
 57     };
 58   </script>
```

5. Nginxサービスの再起動
    > $ systemctl restart nginx.service

---

## 参考

* Ubuntu関連
  * [Ubunutu - パッケージ検索](https://packages.ubuntu.com/ja/)
* Let's Encrypt関連
  * [certbot](https://certbot.eff.org/)
  * [個人ブログ - Let's Encrypt ワイルドカード自動更新（ConoHa）](https://www.eastforest.jp/vps/6149)
