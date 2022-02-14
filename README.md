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
# 実行例
## 全チェック
goavlはカレントディレクトリ以下にあるdesignパッケージ（goファイル）を抽出し、そのファイルに対してチェックを行います。より詳細な例が知りたい方は、[example.md](./doc/example.md)をご確認ください。
```
$ goavl 
[NC001] test/sample/action.go:7    Resource("operandsNG") is not snake case ('operands_ng')
[NC002] test/sample/action.go:9    Action("add-ng") is not snake case ('add_ng')
[UF002] test/sample/action.go:11   Not exist Description() in Action().
[NC004] test/sample/attribute.go:15   Attribute("this-is-ng") is not snake case ('this_is_ng')
[NC004] test/sample/attribute.go:16   Attribute("NgCase") is not snake case ('ng_case')
[NC004] test/sample/attribute.go:17   Attribute("ngCase") is not snake case ('ng_case')
[UF001] test/sample/attribute.go:15   Not exist Example() in Attribute().
[FC023] test/sample/bug.go:10   Attributes() has View(). View() can be used in MediaType, Response
```
## ファイル1件のチェック
-f（--file）オプションを使うと、チェック対象のファイルを指定できます。複数ファイルの指定は、出来ません。
```
$ goavl --file test/sample/goa.go
```
## チェック項目の確認
サブコマンドlistを使用すると、チェック項目一覧が表示されます。表示は、"Inspection ID:チェック内容"の形式です。
```
$ goavl list
NC001: Resource() argument name checker
NC002: Action() argument name checker
NC003: Routing() argument name checker
NC004: Attribute() variable and argument name checker
UF001: Checker whether the example of Attribute() is written
UF002: Check whether Description() is written
FC001: Attribute can be used in: View, Type, Attribute, Attributes
FC002: Default can be used in: Attribute
FC003: Enum can be used in: Attribute, Header, Param, HashOf, ArrayOf
FC004: Example can be used in: Attribute, Header, Param, HashOf, ArrayOf
FC005: Format can be used in: Attribute, Header, Param, HashOf, ArrayOf
(省略)
```

## チェック項目の除外
指摘を許容（無視）したい場合は、--excludeオプションを使用してください。カンマ区切りで、除外したい"Inspection ID"を複数指定できます。
```
$ goavl --exclude=FC001,NC003
```

# ライセンス
goavlプロジェクトは、複合ライセンスです。
- MITライセンス（[casee*.go](./internal/utils/strutils/casee.go) および [camelcase*.go](./internal/utils/strutils/camelcase.go)）
- [Apache License Version 2.0](./LICENSE)（上記以外のコード全て）

MITライセンスのソースコードの作者は、[pinzolo氏](https://github.com/pinzolo)および[Fatih Arslan氏](https://github.com/fatih)です。それぞれの作者が書かれたコードには、MITライセンス全文およびCopyrightが明示されています。

# 名前の由来
初期名称は"goalinter-v1"。linterのVersion 1に見えるため、"goav1linter"に改名。その後、"1"と"l"が似ていたため、それらを統合して名前を短くしました（= "goavl"）