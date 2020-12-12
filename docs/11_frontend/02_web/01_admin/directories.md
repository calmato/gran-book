# Nuxt.js - ディレクトリ構成

[root](./../../../../README.md)
/ [11_frontend](./../../README.md)
/ [02_web](./../README.md)
/ [01_admin](./README.md) 
/ [ディレクトリ構成](./directories.md)

## ディレクトリ構成

<pre>
.
├── app
│   ├── assets
│   ├── components
│   │   ├── atoms
│   │   ├── molecules
│   │   ├── organisms
│   │   └── templates
│   ├── content
│   ├── layouts
│   ├── middleware
│   ├── modules
│   ├── pages
│   ├── plugins
│   ├── static
│   ├── store
│   └── types
│       ├── forms
│       ├── requests
│       ├── responses
│       └── store
└── spec
    ├── components
    │   ├── atoms
    │   ├── molecules
    │   ├── organisms
    │   └── templates
    ├── helpers
    │   ├── responses
    │   └── store
    ├── pages
    └── store
</pre>

|    ディレクトリ名    |                                                        説明                                                        |
| :------------------- | :----------------------------------------------------------------------------------------------------------------- |
| assets               | ディレクトリにはスタイルや画像、フォントなどコンパイルされていないアセットなどを配置                               |
| components           | UI設計用のパーツ                                                                                                   |
| components/atoms     | 最も小さい粒度の要素                                                                                               |
| components/molecules | 2つ以上のAtomsを組み合わせたシンプルなUI要素                                                                       |
| components/organisms | 切り離して単体でも機能する分子の集まり                                                                             |
| components/templates | Organismsを組み合わせて画面の作成                                                                                  |
| content              | Markdown、JSON、YAML、CSVファイルをフェッチ できるモジュールを提供                                                 |
| layouts              | Header,Footerなどの共通のUI要素を作成, エラ〜ページもここに配置                                                    |
| middleware           | 認証処理等のページがレンダリングされる前(SSR処理などが行われる前)に行う処理を記述                                  |
| modules              |                                                                                                                    |
| pages                | ルーティング定義とページへの要素追加を行う                                                                         |
| plugins              | 外部ライブラリの処理等の、各ページで共通化された処理を記述                                                         |
| static               | 直接サーバのルートに配置され、名前を保持しなければいけないファイル、もしくは変更されない可能性の高いファイルを配置 |
| store                | アプリの状態管理                                                                                                   |
| types                | アプリ内で使用する関数,変数の型定義                                                                                |
| types/forms          | コンポーネント内で使用するFormの型定義                                                                             |
| types/request        | APIへのリクエストの型定義                                                                                          |
| types/response       | APIからのレスポンスの型定義                                                                                        |
| types/store          | 状態管理用のVuexの型定義                                                                                           |
| spec                 | 単体テストを配置                                                                                                   |
| spec/helpers         | 単体テストで使用するモック等の定義                                                                                 |
