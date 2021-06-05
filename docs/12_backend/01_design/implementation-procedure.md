# バックエンド - 実装手順メモ

[root](./../../../README.md) 
/ [12_backend](./../README.md) 
/ [01_design](./README.md) 
/ [実装手順メモ](./implementation-procedure.md)

## 実装概要

1. データベースの作成/更新
2. Protobufでマイクロサービス間の通信を定義
3. マイクロサービスの実装
4. SwaggerでRest APIの仕様を定義
5. BFF(API Gateway)の実装

## マイクロサービスの実装

1. Domain層の実装
  1. エンティティを作成
    * internal/domain/<domain>/entity.go
  2. ドメインサービスのインターフェースを作成
    * internal/domain/<domain>/service.go
    * internal/domain/<domain>/repository.go
    * internal/domain/<domain>/validation.go

2. Infrastructure層の実装
  1. (1-2)で作成したインターフェースの詳細を実装 (Service以外)
    * internal/infrastructure/repository/<domain>.go
    * internal/infrastructure/validation/<domain>.go
  2. (2-1)までに作成したインターフェースを利用し、ドメインサービスの詳細を実装
    * internal/infrastructure/service/<domain>.go

3. Application層の実装
  1. 外部から受け取った値をAPI内部で使えるようにするために詰め替える用の Input オブジェクトの作成
    * internal/application/input/<domain>.go
  2. (必要あれば、外部へレスポンスとして返す用の Output オブジェクトも作成)
    * internal/application/output/<domain>.go
  3. リクエスト値のバリデーション用の リクエストバリデーション の実装
    * internal/application/validation/<domain>.go
    * (-> リクエスト値の制約 (バリデーション) は、Inputオブジェクトにタグとして記述する)
  4. ユースケース (ユーザを登録する...みたいな) を表現するためのインターフェースを作成
    * internal/application/<domain>.go
    * (-> ここで定義するメソッドは、基本的にはInterface層のメソッドと1対1対応するようになるはず)

4. Interface層の実装
  1. Protobufで生成されたコードを参考に、外部APIからのリクエスト/レスポンス処理の実装
    * internal/interface/grpc/v1/<domain>.go

## BFF(API Gateway)の実装

1. システム内で使用するオブジェクトの型定義
  1. Input/Output (BFF <-> マイクロサービス間の型定義)
    * src/types/input/index.ts
    * src/types/input/<domain>.ts
      * -> APIへ投げるリクエストを定義 (.protoのリクエストで書いた値と同じになるように...)
    * src/types/output/index.ts
    * src/types/output/<domain>.ts
      * -> APIから受け取るレスポンスを定義 (.protoのレスポンスで書いた値と同じになるように...)
  2. Request/Response (クライアント <-> BFF間の型定義)
    * src/types/request/index.ts
    * src/types/request/<domain>.ts
      * -> クライアントから受け取るリクエストを定義 (swaggerのリクエストで書いた値と同じになるように...)
    * src/types/response/index.ts
    * src/types/response/<domain>.ts
      * -> クライアントへ返すレスポンスを定義 (swaggerのレスポンスで書いた値と同じになるように...)

2. マイクロサービスへ接続する処理を実装
  * src/api/index.ts
  * src/api/<domain>.ts

3. クライアントから受けたリクエストの処理を実装
  * src/routes/index.ts
  * src/routes/v1/<domain>.ts
  * src/index.ts

---

## 参考

* [ドメインサービスとアプリケーションサービスの違い](https://codezine.jp/article/detail/10318)
