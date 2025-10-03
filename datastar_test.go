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
	t.Run(`should output data-bind="foo"`, func(t *testing.T) {
		n := Input(ds.Bind("foo"))
		assert.Equal(t, `<input data-bind="foo">`, n)
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

func TestEffect(t *testing.T) {
	t.Run(`should output data-effect="$foo = $bar + $baz"`, func(t *testing.T) {
		n := Div(ds.Effect("$foo = $bar + $baz"))
		assert.Equal(t, `<div data-effect="$foo = $bar + $baz"></div>`, n)
	})
}

func TestIgnore(t *testing.T) {
	t.Run(`should output data-ignore`, func(t *testing.T) {
		n := Div(ds.Ignore())
		assert.Equal(t, `<div data-ignore></div>`, n)
	})

	t.Run(`should output data-ignore__self=""`, func(t *testing.T) {
		n := Div(ds.Ignore(ds.ModifierSelf))
		assert.Equal(t, `<div data-ignore__self></div>`, n)
	})
}

func TestIgnoreMorph(t *testing.T) {
	t.Run(`should output data-ignore-morph`, func(t *testing.T) {
		n := Div(ds.IgnoreMorph())
		assert.Equal(t, `<div data-ignore-morph></div>`, n)
	})
}

func TestIndicator(t *testing.T) {
	t.Run(`should output data-indicator="fetching"`, func(t *testing.T) {
		n := Button(ds.Indicator("fetching"))
		assert.Equal(t, `<button data-indicator="fetching"></button>`, n)
	})

	tests := []struct {
		name     string
		modifier ds.Modifier
		expected string
	}{
		{name: `should output data-indicator__case.camel="fetching"`, modifier: ds.ModifierCamel, expected: `<button data-indicator__case.camel="fetching"></button>`},
		{name: `should output data-indicator__case.kebab="fetching"`, modifier: ds.ModifierKebab, expected: `<button data-indicator__case.kebab="fetching"></button>`},
		{name: `should output data-indicator__case.snake="fetching"`, modifier: ds.ModifierSnake, expected: `<button data-indicator__case.snake="fetching"></button>`},
		{name: `should output data-indicator__case.pascal="fetching"`, modifier: ds.ModifierPascal, expected: `<button data-indicator__case.pascal="fetching"></button>`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := Button(ds.Indicator("fetching", ds.ModifierCase, test.modifier))
			assert.Equal(t, test.expected, n)
		})
	}
}

func TestText(t *testing.T) {
	t.Run(`should output data-text="$foo"`, func(t *testing.T) {
		n := Div(ds.Text("$foo"))
		assert.Equal(t, `<div data-text="$foo"></div>`, n)
	})
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
