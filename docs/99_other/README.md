# 仕様書

[root](./../../README.md) 
/ [99_other](./README.md)

## Githubの使用ルール

## ブランチルール

* ブランチ名の付け方: {Prefix}/{システム名}/{実装概要}
  * e.g.) feature/user_api/create_users

* Prefix一覧

| Branch Name |        Description         |
| :---------- | :------------------------- |
| master      | Default                    |
| develop     | 開発用ブランチ             |
| release     | 本番リリース用ブランチ     |
| feature     | 機能の実装・編集用ブランチ |
| bugfix      | バグ修正用ブランチ         |
| refactor    | コードリファクタ用ブランチ |
| test        | テスト用ブランチ           |

## コミットルール

* Commit Massage Prefix一覧

* Format: <type>(<scope>): <subject>

```txt
feat: add hat wobble
^--^  ^------------^
|                               |
| +-> Summary in past tense. |
|                               |
+-------> Type: feat, fix, refactor, test, docs, or style.
```

|  Prefix  |                            Description                             |
| :------- | :----------------------------------------------------------------- |
| feat     | new feature for the user, not a new feature for build script       |
| fix      | bug fix for the user, not a fix to a build script                  |
| refactor | efactoring production code, eg. renaming a variable                |
| test     | adding missing tests, refactoring tests; no production code change |
| docs     | changes to the documentation                                       |
| style    | formatting, missing semi colons, etc; no production code change    |

## プルリクルール

* Label一覧

|       Labels        |            Description            |
| :------------------ | :-------------------------------- |
| feature             | 機能の実装・編集                  |
| bugfix              | バグの修正                        |
| refactor            | コードのリファクタ (機能面)       |
| release             | リリース準備                      |
| sytle               | コードのリファクタ (インデント等) |
| test                | テストの作成・編集                |
| documentation       | 仕様書の更新                      |
| WIP                 | 未完成状態                        |
| design              | UI/UX関連の実装・編集             |
| frontend (admin)    | 管理コンソールの実装・編集        |
| frontend (native)   | ネイティブアプリの実装・編集      |
| backend (auth api)  | Auth APIの実装・編集              |
| backend (book api)  | Book APIの実装・編集              |
| backend (other api) | 外部APIの実装・編集               |
| backend (store api) | Store APIの実装・編集             |
| backend (user api)  | User APIの実装・編集              |
| database            | データベース関連の実装・編集      |
| infrastructure      | インフラ関連の実装・編集          |
| monitorinig         | モニタリング関連の実装・編集      |
| workflows           | CI/CD関連の実装・編集             |
