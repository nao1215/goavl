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

// TestMedia is media type for test
var TestMedia = MediaType("application/vnd.test_media", func() {
	Attribute("AbcDefID")
	Attribute("zzzXXX-ss", String, func() {
		NoExample()
	})
})
