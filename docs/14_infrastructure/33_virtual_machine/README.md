# インフラストラクチャ - 仮想マシン

[root](./../../../README.md) 
/ [14_infrastructure](./../README.md) 
/ [33_virtual_machine](./README.md)

## 実装メモ

### GCE初期設定

1. [ubuntu-setup.sh](./../../../infra/server/bin/ubuntu-setup.sh) をscpで転送
2. root権限に昇格
3. スクリプトの実行

### Cloud SDKの設定

1. 初期化
  > $ gcloud init

2. GKE用の証明書取得
  > $ gcloud container clusters get-credentials <gke-cluster-name> \
  >   --zone <zone> \
  >   --project<project-id>

### Let's Encryptの設定

1. テスト実行
  > $ certbot certonly \
  >   --dry-run \
  >   --manual \
  >   --agree-tos \
  >   --no-eff-email \
  >   --manual-public-ip-logging-ok \
  >   --preferred-challenges dns \
  >   -m "<email>" \
  >   -d "<domain>" \

2. 証明書の発行
  > $ certbot certonly \
  >   --manual \
  >   --agree-tos \
  >   --no-eff-email \
  >   --manual-public-ip-logging-ok \
  >   --preferred-challenges dns \
  >   -m "<email>" \
  >   -d "<domain>" \

3. 証明書更新のテスト実行
  > $ certbot renew

---

## 参考

* Ubuntu関連
  * [Ubunutu - パッケージ検索](https://packages.ubuntu.com/ja/)
* Let's Encrypt関連
  * [certbot](https://certbot.eff.org/)
  * [個人ブログ - Let's Encrypt ワイルドカード自動更新（ConoHa）](https://www.eastforest.jp/vps/6149)
