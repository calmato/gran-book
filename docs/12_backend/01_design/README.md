# バックエンド - 設計

[root](./../../../README.md) 
/ [12_backend](./../README.md) 
/ [01_design](./README.md)

# 各種ドキュメント

* [API設計 - 実装手順](./implementation-procedure.md)
* [API設計 - ドメイン・URL設計](./domain-url.md)
* [API設計 - リクエスト・レスポンス設計](./request-response.md)
* [Golang - ディレクトリ構成](./directories-for-golang.md)
* [Node.js - ディレクトリ構成](./directories-for-node.md)

![全体構成](./images/architecture.jpeg)

## 内部API

* APIとしては以下に分割
  * [認証用API](./../31_auth_api/README.md)
  * [ユーザ管理用API](./../32_user_api/README.md)
  * [書籍管理用API](./../33_book_api/README.md)
  * [ECサイト用API](./../34_store_api/README.md)
  * [サポート用API](./../35_information_api/README.md)

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

---

## 参考

* TypeScript
  * [公式Docs - TypeScript](typescriptlang.org)
  * [公式Docs - Express](https://expressjs.com/ja/)
  * [MDN Web Docs - Express](https://developer.mozilla.org/ja/docs/Learn/Server-side/Express_Nodejs/Introduction)

* Golang
  * [公式Docs - Golang](https://golang.org/)
  * [A Tour of Go](https://go-tour-jp.appspot.com/welcome/1) **☆**
  * [Qiita - 他言語プログラマがgolangの基本を押さえる為のまとめ](https://qiita.com/tfrcm/items/e2a3d7ce7ab8868e37f7)

* gRPC
  * [公式Docs - gRPC](https://grpc.io/)
  * [Microsoft Docs - gRPC](https://docs.microsoft.com/ja-jp/dotnet/architecture/cloud-native/grpc)
  * [Google Developer - Protobuf](https://developers.google.com/protocol-buffers) **☆**
  * [Wantedly Engineer Blog - gRPC Internal - gRPC の設計と内部実装から見えてくる世界]()
  * [Qiita - gRPCって何？](https://qiita.com/oohira/items/63b5ccb2bf1a913659d6)
  * [Qiita - goでgRPCの4つの通信方式やってみた](https://qiita.com/tomo0/items/310d8ffe82749719e029) **☆**

* クリーンアーキテクチャ
  * [Qiita - 図解クリーンアーキテクチャ](https://qiita.com/kz_12/items/bc79102247b86626fc72)
  * [Recruit Tech Blog - Goのpackage構成と開発のベタープラクティス](https://engineer.recruit-lifestyle.co.jp/techblog/2018-03-16-go-ddd/)

* Swagger
  * [公式Docs - Swagger](https://swagger.io/) **☆**
  * [Qiita - SwaggerでRESTful APIの管理を楽にする](https://qiita.com/disc99/items/37228f5d687ad2969aa2)
