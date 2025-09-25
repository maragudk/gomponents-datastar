package datastar_test

import (
	"testing"

	. "maragu.dev/gomponents/html"

	ds "maragu.dev/gomponents-datastar"
	"maragu.dev/gomponents-datastar/internal/assert"
)

func TestAttr(t *testing.T) {
	t.Run(`should output data-attr="{title: $title}"`, func(t *testing.T) {
		n := Div(ds.Attr("title", "$title"))
		assert.Equal(t, `<div data-attr="{title: $title}"></div>`, n)
	})

	t.Run(`should output data-attr="{title: $title, id: $id}"`, func(t *testing.T) {
		n := Div(ds.Attr("title", "$title", "id", "$id"))
		assert.Equal(t, `<div data-attr="{title: $title, id: $id}"></div>`, n)
	})
}

func TestBind(t *testing.T) {
	t.Run(`should output data-bind="hat"`, func(t *testing.T) {
		n := Input(ds.Bind("hat"))
		assert.Equal(t, `<input data-bind="hat">`, n)
	})
}

func TestClass(t *testing.T) {
	t.Run(`should output data-class="{hidden: $hidden}"`, func(t *testing.T) {
		n := Div(ds.Class("hidden", "$hidden"))
		assert.Equal(t, `<div data-class="{hidden: $hidden}"></div>`, n)
	})

	t.Run(`should output data-class="{hidden: $hidden, font-bold: $bold}"`, func(t *testing.T) {
		n := Div(ds.Class("hidden", "$hidden", "font-bold", "$bold"))
		assert.Equal(t, `<div data-class="{hidden: $hidden, font-bold: $bold}"></div>`, n)
	})
}

func TestText(t *testing.T) {
	t.Run(`should output data-text="$foo"`, func(t *testing.T) {
		n := Div(ds.Text("$foo"))
		assert.Equal(t, `<div data-text="$foo"></div>`, n)
	})
}
