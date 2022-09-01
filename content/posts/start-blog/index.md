---
title: "Hugo+GAEでブログを作った"
date: 2020-08-19T09:48:26+09:00
draft: false
summary: "静的サイトジェネレータHugoとGoogle App Engineを使って，速い・簡単・安いの3拍子揃った(?)技術ブログを作りました．"
images:
  - posts/start-blog/og-image.png
categories:
  - development
tags:
  - blog
  - hugo
---

静的サイトジェネレータ**Hugo**と**Google App Engine**を使って，速い・簡単・安いの3拍子揃った(?)技術ブログを作りました．

## 構成

ブログを作るにあたり，最低限必要な機能の洗い出しから始めました．

- Markdownで編集
- Git/GitHubで管理
- シンタックスハイライト
- SSR or SSG (検索流入のため)
- 独自ドメイン + SSL対応
- 低コスト
- 低レイテンシ
- 記事のタグ付け，タグ検索
- アクセス解析

ブログ記事をMarkdownで書きGitで管理したいのはエンジニアっぽい気がするからです．
そしてMarkdownは普段使っていて書きやすいから．
技術ブログなのでプログラムを載せるときにハイライトを付けてくれると見栄えが良いですね．

調べたところHugoが上記の条件をクリアできそうで，なおかつ最近書いているGolangで作られたものだったので，Hugoを使うことにしました．

構成は以下のようなものを考えました．

![architecture](architecture.png)

静的サイトのホスティングであればNetlifyやGitHub Pagesも候補に入りますが，クラウドの勉強がしたかったのと，GitHub Pagesはすでにポートフォリオサイトに使用しているため，無料枠が豊富で多少経験のあるGCPのサービスの中から料金や機能を考慮してGAEを選択しました．

独自ドメイン`blog.fukkatsuso.com`の設定とSSL対応も行っています．
ドメインのわかりやすさと将来性を考えてサブドメインでの運用が適切と判断し，Cloud DNSを挟んだ構成にしました．[^subdomain]
[^subdomain]: [GAE でカスタムドメインにサブドメインを使う場合はちゃんとサブドメインをゾーン管理しなければならない](https://blog.kakakikikeke.com/2019/02/how-to-set-custom-domain-on-gae.html)

新しく記事を書いたらGitHubのmasterブランチにプッシュすれば記事の公開ができます．
masterブランチへのプッシュをトリガーに，サイトのビルドからGAEへのデプロイまでを自動で実行するワークフローを作り，快適な執筆環境を用意しました．

費用をざっくり計算したところ1日1000アクセス以内なら少なく見積もって月0.4ドル，多めに見積もっても1ドルかからない程度なので，secretなキーが漏洩しない限り当面は心配ないと思います．
むしろアクセスが少ない方を心配するべきです．

デザインのテーマは[Hugo Notepadium](https://themes.gohugo.io/hugo-notepadium/)を選択しました．
シンプルなデザインで，特に記事一覧画面とカテゴリやタグの表示が気に入っています．

## 工夫点, 苦労した点

### デプロイ作業

「CUI最高!GUIめんどくさい!」になりつつあるので，プロジェクトの作成からGAEにデプロイするまでの準備を，GCPのコンソールでボタンポチポチせずCloud Shell上でコマンドを叩くだけで完結させることにしました．

ここでは主にGitHub Actionsからのデプロイを許可するための権限周りで苦労しました．

gcloudコマンドやCloud IAMのドキュメント，GCP公式のActionsのリポジトリ，エラーメッセージを見れば，どのAPIを有効化したりロールを付与したりすればいいのか理解できるはずなんですが，そもそもどのドキュメントを読めばいいかがわからなかったのが原因です．

参考にしたドキュメントはこちらです．

- [workflowの書き方と必要なロール](https://github.com/GoogleCloudPlatform/github-actions/tree/master/appengine-deploy)
- [gcloudコマンド](https://cloud.google.com/sdk/gcloud/reference?hl=ja)

プロジェクトの作成からデプロイの準備までの一連の流れは当ブログのリポジトリの[README](https://github.com/Fukkatsuso/blog/blob/master/README.md)に記しました．

ドキュメント読んで実行してエラー読んで修正を繰り返しただけですが，躓いた方の参考になると嬉しいです．

### カスタム404ページ

Hugoはビルドを実行すると，生成したサイトの静的ファイルをpublicディレクトリに出力します．ホスティングする際はリクエストに合わせてpublicディレクトリ内のファイルを提供するよう設定すれば良いです．

GAEは静的サイトのハンドラをapp.yamlに記述できますが，ルーティングの正規表現に合致したパスが指すファイルが存在しない場合は即デフォルトの404エラーを返す仕様になっています．

そのため以下のような記述で，ファイルが見つからないパスを自分で書いたアプリケーションに渡し，カスタマイズした404ページを返す処理を実現しようとすることはできません．

```yaml
# app.yaml
handlers:
- url: /(.*\.css)
  static_files: public/\1
  upload: public/(.*\.css)

- url: /(.*\.(bmp|gif|ico|jpeg|jpg|png))
  static_files: public/\1
  upload: public/(.*\.(bmp|gif|ico|jpeg|jpg|png))

- url: /(.*\.xml)
  static_files: public/\1
  upload: public/(.*\.xml)

- url: /(.+)/
  static_files: public/\1/index.html
  upload: public/(.+)/index.html

- url: /(.+)
  static_files: public/\1/index.html
  upload: public/(.+)/index.html

- url: /
  static_files: public/index.html
  upload: public/index.html

- url: /.*
  script: auto
```

カスタム404ページを返すようにするには，上に書いた処理を自分で実装すれば良いです．Golangでの例:

```go
func staticFileHandler(w http.ResponseWriter, r *http.Request) {
  if _, err := os.Stat("public/" + r.URL.Path); os.IsNotExist(err) {
    http.ServeFile(w, r, "public/404.html")
  } else {
    http.ServeFile(w, r, "public/"+r.URL.Path)
  }
}

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Defaulting to port %s", port)
  }

  http.HandleFunc("/", staticFileHandler)

  log.Printf("Listening on port %s", port)
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal(err)
  }
}
```

参考: [3. Hugoを使いブログをGAEで管理 - Puliyoブログ](https://blog.puliyo.com/jp/posts/hugo-on-gae-3/)

app.yamlのハンドラはこれだけで十分になります．

```yaml
handlers:
- url: /.*
  script: auto
```

ちなみにgo1.6より前の`http.ServeFile`には**ディレクトリトラバーサル脆弱性**なるセキュリティホールがあったそうで，セキュリティの勉強にもなりました．
古いバージョンをお使いの場合はご注意ください．[^serve_file]
[^serve_file]: [net/httpで安全に静的ファイルを返す - Shogo's Blog](https://shogo82148.github.io/blog/2016/04/13/serving-static-files-in-golang/)

## 今後の改善点

- SNSシェアボタン
- 関連記事の表示

などを追加するともっとブログっぽさが出る気がします．

キャッシングについても調べておきたいところです．

## 所感

環境構築から1週間ほどでブログのベースを作ることができました．
ローカルサーバにはホットリロード機能がついているので，編集してセーブすると一瞬でビルドしてくれてページを確認できるのも非常に快適です．

定期的に更新できるように開発の方も力を入れていきたいと思います．
