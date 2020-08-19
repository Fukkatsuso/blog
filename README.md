# blog
HugoでMarkdownから静的サイトを生成し，GAEでホスティングするブログ

## Hugoサーバを立ち上げる
```sh
docker build -t hugo .
sudo docker run --rm -it -v `pwd`:/go/src/github.com/Fukkatsuso/blog -p 1313:1313 hugo \
  hugo server --bind=0.0.0.0 -D
```

## デザインテーマ変更
例: hugo-notepadium
```sh
sudo docker run --rm -it -v `pwd`:/go/src/github.com/Fukkatsuso/blog hugo \
  git submodule add https://github.com/cntrump/hugo-notepadium.git themes/hugo-notepadium
```
```toml
# config.toml
theme = "hugo-notepadium"
```

## 記事作成
```sh
sudo docker run --rm -it -v `pwd`:/go/src/github.com/Fukkatsuso/blog hugo \
  hugo new posts/my-first-post.md
```
or
```sh
sudo docker run --rm -it -v `pwd`:/go/src/github.com/Fukkatsuso/blog hugo \
  hugo new posts/my-first-post/index.md
```

## GAEへデプロイ
### Cloud Shell上での準備
1. プロジェクト, GAEアプリの作成
```sh
gcloud projects create --name blog
gcloud config set project blog-XXXXXX
gcloud app create
```

2. APIを有効化(Cloud Buildのために課金を有効にする)
```sh
gcloud services enable appengine.googleapis.com

gcloud alpha billing accounts list
gcloud alpha billing projects link blog-XXXXXX --billing-account YYYYYY-ZZZZZZ-AAAAAA
gcloud services enable cloudbilling.googleapis.com
gcloud services enable cloudbuild.googleapis.com
```

3. サービスアカウント, サービスアカウントキーの作成
```sh
gcloud iam service-accounts create SA-NAME \
  --description="used by GitHub Actions" \
  --display-name="SA-NAME"
gcloud iam service-accounts list

gcloud iam service-accounts keys create ~/blog/SA-NAME/key.json \
  --iam-account SA-NAME@blog-XXXXXX.iam.gserviceaccount.com
```

4. role付与
```sh
gcloud projects add-iam-policy-binding blog-XXXXXX \
  --member='serviceAccount:SA-NAME@blog-XXXXXX.iam.gserviceaccount.com' \
  --role='roles/compute.storageAdmin'
gcloud projects add-iam-policy-binding blog-XXXXXX \
  --member='serviceAccount:SA-NAME@blog-XXXXXX.iam.gserviceaccount.com' \
  --role='roles/cloudbuild.builds.editor'
gcloud projects add-iam-policy-binding blog-XXXXXX \
  --member='serviceAccount:SA-NAME@blog-XXXXXX.iam.gserviceaccount.com' \
  --role='roles/appengine.deployer'
gcloud projects add-iam-policy-binding blog-XXXXXX \
  --member='serviceAccount:SA-NAME@blog-XXXXXX.iam.gserviceaccount.com' \
  --role='roles/appengine.appAdmin'
gcloud projects add-iam-policy-binding blog-XXXXXX \
  --member='serviceAccount:SA-NAME@blog-XXXXXX.iam.gserviceaccount.com' \
  --role='roles/cloudbuild.builds.builder'
```

### GitHubのSecrets
- GCP_PROJECT: プロジェクトID
- GCP_SA_EMAIL: サービスアカウントemail
- GCP_SA_KEY: サービスアカウントのJSON鍵をBase64エンコード
  ```sh
  # Cloud Shell
  openssl base64 -in ~/blog/SA-NAME/key.json
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