[![Build](https://github.com/nao1215/goavl/actions/workflows/build.yml/badge.svg)](https://github.com/nao1215/goavl/actions/workflows/build.yml)
[![UnitTest](https://github.com/nao1215/goavl/actions/workflows/unit_test.yml/badge.svg)](https://github.com/nao1215/goavl/actions/workflows/unit_test.yml)

# goavl: Goa framework (ver1) linter
goavlは、[goa version1（フォーク版）](https://github.com/shogo82148/goa-v1)のlinterです。開発目的は、goaを用いたWeb APIデザイン設計速度を早め、チーム内でのデザイン差異を極力小さくする事です。

goaは、DSLで記述されたデザインをもとに、Web APIホスティングに必要なベース処理（ルーティング、コントローラ、Swaggerなど）を生成します。しかしながら、課題が1点あります。それは、デザインから各種ファイルの生成に失敗した場合、エラー箇所を即座に特定できない点です。  

例えば、近年のコンパイラはエラー箇所の行数を示します。しかし、goaはそのような情報を一切出力しません。そこで、goaデザインファイルの問題箇所を即座に特定するためのLinter開発に着手しました。

goavlは、**goa version 1 のみをサポート**し、現行のversion 3 はサポートしない予定です。その理由は、フォークしたgoa version 1 を私が利用しているからです。何故、フォーク版を使用しているかの背景は、[他のサイト](https://furusax0621.hatenablog.com/entry/2021/12/13/000000)で説明しています（記事の作者とgoavl開発者は別人のため注意）


# インストール方法
### Step.1 前準備
現在は、" $ go install"によるインストールのみをサポートしています。そのため、golangの開発環境をシステムにインストールしていない場合、[golang公式サイト](https://go.dev/doc/install)からgolangをインストールしてください。

### Step2. インストール
```
$ go install github.com/nao1215/goavl@latest
```
### 実行例
goavlはカレントディレクトリ以下にあるdesignパッケージ（goファイル）を抽出し、そのファイルに対してチェックを行います。より詳細な例が知りたい方は、[example.md](./doc/example.md)をご確認ください。
```
$ goavl 
[WARN] test/sample/goa.go:8    Resource("operandsNG") is not snake case ('operands_ng')
[WARN] test/sample/goa.go:9    Action("add-Ng") is not snake case ('add_ng')
[WARN] test/sample/goa.go:10   Routing(GET("add_ng/:left/:right")) is not chain case ('add-ng/:left/:right')
[WARN] test/sample/goa.go:22   Attribute("AbcDefID") is not snake case ('abc_def_id')
[WARN] test/sample/goa.go:23   Attribute("zzzXXX-ss") is not snake case ('zzz_xxx_ss')
[WARN] test/sample/goa.go:24   NoExample() in Attribute(). NoExample() is not user(client) friendly
[WARN] test/sample/goa.go:26   Not exist Example() in Attribute().
```

# ライセンス
goavlプロジェクトは、複合ライセンスです。
- MITライセンス（[casee*.go](./internal/utils/strutils/casee.go) および [camelcase*.go](./internal/utils/strutils/camelcase.go)）
- [Apache License Version 2.0](./LICENSE)（上記以外のコード全て）

MITライセンスのソースコードの作者は、[pinzolo氏](https://github.com/pinzolo)および[Fatih Arslan氏](https://github.com/fatih)です。それぞれの作者が書かれたコードには、MITライセンス全文およびCopyrightが明示されています。

# 名前の由来
初期名称は"goalinter-v1"。linterのVersion 1に見えるため、"goav1linter"に改名。その後、"1"と"l"が似ていたため、それらを統合して名前を短くしました（= "goavl"）