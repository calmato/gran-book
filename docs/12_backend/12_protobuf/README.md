# バックエンド - Protocol Buffers

[root](./../../../README.md) 
/ [12_backend](./../README.md) 
/ [12_protobuf](./README.md)

# 各種ドキュメント

## 標準メソッド命名ルール

| Standard Method |        HTTP Mapping         | HTTP Request Body |   HTTP Response Body    |
| :-------------- | :-------------------------- | :---------------- | :---------------------- |
| List            | GET <collection URL>        | N/A               | Resource* list          |
| Get             | GET <resource URL>          | N/A               | Resource*               |
| Create          | POST <collection URL>       | Resource          | Resource*               |
| Update          | PUT or PATCH <resource URL> | Resource          | Resource*               |
| Delete          | DELETE <resource URL>       | N/A               | google.protobuf.Empty** |

### Example

|   Method   |    HTTP Mapping    |  Request Message  | Response Message  |
| :--------- | :----------------- | :---------------- | :---------------- |
| ListBooks  | GET: /books        | EmptyRequest      | ListBooksResponse |
| GetBook    | GET: /books/:id    | EmptyRequest      | BookResponse      |
| CreateBook | POST: /books       | CreateBookRequest | BookResponse      |
| UpdateBook | PATCH: /books/:id  | UpdateBookRequest | BookResponse      |
| DeleteBook | DELETE: /books/:id | EmptyRequest      | EmptyResponse     |

---

## 参考

* [Cloud API - API設計ガイド](https://cloud.google.com/apis/design)
* [Cloud API - 標準メソッド](https://cloud.google.com/apis/design/standard_methods)
