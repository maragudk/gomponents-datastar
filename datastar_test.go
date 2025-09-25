package datastar_test

import (
	"fmt"
	"testing"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	ds "maragu.dev/gomponents-datastar"
	"maragu.dev/gomponents-datastar/internal/assert"
)

func TestAttributes(t *testing.T) {
	attrs := map[string]func(string, string) g.Node{
		"attr": ds.Attr,
	}

	for name, attr := range attrs {
		t.Run(fmt.Sprintf(`should output data-%v-hat="party"`, name), func(t *testing.T) {
			n := Div(attr("hat", "party"))
			assert.Equal(t, fmt.Sprintf(`<div data-%v-hat="party"></div>`, name), n)
		})
	}
}
