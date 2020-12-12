# React Native - ディレクトリ構成

[root](./../../../README.md)
/ [11_frontend](./../README.md)
/ [01_native](./README.md)
/ [ディレクトリ構成](./directories.md)

## ディレクトリ構成

<pre>
.
├── assets
└── src
    ├── components
    │   ├── atoms
    │   ├── molecules
    │   ├── organisms
    │   └── templates
    ├── layouts
    ├── lib
    │   ├── axios
    │   ├── firebase
    │   ├── hooks
    │   └── local-storage
    ├── screens
    │   ├── Account
    │   ├── Auth
    │   ├── Boarding
    │   ├── Home
    │   ├── Message
    │   ├── Search
    │   └── Store
    ├── store
    │   ├── containers
    │   ├── models
    │   ├── modules
    │   ├── selectors
    │   └── usecases
    └── types
        ├── forms
        ├── models
        ├── requests
        ├── responses
        └── selectors
</pre>

|    ディレクトリ名    | 関連ライブラリ |                                    説明                                    |
| :------------------- | :------------- | :------------------------------------------------------------------------- |
| components           |                | UI設計用のパーツ                                                           |
| components/atoms     |                | 最も小さい粒度の要素                                                       |
| components/molecules |                | 2つ以上のAtomsを組み合わせたシンプルなUI要素                               |
| components/organisms |                | 切り離して単体でも機能する分子の集まり                                     |
| components/templates |                | Organismsを組み合わせて画面の作成                                          |
| layouts              |                | Header,Footerなどの共通のUI要素を作成                                      |
| lib                  |                | アプリ内で使うライブラリの定義                                             |
| screens              |                | Screen(ルーティング周り)の設定                                             |
| screens/<Directory>  |                | Footerの要素ごとにScreenを構成                                             |
| store                |                | アプリの状態管理用                                                         |
| store/containers     | react-redux    | ComponentsとStoreを繋ぐ役割                                                |
| store/models         | redux          | モデルの定義, Vuex Mutationsのイメージ                                     |
| store/modules        | redux          | Actionの定義, アプリケーションのデータフローを表現, Vuex Actionsのイメージ |
| store/selectors      | reselect       | Componentsで表示する用のデータを生成, Vuex Gettersのイメージ               |
| store/usecases       | redux-thunk    | ユースケースを表現, Actionsの詳細を記述する感じ                            |
| types                |                | アプリ内で使用する関数,変数の型を定義                                      |
| types/forms          |                | コンポーネント内で使用するFormの型定義                                     |
| types/models         |                | 状態管理用のモデルの型定義                                                 |
| types/request        |                | APIへのリクエストの型定義                                                  |
| types/response       |                | APIからのレスポンスの型定義                                                |
| types/selectors      |                | Selectorsで取得する値の型定義                                              |
