package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = Resource("operands", func() {
	Action("add", func() {
		Routing(GET("add-ok/:left/:right"))
	})
	Action("delete", func() {
		Routing(GET("delete_ng/:left-ng/:right"))
	})
	Action("post", func() {
		Routing(GET("postNg/:left/:right"))
	})
})
