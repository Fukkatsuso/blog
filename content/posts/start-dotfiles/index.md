---
title: "dotfilesの作成が一段落したのでまとめる"
date: 2020-09-08T22:20:38+09:00
draft: false
summary: "bashやgitの設定，macの初期設定などをまとめた，いわゆるdotfilesを作りました．"
categories:
  - development
tags:
  - dotfiles
  - GitHubActions
---

bashやgitの設定，macの初期設定などをまとめた，いわゆる[dotfiles](https://github.com/Fukkatsuso/dotfiles)を作りました．
ついでにGitHub Actionsでテストなどを実行するCIも構築してみました．

## 動機
dotfilesを作る理由なんて限られてますが，初期設定の度に前の環境を思い出してポチポチするのが手間だからというのと，設定のコード管理ができるからです．
つまり作りたいから作った感じです．

## 開発環境
使用しているマシンとツールを先に示しておきます．
- macbook pro (2019, macOS Mojave)
- Git
- Bash
- iTerm2
- VSCode

zshとかfishとか高度そうなシェルは使わず，まずは基本を抑えてからがいいと考えてBashにしています．
macOS Catalinaからデフォルトシェルがzshになったという噂は聞いていますが，そのときはそのときで対応するつもりです．

デフォルトのターミナルはデフォルトのものからiTerm2になんとなく変えてみました．
コマンドを叩くときは大抵エディタでプログラムを書きますし，コーディングに愛用しているVSCodeであればターミナルも同じ画面で使用できるのでiTerm2単体ではほとんど使っていません．

## ファイル構成
現在のメインマシンであるmacでの使用を前提に作っています．
dotfilesの中身はこんな感じです．
```sh
.
├── .bash_profile
├── .bashrc
├── .gitconfig
├── iterm2
│   ├── com.googlecode.iterm2.plist # iTerm2の設定
│   ├── iceberg.itermcolors # カラーテーマ
│   └── init.sh
├── vscode
│   ├── init.sh
│   ├── install.sh
│   ├── settings.json # 言語やエディタの設定
│   ├── snippets # スニペット
│   │   ├── go.json
│   │   ├── go.mod.json
│   │   └── latex.json
│   └── vscode_extensions # 拡張機能のリスト
├── init.sh
├── install.sh
├── macos.sh
└── .github
    └── workflows
        ├── lint.yml
        └── macos.yml
```

`.bash_profile`, `.bashrc`はBashの設定ファイル，`.gitconfig`はgitの設定ファイルです．
ホームディレクトリにシンボリックリンクを貼ることで，dotfilesのリポジトリ内でファイルを管理しつつ，システムから設定を参照することができます．

シンボリックリンクを貼る処理は`init.sh`に書いており，運用の際はこれを一番始めに実行することにします．
```sh
# 例: .bash_profileのシンボリックリンク
CURRENT_DIR=$(cd "$(dirname "$0")"; pwd)
ln -fs "$CURRENT_DIR/.bash_profile" ~
```

`install.sh`は主にアプリケーションのインストールを実行します．
基本的にHomebrewでインストールしています．

`macos.sh`はmacの初期設定用のスクリプトで，Dock，Finder，Safari，トラックパッドなどの設定を記述しています．

`iterm2`, `vscode`ディレクトリはその名の通りで特定のアプリケーションの設定を行います．
これも同様に`init.sh`と`install.sh`が含まれており，リポジトリのルートにあるものとは独立させています．

## 使い方
### 実行
`init.sh`, `install.sh`の順に実行するだけです．
```sh
bash init.sh
bash install.sh
```

`vscode`や`iterm2`内のファイルも同様です．
```sh
bash vscode/init.sh
bash vscode/install.sh
```
```sh
bash iterm2/init.sh
```

注意点として，`. init.sh`や`source init.sh`のような実行方法だと，カレントディレクトリのパスを正しく取得できなくなるため失敗します．
上記のようにbashコマンドの引数にスクリプトを指定する方法か，実行権限を付与して実行する方法でないとうまくいきません．
僕はこれで1時間潰しました...
- 参考: [bashシェルスクリプトの実行方法・実行環境](https://www.kenschool.jp/blog/?p=4499)

### VSCodeの拡張機能
既にVSCodeを利用していてインストール済みの拡張機能をリストアップ・記録する場合は，次のようなコマンドを実行します．
```sh
code --list-extensions > vscode/vscode_extensions
```

拡張機能のインストールは`vscode/vscode_extensions`のファイルを参照して行います．

今見てみると，vscodeでディレクトリを切っているのでファイル名は`vscode_extensions`ではなく`extensions`の方がスマートですね．

## GitHub Actionsでスクリプトの文法チェックとテスト
GitHubのmasterブランチへのプッシュをトリガーとして，スクリプトの文法チェックとmacOS上でのテストを実行するワークフローを作りました．

```yml
# .github/workflows/lint.yml
name: lint

on:
  push:
    branches:
      - master
    paths:
      - '**.sh'

jobs:
  lint:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v2

      - name: Run ShellCheck
        uses: reviewdog/action-shellcheck@v1
        env:
          shellcheck_flags: --exclude=SC1090
```
shellcheckというツールでスクリプトの文法チェックを行うアクション(reviewdog/action-shellcheck@v1)を使わせてもらい，ファイル名が`.sh`で終わるファイルに変更があった際にチェックを走らせるようにしました．

```yml
# .github/workflows/macos.yml
name: macOS

on:
  push:
    branches:
      - master
    paths:
      - '**'
      - '!README.md'
      - '!.gitignore'

jobs:
  macos-test:
    runs-on: macos-latest

    steps:
      - uses: actions/checkout@v2
      - run: bash init.sh
      - run: bash install.sh
      - run: bash iterm2/init.sh
      - run: bash vscode/init.sh
      - run: bash vscode/install.sh
      - run: bash macos.sh
```
`macos.yml`では，一連のスクリプトを実行してmacOS上での動作確認を取るようにしています．

## その他自動化できなかった/しなかった項目
できると嬉しいが急を要さないもの，検索しても見当たらなかった設定項目などを今後の改善点として．

### google IMEの設定
macのセットアップに関する記事を見ると，日本語入力システムには大抵google IMEを使っているようでした．
macデフォルトのと比べてもgoogle IMEの方が使いやすいと感じたのでこれを使っています．

日本語入力の環境設定（句読点を"，．"にして，記号や数字の半角・全角の設定）がGUIでできるのはもちろんですが，これをCLIで設定する方法が見つからなかったのでTODOリスト入りとなっています．

### TouchBarの表示をファンクションキーにする
2019年のmacbook proには本来ファンクションキーがある場所にTouchBarというインターフェースが付いており，ソフトウェアキーが配置できるスペースがあります．

デフォルトだとここに音量や画面輝度の調節キーがあるのですが，見慣れないキーが気になって遊んでしまいそうになりますし，なんだかんだ一番しっくりくるファンクションキーを表示する設定にしています．

これをCLIで設定する方法もパッと見つけられず，できませんでした．
よく探せば見つかりそうな気がする設定項目No.1ではあります．

### セットアップのワンライナー化
現状だとスクリプトを1つずつ指定して実行するフローになっていますが，OS，目的ごとに，たった1つのコマンドで必要なスクリプト全てを実行できたら素晴らしいと思います．

makefileを書いたり全体の統括的なスクリプトを用意したりして実現する案と，依存関係のシンプルさを重視してワンライナー化は諦める方針とで考え中です．

### pmset
`pmset`コマンドでmacの電源管理設定ができます．

スリープまでの時間，ディスプレイをオフにするまでの時間，スタンバイモードの有効化/無効化などなど，バッテリーの持ちに関わる部分なので，普段の使用状況を鑑みながら最適化していきたいところです．

## まとめ
エイリアスやキーボード設定でオレオレ仕様にするのは好みではないので必要最低限の設定+αという感じです．

早く実機で試してみたい...
