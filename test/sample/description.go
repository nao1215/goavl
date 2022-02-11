package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

// DescriptionMedia is media type for sample
var DescriptionMedia = MediaType("application/vnd.description_media", func() {
	Attribute("ok")
	Attribute("ok", String, func() {
		Description("ok")
	})
	Attribute("ng", String, func() {
		Example("ng")
	})
})
