## DB設計

### User
|        Field         |      Explanation      |     DataType      |    Validation                          |
| :------------------- | :----------------- | :---------------- |:-------------------------------------- |
| user_id               | ユーザーUID          | INT        | ・Required  (FK)                                               |
| user_name             | ニックネーム           | VARCHAR(32)      | ・Required(NN)<br/>・Length: n <= 32                             |
| sex                  | 性別                | Boolean       | ・Required(NN)                                                   |
| email                | メールアドレス         | VARCHAR(256)     | ・Required(NN)<br/>・Length: n <= 256<br/>・Format: Email        |
| thumbnail            | サムネイル            | BLOB     | ・Format: base64                                             |
| password             | パスワード            | VARCHAR(32)    | ・Required(NN)<br/>・Length: 6 <= n <= 32<br/>・Format: Password |
| user_follower               | フォローワー         | User   | Length(Array): 0 <= n <= 9999                      |
| user_follow               | フォローしているユーザー          | User   | Length(Array): 0 <= n <= 9999                      |
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

### Book   
|        Field         |      Japanese      |                          Validation                          |
| :------------------- | :----------------- | :----------------------------------------------------------- |
| Book_id               | ユーザーUID          | INT        |・Required  (FK)                                               |


### Sale    
|        Field         |      Japanese      |                          Validation                          |
| :------------------- | :----------------- | :----------------------------------------------------------- |
