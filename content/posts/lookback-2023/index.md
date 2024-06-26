---
title: "2023 年の振り返りと来年の目標"
date: 2023-12-31T22:00:00+09:00
draft: false
summary: "2023 年の目標を振り返り、2024 年の目標を立ててみました。"
images:
  - posts/lookback-2023/og-image.png
categories:
  - log
tags:
  - lookback
---

2023 年の目標を振り返り、2024 年の目標を立ててみました。

## 目標の達成度

達成した目標にチェックマークを付けています。

- [ ] 新規にアプリケーションを開発する
  - 誰かに使ってもらい、感想を聞く
- [ ] 「クリーンアーキテクチャ」を読む
- [x] インフラ方面の知識を得る
  - AWS, GCP の活用法
  - kubernetes
- [ ] 新しいプログラミング言語を習得する
  - Rust など
  - サーバーサイドで使えるものだと良し
- [x] ISUCON に挑戦する
- [x] 資産運用を始める
  - 積立 NISA
  - iDeCo

今年の目標は半分達成できたかどうか、くらいでした。

一番達成できたと言えるのは「ISUCON に挑戦する」という目標です。
詳しくは [参加記](https://blog.fukkatsuso.com/posts/isucon13) に書きましたが、23 卒の同期とチームを組んで半年間、毎週末 ISUCON 練習に勤しむほど熱中していました。

その反面、個人開発にほぼ時間を割かなくなり、それ関連の目標は全く着手していません…

「資産運用を始める」「インフラ方面の知識を得る」は、お給料の一部を積立 NISA に回すようにしたのと、配属先がインフラの部署になったのとで、人生の流れに乗って達成できたと思います。

「新しいプログラミング言語を習得する」は仕事中に Perl で書かれたスクリプトを読んだくらいで、習得できたかどうかは微妙なところです。

## 2023 年の出来事

今年の大きな変化としては、大学院を修了して新社会人になったことでした。

### 大学院修了

大学院での研究はずっと苦しかった記憶しかなく、3 月に無事修了できたことだけでホッとしました。

立派といえる研究もできていないですが、その反省として、もし研究室に入る前の自分にアドバイスするとしたら「夢中になって論文を読めるような分野を選べ」「論文を読んで、その分野の知識と研究の仕方を学べ」ということでしょうか。

これができたらもう何も心配いらないです。
研究を進めるうえでどんな情報が必要かを理解して、夢中になって情報収集できれば、苦しいことも引きずらず前向きに研究に取り組める 2 年間が得られていたと思います。

### 福岡から東京へ引っ越し

就職のため引っ越しました。

初めての引っ越しだったというのもありますが、将来の引っ越しも考え知見として貯めておくためにも、
お部屋探しから諸々の手続き、移動、掃除、家具・家電購入、荷造り、荷解きまでの全てを Notion で管理して、万全の準備を整え実行しました。

そのおかげで引っ越しの段取りはほぼ完璧に実行できたと思います。
が、問題は新居に住み始めてからで、マンションの隣人がゲームをして叫ぶ声が平日の夜や週末に響く始末で、ハズレくじを引かされました…

そんな部屋を選んでしまった原因はいくつかあります（以下恨み節）.
内見が昼間の 1 時間ほどで隣人が静かな時間帯だったため、隣人問題を全く意識しなかったこと。
築 1 年で敷金・礼金ゼロという好条件に、一度は「裏があるんじゃないか」と疑いながらも飛びついてしまったこと。
そしてお部屋探しに費やしたのはたった 1 日と、じっくり内見・熟考できる時間的余裕を作っていなかったことです。

内見の日、不動産の営業の人に、金額的な条件の良さに対する疑念を投げかけたところ、「大手の管理会社が広告費として条件を良くしている」などという答えが返ってきたのをよく覚えています。
愚かな自分はそれを信用してしまい、また部屋の中だけ見ればとても良い感じだったので、疑いを捨てて契約まで至りました。

次引っ越すときは敷金・礼金くらいでケチらず、同じ部屋でも複数回内見して隣人問題を回避します！!!

### 社会人の仲間入り

そんな苦労を経て、4 月から晴れてエンジニアとして DeNA に新卒入社しました。

（参考：[23 新卒ソフトウェアエンジニア職の就活を振り返る](https://blog.fukkatsuso.com/posts/job-hunting-newgrad)）

入社から 2 週間は全体研修で会社のことや社会人としての基礎スキルを学び、それから 3 ヶ月にわたってエンジニア研修を受けました。

エンジニア研修では、変化し続ける業界で変わらないものを学び、変わるものを学び続けるエンジニアになることをテーマに、カリキュラムが組まれていました。

**変わらないもの**としては、コンピュータサイエンスに関して特にアルゴリズムを、講義＆演習の形式で学びました。
大学で CS を専攻していたり、競技プログラミングを経験していた人にとっては既知の内容が多かったです。
ただその点に関しては、演習で経験者と未経験者がバランスよくチーム分けされ、互いに教え合うことで自分たちなりの答えを出すという形式だったおかげで、教える側になっても退屈ではなく、どの程度・どのような伝え方で伝えるかを考えながら取り組む機会になったと思います。

**変わるもの**としては、変化が大きい Web アプリケーション開発に関して、チーム別の開発を通じて学びました。
5 月は Golang と React を用いたサンプルアプリケーションを開発し、6 月は「デザインスプリント」というプロトタイピング手法で考えたアプリケーションをスクラムで開発しました。
技術的に学ぶことはそれほど多くなかったですが、チーム開発の経験はとても濃く、総合的な学びは大きかったです（研修の意図を強く感じました）.

そして 7 月の上旬に配属発表があり、自分は IT 基盤部という部署で**インフラエンジニア (SRE)** としてキャリアをスタートすることが決まりました。

正直に言うと、当時の第一志望はアプリケーション開発側の部署だったので、配属発表の瞬間は気持ちが落ち込みました。
そうは言っても第二志望の部署でしたし、自分を客観的に見ると SRE の方が適正があり、配属先のチームは元々興味のある大規模ゲーム領域のインフラを担当しているところだったので、今となっては配属結果に満足しています。

12 月現在、配属されて半年ほど経ったところですが、特に大きな不満もなく毎日前向きに仕事ができています。
毎週末の ISUCON 練習が偶然仕事と繋がっており、週末遊んでいるうちにも自然とスキルアップできるのでとても楽しいです。
ソフトスキルの面でもアウトプット力や人に頼る力が向上している実感があり、順調に成長できている（はず）と思います。

ただ、インフラ運用の全体像を理解するにはまだまだ時間がかかりそうです。
「所属チームでインフラエンジニアとして一人前になるには 3 年かかる」と言われているのですが、その通りです。
早くても 2 年ですね。
それでも決して悪い意味に捉えておらず、じっくり理解しながら進めることを重要視している部署なので、そういうものだと思って特に不満はありません。

## 今年買ってよかったもの

新社会人の新生活ということで、大きめの買い物をちょこちょこやっていました。

- SwitchBot カーテン
- Eufy Clean G40+
- FlexiSpot E7 Pro
- FlexiSpot モニターアーム F8LD
- サリダチェア YL8
- Acer ゲーミングモニター
- GOKUMIN プレミアムスプリングマットレス
- ブレインスリープ コンフォーター パーフェクト ウォーム
- HHKB Professional HYBRID Type-S
- 電動コーヒーミル

カーテンの自動開閉とロボット掃除機を導入し、生活環境を快適にしました。

FlexiSpot のデスクは、仕事がリモートワークメインのため、昇降デスクで仕事ができると便利そうだということで購入しました。
長いミーティングのときには立って参加するようにしており、それなりに活躍しています。

ブレインスリープは職場の先輩にオススメされて買ってみたのですが、これ一枚で冬を越せる、というか他の掛け布団が要らなくなるくらい、暖かくて蒸れない最高の掛け布団でした。
3 万円以上しますがとても良い買い物でした。

HHKB は Amazon のブラックフライデーで買ってみました。
MacBook のキーボードとは配列やキーストロークが異なり、慣れるまで時間がかかりましたが、慣れてからは「タイピングが楽しくなるキーボード」という評価です。

最後に書いたコーヒーミルは、仕事前や休憩中にコーヒーの風味をもっと楽しみたくなり、豆から挽こうと思い立って手頃なものを購入しました。
以前はドリップバッグで淹れていましたが、やはり豆から挽くと香りと味がしっかり感じられて良いですね。

## 娯楽

### ゲーム

「ゼルダの伝説 ティアーズ オブ ザ キングダム」をプレイしました。

実はまだクリアしておらず、ポケモン SV の追加コンテンツも未消化と、今年はあまりゲームをしなかった年でした。

### アニメ

観た中で特に良かったな〜という作品です。

- 進撃の巨人 The Final Season 完結編
- Dr.STONE NEW WORLD
- 呪術廻戦第 2 期
- 無職転生Ⅱ
- ウマ娘　プリティーダービー Season 3
- 陰の実力者になりたくて！　2nd season

### 映画

マッドハウス 50 周年と作品公開 25 周年の記念として、「パーフェクトブルー」の 4K リマスター版が 9 月に上映されていたので、映画館に行って観てみました。

今敏監督の「千年女優」「パプリカ」「妄想代理人」は観たことがあり、この作品は初めてだったのですが、あの空想と現実が入り乱れる展開はやはり混乱させられます。
映画が終わってもなお整理がつかない感覚を、映画館で体感できて面白かったです。

## 2023 年の総括

学生から社会人になり、エンジニアキャリアを順調に走り出した 1 年でした。

個人開発は全くと言っていいほど取り組みませんでしたが、それは ISUCON に夢中になった結果なので問題ありません。

ただ物足りなさはあります。
来年はまずインプットから増やしていきたいですね。

## 2024 年の目標

今の調子だと、完全なプライベートの時間はあまり作れないだろうと予想しています。
その中で達成できそうなことを挙げてみました。

- 技術書を 3 冊読む
- プライベートで新しいプログラミング言語（Rust など）に触れる
- ISUCON に出場して入賞する
- iDeCo を始める

ISUCON は本当に面白いので力を入れていきたいです。
そして「インプットの習慣化」を来年の大きなテーマとして頑張っていきます。
