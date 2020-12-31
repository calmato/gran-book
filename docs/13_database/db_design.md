## DB設計 - 全体構成

[root](./../../README.md) 
/ [13_database](./db_design.md)

### 共通設計

* エンジン: InnoDB
* 文字コード: utf-8mb4

### users

|       Field       |   DataType   | PK/FK | NotNull | Unsigned | Other | Default |         Explanation          |               Validation               |
| :---------------- | :----------- | :---: | :-----: | :------: | :---: | :------ | :--------------------------- | :------------------------------------- |
| id                | VARCHAR(32)  |  PK   |  true   |         |  UQ   |         | ユーザーID                   |                                        |
| role              | TYNYINT(2)   |       |  true   |  true     |       |         | 0: 利用者 / 1: 開発者 / 2: 管理者　|                                        |
| username          | VARCHAR(32)  |       |  true   |          |       |         | ニックネーム                 |                                        |
| gender            | TYNYINT(2)       |       |  true   |   true   |       | 0       | 性別                         |                                        |
| email             | VARCHAR(256) |       |  false  |          |  UQ   |         | メールアドレス               | Format: Email                          |
| thumbnail_url     | VARCHAR(256) |       |  false  |          |       |         | サムネイル                   | Format: base64                         |
| self_introduction | VARCHAR(256) |       |  false  |          |       |         | 自己紹介                     |                                        |
| last_name         | VARCHAR(16)  |       |  false  |          |       |         | 名字（漢字)                  |                                        |
| first_name        | VARCHAR(16)  |       |  false  |          |       |         | 名前（漢字)                  |                                        |
| lastname_kana     | VARCHAR(32)  |       |  false  |          |       |         | 名前（かな)                  |                                        |
| firstname_kana    | VARCHAR(32)  |       |  false  |          |       |         | 名前（かな)                  |                                        |
| postal_code       | VARCHAR(16)  |       |  false  |          |       |         | 郵便番号                     | xxx-xxxx　ハイフン無し                 |
| prefectures       | VARCHAR(32)  |       |  false  |          |       |         | 都道府県                     |                                        |
| cities            | VARCHAR(32)  |       |  false  |          |       |         | 市町村                       |                                        |
| addressLine1      | VARCHAR(64)  |       |  false  |          |       |         | 番地・丁目                   |                                        |
| addressLine2      | VARCHAR(64)  |       |  false  |          |       |         | マンション・ビル名・部屋番号        |                                        |
| phoneNumber       | VARCHAR(16)  |       |  false  |          |       |         | 電話番号                     |                       |
| instance_id       | VARCHAR(256) |       |  false  |          |       |         | 端末ID                       | NNにするか検討中                       |
| created_at        | DATETIME     |       |  true   |          |       |         | 登録日時                     |                                        |
| updated_at        | DATETIME     |       |  true   |          |       |         | 更新日時                     |                                        |

### books

|   Field       |  DataType   | PK/FK | NotNull | Unsigned |     Other     | Default | Explanation  |                   Validation                    |
| :---------    | :---------  | :---: | :-----: | :------: | :-----------: | :------ | :----------  | :---------------------------------------------- |
| id            | VARCHAR(32) |  PK   |  true   |          |      UQ       |         | ID           |                                                 |
| publisher_id  | VARCHAR(32) |  FK   |  true   |          |      UQ       |         | 出版社ID      |                                                 |
| title         | VARCHER(32) |       |  true   |          |               |         | タイトル      |                                                 |
| description   | TEXT        |       |  true   |          |               |         | 説明          |                                                 |
| isbn          | VARCHAR(16) |       |  true   |          |               |         | ID            |                                                 |
| thumbnail_url | VARCHAR(256)|       |  false  |          |               |         | サムネイル     |                                                 |
| published_at  | DATETIME    |       |  true   |          |               |         | 発売日         |                                                 |
| created_at    | DATETIME    |       |  true   |          |               |         | 登録日時       |                                                 |
| updated_at    | DATETIME    |       |  true   |          |               |         | 更新日時       |                                                 |

### authors

|   Field       |  DataType   | PK/FK | NotNull | Unsigned |     Other     | Default | Explanation  |                   Validation                    |
| :---------    | :---------  | :---: | :-----: | :------: | :-----------: | :------ | :----------  | :---------------------------------------------- |
| id            | VARCHAR(32)  |  PK   |  true   |         |      UQ       |         | ID           |                                                 |
| author        | VARCHAR(32) |       |  true   |          |               |         | 著書　        |                                                 |
| created_at    | DATETIME    |       |  true   |          |               |         | 登録日時       |                                                 |
| updated_at    | DATETIME    |       |  true   |          |               |         | 更新日時       |                                                 |

### categories
|   Field       |  DataType   | PK/FK | NotNull | Unsigned |     Other     | Default | Explanation  |                   Validation                    |
| :---------    | :---------  | :---: | :-----: | :------: | :-----------: | :------ | :----------  | :---------------------------------------------- |
| id            | VARCHAR(32)  |  PK   |  true   |         |      UQ       |         | ID           |                                                 |
| category      | VARCHAR(32) |       |  true   |          |               |         | カテゴリー    |                                                 |
| created_at    | DATETIME    |       |  true   |          |               |         | 登録日時       |                                                 |
| updated_at    | DATETIME    |       |  true   |          |               |         | 更新日時       |                                                 |

