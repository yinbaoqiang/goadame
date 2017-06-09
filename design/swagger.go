package design

import "github.com/goadesign/goa/design/apidsl"

var _ = apidsl.Resource("swagger", func() {
	apidsl.Origin("*", func() {
		apidsl.Methods("GET", "OPTIONS")
	})
	apidsl.Files("/swagger.json", "swagger/swagger.json")
})
var _ = apidsl.Resource("public", func() {
	apidsl.Origin("*", func() {
		apidsl.Methods("GET", "OPTIONS")
	})
	apidsl.Files("/ui/*filepath", "dist")
})
