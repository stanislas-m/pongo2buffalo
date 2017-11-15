package pongo2

import (
	"github.com/stretchr/testify/require"
	"testing"
	"fmt"
)

func Test_Render_Simple_HTML(t *testing.T) {
	r := require.New(t)

	input := `<p>Hi</p>`
	s, err := BuffaloRenderer(input, nil, nil)
	r.NoError(err)
	r.Equal(input, s)
}

func Test_Render_Variable(t *testing.T) {
	r := require.New(t)

	input := `{{ greet }} {{ name }}!`
	ctxData := map[string]interface{}{
		"greet": "Hello",
		"name": "Stan",
	}
	s, err := BuffaloRenderer(input, ctxData, nil)
	r.NoError(err)
	r.Equal("Hello Stan!", s)
}

func Test_Render_Condition(t *testing.T) {
	r := require.New(t)

	input := `{% if is_allowed %}I'm allowed!{% endif %}`
	ctxData := map[string]interface{}{
		"is_allowed": true,
	}
	s, err := BuffaloRenderer(input, ctxData, nil)
	r.NoError(err)
	r.Equal("I'm allowed!", s)
}

func Test_Render_Helper(t *testing.T) {
	r := require.New(t)

	input := `{{ hello_helper(user) }}`
	ctxData := map[string]interface{}{
		"user": "Mark",
	}
	helpers := map[string]interface{}{
		"hello_helper": func (s string) string {
			return fmt.Sprintf("Hello %s!", s)
		},
	}
	s, err := BuffaloRenderer(input, ctxData, helpers)
	r.NoError(err)
	r.Equal("Hello Mark!", s)
}