### publishers
|   Field       |  DataType   | PK/FK | NotNull | Unsigned |     Other     | Default | Explanation  |                   Validation                    |
| :---------    | :---------  | :---: | :-----: | :------: | :-----------: | :------ | :----------  | :---------------------------------------------- |
| id            | VARCHAR(32)  |  PK   |  true   |         |      UQ       |         | ID           |                                                 |
| publisher     | VARCHAR(32) |       |  true   |          |               |         | 出版社　　    |                                                 |
| created_at    | DATETIME    |       |  true   |          |               |         | 登録日時       |                                                 |
| updated_at    | DATETIME    |       |  true   |          |               |         | 更新日時       |                                                 |

### authors_books
|   Field       |  DataType   | PK/FK | NotNull | Unsigned |     Other     | Default | Explanation  |                   Validation                    |
| :---------    | :---------  | :---: | :-----: | :------: | :-----------: | :------ | :----------  | :---------------------------------------------- |
| id            | VARCHAR(32) |  PK   |  true   |         |      UQ       |         | ID           |                                                 |
| book_id       | VARCHAR(32) |  FK   |  true   |          |     UQ        |         | 本ID         |                                                　 |
| author_id     | VARCHAR(32) |  FK   |  true   |          |     UQ        |         | 著書ID  　    |                                                　 |
| created_at    | DATETIME    |       |  true   |          |               |         | 登録日時       |                                                 |
| updated_at    | DATETIME    |       |  true   |          |               |         | 更新日時       |                                                 |

### users_books
|   Field       |  DataType   | PK/FK | NotNull | Unsigned |     Other     | Default | Explanation  |                   Validation                    |
| :---------    | :---------  | :---: | :-----: | :------: | :-----------: | :------ | :----------  | :---------------------------------------------- |
| id            | VARCHAR(32)  |  PK   |  true   |         |      UQ       |         | ID           |                                                 |
| user_id       | VARCHAR(32)  |  FK   |  true   |         |      UQ       |         | ユーザーID  　 |                                                |
| book_id       | VARCHAR(32)  |  FK   |  true   |         |      UQ       |         | 本ID       　 |                                                |
| status        | TYNYINT(5)   |       |  true   |          |               |         | 本のステータス(0: 未読 / 1: 読んだ/2: 読んでる/3: 積読/4: 手放したい/5: 欲しい)| |   
| impression    | TEXT        |       |  false   |          |               |         | 感想          |                                                 |
| created_at    | DATETIME    |       |  true   |          |               |         | 登録日時       |                                                 |
| updated_at    | DATETIME    |       |  true   |          |               |         | 更新日時       |                                                 |

### categories_books
|   Field       |  DataType   | PK/FK | NotNull | Unsigned |     Other     | Default | Explanation  |                   Validation                    |
| :---------    | :---------  | :---: | :-----: | :------: | :-----------: | :------ | :----------  | :---------------------------------------------- |
| id            | VARCHAR(32) |  PK   |  true   |          |      UQ       |         | ID           |                                                 |
| category_id   | VARCHAR(32)  |  FK   |  true   |         |      UQ       |         | カテゴリーID  |                                                |
| book_id       | VARCHAR(32)  |  FK   |  true   |         |     UQ          |         | 本ID       　 |                                                |
| created_at    | DATETIME    |       |  true   |          |               |         | 登録日時       |                                                 |
| updated_at    | DATETIME    |       |  true   |          |               |         | 更新日時       |                                                 |

### follows

|    Field    |  DataType  | PK/FK | NotNull | Unsigned | Other |         Default          |        Explanation       | Validation |
| :---------- | :--------- | :---: | :-----: | :------: | :---: | :----------------------- | :----------------------  | ---------- |
| id          | VARCHAR(32)|  PK   |  true   |           |  UQ   |                          | ID                       |　          |
| follow_id | VARCHAR(32) |  FK   |  true   |          |  UQ   |                          | フォローするユーザーID     |            |
| follower_id | VARCHAR(32) |  FK   |  true   |          |  UQ   |                          | フォローされるユーザーID   |            |
| created_at  | DATETIME   |       |  true   |          |       |                          | 登録日時                  |            |
| updated_at  | DATETIME   |       |  true   |          |       |                          | 更新日時                  |            |

### followers

|    Field    |  DataType  | PK/FK | NotNull | Unsigned | Other |         Default          | Explanation | Validation |
| :---------- | :--------- | :---: | :-----: | :------: | :---: | :----------------------- | :---------- | ---------- |
| id          | VARCHAR(32)|  PK   |  true   |          |  UQ   |                          | ID                       |　          |
| follow_id | VARCHAR(32) |  FK   |  true   |         |  UQ  |                          | フォローするユーザーID     |            |
| follower_id | VARCHAR(32) |  FK   |  true   |         |  UQ  |                          | フォローされるユーザーID   |            |
| created_at  | DATETIME   |       |  true   |          |       |                          | 登録日時                  |            |
| updated_at  | DATETIME   |       |  true   |          |       |                          | 更新日時                  |            |

### Sale

| Field | DataType | PK/FK | NotNull | Unsigned | Other | Default | Explanation | Validation |
| :---- | :------- | :---: | :-----: | :------: | :---: | :------ | :---------- | ---------- |
