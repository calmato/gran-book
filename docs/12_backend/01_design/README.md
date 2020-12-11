# バックエンド - 設計

[root](./../../../README.md) 
/ [12_backend](./../README.md) 
/ [01_design](./README.md)

# 各種ドキュメント

## 内部API

* 以下4つに分割
  * 認証用API (Firebase Authentication)
  * [ユーザ管理用API](./../21_user_api/README.md)
  * [書籍管理用API](./../22_book_api/README.md)
  * [ECサイト用API](./../23_store_api/README.md)

### 認証用API

* 使用するSaaS
  * Firebase Authentication

* 対応させる認証方法
  * Apple
  * Facebook
  * Google
  * Twitter

### ユーザ管理用API

* 使用言語: Golang

### 書籍管理用API

* 使用言語: Golang

### ECサイト用API

* 使用言語: Golang

---

## 外部API

* 以下のサービスを採用
  * 書籍検索関連のAPI
    * [Google Books API](./../31_google_books_api/README.md)
  * 支払い関連のAPI
    * [Stripe](./../32_stripe/README.md)

* その他検討対象のサービス
  * 書籍検索関連のAPI
    * Amazon Advertising API
  * 支払い関連のAPI
    * PayPal
