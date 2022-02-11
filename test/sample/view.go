package design

import (
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

// ViewMedia is media type for test
var ViewMedia = MediaType("application/vnd.view_media", func() {
	View("ok")
	Attributes(func() {
		Attribute("test")
		View("ng")
	})
})

// ViewType is type
var ViewType = Type("ViewType", func() {
	View("ng")
})
