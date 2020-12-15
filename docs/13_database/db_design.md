## DB設計 - 全体構成

[root](./../../README.md) 
/ [13_database](./db_design.md)

### 共通設計

* エンジン: InnoDB
* 文字コード: utf-8mb4

### users

|       Field       |   DataType   | PK/FK | NotNull | Unsigned | Other | Default |         Explanation          |               Validation               |
| :---------------- | :----------- | :---: | :-----: | :------: | :---: | :------ | :--------------------------- | :------------------------------------- |
| id                | BIGINT(20)   |  PK   |  true   |   true   |  AI   |         | ユーザーID                   |                                        |
| username          | VARCHAR(32)  |       |  true   |          |       |         | ニックネーム                 |                                        |
| gender            | INT(2)       |       |  true   |   true   |       | 0       | 性別                         |                                        |
| email             | VARCHAR(256) |       |  false  |          |  UQ   |         | メールアドレス               | Format: Email                          |
| thumbnail_url     | BLOB         |       |  false  |          |       |         | サムネイル                   | Format: base64                         |
| password_digest   | VARCHAR(256) |       |  true   |          |       |         | パスワード                   | Length: 6 <= n <= 32, Format: Password |
| self_introduction | VARCHAR(256) |       |  false  |          |       |         | 自己紹介                     |                                        |
| last_name         | VARCHAR(16)  |       |  false  |          |       |         | 名字（漢字)                  |                                        |
| first_name        | VARCHAR(16)  |       |  false  |          |       |         | 名前（漢字)                  |                                        |
| lastname_kana     | VARCHAR(32)  |       |  false  |          |       |         | 名前（かな)                  |                                        |
| firstname_kana    | VARCHAR(32)  |       |  false  |          |       |         | 名前（かな)                  |                                        |
| postal_code       | CHAR(8)      |       |  false  |          |       |         | 郵便番号                     | xxx-xxxx　ハイフン付き                 |
| prefectures       | VARCHAR(32)  |       |  false  |          |       |         | 都道府県                     |                                        |
| cities            | VARCHAR(32)  |       |  false  |          |       |         | 市町村                       |                                        |
| address_chrome    | VARCHAR(64)  |       |  false  |          |       |         | 番地・丁目                   |                                        |
| other_address     | VARCHAR(64)  |       |  false  |          |       |         | マンション・ビル名・部屋番号 |                                        |
| instance_id       | BLOB         |       |  false  |          |       |         | 端末ID                       | NNにするか検討中                       |
| created_at        | DATETIME     |       |  true   |          |       |         | 登録日時                     |                                        |
| updated_at        | DATETIME     |       |  true   |          |       |         | 更新日時                     |                                        |

### books

|   Field    |  DataType  | PK/FK | NotNull | Unsigned |     Other     | Default | Explanation |                   Validation                    |
| :--------- | :--------- | :---: | :-----: | :------: | :-----------: | :------ | :---------- | :---------------------------------------------- |
| id         | BIGINT(20) |  PK   |  true   |   true   |      AI       |         | ID          |                                                 |
| book_id    | BIGINT(20) |  FK   |  true   |   true   | UQ(Composite) |         | 書籍ID      |                                                 |
| user_id    | BIGINT(20) |  FK   |  true   |   true   | UQ(Composite) |         | ユーザーID  |                                                 |
| status     | INT(2)     |  FK   |  true   |   true   |               | 0       | 本の状態    | (読んだ / 読んでいる /積読 /手放したい /欲しい) |
| created_at | DATETIME   |       |  true   |          |               |         | 登録日時    |                                                 |
| updated_at | DATETIME   |       |  true   |          |               |         | 更新日時    |                                                 |

### follows

|    Field    |  DataType  | PK/FK | NotNull | Unsigned | Other |         Default          | Explanation | Validation |
| :---------- | :--------- | :---: | :-----: | :------: | :---: | :----------------------- | :---------- | ---------- |
| id          | BIGINT(20) |  PK   |  true   |   true   |  AI   | フォロー管理ID           |             |            |
| follower_id | BIGINT(20) |  FK   |  true   |   true   |       | フォローするユーザーID   |             |            |
| followed_id | BIGINT(20) |  FK   |  true   |   true   |       | フォローされるユーザーID |             |            |
| created_at  | DATETIME   |       |  true   |          |       | 登録日時                 |             |            |
| updated_at  | DATETIME   |       |  true   |          |       | 更新日時                 |             |            |

### followers

|    Field    |  DataType  | PK/FK | NotNull | Unsigned | Other |         Default          | Explanation | Validation |
| :---------- | :--------- | :---: | :-----: | :------: | :---: | :----------------------- | :---------- | ---------- |
| id          | BIGINT(20) |  PK   |  true   |   true   |  AI   | フォロワー管理ID         |             |            |
| follower_id | BIGINT(20) |  FK   |  true   |   true   |       | フォローするユーザーID   |             |            |
| followed_id | BIGINT(20) |  FK   |  true   |   true   |       | フォローされるユーザーID |             |            |
| created_at  | DATETIME   |       |  true   |          |       | 登録日時                 |             |            |
| updated_at  | DATETIME   |       |  true   |          |       | 更新日時                 |             |            |

### Sale

| Field | DataType | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :---- | :------- | :---: | :-----: | :------: | :---: | :------ | :---------- | ---------- |
