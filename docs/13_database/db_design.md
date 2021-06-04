## DB設計 - 全体構成

[root](./../../README.md) 
/ [13_database](./db_design.md)

### 共通設計

* エンジン: InnoDB
* 文字コード: utf8mb4

## DB設計 - テーブル設計

### User DB

#### users

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
| instance_id       | VARCHAR(256) |       |  false  |          |       | NULL    | 端末ID                            |                        |
| created_at        | DATETIME     |       |  true   |          |       |         | 登録日時                          |                        |
| updated_at        | DATETIME     |       |  true   |          |       |         | 更新日時                          |                        |
| deleted_at        | DATETIME     |       |  false  |          |       | NULL    | 削除日時                          |                        |

#### relationships

|    Field    |  DataType   | PK/FK | NotNull | Unsigned | Other | Default |       Explanation        | Validation |
| :---------- | :---------- | :---: | :-----: | :------: | :---: | :------ | :----------------------- | ---------- |
| id          | VARCHAR(36) |  PK   |  true   |          |  AI   |         | ID                       |            |
| follow_id   | VARCHAR(36) |  FK   |  true   |          |  UQ   |         | フォローするユーザーID   |            |
| follower_id | VARCHAR(36) |  FK   |  true   |          |  UQ   |         | フォローされるユーザーID |            |
| created_at  | DATETIME    |       |  true   |          |       |         | 登録日時                 |            |
| updated_at  | DATETIME    |       |  true   |          |       |         | 更新日時                 |            |

### Book DB

#### authors

|   Field    |  DataType   | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :--------- | :---------- | :---: | :-----: | :------: | :---: | :------ | :---------- | :--------- |
| id         | BIGINT(20)  |  PK   |  true   |   true   |  AI   |         | ID          |            |
| name       | VARCHAR(32) |       |  true   |          |  UQ   | ''      | 著者        |            |
| name_kana  | VARCHAR(64) |       |  true   |          |       | ''      | 著者 (かな) |            |
| created_at | DATETIME    |       |  true   |          |       |         | 登録日時    |            |
| updated_at | DATETIME    |       |  true   |          |       |         | 更新日時    |            |

#### books

|      Field       |   DataType   | PK/FK | NotNull | Unsigned | Other | Default |   Explanation   | Validation |
| :--------------- | :----------- | :---: | :-----: | :------: | :---: | :------ | :-------------- | :--------- |
| id               | BIGINT(20)   |  PK   |  true   |   true   |  AI   |         | ID              |            |
| title            | VARCHER(64)  |       |  true   |          |  UQ   | ''      | タイトル        |            |
| title_kana       | VARCHER(128) |       |  true   |          |       | ''      | タイトル (かな) |            |
| description      | TEXT(2000)   |       |  false  |          |       | NULL    | 説明            |            |
| isbn             | VARCHAR(13)  |       |  true   |          |       | ''      | ID              |            |
| publisher        | VARCHAR(64)  |       |  true   |          |       | ''      | 出版社          |            |
| published_on     | VARCHAR(16)  |       |  false  |          |       | ''      | 出版日          |            |
| thumbnail_url    | TEXT(8192)   |       |  false  |          |       | ''      | サムネイル      |            |
| rakuten_url      | TEXT(8192)   |       |  false  |          |       | ''      | 楽天 URL        |            |
| rakuten_size     | VARCHAR(64)) |       |  false  |          |       | ''      | 楽天 書籍サイズ |            |
| rakuten_genre_id | VARCHAR(64)) |       |  false  |          |       | '000'   | 楽天 カテゴリID |            |
| created_at       | DATETIME     |       |  true   |          |       |         | 登録日時        |            |
| updated_at       | DATETIME     |       |  true   |          |       |         | 更新日時        |            |

#### authors_books

|   Field    |  DataType  | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :--------- | :--------- | :---: | :-----: | :------: | :---: | :------ | :---------- | :--------- |
| id         | BIGINT(20) |  PK   |  true   |   true   |  AI   |         | ID          |            |
| book_id    | BIGINT(20) |  FK   |  true   |   true   |  UQ   |         | 書籍ID      |            |
| author_id  | BIGINT(20) |  FK   |  true   |   true   |  UQ   |         | 著書ID      |            |
| created_at | DATETIME   |       |  true   |          |       |         | 登録日時    |            |
| updated_at | DATETIME   |       |  true   |          |       |         | 更新日時    |            |

#### boolshelves

