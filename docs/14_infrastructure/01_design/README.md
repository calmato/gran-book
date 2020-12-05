# インフラストラクチャ - 設計

[root](./../../../README.md) 
/ [14_infrastructure](./../README.md) 
/ [01_design](./README.md)

# 各種ドキュメント

## インフラ構成

* [GCP - インフラ構成](./../11_gcp/README.md)

## アーキテクチャ

### Frontend

* ネイティブアプリ
* 管理用Webアプリ

### Backend

* 認証用API (Firebase Authentication)
* ユーザ管理用API
* 書籍管理用API
* ECサイト用API

### Database

* 認証用DB (Firebase Authentication)
* ユーザー管理用DB (MySQL)
* 書籍管理用DB (MySQL)
* ECサイト用DB (MySQL)
* メッセージ管理用DB (NoSQL)

### Infrastructure

* L7 ロードバランサー
* サムネイル用ストレージ (Object Storage)

### Other

* 書籍検索API (Google Books API)
