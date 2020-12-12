# Golang - ディレクトリ構成

[root](./../../../README.md) 
/ [12_backend](./../README.md) 
/ [01_design](./README.md) 
/ [ディレクトリ構成](./directories-for-golang.md) 

## ディレクトリ構成

<pre>
.
├── bin
├── cmd
├── config
├── internal
│   ├── application
│   │   ├── request
│   │   ├── response
│   │   └── validation
│   ├── domain
│   ├── infrastructure
│   │   ├── api
│   │   ├── notification
│   │   ├── repository
│   │   ├── service
│   │   ├── storage
│   │   └── validation
│   └── interface
│       └── handler
├── lib
├── mock
└── registry
</pre>

|       ディレクトリ名        |                                       説明                                       |
| :-------------------------- | :------------------------------------------------------------------------------- |
| bin                         | システムで使用するバイナリファイルを配置                                         |
| cmd                         | 実行用ファイルを配置                                                             |
| config                      | システムの設定ファイルを配置                                                     |
| internal                    | アプリケーションの実装用ディレクトリ, 他のアプリで共有するものはここに配置しない |
| application                 | クリーンアーキテクチャのアプリケーション層。ドメイン層のみ参照可能               |
| application/request         | リクエストの構造体を定義                                                         |
| application/response        | レスポンスの構造体を定義                                                         |
| application/validation      | リクエスト値のバリデーションチェック。エラーがある場合400系を返す                |
| domain                      | クリーンアーキテクチャのドメイン層。ドメインごとにディレクトリを作成             |
| infrastructure              | クリーンアーキテクチャのインフラ層。ドメイン層の詳細を実装等                     |
| infrastructure/api          | 外部APIへのリクエスト処理等                                                      |
| infrastructure/notification | アプリへのプッシュ通知等の実装                                                   |
| infrastructure/repository   | データストアの操作処理を実装                                                     |
| infrastructure/service      | ドメイン層のサービスの詳細を実装                                                 |
| infrastructure/storage      | ストレージへの書き込み/参照操作を実装                                            |
| infrastructure/validation   | ドメイン層バリデーションチェック。エラーがある場合、基本的に500系を返す          |
| interface                   | クリーンアーキテクチャのインターフェース層。インフラ層以外を参照可能             |
| interface/handler           | リクエストの取得とレスポンスの生成を行う                                         |
| lib                         | システム内で使用する共通のライブラリを実装。外部ライブラリの設定等もここでする   |
| mock                        | テスト用のモックを配置                                                           |
| registry                    | DI 依存性の注入                                                                  |
