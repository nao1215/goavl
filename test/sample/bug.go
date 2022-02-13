package design

import (
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var BugType = Type("bug", func() {
	Attributes(func() {
		View("bug")
	})
	View("bug")
})

var BugMedia = MediaType("application/vnd.bug_media", func() {
	View("ng")
	Attributes(func() {
		View("ng")
	})
})
