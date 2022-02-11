package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = Resource("operandsNG", func() {
	Action("add_ok", func() {})
	Action("add-ng", func() {})
	Action("addNg", func() {})
	Action("AddNg", func() {})
})
