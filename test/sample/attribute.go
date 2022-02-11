package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

// NameMedia is media type for sample
var NameMedia = MediaType("application/vnd.name_media", func() {
	Attribute("this_is_ok")
	Attribute("this-is-ng", String, func() {})
	Attribute("NgCase", String, func() {})
	Attribute("ngCase", String, func() {})
})

// SampleMedia is media type for sample
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
