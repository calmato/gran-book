## DB設計

[root](./../../README.md) 
/ [13_database](./db_design.md)


### users

|       Field       |         Explanation          |   DataType    |                            Validation                            |
| :---------------- | :--------------------------- | :------------ | :--------------------------------------------------------------- |
| id                | ユーザーID                   | INT           | NN/PK/AI                                                 |
| username          | ニックネーム                 | VARCHAR(32)   | NN<br/>・Length: n <= 32                             |
| gender               | 性別                         | Boolean       | NN                                                   |
| email             | メールアドレス               | VARCHAR(256)  | NN<br/>・Length: n <= 256<br/>・Format: Email        |
| thumbnail_url         | サムネイル                   | BLOB          | ・Format: base64                                                 |
| password_digest          | パスワード                   | VARCHAR(256)   | NN<br/>・Length: 6 <= n <= 32<br/>・Format: Password |
| self_introduction | 自己紹介                     | VARCHAR(1000) | Length: 0 <= n <= 1000                                           |
| last_name         | 名字（漢字)                  | VARCHAR(16)   | Length: 0 <= n <= 32                                             |
| first_name        | 名前（漢字)                  | VARCHAR(16)   | Length: 0 <= n <= 32                                             |
| lastname_kana    | 名前（かな)        | VARCHAR(16)   | Length: 0 <= n <= 32                                             |
| firstname_kana   | 名前（かな)        | VARCHAR(16)   | Length: 0 <= n <= 32                                             |
| phone             | 電話番号                     | VARCHAR(20)   | Length: 0 <= n <= 20 ハイフンなし、数字のみ                                            |
| postal_code       | 郵便番号                     | CHAR(8)   | xxx-xxxx　ハイフン付き、8文字                                          |
| prefectures       | 都道府県                     | VARCHAR(16)   | Length: 0 <= n <= 16                                             |
| cities            | 市町村                       | VARCHAR(256)  | Length: 0 <= n <= 256                                            |
| address_chrome     | 番地・丁目 | VARCHAR(256)  | Length: 0 <= n <= 256                                            |
| other_address     | マンション・ビル名・部屋番号 | VARCHAR(256)  | Length: 0 <= n <= 256                                            |
| created_at     | 登録日時 | TIMESTAMP  |    NN                                         |
| created_at     | 登録日時 | TIMESTAMP  |    NN                                        |
| instance_id       | 端末ID                       | BLOB          |　NNにするか検討                                                       |

### books

|    Field    |                       Explanation                       |  DataType   |    Validation    |
| :---------- | :------------------------------------------------------ | :---------- | :--------------- |
| id     | 本ID                                                    | INT         | ・Required  PK/AI |
| user_id     | ユーザーID                                              | INT         | FK               |
| status | 本の状態(読んだ / 読んでいる /積読 /手放したい /欲しい) | VARCHAR(10) | NN               |
| created_at     | 登録日時 | TIMESTAMP  |    NN                                         |
| created_at     | 登録日時 | TIMESTAMP  |    NN                                        |

### follws

|      Field       |       Explanation        | DataType | Validation |
| :--------------- | :----------------------- | :------- | :--------- |
| id        | フォローID               | INT      | PK/AI         |
| follower_id   | フォローするユーザーID   | INT      | FK         |
| followed_id | フォローされるユーザーID | INT      | FK         |

### followers

|      Field       |       Explanation        |           DataType            | Validation |
| :--------------- | :----------------------- | :---------------------------- | :--------- |
| follow_id   | フォローするユーザーID   | INT   | PK/AI         |
| followed_id | フォローされるユーザーID | INT  | FK         |
| created_at     | 登録日時 | TIMESTAMP  |    NN                                         |
| created_at     | 登録日時 | TIMESTAMP  |    NN                                        |

### Sale

| Field | Explanation | DataType | Validation |
| :---- | :---------- | :------- | :--------- |
