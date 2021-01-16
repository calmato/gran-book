# インフラストラクチャ - やったことメモ

[root](./../../../README.md) 
/ [14_infrastructure](./../README.md) 
/ [01_design](./README.md) 
/ [note](./note.md)

## GCPのプロジェクト作成

1. GCPのプロジェクト作成
2. Compute Engine APIの有効化
3. Cloud SQL Admin APIの有効化

## Firebaseのプロジェクト作成

1. Firebaseのプロジェクト作成
2. Authenticationの開始
3. Cloud Firestoreの開始
4. Firebase Admin SDKの作成
5. マイアプリ(Web App)の追加

## サービスアカウントの作成

1. コンテナ実行用サービスアカウントの作成
  * 付与するロール
    * Cloud SQL クライアント
    * Firebase Admin SDK 管理者サービス エージェント
    * Firebase Authentication 管理者
2. GitHub Actions実行用サービスアカウントの作成
  * 付与するロール
    * 編集者
3. Terraform実行用サービスアカウントの作成
  * 付与するロール
    * Cloud SQL 管理者
    * Compute インスタンス管理者（v1）
    * Compute ネットワーク管理者
    * Kubernetes Engine 管理者
    * サービス アカウント ユーザー
    * ストレージ管理者

## Cloud SQLの設定

1. Cloud SQLのダッシュボードよりユーザ作成
2. Cloud SQL Proxyのインストール
3. 以下コマンドを実行
  > $ ./cloud_sql_proxy \
  > -instances=<INSTANCE_CONNECTION_NAME>=tcp:3326 \
  > -credential_file=<PATH_TO_KEY_FILE>
4. MySQLクライアントを使用して接続して、SQLの実行
  > $ mysql -h 127.0.0.1 -P 3326 -u<USERNAME> -p<PASSWORD> < infra/mysql/sql/*.sql

---

## 参考

* [Cloud SQL - MySQL クライアントの接続](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy#connecting-client)
