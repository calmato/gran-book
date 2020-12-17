# バックエンド - API ドメイン・URL設計

[root](./../../../README.md) 
/ [12_backend](./../README.md) 
/ [01_design](./README.md) 
/ [ドメイン・URL設計](./domain-url.md)

## 基本方針

* RestfulなAPIの設計にする

### 命名規則

* ケバブケース
* 小文字のみ
* ドメイン名はサブドメインを利用して分割

* e.g.) https://domain-name.com/v1/books/{id}?sort=createdAt

### メソッドの使い分け

| Methods | Description |
| :--- | :--- |
| GET | リソースの詳細・一覧取得 |
| POST | リソースの新規登録 |
| PATCH | リソースの更新 |
| PUT | リソース(画像)の置き換え |
| DELETE | リソースの削除 |