|   Field    |  DataType   | PK/FK | NotNull | Unsigned | Other | Default |  Explanation   | Validation |
| :--------- | :---------- | :---: | :-----: | :------: | :---: | :------ | :------------- | :--------- |
| id         | BIGINT(20)  |  PK   |  true   |   true   |  AI   |         | ID             |            |
| book_id    | BIGINT(20)  |  FK   |  true   |   true   |  UQ   |         | 書籍ID         |            |
| user_id    | VARCHAR(36) |  FK   |  true   |          |  UQ   |         | ユーザID       |            |
| status     | TINYINT(4)  |       |  true   |   true   |       | '0'     | 読書ステータス |            |
| read_on    | DATE        |       |  false  |          |       |         | 読んだ日       |            |
| created_at | DATETIME    |       |  true   |          |       |         | 登録日時       |            |
| updated_at | DATETIME    |       |  true   |          |       |         | 更新日時       |            |

#### reviews

|   Field    |  DataType   | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :--------- | :---------- | :---: | :-----: | :------: | :---: | :------ | :---------- | :--------- |
| id         | BIGINT(20)  |  PK   |  true   |   true   |  AI   |         | ID          |            |
| book_id    | BIGINT(20)  |  FK   |  true   |   true   |  UQ   |         | 書籍ID      |            |
| user_id    | VARCHAR(36) |  FK   |  true   |          |  UQ   |         | ユーザID    |            |
| score      | TINYINT(4)  |       |  true   |   true   |       | '0'     | 評価        |            |
| impression | TEXT(2000)  |       |  false  |          |       |         | 感想        |            |
| created_at | DATETIME    |       |  true   |          |       |         | 登録日時    |            |
| updated_at | DATETIME    |       |  true   |          |       |         | 更新日時    |            |

### Information DB

#### inquiries

|    Field    |   DataType   | PK/FK | NotNull | Unsigned | Other | Default |   Explanation    |  Validation   |
| :---------- | :----------- | :---: | :-----: | :------: | :---: | :------ | :--------------- | :------------ |
| id          | BIGINT(20)   |  PK   |  true   |   true   |  AI   |         | ID               |               |
| sender_id   | VARCHAR(36)  |  FK   |  true   |          |       |         | 送信者ID         |               |
| admin_id    | VARCHAR(36)  |  FK   |  true   |          |       |         | 対応者ID         |               |
| subject     | VARCHAR(45)  |       |  true   |          |       | ''      | お問い合わせ種別 |               |
| description | TEXT(2000)   |       |  true   |          |       | ''      | 説明             |               |
| email       | VARCHAR(256) |       |  true   |          |       | ''      | メールアドレス   | Format: Email |
| is_replied  | TINYINT(1)   |       |  true   |   true   |       | 0       | 対応フラグ       |               |
| created_at  | DATETIME     |       |  true   |          |       |         | 登録日時         |               |
| updated_at  | DATETIME     |       |  true   |          |       |         | 更新日時         |               |

#### categories

|   Field    |  DataType   | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :--------- | :---------- | :---: | :-----: | :------: | :---: | :------ | :---------- | :--------- |
| id         | BIGINT(20)  |  PK   |  true   |   true   |  AI   |         | ID          |            |
| name       | VARCHAR(64) |       |  true   |          |  UQ   | ''      | カテゴリー  |            |
| created_at | DATETIME    |       |  true   |          |       |         | 登録日時    |            |
| updated_at | DATETIME    |       |  true   |          |       |         | 更新日時    |            |

#### notifications

|    Field    |  DataType   | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :---------- | :---------- | :---: | :-----: | :------: | :---: | :------ | :---------- | :--------- |
| id          | BIGINT(20)  |  PK   |  true   |   true   |  AI   |         | ID          |            |
| author_id   | VARCHAR(36) |  FK   |  false  |          |       |         | 作成者ID    |            |
| editor_id   | VARCHAR(36) |  FK   |  false  |          |       |         | 編集者ID    |            |
| category_id | BIGINT(20)  |  FK   |  false  |   true   |       |         | カテゴリID  |            |
| title       | VARCHAR(64) |       |  true   |          |       | ''      | タイトル    |            |
| description | TEXT(2000)  |       |  true   |          |       | ''      | 説明        |            |
| importance  | TINYINT(4)  |       |  true   |          |       | 0       | 重要度      |            |
| created_at  | DATETIME    |       |  true   |          |       |         | 登録日時    |            |
| updated_at  | DATETIME    |       |  true   |          |       |         | 更新日時    |            |

### Store DB

#### Sale

| Field | DataType | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :---- | :------- | :---: | :-----: | :------: | :---: | :------ | :---------- | ---------- |
