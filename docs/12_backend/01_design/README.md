# バックエンド - 設計

[root](./../../../README.md) 
/ [12_backend](./../README.md) 
/ [01_design](./README.md)

# 各種ドキュメント

## 内部API

* 以下4つに分割
  * [認証用API](./../21_auth_api/README.md)
  * [ユーザ管理用API](./../22_user_api/README.md)
  * [書籍管理用API](./../23_book_api/README.md)
  * [ECサイト用API](./../24_store_api/README.md)

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
    * 楽天Book API
  * 支払い関連のAPI
    * PayPal
