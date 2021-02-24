# Gran Book

<!-- CI/CDのバッチを貼り付け -->
読書管理アプリ

## 各種設定

<!-- 使用するミドルウェア,言語のバージョン等記載 -->
<details open>
<summary>依存関係</summary>

* Frontend
  * Node: 15.4.0
  * Expo: 4.0.13
* Backend
  * Golang: 1.15.6
  * Node: 15.4.0
  * MySQL: 8.0.22

</details>

<!-- 環境構築手順を記載 -->
<details open>
<summary>環境構築</summary>

### リポジトリのダウンロード

> $ git clone https://github.com/calmato/gran-book.git

> $ cd ./presto-pay

### コンテナの初期設定

* コンテナの作成

> $ make setup

* Firebase Admin SDKを `secretディレクトリ` にコピペ

* .envファイルの編集

* secretディレクトリのファイルを.envへ書き込み

> $ var=$(jq -c < ./secrets/[ファイル名])

> $ sed -i -e "/^GCP_SERVICE_KEY_JSON=.*/d" .env

> $ echo -E "GCP_SERVICE_KEY_JSON=${var}" >> .env

### コンテナの起動

> $ make start

</details>

## その他

<!-- Makefileとしてまとめたコマンドを記載 -->
<details>
<summary>コマンド一覧</summary>

|      Commands      |              Description               |
| :----------------- | :------------------------------------- |
| make setup         | * 初回のみ実行                         |
| make build         | * コンテナの再構築                     |
| make install       | * コンテナ内にライブラリをインストール |
| make start         | * コンテナを起動                       |
| make start-native  | * ネイティブアプリ関連のコンテナを起動 |
| make start-admin   | * 管理者コンソール関連のコンテナを起動 |
| make start-api     | * API関連のコンテナを起動              |
| make start-swagger | * Swaggerのコンテナを起動              |
| make stop          | * コンテナの停止                       |
| make down          | * コンテナの削除                       |
| make logs          | * コンテナのログを取得                 |

</details>

<!-- docs配下のドキュメントをツリー型で記載 -->
<details>
<summary>各種ドキュメント</summary>

* [01_specification](./docs/01_specification/README.md)
* [02_design](./docs/02_design/README.md)
* [11_frontend](./docs/11_frontend/README.md)
  * [01_native](./docs/11_frontend/01_native/README.md)
  * [02_web](./docs/11_frontend/02_web/README.md)
    * [01_admin](./docs/11_frontend/02_web/01_admin/README.md)
* [12_backend](./docs/12_backend/README.md)
  * [01_design](./docs/12_backend/01_design/README.md)
  * [11_swagger](./docs/12_backend/11_swagger/README.md)
  * [12_protobuf](./docs/12_backend/12_protobuf/README.md)
  * [21_auth_api](./docs/12_backend/21_auth_api/README.md)
  * [22_user_api](./docs/12_backend/22_user_api/README.md)
  * [23_book_api](./docs/12_backend/23_book_api/README.md)
  * [24_store_api](./docs/12_backend/24_store_api/README.md)
  * [31_google_book_api](./docs/12_backend/31_google_book_api/README.md)
  * [32_stripe](./docs/12_backend/32_stripe/README.md)
* [13_database](./docs/13_database/README.md)
  * [01_user_db](./docs/13_database/01_user_db/README.md)
  * [02_book_db](./docs/13_database/02_book_db/README.md)
  * [03_store_db](./docs/13_database/03_store_db/README.md)
  * [11_auth_db](./docs/13_database/11_auth_db/README.md)
  * [12_message_db](./docs/13_database/12_message_db/README.md)
* [14_infrastructure](./docs/14_infrastructure/README.md)
  * [01_design](./docs/14_infrastructure/01_design/README.md)
  * [11_gcp](./docs/14_infrastructure/11_gcp/README.md)
  * [12_firebase](./docs/14_infrastructure/12_firebase/README.md)
  * [21_reverse-proxy](./docs/14_infrastructure/21_reverse-proxy/README.md)
  * [31_docker](./docs/14_infrastructure/31_docker/README.md)
  * [32_kubernetes](./docs/14_infrastructure/32_kubernetes/README.md)
  * [33_virtual_machine](./docs/14_infrastructure/33_virtual_machine/README.md)
  * [41_prometheus](./docs/14_infrastructure/41_prometheus/README.md)
  * [42_grafana](./docs/14_infrastructure/42_grafana/README.md)
  * [43_fluentd](./docs/14_infrastructure/43_fluentd/README.md)
  * [51_github-actions](./docs/14_infrastructure/51_github-actions/README.md)
  * [52_terraform](./docs/14_infrastructure/52_terraform/README.md)
* [99_other](./docs/99_other/README.md)
</details>
