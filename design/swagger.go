package design

import "github.com/goadesign/goa/design/apidsl"

var _ = apidsl.Resource("swagger", func() {
	apidsl.Origin("*", func() {
		apidsl.Methods("GET", "OPTIONS")
	})
	apidsl.Files("/swagger.json", "public/swagger/swagger.json")
})
