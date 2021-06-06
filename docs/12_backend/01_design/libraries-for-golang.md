# Golang - ライブラリ関連

[root](./../../../README.md) 
/ [12_backend](./../README.md) 
/ [01_design](./README.md) 
/ [ディレクトリ関連](./libraries-for-golang.md) 

## 使用するライブラリ

### ライブラリ一覧

* バリデーション
  * [validator](https://github.com/go-playground/validator)

* データベース
  * MySQL
  * Firebase Firestore
  * ORM: [Gorm](https://gorm.io/ja_JP/)

* ストレージ
  * Firebase Firestore

* ログ
  * [zap](https://github.com/uber-go/zap)

* エラーハンドリング
  * [xerrors](https://github.com/golang/xerrors)

* 通信
  * [grpc](https://github.com/grpc/grpc-go)

* Liter
  * [golangci-lint](https://golangci-lint.run/)

* Test
  * 標準ライブラリ
  * [mock](https://github.com/golang/mock)

---

### バリデーション

* READMEに↓の記載があるので、v10以上のものを使用する
  > v9 has entered maintenance status as of 2019-12-24. Please make all new functionality PR's against master.

* 実装時の参考になりそうなサイト
  * [Go Docs - validator](https://pkg.go.dev/github.com/go-playground/validator)
  * [Qiita - go-playground/validator リクエストパラメータ向けValidationパターンまとめ](https://qiita.com/RunEagler/items/ad79fc860c3689797ccc)

### データベース関連

|   Name    |    自動生成の経路    | AutoMigration | Schemaからのリバース | Relation機能のライブラリ提供 |
| :-------- | :------------------- | :------------ | -------------------- | ---------------------------- |
| GORM      | Struct -> Schema生成 | ○             | -                    | ○                            |
| SQLBoiler | Schema -> Struct生成 | -             | ○                    | -                            |
| XORM      | Schema -> Struct生成 | ○             | -                    | -                            |

* Relation機能が提供されているGormを採用

* Gorm: v1とv2が存在するため、v2を使用する (**ネットのやってみた系は基本的にv1で実装されてる!!**)
  * [リリースノート](https://gorm.io/docs/v2_release_note.html)
  * [Techtouch Developers Blog - GORM v2 触ってみた Major Features 編](https://tech.techtouch.jp/entry/gorm-v2#GORM-v2-%E3%83%AA%E3%83%AA%E3%83%BC%E3%82%B9%E3%83%8E%E3%83%BC%E3%83%88)

* 実装時の参考になりそうなサイト
  * [Gorm](https://gorm.io/ja_JP/)

### ログ関連

|        library        | star  | latest commit |
| :-------------------- | ----: | :------------ |
| sirupsen/logrus       | 15.2k | last month    |
| uber-go/zap           | 10.1k | 2 months ago  |
| rs/zerolog            |  3.5k | last month    |
| inconshreveable/log15 |   966 | 6 months ago  |
| apex/log              |   948 | last month    |

(2020/06時点)

* logrus
  * v2予定なし
* zap -> **採用**
  * json形式

---

## 参考

* [GitHub - Logrus](https://github.com/sirupsen/logrus)
* [GitHub - mock](https://github.com/golang/mock)
* [GitHub - GORM](https://github.com/go-gorm/gorm)
* [GitHub - grpc](https://github.com/grpc/grpc-go)
* [GitHub - zap](https://github.com/uber-go/zap)
* [公式Docs - golangci-lint](https://golangci-lint.run/)
* [公式Docs - gorm](https://gorm.io/ja_JP/)
* [Go Docs - gorm](https://pkg.go.dev/gorm.io/gorm)
* [Go Docs - validator](https://pkg.go.dev/github.com/go-playground/validator)
* [Go Docs - xerrors](https://pkg.go.dev/golang.org/x/xerrors)
* [Go Docs - zap](https://pkg.go.dev/go.uber.org/zap)
* [Qiita - go-playground/validator リクエストパラメータ向けValidationパターンまとめ](https://qiita.com/RunEagler/items/ad79fc860c3689797ccc)
* [Future Tech Blog - Go言語のDBレイヤーライブラリの評価](https://future-architect.github.io/articles/20190926/)
* [onemuri.space - golang で構造化ロギング](https://onemuri.space/note/4xhm6jmts/)
* [Techtouch Developers Blog - GORM v2 触ってみた Major Features 編](https://tech.techtouch.jp/entry/gorm-v2#GORM-v2-%E3%83%AA%E3%83%AA%E3%83%BC%E3%82%B9%E3%83%8E%E3%83%BC%E3%83%88)
