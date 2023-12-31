# blog

[![Deploy](https://github.com/Fukkatsuso/blog/workflows/Deploy/badge.svg)](https://github.com/Fukkatsuso/blog/actions?query=workflow%3ADeploy)
[![Lint](https://github.com/Fukkatsuso/blog/workflows/Lint/badge.svg)](https://github.com/Fukkatsuso/blog/actions?query=workflow%3ALint)

Hugo で Markdown から静的サイトを生成し、GAE でホスティングするブログ

## Hugo サーバを立ち上げる

```sh
docker-compose up
```

## デザインテーマ変更

例：hugo-notepadium

```sh
docker-compose run --rm hugo \
  git submodule add https://github.com/cntrump/hugo-notepadium.git themes/hugo-notepadium
```

```toml
# config.toml
theme = "hugo-notepadium"
```

## 記事作成

```sh
docker-compose run --rm hugo hugo new posts/my-first-post.md
```

or

```sh
docker-compose run --rm hugo hugo new posts/my-first-post/index.md
```

## OGP 画像生成

1. og-image.config.json に画像生成のための設定を追記

1. markdown の frontmatter に以下を記入
  
    ```yml
    images:
      - posts/<記事タイトル>/og-image.png
    ```

1. ローカルで画像生成を試す場合のコマンド

    ```sh
    docker build -t og-image og-image/
    docker run --rm \
      -v `pwd`:/go/src/github.com/Fukkatsuso/blog \
      -w /go/src/github.com/Fukkatsuso/blog \
      og-image og-image.config.json
    ```

## textlint をかける

必要なツールのインストール

```sh
npm install -g \
  textlint \
  textlint-rule-preset-ja-spacing \
  textlint-rule-preset-jtf-style
```

taichi.vscode-textlint（VSCode の拡張機能）をインストールして settings.json に以下を追記

```json
{
  "textlint.run": "onSave",
  "textlint.languages": [
      "markdown"
  ]
}
```

これにより、ファイル保存時に自動でチェックしてくれる。

## GAE へデプロイ

### Cloud Shell 上での準備

1. プロジェクト、GAE アプリの作成

    ```sh
    export PROJECT_ID=blog-XXXXXX
    gcloud projects create --name ${PROJECT_ID}
    gcloud config set project ${PROJECT_ID}
    gcloud app create
    ```

1. API を有効化（Cloud Build のために課金を有効にする）

    ```sh
    gcloud services enable appengine.googleapis.com

    gcloud alpha billing accounts list
    gcloud alpha billing projects link ${PROJECT_ID} --billing-account YYYYYY-ZZZZZZ-AAAAAA
    gcloud services enable cloudbilling.googleapis.com
    gcloud services enable cloudbuild.googleapis.com
    ```

1. サービスアカウント、サービスアカウントキーの作成

    ```sh
    export SA_NAME=githubactions
    gcloud iam service-accounts create ${SA_NAME} \
      --description="used by GitHub Actions" \
      --display-name="${SA_NAME}"
    gcloud iam service-accounts list

    export IAM_ACCOUNT=${SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com

    gcloud iam service-accounts keys create ~/${PROJECT_ID}/${SA_NAME}/key.json \
      --iam-account ${IAM_ACCOUNT}
    ```

1. role 付与

    ```sh
    gcloud projects add-iam-policy-binding ${PROJECT_ID} \
      --member="serviceAccount:${IAM_ACCOUNT}" \
      --role='roles/compute.storageAdmin'
    gcloud projects add-iam-policy-binding ${PROJECT_ID} \
      --member="serviceAccount:${IAM_ACCOUNT}" \
      --role='roles/cloudbuild.builds.editor'
    gcloud projects add-iam-policy-binding ${PROJECT_ID} \
      --member="serviceAccount:${IAM_ACCOUNT}" \
      --role='roles/appengine.deployer'
    gcloud projects add-iam-policy-binding ${PROJECT_ID} \
      --member="serviceAccount:${IAM_ACCOUNT}" \
      --role='roles/appengine.appAdmin'
    gcloud projects add-iam-policy-binding ${PROJECT_ID} \
      --member="serviceAccount:${IAM_ACCOUNT}" \
      --role='roles/cloudbuild.builds.builder'
    gcloud iam service-accounts add-iam-policy-binding ${PROJECT_ID}@appspot.gserviceaccount.com \
      --member="serviceAccount:${IAM_ACCOUNT}" \
      --role='roles/iam.serviceAccountUser'
    ```

### GitHub の Secrets

- GCP_PROJECT: プロジェクト ID
- GCP_SA_KEY: サービスアカウントの JSON 鍵を Base64 エンコード

    ```sh
    # Cloud Shell
    openssl base64 -in ~/${PROJECT_ID}/${SA_NAME}/key.json
    ```

### GitHub へ Push

- main.go と app.yaml を忘れずに
- 公開記事は `draft: false` にする
- master ブランチへの Push で自動的に GAE へデプロイしてくれる

## 参考

- [Hugo クイックスタート](https://gohugo.io/getting-started/quick-start/)
- デザインテーマ
  - [Hugo Notepadium](https://themes.gohugo.io/hugo-notepadium/)
- [gcloud リファレンス](https://cloud.google.com/sdk/gcloud/reference?hl=ja)
- [app.yaml リファレンス](https://cloud.google.com/appengine/docs/standard/go/config/appref?hl=ja)
- [markdownlint Rules](https://github.com/DavidAnson/markdownlint/blob/main/doc/Rules.md)
