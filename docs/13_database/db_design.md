## DB設計 - 全体構成

[root](./../../README.md) 
/ [13_database](./db_design.md)

### 共通設計

* エンジン: InnoDB
* 文字コード: utf8mb4

### users

|       Field       |   DataType   | PK/FK | NotNull | Unsigned | Other | Default |            Explanation            |       Validation       |
| :---------------- | :----------- | :---: | :-----: | :------: | :---: | :------ | :-------------------------------- | :--------------------- |
| id                | VARCHAR(36)  |  PK   |  true   |          |       |         | ユーザーID                        |                        |
| username          | VARCHAR(32)  |       |  true   |          |       | ''      | ニックネーム                      |                        |
| gender            | TYNYINT(2)   |       |  true   |   true   |       | '0'     | 性別                              |                        |
| email             | VARCHAR(256) |       |  false  |          |  UQ   | NULL    | メールアドレス                    | Format: Email          |
| role              | TYNYINT(2)   |       |  true   |   true   |       | '0'     | 0: 利用者 / 1: 管理者 / 2: 開発者 |                        |
| thumbnail_url     | TEXT(8192)   |       |  false  |          |       | NULL    | サムネイル                        | Format: base64         |
| self_introduction | VARCHAR(256) |       |  false  |          |       | NULL    | 自己紹介                          |                        |
| last_name         | VARCHAR(16)  |       |  false  |          |       | NULL    | 名字（漢字)                       |                        |
| first_name        | VARCHAR(16)  |       |  false  |          |       | NULL    | 名前（漢字)                       |                        |
| last_name_kana    | VARCHAR(32)  |       |  false  |          |       | NULL    | 名前（かな)                       |                        |
| fir_stname_kana   | VARCHAR(32)  |       |  false  |          |       | NULL    | 名前（かな)                       |                        |
| postal_code       | VARCHAR(16)  |       |  false  |          |       | NULL    | 郵便番号                          | xxx-xxxx　ハイフンあり |
| prefecture        | VARCHAR(32)  |       |  false  |          |       | NULL    | 都道府県                          |                        |
| city              | VARCHAR(32)  |       |  false  |          |       | NULL    | 市町村                            |                        |
| address_line1     | VARCHAR(64)  |       |  false  |          |       | NULL    | 番地・丁目                        |                        |
| address_line2     | VARCHAR(64)  |       |  false  |          |       | NULL    | マンション・ビル名・部屋番号      |                        |
| phoneNumber       | VARCHAR(16)  |       |  false  |          |       | NULL    | 電話番号                          |                        |
| instance_id       | VARCHAR(256) |       |  false  |          |       | NULL    | 端末ID                            | NNにするか検討中       |
| created_at        | DATETIME     |       |  true   |          |       |         | 登録日時                          |                        |
| updated_at        | DATETIME     |       |  true   |          |       |         | 更新日時                          |                        |

### books

|     Field     |  DataType   | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :------------ | :---------- | :---: | :-----: | :------: | :---: | :------ | :---------- | :--------- |
| id            | BIGINT(20)  |  PK   |  true   |   true   |  AI   |         | ID          |            |
| publisher_id  | BIGINT(20)  |  FK   |  true   |   true   |  UQ   |         | 出版社ID    |            |
| title         | VARCHER(32) |       |  true   |          |  UQ   | ''      | タイトル    |            |
| description   | TEXT(1000)  |       |  true   |          |       | NULL    | 説明        |            |
| isbn          | VARCHAR(16) |       |  true   |          |       | ''      | ID          |            |
| thumbnail_url | TEXT(8192)  |       |  false  |          |       | NULL    | サムネイル  |            |
| published_at  | DATE        |       |  true   |          |       |         | 発売日      |            |
| created_at    | DATETIME    |       |  true   |          |       |         | 登録日時    |            |
| updated_at    | DATETIME    |       |  true   |          |       |         | 更新日時    |            |

### authors

|   Field    |  DataType   | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :--------- | :---------- | :---: | :-----: | :------: | :---: | :------ | :---------- | :--------- |
| id         | BIGINT(20)  |  PK   |  true   |   true   |  AI   |         | ID          |            |
| name       | VARCHAR(32) |       |  true   |          |  UQ   | ''      | 著者        |            |
| created_at | DATETIME    |       |  true   |          |       |         | 登録日時    |            |
| updated_at | DATETIME    |       |  true   |          |       |         | 更新日時    |            |

### categories

