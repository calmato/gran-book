## DB設計

[root](./../../README.md) 
/ [13_database](./db_design.md)


### users
|        Field         |      Explanation      |     DataType      |    Validation                          |
| :------------------- | :----------------- | :---------------- |:-------------------------------------- |
| id               | ユーザーID          | INT        | ・Required  (FK)                                               |
| username             | ニックネーム           | VARCHAR(32)      | ・Required(NN)<br/>・Length: n <= 32                             |
| sex                  | 性別                | Boolean       | ・Required(NN)                                                   |
| email                | メールアドレス         | VARCHAR(256)     | ・Required(NN)<br/>・Length: n <= 256<br/>・Format: Email        |
| thumbnail            | サムネイル            | BLOB     | ・Format: base64                                             |
| password             | パスワード            | VARCHAR(32)    | ・Required(NN)<br/>・Length: 6 <= n <= 32<br/>・Format: Password |
| self_introduction               | 自己紹介          | VARCHAR(1000)    | Length: 0 <= n <= 1000                      |
| last_name       | 名字（漢字)      | VARCHAR(16)      | Length: 0 <= n <= 16            |
| first_name       | 名前（漢字)      | VARCHAR(16)      | Length: 0 <= n <= 16            |
| last_name_kana       | 配送の際の名前（かな)      | VARCHAR(16)      | Length: 0 <= n <= 16            |
| first_name_kana      | 配送の際の名前（かな)      | VARCHAR(16)      | Length: 0 <= n <= 16            |
| phone                | 電話番号                | VARCHAR(20)        | Length: 0 <= n <= 20            |
| postal_code          | 郵便番号                | VARCHAR(32)        | Length: 0 <= n <= 32           |
| prefectures          | 都道府県                | VARCHAR(16)        | Length: 0 <= n <= 16           |
| cities               | 市町村                 |  VARCHAR(256)      | Length: 0 <= n <= 256          |
| other_address        | マンション・ビル名・郵便番号   | VARCHAR(256)      | Length: 0 <= n <= 256           |
| instance_id           | 端末ID             | BLOB     |・Required                                                   |

### books   
|        Field         |      Explanation      |     DataType      |    Validation                          |
| :------------------- | :----------------- | :---------------- |:-------------------------------------- |
| book_id               | 本ID          | INT        |・Required  (FK)              | 
| user_id            | ユーザーID        | INT        | FK   |
| book_status         | 本の状態(読んだ / 読んでいる /積読 /手放したい /欲しい)            | VARCHAR(10)            |  NN     |

### follws
|        Field         |      Explanation      |     DataType      |    Validation                          |
| :------------------- | :----------------- | :---------------- |:-------------------------------------- |
| follow_id            | フォローID        | INT        | FK   |
| follow_user_id            | フォローするユーザーID        | INT        | FK   |
| followed_user_id            | フォローされるユーザーID        | INT        | FK   |

### followers
|        Field         |      Explanation      |     DataType      |    Validation                          |
| :------------------- | :----------------- | :---------------- |:-------------------------------------- |
| follow_user_id            | フォローするユーザーID        | Length(Array): 0 <= n <= 9999         | FK   |
| followed_user_id            | フォローされるユーザーID        | Length(Array): 0 <= n <= 9999         | FK   |

### Sale    
|        Field         |      Japanese      |                          Validation                          |
| :------------------- | :----------------- | :----------------------------------------------------------- |
