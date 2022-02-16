# goavl: サンプル集
本ファイルでは、goavlがどのような内容をチェックし、どのような結果を返すかを例示します。

# 命名規則チェック

|関数名|チェック対象|命名規則|
|:--|:--|:--|
|Action()|第一引数|スネークケース|
|Attribute()|第一引数|スネークケース|
|Resource()|第一引数|スネークケース|
|Routing()|第一引数|チェインケース（ケバブケース）|


## Resource()の引数
Resource()の第一引数は、スネークケースである事が期待値です。
``` test/sample/resource.go
var _ = Resource("operands_ok", func() {})
var _ = Resource("operandsNg", func() {})
var _ = Resource("OperandsNg", func() {})
var _ = Resource("operands-Ng", func() {})
var _ = Resource("operands-Ng_case", func() {})
```
```
$ goavl 
[WARN] test/sample/resource.go:9    Resource("operandsNg") is not snake case ('operands_ng')
[WARN] test/sample/resource.go:10   Resource("OperandsNg") is not snake case ('operands_ng')
[WARN] test/sample/resource.go:11   Resource("operands-Ng") is not snake case ('operands_ng')
[WARN] test/sample/resource.go:12   Resource("operands-Ng_case") is not snake case ('operands_ng_case')
```

## Action()の引数
Resource()の第一引数は、スネークケースである事が期待値です。
``` test/sample/action.go
var _ = Resource("operands", func() {
	Action("add_ok", func() {})
	Action("add-ng", func() {})
	Action("addNg", func() {})
	Action("AddNg", func() {})
})
```
```
$ goavl 
[WARN] test/sample/action.go:10   Action("add-ng") is not snake case ('add_ng')
[WARN] test/sample/action.go:11   Action("addNg") is not snake case ('add_ng')
[WARN] test/sample/action.go:12   Action("AddNg") is not snake case ('add_ng')
```

## Routing()の引数
Routing()の第一引数は、チェインケース（ケバブケース）である事が期待値です。
``` test/sample/resource.go
var _ = Resource("operands", func() {
	Action("add", func() {
		Routing(GET("add-ok/:left/:right"))
	})
	Action("delete", func() {
		Routing(DELETE("delete_ng/:left-ng/qright"))
	})
	Action("post", func() {
		Routing(POST("/postNg/abc.php"))
	})
	Action("put", func() {
		Routing(PUT(""))
	})
})
```
```
$ goavl 
[WARN] test/sample/routing.go:12   Routing(DELETE("delete_ng/:left-ng/:right")) is not chain case ('delete-ng/:left-ng/:right')
[WARN] test/sample/routing.go:15   Routing(POST("/postNg/abc.php")) is not chain case ('/post-ng/abc.php')
```

## Attribute()の引数
Attribute()の第一引数は、スネークケースである事が期待値です。
```
var NameMedia = MediaType("application/vnd.name_media", func() {
	Attribute("this_is_ok")
	Attribute("this-is-ng", String, func() {})
	Attribute("NgCase", String, func() {})
	Attribute("ngCase", String, func() {})
})
```
```
$ goavl 
[WARN] test/sample/attribute.go:11   Attribute("this-is-ng") is not snake case ('this_is_ng')
[WARN] test/sample/attribute.go:12   Attribute("NgCase") is not snake case ('ng_case')
[WARN] test/sample/attribute.go:13   Attribute("ngCase") is not snake case ('ng_case')
```

# 構文チェック
- Description()を必ず記載する事
- Example()を必ず記載する事
- Attribute()：使用可能な関数内で使用されているか
- Default()：使用可能な関数内で使用されているか
- Enum()：使用可能な関数内で使用されているか
- Example()：使用可能な関数内で使用されているか
- Format()：使用可能な関数内で使用されているか
- Header()：使用可能な関数内で使用されているか
- MaxLength()：使用可能な関数内で使用されているか
- Maximum()：使用可能な関数内で使用されているか
- Member()：使用可能な関数内で使用されているか
- MinLength()：使用可能な関数内で使用されているか
- Minimum()：使用可能な関数内で使用されているか
- MultipartForm()：使用可能な関数内で使用されているか
- NoExample()：使用可能な関数内で使用されているか
- Param()：使用可能な関数内で使用されているか
- Pattern()：使用可能な関数内で使用されているか
- ReadOnly()：使用可能な関数内で使用されているか
- Required()：使用可能な関数内で使用されているか
- TypeName()：使用可能な関数内で使用されているか
- URL()：使用可能な関数内で使用されているか
- View()：使用可能な関数内で使用されているか

## Description()の有無チェック
API(), Resource(), Action(), MediaType(), Attribute(), Response(), ResponseTemplate()の中にDescription()がない場合は、警告が出ます。
```
var DescriptionMedia = MediaType("application/vnd.description_media", func() {
	Attribute("ok")
	Attribute("ok", String, func() {
		Description("ok")
	})
	Attribute("ng", String, func() {
		Example("ng")
	})
})
```
```
$ goavl
[WARN] test/sample/description.go:14   Not exist Description() in Attribute().
```
## Example()の有無チェック
Example()は、クライアント開発者にAPI仕様を伝えるために重要です。そのため、Example()が記載されていない場合は、以下の指摘が発生します。
```
var SampleMedia = MediaType("application/vnd.sample_media", func() {
	Attribute("ok")
	Attribute("ok", String, func() {
		Example("OK")
	})
	Attribute("ng", String, func() {
		NoExample()
	})
	Attribute("ng", String, func() {
		Description("no example in Attribute()")
	})
})
```
```
$ ./goavl 
[WARN] test/sample/attribute.go:23   NoExample() in Attribute(). NoExample() is not user(client) friendly
[WARN] test/sample/attribute.go:25   Not exist Example() in Attribute().
```

## View()の使用箇所チェック
View()は、MediaType()とResponse()内でのみ使用できます。それ以外の箇所で使用している場合は、以下の指摘が発生します。
```
var ViewMedia = MediaType("application/vnd.view_media", func() {
	View("ok")
	Attributes(func() {
		Attribute("test")
		View("ng")
	})
})

var ViewType = Type("ViewType", func() {
	View("ng")
})
```
```
$ ./goavl
[WARN] test/sample/view.go:10   Attributes() has View(). View() can be used in MediaType() or Response()
[WARN] test/sample/view.go:17   Type() has View(). View() can be used in MediaType() or Response()
```