|   Field    |  DataType   | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :--------- | :---------- | :---: | :-----: | :------: | :---: | :------ | :---------- | :--------- |
| id         | BIGINT(20)  |  PK   |  true   |   true   |  AI   |         | ID          |            |
| name       | VARCHAR(32) |       |  true   |          |  UQ   | ''      | カテゴリー  |            |
| created_at | DATETIME    |       |  true   |          |       |         | 登録日時    |            |
| updated_at | DATETIME    |       |  true   |          |       |         | 更新日時    |            |

### publishers

|   Field    |  DataType   | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :--------- | :---------- | :---: | :-----: | :------: | :---: | :------ | :---------- | :--------- |
| id         | BIGINT(20)  |  PK   |  true   |   true   |  AI   |         | ID          |            |
| name       | VARCHAR(32) |       |  true   |          |  UQ   | ''      | 出版社      |            |
| created_at | DATETIME    |       |  true   |          |       |         | 登録日時    |            |
| updated_at | DATETIME    |       |  true   |          |       |         | 更新日時    |            |

### authors_books

|   Field    |  DataType  | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :--------- | :--------- | :---: | :-----: | :------: | :---: | :------ | :---------- | :--------- |
| id         | BIGINT(20) |  PK   |  true   |   true   |  AI   |         | ID          |            |
| book_id    | BIGINT(20) |  FK   |  true   |   true   |  UQ   |         | 本ID        |            |
| author_id  | BIGINT(20) |  FK   |  false  |   true   |  UQ   | NULL    | 著書ID      |            |
| created_at | DATETIME   |       |  true   |          |       |         | 登録日時    |            |
| updated_at | DATETIME   |       |  true   |          |       |         | 更新日時    |            |

### books_users

|   Field    |  DataType   | PK/FK | NotNull | Unsigned | Other | Default |                                 Explanation                                  | Validation |
| :--------- | :---------- | :---: | :-----: | :------: | :---: | :------ | :--------------------------------------------------------------------------- | :--------- |
| id         | BIGINT(20)  |  PK   |  true   |   true   |  AI   |         | ID                                                                           |            |
| user_id    | BIGINT(20)  |  FK   |  false  |   true   |  UQ   |         | ユーザーID                                                                   |            |
| book_id    | VARCHAR(36) |  FK   |  true   |   true   |  UQ   |         | 本ID                                                                         |            |
| status     | TYNYINT(4)  |       |  true   |   true   |       | '0'     | 本のステータス(0:未読, 1:読んだ, 2:読んでる, 3:積読, 4:手放したい, 5:欲しい) |            |
| impression | TEXT(1000)  |       |  false  |          |       | NULL    | 感想                                                                         |            |
| created_at | DATETIME    |       |  true   |          |       |         | 登録日時                                                                     |            |
| updated_at | DATETIME    |       |  true   |          |       |         | 更新日時                                                                     |            |

### books_categories

|    Field    |  DataType  | PK/FK | NotNull | Unsigned | Other | Default | Explanation  | Validation |
| :---------- | :--------- | :---: | :-----: | :------: | :---: | :------ | :----------- | :--------- |
| id          | BIGINT(20) |  PK   |  true   |   true   |  AI   |         | ID           |            |
| category_id | BIGINT(20) |  FK   |  false  |   true   |  UQ   | NULL    | カテゴリーID |            |
| book_id     | BIGINT(20) |  FK   |  true   |   true   |  UQ   |         | 本ID         |            |
| created_at  | DATETIME   |       |  true   |          |       |         | 登録日時     |            |
| updated_at  | DATETIME   |       |  true   |          |       |         | 更新日時     |            |

### follows

|    Field    |  DataType   | PK/FK | NotNull | Unsigned | Other | Default |       Explanation        | Validation |
| :---------- | :---------- | :---: | :-----: | :------: | :---: | :------ | :----------------------- | ---------- |
| id          | VARCHAR(36) |  PK   |  true   |          |  AI   |         | ID                       |            |
| follow_id   | VARCHAR(36) |  FK   |  true   |          |  UQ   |         | フォローするユーザーID   |            |
| follower_id | VARCHAR(36) |  FK   |  true   |          |  UQ   |         | フォローされるユーザーID |            |
| created_at  | DATETIME    |       |  true   |          |       |         | 登録日時                 |            |
| updated_at  | DATETIME    |       |  true   |          |       |         | 更新日時                 |            |

### Sale

| Field | DataType | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :---- | :------- | :---: | :-----: | :------: | :---: | :------ | :---------- | ---------- |
