# blog
[![Deploy](https://github.com/Fukkatsuso/blog/workflows/Deploy/badge.svg)](https://github.com/Fukkatsuso/blog/actions?query=workflow%3ADeploy)

HugoでMarkdownから静的サイトを生成し，GAEでホスティングするブログ

## Hugoサーバを立ち上げる
```sh
docker-compose up
```

## デザインテーマ変更
例: hugo-notepadium
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

## GAEへデプロイ
### Cloud Shell上での準備
1. プロジェクト, GAEアプリの作成
```sh
export PROJECT_ID=blog-XXXXXX
gcloud projects create --name ${PROJECT_ID}
gcloud config set project ${PROJECT_ID}
gcloud app create
```

2. APIを有効化(Cloud Buildのために課金を有効にする)
```sh
gcloud services enable appengine.googleapis.com

gcloud alpha billing accounts list
gcloud alpha billing projects link ${PROJECT_ID} --billing-account YYYYYY-ZZZZZZ-AAAAAA
gcloud services enable cloudbilling.googleapis.com
gcloud services enable cloudbuild.googleapis.com
```

3. サービスアカウント, サービスアカウントキーの作成
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

4. role付与
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
```

### GitHubのSecrets
- GCP_PROJECT: プロジェクトID
- GCP_SA_KEY: サービスアカウントのJSON鍵をBase64エンコード
  ```sh
  # Cloud Shell
  openssl base64 -in ~/${PROJECT_ID}/${SA_NAME}/key.json
  ```

### GitHubへPush
- main.goとapp.yamlを忘れずに
- masterブランチへのPushで自動的にGAEへデプロイしてくれる

## 参考
- [Hugo クイックスタート](https://gohugo.io/getting-started/quick-start/)
- デザインテーマ
  - [Hugo Notepadium](https://themes.gohugo.io/hugo-notepadium/)
- [gcloud リファレンス](https://cloud.google.com/sdk/gcloud/reference?hl=ja)
- [app.yaml リファレンス](https://cloud.google.com/appengine/docs/standard/go/config/appref?hl=ja)