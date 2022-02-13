package design

import (
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var BugMedia = MediaType("application/vnd.bug_media", func() {
	Attributes(func() {
		Required("bug")
		View("test")
	})
	Required("bug")
})
