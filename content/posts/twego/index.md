---
title: "ツイートをリアルタイム検索してターミナルに流せるCLIツール「twego」を開発した"
date: 2022-04-10T17:20:00+09:00
draft: false
summary: "Twitterの Filtered stream API を利用して，ツイートをリアルタイムでターミナルに流してくれるCLIを作りました．"
categories:
  - development
tags:
  - golang
  - cli
---

## 作ったもの

GitHubリポジトリ: <https://github.com/Fukkatsuso/twego>

![twego-demo](https://github.com/Fukkatsuso/twego/blob/main/docs/twego-demo.gif?raw=true)

Twitterの [Filtered stream API](https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/introduction) を利用して，ツイートをリアルタイムでターミナルに流してくれるCLIを作りました．

ツール名は `twego` です．
**tweet** を流す **golang** 製のCLIツールという意味で命名しました．

機能はシンプルで以下の通りです．

- `twego auth`: Twitter Bearer Token をCLIツールに渡す．APIを叩くのに必要な認証で使う
- `twego rules add|delete`: ツイートの検索ルールを設定する
- `twego stream`: リアルタイムでルールにマッチするツイートを時刻とともに出力する

## 背景

twegoを作った背景として，スポーツ観戦や音楽ライブ鑑賞などの最中にTwitterで他の人のリアルタイムな反応を見ることが多かったのですが，「Twitterアプリを起動」→「キーワードを設定して検索」→「最新の検索結果をスクロールして読んでは再読み込みを何度も繰り返す」という手間のかかる手順が面倒になってきていました．

凝ったデザインは必要なく，CLIツールの開発が未経験だったこともあり，極力シンプルな機能に抑えたCLIツールとしてサクッと開発してみることにしました．

## 技術選定

ある程度馴染みのあるGolangで作るということは決めていました．
CLIツール開発は初めてだったので，できるだけ「お作法」に従ったほうが色々と楽だろうと考え，一番使われていそうなフレームワークとして [cobra](https://github.com/spf13/cobra) を利用することにしました．

また今回のCLIツールはホスト上でビルドしてバイナリを動かす形で利用する予定でしたが，開発時になるべくローカル環境を汚したり変更したりしたくないという理由から，Dockerコンテナ上で開発しました．
最終的にはホスト上でもDockerコンテナ上でも動かせるようにしています．

## 工夫点など

最近「Go言語による並行処理」という本を読んで並行処理と少しだけ仲良くなっていた矢先のことだったので，ツイートを受信しつつ逐一出力するのに `channel` を活用してみました．

```go
// ツイートの受信
func GetTweetStream(done <-chan struct{}, bearerToken string) <-chan Tweet {
  stream := make(chan Tweet)

  go func() {
    defer close(stream)

    // ...

    decoder := json.NewDecoder(resp.Body)
    for {
      decoded := make(chan Tweet)
      go func() {
        defer close(decoded)
        var tweet Tweet
        if err := decoder.Decode(&tweet); err != nil {
          fmt.Println(err)
          return
        }
        decoded <- tweet
      }()

      select {
      case <-done:
        return
      case tweet, ok := <-decoded:
        if !ok {
          fmt.Println("cannot read the decoded response")
          return
        }
        select {
        case <-done:
          return
        case stream <- tweet:
        }
      }
    }
	}()

	return stream
}
```

```go
// ツイートの出力
done := make(chan struct{})
defer close(done)

stream := GetTweetStream(done, bearerToken)

w := tabwriter.NewWriter(os.Stdout, 0, 2, 0, ' ', 0)
for {
  select {
  case tweet, ok := <-stream:
    if !ok {
      return errors.New("stream is closed")
    }

    now := time.Now().Format("2006/01/02 15:04:05")
    texts := strings.Split(tweet.Data.Text, "\n")
    for i, text := range texts {
      if i == 0 {
        fmt.Fprintln(w, now, "\t", text)
      } else {
        fmt.Fprintln(w, "\t", text)
      }
    }
    w.Flush()
  }
}
```

検索ルールやツイートの出力の際は `text/tabwriter` を使って見た目をキレイに揃えるようにして，使い心地アップを図っています．

```bash
$ twego rules list
ID                    VALUE                 TAG
1234567890123456789   twitter -is:retweet   twitter
0123456789012345678   golang -is:retweet    golang

$ twego stream
2022/04/08 19:40:31  Lorem ipsum ...
2022/04/08 19:40:41  Ad qui ...
                     Usu zril ...
```

仕上げにドキュメントとhelpの体裁を整え，Dockerでも動かせることとシンプルで使いやすいことをアピールしました（伝わるかは別）．

## まとめ

実際にPCで作業しながらtwegoでスポーツの試合進行をチェックしてみた感触も，なかなか良かったです．
プロ野球観戦に使おうかなと考えています．

CLIツール開発に対してはフラグ管理など大変そうなイメージがあったのですが，フレームワークのおかげで簡単に開発できました．
これで開発の選択肢も1つ増えましたし，また何かのCLIツールを作ってみたいと思います．

完全に自分向けとして作ったツールですが，お試しでも良いので使っていただけると嬉しいです．
