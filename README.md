[![Build](https://github.com/nao1215/goavl/actions/workflows/build.yml/badge.svg)](https://github.com/nao1215/goavl/actions/workflows/build.yml)
[![UnitTest](https://github.com/nao1215/goavl/actions/workflows/unit_test.yml/badge.svg)](https://github.com/nao1215/goavl/actions/workflows/unit_test.yml)  
[[日本語](./doc/ja/README.md)]
# goavl: Goa framework (ver1) linter
goavl is linter of [goa version1 (fork version)] (https://github.com/shogo82148/goa-v1). The purpose of development is to speed up Web API design using goa and minimize design differences within the team.

Goa generate the base-files(routing, controller, Swagger, etc.) required for Web API hosting that based on the design written in DSL. However, it is difficult to use. If the design fails to generate various files, the error location cannot be identified immediately (error message is not user friendly). For example, modern compilers show the number of lines in error. However, goa does not output any such information.

So, we started developing Linter to instantly identify the problem areas in the goa design file. goavl will only support **goa version 1**, will not support the current version (it's means version 3). The reason is that I'm using forked goa version 1.


# How to install
### Step.1 Install golang
goavl only supports installation with "$ go install". If you does not have the golang development environment installed on your system, please install golang from the [golang official website] (https://go.dev/doc/install).

### Step2. Install goavl
```
$ go install github.com/nao1215/goavl@latest
```
# How to use
## check all
goavl extracts the design package (go file) under the current directory and checks the files. If you want to know a more detailed example, please check [example.md] (./doc/ja/example.md).
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
## Check only one file
You can use the -f (--file) option to specify the file to check. You can not specify multiple files.
```
$ goavl --file test/sample/goa.go
```
## Confirmation of check items (inspection IDs)
If you use the subcommand list, goavl print the list of check items. The output is in the form of "Inspection ID: Check Contents".
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

## Exclusion of check items (inspection IDs)
If you want to ignore the indication, you use the --exclude option. You can specify multiple "Inspection IDs" that you want to exclude, separated by commas.
```
$ goavl --exclude=FC001,NC003
```
# LICENSE
The goavl project is the mixed-license.
- MIT License（[casee*.go](./internal/utils/strutils/casee.go) and [camelcase*.go](./internal/utils/strutils/camelcase.go)）
- [Apache License Version 2.0](./LICENSE)（All codes other than the above）

The authors of the MIT license source code are [pinzolo] (https://github.com/pinzolo) and [Fatih Arslan] (https://github.com/fatih). The code written by each author clearly states the full MIT license and Copyright.

# Origin of the name
The initial name is "goal inter-v1". Renamed to "goav1linter" because it looks like version 1 of linter. After that, "1" and "l" were similar, so I integrated them and shortened the name (= "goavl").