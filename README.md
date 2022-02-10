# goavl: Goa framework (ver1) linter
goavlは、[goa version1（フォーク版）](https://github.com/shogo82148/goa-v1)のlinterです。開発目的は、goaを用いたWeb APIデザイン設計速度を早め、チーム内でのデザイン差異を極力小さくする事です。

goaは、DSLに記述されたデザインをもとに、Web APIホスティングに必要なベース処理（ルーティング、コントローラ、Swaggerなど）を生成します。しかしながら、課題が1点あります。それは、デザインから各種ファイルの生成に失敗した場合、エラー箇所を即座に特定できない点です。  

例えば、近年のコンパイラはエラー箇所の行数を示します。しかし、goaはそのような情報を一切出力しません。そこで、goaデザインファイルの問題箇所を即座に特定するためのLinter開発に着手しました。

goavlは、**goa version 1 のみをサポート**し、現行のversion 3 はサポートしない予定です。その理由は、私がフォークしたgoa version 1 を利用しているからです。何故、フォーク版を使用しているかの背景は、[他のサイト](https://furusax0621.hatenablog.com/entry/2021/12/13/000000)で説明しています（記事の作者とgoavl開発者は別人のため注意）


# インストール方法
### Step.1 前準備
現在は、" $ go install"によるインストールのみをサポートしています。そのため、golangの開発環境をシステムにインストールしていない場合、[golang公式サイト](https://go.dev/doc/install)からgolangをインストールしてください。

### Step2. インストール
```
$ go install github.com/nao1215/goavl@latest
```

# 開発進捗
### 作成完了：命名規則チェック（チェック対象一覧）
- Resource()の引数
- Action()の引数
- Routing()の引数

### 作成予定
- Type()の変数名、引数の命名規則チェック
- MediaType()の変数名、引数の命名規則チェック
- View()の構文チェック（使用できない関数内での呼び出しがないかどうか）

### 実行例
```
$ cat test/sample/goa.go   ※ チェック対象のファイルの中身を表示
package design

import (
        . "github.com/shogo82148/goa-v1/design"
        . "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = Resource("operandsNG", func() {
        Action("add-Ng", func() {
                Routing(GET("add_ng/:left/:right"))
                Description("add returns the sum of the left and right parameters in the response body")
                Params(func() {
                        Param("left", Integer, "Left operand")
                        Param("right", Integer, "Right operand")
                })
                Response(OK, "text/plain")
        })
})

$ goavl
/home/nao/.go/src/github.com/nao1215/goavl/test/sample/goa.go:8 Resource("operandsNG") is not snake case ('operands_ng')
/home/nao/.go/src/github.com/nao1215/goavl/test/sample/goa.go:9 Action("add-Ng") is not snake case ('add_ng')
/home/nao/.go/src/github.com/nao1215/goavl/test/sample/goa.go:10 Routing(GET("add_ng/:left/:right")) is not chain case ('add-ng/:left/:right')
```

# ライセンス
goavlプロジェクトは、複合ライセンスです。
- MITライセンス（[casee*.go](./internal/utils/strutils/casee.go) および [camelcase*.go](./internal/utils/strutils/camelcase.go)）
- [Apache License Version 2.0](./LICENSE)（上記以外のコード全て）

MITライセンスのソースコードの作者は、[pinzolo氏](https://github.com/pinzolo)および[Fatih Arslan氏](https://github.com/fatih)です。それぞれの作者が書かれたコードには、MITライセンス全文およびCopyrightが明示されています。

# 名前の由来
初期名称は"goalinter-v1"。linterのVersion 1に見えるため、"goav1linter"に改名。その後、"1"と"l"が似ていたため、それらを統合して名前を短くしました（= "goavl"）