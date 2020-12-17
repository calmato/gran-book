# バックエンド - API リクエスト・レスポンス設計

[root](./../../../README.md) 
/ [12_backend](./../README.md) 
/ [01_design](./README.md) 
/ [リクエスト・レスポンス設計](./request-response.md)

## 基本方針

* JSON形式でのリクエスト・レスポンスを作成

---

## リクエスト設計

### ヘッダー

|    Element    |     Type      |     Example      |      Description       |
| :------------ | :------------ | :--------------- | :--------------------- |
| Accept        | RequestHeader | application/json | レスポンスボディの形式 |
| Content-Type  | EntityHeader  | application/json | リクエストボディの形式 |
| Authorization | RequestHeader |                  | 認証キー               |

---

## レスポンス設計

### ヘッダー

|             Element              |      Type      |              Example              |                                                                             Description                                                                              |
| :------------------------------- | :------------- | :-------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Access-Control-Allow-Methods     | ResponseHeader | GET,POST,PATCH,PUT,DELETE,OPTIONS | プリフライトリクエストのレスポンスの中で、リソースにアクセスするときに利用できる1つまたは複数のメソッドを指定                                                        |
| Access-Control-Allow-Headers     | ResponseHeader | *                                 | プリフライトリクエストへのレスポンスで、実際のリクエストの間に使用できる HTTP ヘッダーを示すために使用                                                               |
| Access-Control-Allow-Origin      | ResponseHeader | *                                 | 指定されたオリジンからのリクエストを行うコードでレスポンスが共有できるかどうか                                                                                       |
| Access-Control-Max-Age           | ResponseHeader | 600                               | プリフライトリクエストの結果 (つまり Access-Control-Allow-Methods および Access-Control-Allow-Headers ヘッダーに含まれる情報) をキャッシュすることができる時間の長さ |
| Access-Control-Allow-Credentials | ResponseHeader | false                             | リクエストの資格情報モード (Request.credentials) が include である場合に、レスポンスをフロントエンドの JavaScript コードに公開するかどうかをブラウザーに指示         |

### ステータスコード

| StatusCode |      Description      |                                                                                     Notes                                                                                      |
| ---------: | :-------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|        200 | OK                    | Not an error. Returned on success.                                                                                                                                             |
|        400 | Bad Request           | Bad Request/Invalid Arguments.                                                                                                                                                 |
|        401 | Unauthorized          | Unauthenticated. The request does not have valid authentication credentials for the operation.                                                                                 |
|        403 | Forbidden             | Permission denied/forbidden. This response must not be used for rejections caused by exhausting some resource. Forbidden must not be used if the caller can not be identified. |
|        404 | Not Found             | Not found.                                                                                                                                                                     |
|        409 | Conflict              | The operation was aborted, typically due to a concurrency issue such as a sequencer check failure or transaction abort.                                                        |
|        500 | Internal Server Error | Internal errors. This means that some invariants expected by the underlying system have been broken. This error code is reserved for serious errors.                           |
|        501 | Not Implemented       | The operation is not implemented or is not supported/enabled in this service.                                                                                                  |
|        503 | Service Unavailable   | The service is currently unavailable. This is most likely a transient condition, which can be corrected by retrying with a backoff.                                            |
|        504 | Gateway Timeout       | The deadline expired before the operation could complete.                                                                                                                      |

### エラーレスポンス

* レスポンス構造

```json
{
  "errors": [
    "field": "email",
    "reason": "is required",
    "message": "email is required"
  ],
  "code": 400,
  "message": "Bad Request/Invalid Arguments."
}
```

|    Element     |                    Description                    |
| :------------- | :------------------------------------------------ |
| status         | HTTPステータスコード                              |
| code           | エラーコード                                      |
| message        | エラーの概要                                      |
| errors         | エラーの詳細一覧                                  |
| errors.field   | エラーの対象                                      |
| errors.reason  | エラーの原因                                      |
| errors.message | エラーメッセージ(fieldとreasonを組み合わせたもの) |

* エラーコード

| ErrorCode | StatusCode |         Summary          |                                      Description                                       |
| --------: | ---------: | :----------------------- | :------------------------------------------------------------------------------------- |
|       101 |        400 | InvalidRequestValidation | API リクエストが無効であるか、形式が正しくありません。                                 |
|       102 |        400 | UnableParseJSON          | JSON型から構造体への変換エラー。                                                       |
|       103 |        400 | UnableConvertBase64      | Byte64型への変換エラー。                                                               |
|       111 |        401 | Unauthorized             | リクエストを行う権限がユーザーにありません。                                           |
|       112 |        401 | Expired                  | セッションが有効期限切れです。                                                         |
|       121 |        403 | Forbidden                | リクエストされた操作は禁止されているため、完了できません。                             |
|       131 |        404 | NotFound                 | リクエストに関連付けられたリソースが見つかりません。                                   |
|       132 |        404 | NotExistsInDatastore     | データストアに対象のレコードが存在しません。                                           |
|       132 |        404 | NotExistsInStorage       | ストレージに対象のファイルが存在しません。                                             |
|       141 |        409 | Conflict                 | リクエストされた操作は既存の項目と競合するため、API リクエストを完了できませんでした。 |
|       151 |        500 | InternalServerError      | 内部エラーのためリクエストは失敗しました。                                             |
|       152 |        500 | ErrorInDatastore         | データストアでのエラー                                                                 |
|       153 |        500 | ErrorInStorage           | ストレージでのエラー                                                                   |
|       154 |        500 | ErrorInOtherAPI          | 他のAPIでのエラー                                                                      |
|       155 |        500 | Unknown                  | 不明なエラー                                                                           |

---

## 参考

* [Google Maps Booking API - Status Response Codes](https://developers.google.com/maps-booking/reference/rest-api-v3/status_codes)
* [Cloud Storage - HTTP status and error codes for JSON](https://cloud.google.com/storage/docs/json_api/v1/status-codes)
* [キャンペーンマネージャー360 - エラーメッセージ](https://developers.google.com/doubleclick-advertisers/core_errors?hl=ja)
