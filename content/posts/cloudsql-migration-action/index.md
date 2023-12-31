---
title: "GitHub Actions で Cloud SQL のマイグレーションを自動化する"
date: 2021-07-14T23:22:00+09:00
draft: false
summary: "GCP の Cloud SQL を利用したプロジェクトで、GitHub Actions と golang-migrate を使って DB のマイグレーションを自動化してみました。"
images:
  - posts/cloudsql-migration-action/og-image.png
categories:
  - development
tags:
  - GitHubActions
  - GCP
  - DB
---

GCP の Cloud SQL を利用したプロジェクトで、GitHub Actions と golang-migrate を使って DB のマイグレーションを自動化してみました。

## 構成

### 使用したツール・サービス

- GitHub Actions
- golang-migrate
- Cloud SQL Auth Proxy（Docker イメージ）
- Cloud SQL (MySQL5.7, db-f1-micro)

### サービス構成、フロー

![flow](flow.drawio.png)

流れは簡単で、

1. 開発者 (Developer) が GitHub に push
2. GitHub Actions でプロキシを使用して Cloud SQL に接続し、マイグレーション

これだけです。

### ディレクトリ構成

```bash
.
├── .github
│   └── workflows
│       └── migration.yml
└── db
    └── migrations
        ├── 000001_create_users_table.down.sql
        ├── 000001_create_users_table.up.sql
        ├── 000002_create_posts_table.down.sql
        └── 000002_create_posts_table.up.sql
```

`db/migrations` 以下に、golang-migrate の仕様に合わせたマイグレーションファイルを配置しています。
上では例として、users テーブルの作成、posts テーブルの作成をするつもりのファイルを配置しました。

GitHub Actions では `.github/workflows/migration.yml` に書いたワークフローを走らせます。
`migrate up` コマンドで `db/migrations/XXX.up.sql` が実行され、`migrate down` コマンドで `db/migrations/XXX.down.sql` が実行される仕様です。

## GCP

準備として、GCP のサービスアカウントを作成します。

