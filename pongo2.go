package pongo2

import (
	p2 "github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"sync"
)

var cache = map[string]*p2.Template{}
var moot = &sync.Mutex{}

// BuffaloRenderer implements the render.TemplateEngine interface allowing pongo2 to be used as a template engine
// for Buffalo
func BuffaloRenderer(input string, data map[string]interface{}, helpers map[string]interface{}) (string, error) {
	t, err := Parse(input)
	if err != nil {
		return "", err
	}
	ctx := p2.Context{}
	if data != nil {
		// Add data to context
		for k, v := range data {
			ctx[k] = v
		}
	}
	if helpers != nil {
		// Add helpers to context
		for k, v := range helpers {
			ctx[k] = v
		}
	}
	// Execute template
	out, err := t.Execute(ctx)
	if err != nil {
		return "", err
	}
	return out, nil
}

// Parse an input string and return a *pongo2.Template, and caches the parsed template.
func Parse(input string) (*p2.Template, error) {
	moot.Lock()
	defer moot.Unlock()
	if t, ok := cache[input]; ok {
		return t, nil
	}
	t, err := p2.FromString(input)

	if err == nil {
		cache[input] = t
	}

	if err != nil {
		return t, errors.WithStack(err)
	}

	return t, nil
}
