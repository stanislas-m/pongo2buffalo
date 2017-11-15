Pongo2Buffalo
=============

An adapter to use [Pongo2](https://github.com/flosch/pongo2) with [Buffalo](https://github.com/gobuffalo/buffalo).

Usage
-----

Configure the adapter to replace Plush for html files, in the `render.go` file:

```go
package actions

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
	"github.com/stanislas-m/pongo2buffalo"
)

var r *render.Engine
var assetsBox = packr.NewBox("../public/assets")

func init() {
	r = render.New(render.Options{
		TemplateEngines: map[string]render.TemplateEngine{
			"html": pongo2.BuffaloRenderer,
		},

		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{},
	})
}
```