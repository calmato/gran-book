# Nuxt.js - コンポーネント設計

[root](./../../../../README.md)
/ [11_frontend](./../../README.md)
/ [02_web](./../README.md)
/ [01_admin](./README.md) 
/ [ディレクトリ構成](./directories.md)

## アトミックデザインルール

* Atoms
  * Vuetifyにない自作のコンポーネントを作成する際に使用

* Molecules
  * Atoms, Vuetifyのコンポーネントを組み合わせたものを作成する際に使用
  * e.g.) SubmitButton, CloseButton, SocialButton

* Organisms
  * FormやCard, Navigationはここで作成

* Templates
  * UIのワイヤーフレームを定義
  * Vuetifyの `v-app`, `v-content`, `v-container`, `v-row`, `v-col` はここで使用

* Pages
  * Templatesのみ参照

## ハンドラ命名規則

* Atoms ~ Templates
  * on○○
    * Atoms ~ Organsisms -> onClick, onChange
    * Templates -> onClickSubmitForm, onClickCloseDialog
* Pages
  * handle○○ -> handleSubmit, handleUpload

---

## 参考

* [Qiita - TypeScriptのInterfaceとTypeの比較](https://qiita.com/tkrkt/items/d01b96363e58a7df830e#%E6%AF%94%E8%BC%83%E3%81%BE%E3%81%A8%E3%82%81)
* [Zenn - Nuxt(vue) + TypeScriptをはじめるときに知っておきたかった10のこと](https://zenn.dev/nus3/articles/ec0db8857209a509646b)
