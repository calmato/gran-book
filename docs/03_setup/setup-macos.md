# 環境構築 - macOS

[root](./../../README.md) 
/ [03_setup](./README.md) 
/ [macOS - 環境構築](./setup-macos.md)

## macOS (Intel)の環境構築手順

## 動作確認

| item  |           detail            |
| :---- | :-------------------------- |
| OS    | macOS Big Sur (11.4)        |
| CPU   | 2CPU                        |
| MEM   | 8GB                         |
| Disk  | 256GB                       |
| H/W   | MacBook Pro (13-inch, 2016) |
| Shell | zsh                         |

## 概要

1. brewのインストール
2. GitHubリポジトリのクローン
3. コンテナ実行環境の設定
4. ネイティブアプリ起動設定

### brewのインストール

1. Spotlight検索より `Terminal` を開く

2. XCodeをインストール
    > $ xcode-select --install

3. Homebrewをインストール
    > $ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

4. brewコマンドが正常に動作していることを確認
    > $ brew doctor

### GitHubリポジトリのクローン

1. Gitのインストール
    > $ brew install git

2. GitHubリポジトリをクローン
    1. リポジトリをクローンするディレクトリに移動
    2. リポジトリをクローン
        > $ git clone https://github.com/calmato/gran-book.git

### コンテナ実行環境の設定

1. コンテナ実行環境の準備
    1. Docker をインストール
        > $ brew install --cask docker
    2. Docker Compose をインストール
        > $ brew install docker-compose
    3. SpotLight検索より `Docker` を起動

2. コンテナ実行環境の動作確認
    1. Docker がインストールされていることを確認
        > $ docker version
    2. Docker Compose がインストールされていることを確認
        > $ docker-compose version
    3. Docker が起動していることを確認
        > $ docker info

3. コンテナイメージの作成
    1. .envファイルを作成
        * (記述内容は誰かに聞く)
    2. コンテナをビルド
        > $ make setup

4. コンテナの起動
    > $ make start

### ネイティブアプリ起動設定

1. Node.jsのインストール
    1. nodebrewのインストール
        > $ curl -L git.io/nodebrew | perl - setup
    2. PATHを通す
        > $ echo 'export PATH=$HOME/.nodebrew/current/bin:$PATH' >> ~/.zshrc
    3. シェルの再読み込み
        > $ eval $SHELL
    4. Node.jsをインストール (e.g. v16.3.0 のインストール)
        > $ nodebrew install v16.3.0
    5. インストールしたバージョンを使用するよう設定
        > $ nodebrew use v16.3.0
    6. yarnのインストール
        > $ npm install --global yarn

2. Expoのインストール
    > $ yarn global add expo-cli

3. ネイティブアプリの起動検証
    1. リポジトリをクローンしたディレクトリへ移動
    2. nativeディレクトリへ移動
        > $ cd ./native
    3. 起動に必要なライブラリをインストール
        > $ yarn
    4. Expoを起動
        > $ expo start

---

## 参考

* [GitHub - nodebrew](https://github.com/hokaccha/nodebrew)