API の有効化とサービスアカウントへの role 付与などが必要ですが、ここでは割愛します。
Cloud Run と Cloud SQL を使っており、GitHub Actions からマイグレーションを実行するようなプロジェクトにおいてサービスアカウントを作成する方法を[ドキュメント](https://github.com/Fukkatsuso/cryptocurrency-trading-bot/blob/main/doc/gcp_project.md)として書いています。
こちら参考にしてみてください。

サービスアカウントが作成できたらキー（json ファイル）を手元に保存しておきます。
キーは外部に公開しないでください。

Cloud SQL インスタンスも立てておく必要があります。
これも[ドキュメント](https://github.com/Fukkatsuso/cryptocurrency-trading-bot/blob/main/doc/gcp_project.md)に手順を書いています。

## Secrets の設定

GitHub のリポジトリのページから、Settings > Secrets の画面で Repository secrets の設定を行います。

- CLOUDSQL_INSTANCE: Cloud SQL のインスタンス名
- GCP_PROJECT: GCP のプロジェクト ID
- GCP_REGION: Cloud SQL のリージョン
- GCP_SA_KEY: GCP のサービスアカウントキー
- MYSQL_DATABASE: データベース名
- MYSQL_PASSWORD: データベースのパスワード
- MYSQL_USER: データベースのユーザ

これらを設定してください。
ワークフローから `${{ secrets.GCP_PROJECT }}` のように参照でき、なおかつログにはアスタリスクで隠された文字列として表示されるため安全です。

## マイグレーションのワークフロー

ようやく本題のマイグレーションが実行できます。
以下は `.github/workflows/migration.yml` の中身です。

```yml
name: Migration

on:
  push:
    branches:
      - main

env:
  PROXY_IMAGE: gcr.io/cloudsql-docker/gce-proxy
  CLOUDSQL_INSTANCE_CONNECTION_NAME: ${{ secrets.GCP_PROJECT }}:${{ secrets.GCP_REGION }}:${{ secrets.CLOUDSQL_INSTANCE }}
  MYSQL_DSN: mysql://${{ secrets.MYSQL_USER }}:${{ secrets.MYSQL_PASSWORD }}@tcp(127.0.0.1:3306)/${{ secrets.MYSQL_DATABASE }}

jobs:
  migrate-db:
    runs-on: ubuntu-18.04
    defaults:
      run:
        working-directory: db

    steps:
      - uses: actions/checkout@v1

      - name: Start Cloud SQL Proxy
        run: |
          echo '${{ secrets.GCP_SA_KEY }}' > sa_key
          docker pull $PROXY_IMAGE
          docker run -d \
            -v $PWD/sa_key:/config \
            -p 127.0.0.1:3306:3306 \
            $PROXY_IMAGE /cloud_sql_proxy \
            -instances=$CLOUDSQL_INSTANCE_CONNECTION_NAME=tcp:0.0.0.0:3306 \
            -credential_file=/config

      - name: Install migrate
        run: |
          curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | sudo apt-key add -
          echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" | sudo tee /etc/apt/sources.list.d/migrate.list
          sudo apt-get update
          sudo apt-get install -y migrate

      - name: Migrate DB (up)
        run: migrate -path "./migrations/" -database "$MYSQL_DSN" up
```

main ブランチへの push のみトリガーとするようにしています。
これにより、直接 main ブランチへ push したとき、もしくはプルリクなどで main ブランチにマージされたときにワークフローが開始されます。

以下、ステップごとの説明です。

### Step1. Cloud SQL のプロキシを起動

```yml
- name: Start Cloud SQL Proxy
  run: |
    echo '${{ secrets.GCP_SA_KEY }}' > sa_key
    docker pull $PROXY_IMAGE
    docker run -d \
      -v $PWD/sa_key:/config \
      -p 127.0.0.1:3306:3306 \
      $PROXY_IMAGE /cloud_sql_proxy \
      -instances=$CLOUDSQL_INSTANCE_CONNECTION_NAME=tcp:0.0.0.0:3306 \
      -credential_file=/config
```

権限付与した GCP のサービスアカウントのキーを、sa_key という名前のファイルに保存しておきます。
これはプロキシの認証に必要です。

次に Cloud SQL Proxy の公式の Docker イメージを pull し、`docker run` で起動します。
このコマンドは [GCP のドキュメント](https://cloud.google.com/sql/docs/postgres/connect-admin-proxy#connecting-docker)にほとんどそのまま書いてあります。

書き方がわかればこんなもんですが、サービスアカウントキーをシングルクオートで囲っておかないと、`echo` コマンドでリダイレクトする際に失敗してしまうという罠がありました。

### Step2. golang-migrate をインストール

```yml
- name: Install migrate
  run: |
    curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | sudo apt-key add -
    echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" | sudo tee /etc/apt/sources.list.d/migrate.list
    sudo apt-get update
    sudo apt-get install -y migrate
```

ジョブは Ubuntu で動かしているので [Linux 向けのインストール方法](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#linux-deb-package)をもとに記述しています。

これも、パイプラインやリダイレクトの権限周りでエラーが起き、`sudo` を付けたり `tee` コマンドに置き換えたりする必要がありました。

### Step3. マイグレーションを実行

```yml
- name: Migrate DB (up)
  run: |
    migrate -path "./migrations/" -database "$MYSQL_DSN" up
```

golang-migrate は、up や down のコマンドの引数として 1 や 2 といった数字を与えると、現在の DB のバージョンを 1 段階、2 段階上げたり下げたりすることが可能です。
数字を与えない場合、up もしくは down のマイグレーションが全て実行されます。

今回は up (`db/migrations/XXX.up.sql`) を全て実行して DB を最新バージョンにします。

基本的に、コードにコミットされているスクリプトは全て実行して、バージョンを下げたり細かな調整をしたりするのは手動でやるという想定です。

## 最後に

Cloud SQL のマイグレーションを自動化する方法を紹介しました。

Cloud Build を使った自動化についての記事はいくつもありますが、GitHub Actions でやる方法はあまり見かけなかったのでブログ記事に書いてみました。
ちなみに個人で利用する分には無料枠のある GitHub Actions の方がお手軽感もあって好きです。

できてしまえばこんなものかという感じがしますが、成功するまでのドキュメントを読みこんだりする過程で、クラウドサービスの扱いに慣れるきっかけにもなると思います。
何よりも、自動化すれば開発自体に集中できるという点が一番良いですね。

## 参考

- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- [Cloud SQL Auth Proxy について](https://cloud.google.com/sql/docs/postgres/sql-proxy)
- [Cloud SQL Auth Proxy を使用して接続する](https://cloud.google.com/sql/docs/postgres/connect-admin-proxy)
