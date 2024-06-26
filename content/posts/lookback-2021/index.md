---
title: "2021 年の振り返りと来年の目標"
date: 2021-12-31T14:41:23+09:00
draft: false
summary: "2021 年の目標と活動を振り返り、2022 年の目標を立ててみました。"
images:
  - posts/lookback-2021/og-image.png
categories:
  - log
tags:
  - lookback
---

2021 年の目標と活動を振り返り、2022 年の目標を立ててみました。

## 目標の達成度

去年の年末には、こんな目標を掲げていました。

- 新規に Web アプリを開発する
- クリーンアーキテクチャ、ドメイン駆動設計、アジャイル開発など開発手法に関する本を読む
- AtCoder のレート 1200 以上（水色）へ
- M1 の夏はインターンに参加する

1 個ずつ振り返っていきます。

### 新規に Web アプリを開発する

Web アプリと言えるかは微妙なところですが、個人で [仮想通貨の自動取引 bot](https://github.com/Fukkatsuso/cryptocurrency-trading-bot) を開発しました。

昨年少し勉強した GCP やテストの経験を活用して、開発しやすいプロダクトを作ることを目標に、CI/CD やドキュメント整備などに力を入れて作っていきました。

4 月から現在まで長いこと開発が続いていて、取引 bot としての機能は十分満たしていますが、bot を管理するための管理画面の整備は完了していません。
なので目処が立ったタイミングでブログ記事にしようと考えています。

### 開発手法に関する本を読む

ドメイン駆動設計の本を読みました。
振り返ってみると、ドメイン駆動設計そのものよりも、レイヤードアーキテクチャや依存関係逆転の原則についての知識を得たと言うべきかもしれません。
こういうのは今すぐ実践できるものでもないですし、話題に挙がったときの共通言語というか引き出しとして持っておくくらいの意識で勉強しています。

あと Kindle（端末ではなくアプリ）デビューしました。
PC でも読めて省スペースなところが最高ですね。
今はクリーンアーキテクチャの本を積んでいて、春休みとか就活終わった頃に消化したいです。

### AtCoder のレート 1200 以上（水色）へ

3 月いっぱいで競プロを打ち切りにした関係で達成できませんでした。

それまで精進もやるにはやっていましたが、実際のところは卒論との兼ね合いで十分な時間が取れず、1 日 1AC を簡単な問題でばかり済ませてしまっていました。

![atcoder-heatmap](atcoder-heatmap.png)

卒論を終えてからも、緑 diff や水色 diff で肩慣らしして 3 月上旬のコンテストに出たのですが、レートが下がってしまう始末で…

精進にしてもコンテストにしてもこれ以上続けたところで今の自分では伸びそうにないなと感じ、個人開発の時間に充てたほうがよっぽど有意義だということで止めました。
いつかまた、やる気が出てきたら水色を目指します。

### M1 の夏はインターンに参加する

8 社ほど選考を受けましたが残念ながら夏季インターンには参加できず、その分長期インターンにて開発を頑張りました。

早いところで 3 月から選考が始まるのに初動が 5 月の 1on1 面談からで、本格的に選考に進み始めたのが 6 月という手遅れさが原因の 1 つでした。
あとは面接でのアピールが下手ですね。

ただ、大学院が忙しかったこともあり、選考を受けるにつれて「無理にインターンに参加せず、自分のやりたいことを思いっきりやる方向でも全然悪くない」と気持ちに変化が生まれ、そういう意味では自分を振り返るきっかけになったので良かったのかなと思っています。

インターン参加は叶いませんでしたが、選考を受けたおかげでハッカソンに招待いただいたり本選考にスムーズに進めたりとメリットもたくさんあったので、ダメ元で挑戦していくのも悪くないと思えました。

…そんな感じで、今年の目標は半分達成できました。
自分一人で実現できる目標は達成して、誰かと関わったり比較したりするような目標は達成せずということでしょうか。

## 個人開発

今年個人で取り組んだのはこれらになります。

- [仮想通貨取引 bot](https://github.com/Fukkatsuso/cryptocurrency-trading-bot) の開発
- [メシガイド（飲食店検索 LINE bot）](https://github.com/Fukkatsuso/linebot-restaurant-go) の改善

今年は 4 分の 3 くらい仮想通貨取引 bot の開発に費やしました。
予定よりずっと時間がかかってしまいましたが、実現したいことは概ねできているので OK です。
メシガイドの方はリファクタリングしたり、複数画像の GET を Go の並行処理で高速化したりと結構楽しみながら改善を加えていけたと思います。

数としては少ないですが、人に語れるくらいモノを作り込むことって大事なんだなと感じた経験（技育祭とか就活とか）を機に、腰を据えてじっくりやっていった結果そうなっているだけです。

![github-heatmap](github-heatmap.png)

GitHub のコントリビュート数は、去年の 850 から微増の 897 でした。
main ブランチに反映していない分を考慮しても大体 900 です。

できるだけ毎日草を生やすという目標は約 92%達成できました。
来年も継続していきますが、今度は質の部分をレベルアップさせていきたいと思います。

## 長期インターン

株式会社 AIoT さんにて昨年 11 月に始めた長期インターンですが、就活や大学院での忙しさを考慮して 9 月をもって退職しました。

インターンでは主にバックエンドを担当して CRUD API を開発したり、HTML/CSS による帳票 pdf の作成などのフロント方面も触らせてもらったりしました。
言語・フレームワークとしては Go や NestJS をメインに使い、他には Vue.js も少し使っていました。
他にも実務を通じて DB のマイグレーションや git の使い方なども勉強させてもらいました。

そんな感じで色んな技術に触れつつ実務経験ができたのは良かったと思っています。

心残りとしては、フルリモートなこともあってチームで連携しながらの開発経験はあまりできなかったこと、新規開発がほとんどで改善に挑戦する機会が少なかったことがあります。
これから学生のうちにチーム開発できるかは不明ですが、工夫・改善なら個人開発でもできそうなので意識してやっていくつもりです。

## 大学院、就活

4 月から大学院へ進学し、現在 23 卒として就活しています。

大学院では機械学習やソフトウェア開発について勉強しています。
研究の方はというと、分野的には自然言語処理が近いですが、教授や先輩と研究テーマについて議論するということが多かったです（あまり進んでいない）.

就活では IT 業界、特に Web 系企業の選考を受けていて、割合としてはゲーム・エンタメ関連の事業をやっている企業が多いです。
受ける企業はある程度絞り、集中して選考に臨もうというスタンスでやっています。
本格的に本選考に進み始めたのは 11 月頃からで、できれば M2 に進む前の 2 月・3 月までには就活を終わらせたいなと考えています。

## 2021 年の総括

今年は大学院に進んで途中まで長期インターンと両立し、空き時間に個人開発をチマチマ進めつつ本を少し読み、就活も始めるといった年でした。

長期間の活動が多かったせいか、始めたことよりも止めたことの方が多い 1 年だった印象です。

来年は多分、これまで色々手を伸ばしてきたのを大学院修了と就職に向けて収束させていく 1 年なんだろうなと考えています。

## 来年の目標

就活とか研究は「やらなきゃいけない」部類なので目標からは外し、エンジニア関連での目標を立てることにします。

全部できるかはともかく、やりたいことをリストアップしてみました。

- 仮想通貨取引 bot をひとまず完成させる
  - ブログも書く
  - できれば収益を出す
- 新規に Web アプリを開発する
  - 自分が本気で使いたいと思えることはもちろん、誰かに使ってもらえるようなサービスを作りたい
  - CLI ツール等でも可
- 開発手法に関する本を読む
  - クリーンアーキテクチャなど
- 並行処理について勉強する
  - 「Go 言語による並行処理」を読みたい
- インフラ方面の知識を得る
  - AWS, GCP の活用法
  - Terraform
  - kubernetes
- フロントエンドの技術も触ってみる
  - React, Next.js がスタンダードっぽい雰囲気
- イーサリアムについて知る
  - [この本](https://www.amazon.co.jp/dp/4873118964) とか読んでみたい

「大学院修了と就職に向けて収束させていく」と言った割に多方面に手を伸ばす感じになりましたが、新しい知識の吸収は積極的にやっていきたいです。

最低でも目標の 50%は達成できるよう頑張ります。
