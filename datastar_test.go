package datastar_test

import (
	"testing"
	"time"

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

func TestComputed(t *testing.T) {
	t.Run(`should output data-computed-foo="$bar + $baz"`, func(t *testing.T) {
		n := Div(ds.Computed("foo", "$bar + $baz"))
		assert.Equal(t, `<div data-computed-foo="$bar + $baz"></div>`, n)
	})

	tests := []struct {
		name     string
		modifier ds.Modifier
		expected string
	}{
		{name: `should output data-computed-foo__case.camel="$bar + $baz"`, modifier: ds.ModifierCamel, expected: `<div data-computed-foo__case.camel="$bar + $baz"></div>`},
		{name: `should output data-computed-foo__case.kebab="$bar + $baz"`, modifier: ds.ModifierKebab, expected: `<div data-computed-foo__case.kebab="$bar + $baz"></div>`},
		{name: `should output data-computed-foo__case.snake="$bar + $baz"`, modifier: ds.ModifierSnake, expected: `<div data-computed-foo__case.snake="$bar + $baz"></div>`},
		{name: `should output data-computed-foo__case.pascal="$bar + $baz"`, modifier: ds.ModifierPascal, expected: `<div data-computed-foo__case.pascal="$bar + $baz"></div>`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := Div(ds.Computed("foo", "$bar + $baz", ds.ModifierCase, test.modifier))
			assert.Equal(t, test.expected, n)
		})
	}
}

func TestOn(t *testing.T) {
	t.Run(`should output data-on-click="$foo = ''"`, func(t *testing.T) {
		n := Button(ds.On("click", "$foo = ''"))
		assert.Equal(t, `<button data-on-click="$foo = &#39;&#39;"></button>`, n)
	})

	t.Run(`should output data-on-click__window__debounce.500ms.leading="$foo = ''"`, func(t *testing.T) {
		n := Button(ds.On("click", "$foo = ''", ds.ModifierWindow, ds.ModifierDebounce, ds.ModifierDuration(500*time.Millisecond), ds.ModifierLeading))
		assert.Equal(t, `<button data-on-click__window__debounce.500ms.leading="$foo = &#39;&#39;"></button>`, n)
	})
}
