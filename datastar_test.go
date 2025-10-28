package datastar_test

import (
	"fmt"
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
	t.Run(`should output data-computed="{foo: () => $bar + $baz}"`, func(t *testing.T) {
		n := Div(ds.Computed("foo", "$bar + $baz"))
		assert.Equal(t, `<div data-computed="{foo: () =&gt; $bar + $baz}"></div>`, n)
	})

	t.Run(`should output data-computed="{foo: () => $bar + $baz, total: () => $price * $quantity}"`, func(t *testing.T) {
		n := Div(ds.Computed("foo", "$bar + $baz", "total", "$price * $quantity"))
		assert.Equal(t, `<div data-computed="{foo: () =&gt; $bar + $baz, total: () =&gt; $price * $quantity}"></div>`, n)
	})
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

func TestJSONSignals(t *testing.T) {
	t.Run(`should output data-json-signals`, func(t *testing.T) {
		n := Pre(ds.JSONSignals(ds.Filter{}))
		assert.Equal(t, `<pre data-json-signals></pre>`, n)
	})

	t.Run(`should output data-json-signals__terse`, func(t *testing.T) {
		n := Pre(ds.JSONSignals(ds.Filter{}, ds.ModifierTerse))
		assert.Equal(t, `<pre data-json-signals__terse></pre>`, n)
	})

	t.Run(`should output data-json-signals="{include: /user/}"`, func(t *testing.T) {
		n := Pre(ds.JSONSignals(ds.Filter{Include: "/user/"}))
		assert.Equal(t, `<pre data-json-signals="{include: /user/}"></pre>`, n)
	})

	t.Run(`should output data-json-signals="{exclude: /temp$/}"`, func(t *testing.T) {
		n := Pre(ds.JSONSignals(ds.Filter{Exclude: "/temp$/"}))
		assert.Equal(t, `<pre data-json-signals="{exclude: /temp$/}"></pre>`, n)
	})

	t.Run(`should output data-json-signals="{include: /^app/, exclude: /password/}"`, func(t *testing.T) {
		n := Pre(ds.JSONSignals(ds.Filter{Include: "/^app/", Exclude: "/password/"}))
		assert.Equal(t, `<pre data-json-signals="{include: /^app/, exclude: /password/}"></pre>`, n)
	})

	t.Run(`should output data-json-signals__terse="{include: /counter/}"`, func(t *testing.T) {
		n := Pre(ds.JSONSignals(ds.Filter{Include: "/counter/"}, ds.ModifierTerse))
		assert.Equal(t, `<pre data-json-signals__terse="{include: /counter/}"></pre>`, n)
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

func TestOn(t *testing.T) {
	t.Run(`should output data-on:click="$foo = ''"`, func(t *testing.T) {
		n := Button(ds.On("click", "$foo = ''"))
		assert.Equal(t, `<button data-on:click="$foo = &#39;&#39;"></button>`, n)
	})

	t.Run(`should output data-on:click__window__debounce.500ms.leading="$foo = ''"`, func(t *testing.T) {
		n := Button(ds.On("click", "$foo = ''", ds.ModifierWindow, ds.ModifierDebounce, ds.Duration(500*time.Millisecond), ds.ModifierLeading))
		assert.Equal(t, `<button data-on:click__window__debounce.500ms.leading="$foo = &#39;&#39;"></button>`, n)
	})
}

func TestOnIntersect(t *testing.T) {
	t.Run(`should output data-on-intersect="$intersected = true"`, func(t *testing.T) {
		n := Div(ds.OnIntersect("$intersected = true"))
		assert.Equal(t, `<div data-on-intersect="$intersected = true"></div>`, n)
	})

	t.Run(`should output data-on-intersect__once__full="$fullyIntersected = true"`, func(t *testing.T) {
		n := Div(ds.OnIntersect("$fullyIntersected = true", ds.ModifierOnce, ds.ModifierFull))
		assert.Equal(t, `<div data-on-intersect__once__full="$fullyIntersected = true"></div>`, n)
	})

	t.Run(`should output data-on-intersect__half="$halfIntersected = true"`, func(t *testing.T) {
		n := Div(ds.OnIntersect("$halfIntersected = true", ds.ModifierHalf))
		assert.Equal(t, `<div data-on-intersect__half="$halfIntersected = true"></div>`, n)
	})
}

func TestOnInterval(t *testing.T) {
	t.Run(`should output data-on-interval="$count++"`, func(t *testing.T) {
		n := Div(ds.OnInterval("$count++"))
		assert.Equal(t, `<div data-on-interval="$count++"></div>`, n)
	})

	t.Run(`should output data-on-interval__duration.500ms="$count++"`, func(t *testing.T) {
		n := Div(ds.OnInterval("$count++", ds.ModifierDuration, ds.Duration(500*time.Millisecond)))
		assert.Equal(t, `<div data-on-interval__duration.500ms="$count++"></div>`, n)
	})
}

func TestInit(t *testing.T) {
	t.Run(`should output data-init="$count = 1"`, func(t *testing.T) {
		n := Div(ds.Init("$count = 1"))
		assert.Equal(t, `<div data-init="$count = 1"></div>`, n)
	})

	t.Run(`should output data-init__delay.500ms="$count = 1"`, func(t *testing.T) {
		n := Div(ds.Init("$count = 1", ds.ModifierDelay, ds.Duration(500*time.Millisecond)))
		assert.Equal(t, `<div data-init__delay.500ms="$count = 1"></div>`, n)
	})
}

func TestOnSignalPatch(t *testing.T) {
	t.Run(`should output data-on-signal-patch="console.log('Signal patch:', patch)"`, func(t *testing.T) {
		n := Div(ds.OnSignalPatch("console.log('Signal patch:', patch)"))
		assert.Equal(t, `<div data-on-signal-patch="console.log(&#39;Signal patch:&#39;, patch)"></div>`, n)
	})

	t.Run(`should output data-on-signal-patch__debounce.500ms="doSomething()"`, func(t *testing.T) {
		n := Div(ds.OnSignalPatch("doSomething()", ds.ModifierDebounce, ds.Duration(500*time.Millisecond)))
		assert.Equal(t, `<div data-on-signal-patch__debounce.500ms="doSomething()"></div>`, n)
	})
}

func TestOnSignalPatchFilter(t *testing.T) {
	t.Run(`should output data-on-signal-patch-filter="{include: /^counter$/}"`, func(t *testing.T) {
		n := Div(ds.OnSignalPatchFilter(ds.Filter{Include: "/^counter$/"}))
		assert.Equal(t, `<div data-on-signal-patch-filter="{include: /^counter$/}"></div>`, n)
	})

	t.Run(`should output data-on-signal-patch-filter="{exclude: /changes$/}"`, func(t *testing.T) {
		n := Div(ds.OnSignalPatchFilter(ds.Filter{Exclude: "/changes$/"}))
		assert.Equal(t, `<div data-on-signal-patch-filter="{exclude: /changes$/}"></div>`, n)
	})

	t.Run(`should output data-on-signal-patch-filter="{include: /user/, exclude: /password/}"`, func(t *testing.T) {
		n := Div(ds.OnSignalPatchFilter(ds.Filter{Include: "/user/", Exclude: "/password/"}))
		assert.Equal(t, `<div data-on-signal-patch-filter="{include: /user/, exclude: /password/}"></div>`, n)
	})
}

func TestPreserveAttr(t *testing.T) {
	t.Run(`should output data-preserve-attr="open"`, func(t *testing.T) {
		n := Details(ds.PreserveAttr("open"))
		assert.Equal(t, `<details data-preserve-attr="open"></details>`, n)
	})

	t.Run(`should output data-preserve-attr="open class"`, func(t *testing.T) {
		n := Details(ds.PreserveAttr("open", "class"))
		assert.Equal(t, `<details data-preserve-attr="open class"></details>`, n)
	})
}

func TestRef(t *testing.T) {
	t.Run(`should output data-ref="foo"`, func(t *testing.T) {
		n := Div(ds.Ref("foo"))
		assert.Equal(t, `<div data-ref="foo"></div>`, n)
	})

	tests := []struct {
		name     string
		modifier ds.Modifier
		expected string
	}{
		{name: `should output data-ref__case.camel="foo"`, modifier: ds.ModifierCamel, expected: `<div data-ref__case.camel="foo"></div>`},
		{name: `should output data-ref__case.kebab="foo"`, modifier: ds.ModifierKebab, expected: `<div data-ref__case.kebab="foo"></div>`},
		{name: `should output data-ref__case.snake="foo"`, modifier: ds.ModifierSnake, expected: `<div data-ref__case.snake="foo"></div>`},
		{name: `should output data-ref__case.pascal="foo"`, modifier: ds.ModifierPascal, expected: `<div data-ref__case.pascal="foo"></div>`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := Div(ds.Ref("foo", ds.ModifierCase, test.modifier))
			assert.Equal(t, test.expected, n)
		})
	}
}

func TestShow(t *testing.T) {
	t.Run(`should output data-show="$foo"`, func(t *testing.T) {
		n := Div(ds.Show("$foo"))
		assert.Equal(t, `<div data-show="$foo"></div>`, n)
	})
}

func TestSignals(t *testing.T) {
	t.Run(`should output data-signals="{"foo":1}"`, func(t *testing.T) {
		n := Div(ds.Signals(map[string]any{"foo": 1}))
		assert.Equal(t, `<div data-signals="{&#34;foo&#34;:1}"></div>`, n)
	})

	t.Run(`should output data-signals="{"bar":2,"foo":1}"`, func(t *testing.T) {
		n := Div(ds.Signals(map[string]any{"foo": 1, "bar": 2}))
		assert.Equal(t, `<div data-signals="{&#34;bar&#34;:2,&#34;foo&#34;:1}"></div>`, n)
	})

	t.Run(`should output data-signals with nested objects`, func(t *testing.T) {
		n := Div(ds.Signals(map[string]any{"foo": map[string]any{"bar": 1, "baz": 2}}))
		assert.Equal(t, `<div data-signals="{&#34;foo&#34;:{&#34;bar&#34;:1,&#34;baz&#34;:2}}"></div>`, n)
	})

	t.Run(`should output data-signals__ifmissing="{"foo":1}"`, func(t *testing.T) {
		n := Div(ds.Signals(map[string]any{"foo": 1}, ds.ModifierIfMissing))
		assert.Equal(t, `<div data-signals__ifmissing="{&#34;foo&#34;:1}"></div>`, n)
	})

	t.Run(`should output data-signals__case.kebab="{"foo":1}"`, func(t *testing.T) {
		n := Div(ds.Signals(map[string]any{"foo": 1}, ds.ModifierCase, ds.ModifierKebab))
		assert.Equal(t, `<div data-signals__case.kebab="{&#34;foo&#34;:1}"></div>`, n)
	})
}

func TestStyle(t *testing.T) {
	t.Run(`should output data-style="{display: $hiding ? 'none' : 'flex'}"`, func(t *testing.T) {
		n := Div(ds.Style("display", "$hiding ? 'none' : 'flex'"))
		assert.Equal(t, `<div data-style="{display: $hiding ? &#39;none&#39; : &#39;flex&#39;}"></div>`, n)
	})

	t.Run(`should output data-style="{display: $hiding ? 'none' : 'flex', color: $usingRed ? 'red' : 'green'}"`, func(t *testing.T) {
		n := Div(ds.Style("display", "$hiding ? 'none' : 'flex'", "color", "$usingRed ? 'red' : 'green'"))
		assert.Equal(t, `<div data-style="{display: $hiding ? &#39;none&#39; : &#39;flex&#39;, color: $usingRed ? &#39;red&#39; : &#39;green&#39;}"></div>`, n)
	})
}

func TestText(t *testing.T) {
	t.Run(`should output data-text="$foo"`, func(t *testing.T) {
		n := Div(ds.Text("$foo"))
		assert.Equal(t, `<div data-text="$foo"></div>`, n)
	})
}

func ExampleAttr() {
	fmt.Print(Div(ds.Attr("title", "$title")))
	// Output: <div data-attr="{title: $title}"></div>
}

func ExampleAttr_multiple() {
	fmt.Print(Div(ds.Attr("title", "$title", "id", "$id")))
	// Output: <div data-attr="{title: $title, id: $id}"></div>
}

func ExampleBind() {
	fmt.Print(Input(ds.Bind("foo")))
	// Output: <input data-bind="foo">
}

func ExampleClass() {
	fmt.Print(Div(ds.Class("hidden", "$hidden")))
	// Output: <div data-class="{hidden: $hidden}"></div>
}

func ExampleClass_multiple() {
	fmt.Print(Div(ds.Class("hidden", "$hidden", "font-bold", "$bold")))
	// Output: <div data-class="{hidden: $hidden, font-bold: $bold}"></div>
}

func ExampleComputed() {
	fmt.Print(Div(ds.Computed("foo", "$bar + $baz")))
	// Output: <div data-computed="{foo: () =&gt; $bar + $baz}"></div>
}

func ExampleComputed_multiple() {
	fmt.Print(Div(ds.Computed("foo", "$bar + $baz", "total", "$price * $quantity")))
	// Output: <div data-computed="{foo: () =&gt; $bar + $baz, total: () =&gt; $price * $quantity}"></div>
}

func ExampleEffect() {
	fmt.Print(Div(ds.Effect("$foo = $bar + $baz")))
	// Output: <div data-effect="$foo = $bar + $baz"></div>
}

func ExampleIgnore() {
	fmt.Print(Div(ds.Ignore()))
	// Output: <div data-ignore></div>
}

func ExampleIgnore_withModifierSelf() {
	fmt.Print(Div(ds.Ignore(ds.ModifierSelf)))
	// Output: <div data-ignore__self></div>
}

func ExampleIgnoreMorph() {
	fmt.Print(Div(ds.IgnoreMorph()))
	// Output: <div data-ignore-morph></div>
}

func ExampleIndicator() {
	fmt.Print(Button(ds.Indicator("fetching")))
	// Output: <button data-indicator="fetching"></button>
}

func ExampleIndicator_withModifierCase() {
	fmt.Print(Button(ds.Indicator("fetching", ds.ModifierCase, ds.ModifierKebab)))
	// Output: <button data-indicator__case.kebab="fetching"></button>
}

func ExampleJSONSignals() {
	fmt.Print(Pre(ds.JSONSignals(ds.Filter{})))
	// Output: <pre data-json-signals></pre>
}

func ExampleJSONSignals_withFilter() {
	fmt.Print(Pre(ds.JSONSignals(ds.Filter{Include: "/user/"})))
	// Output: <pre data-json-signals="{include: /user/}"></pre>
}

func ExampleJSONSignals_withModifier() {
	fmt.Print(Pre(ds.JSONSignals(ds.Filter{}, ds.ModifierTerse)))
	// Output: <pre data-json-signals__terse></pre>
}

func ExampleOn_click() {
	fmt.Print(Button(ds.On("click", "$foo = ''")))
	// Output: <button data-on:click="$foo = &#39;&#39;"></button>
}

func ExampleOn_withModifiers() {
	fmt.Print(Button(ds.On("click", "$foo = ''", ds.ModifierWindow, ds.ModifierDebounce, ds.Duration(500*time.Millisecond), ds.ModifierLeading)))
	// Output: <button data-on:click__window__debounce.500ms.leading="$foo = &#39;&#39;"></button>
}

func ExampleOnIntersect() {
	fmt.Print(Div(ds.OnIntersect("$intersected = true")))
	// Output: <div data-on-intersect="$intersected = true"></div>
}

func ExampleOnIntersect_withModifiers() {
	fmt.Print(Div(ds.OnIntersect("$fullyIntersected = true", ds.ModifierOnce, ds.ModifierFull)))
	// Output: <div data-on-intersect__once__full="$fullyIntersected = true"></div>
}

func ExampleOnInterval() {
	fmt.Print(Div(ds.OnInterval("$count++")))
	// Output: <div data-on-interval="$count++"></div>
}

func ExampleOnInterval_withDuration() {
	fmt.Print(Div(ds.OnInterval("$count++", ds.ModifierDuration, ds.Duration(500*time.Millisecond))))
	// Output: <div data-on-interval__duration.500ms="$count++"></div>
}

func ExampleInit() {
	fmt.Print(Div(ds.Init("$count = 1")))
	// Output: <div data-init="$count = 1"></div>
}

func ExampleInit_withDelay() {
	fmt.Print(Div(ds.Init("$count = 1", ds.ModifierDelay, ds.Duration(500*time.Millisecond))))
	// Output: <div data-init__delay.500ms="$count = 1"></div>
}

func ExampleOnSignalPatch() {
	fmt.Print(Div(ds.OnSignalPatch("console.log('Signal patch:', patch)")))
	// Output: <div data-on-signal-patch="console.log(&#39;Signal patch:&#39;, patch)"></div>
}

func ExampleOnSignalPatch_withModifiers() {
	fmt.Print(Div(ds.OnSignalPatch("doSomething()", ds.ModifierDebounce, ds.Duration(500*time.Millisecond))))
	// Output: <div data-on-signal-patch__debounce.500ms="doSomething()"></div>
}

func ExampleOnSignalPatchFilter() {
	fmt.Print(Div(ds.OnSignalPatchFilter(ds.Filter{Include: "/^counter$/"})))
	// Output: <div data-on-signal-patch-filter="{include: /^counter$/}"></div>
}

func ExampleOnSignalPatchFilter_withExclude() {
	fmt.Print(Div(ds.OnSignalPatchFilter(ds.Filter{Include: "/user/", Exclude: "/password/"})))
	// Output: <div data-on-signal-patch-filter="{include: /user/, exclude: /password/}"></div>
}

func ExamplePreserveAttr() {
	fmt.Print(Details(ds.PreserveAttr("open")))
	// Output: <details data-preserve-attr="open"></details>
}

func ExamplePreserveAttr_multiple() {
	fmt.Print(Details(ds.PreserveAttr("open", "class")))
	// Output: <details data-preserve-attr="open class"></details>
}

func ExampleRef() {
	fmt.Print(Div(ds.Ref("foo")))
	// Output: <div data-ref="foo"></div>
}

func ExampleRef_withModifierCase() {
	fmt.Print(Div(ds.Ref("foo", ds.ModifierCase, ds.ModifierKebab)))
	// Output: <div data-ref__case.kebab="foo"></div>
}

func ExampleShow() {
	fmt.Print(Div(ds.Show("$foo")))
	// Output: <div data-show="$foo"></div>
}

func ExampleSignals() {
	fmt.Print(Div(ds.Signals(map[string]any{"foo": 1})))
	// Output: <div data-signals="{&#34;foo&#34;:1}"></div>
}

func ExampleSignals_withModifierIfMissing() {
	fmt.Print(Div(ds.Signals(map[string]any{"foo": 1}, ds.ModifierIfMissing)))
	// Output: <div data-signals__ifmissing="{&#34;foo&#34;:1}"></div>
}

func ExampleStyle() {
	fmt.Print(Div(ds.Style("display", "$hiding ? 'none' : 'flex'")))
	// Output: <div data-style="{display: $hiding ? &#39;none&#39; : &#39;flex&#39;}"></div>
}

func ExampleStyle_multiple() {
	fmt.Print(Div(ds.Style("display", "$hiding ? 'none' : 'flex'", "color", "$usingRed ? 'red' : 'green'")))
	// Output: <div data-style="{display: $hiding ? &#39;none&#39; : &#39;flex&#39;, color: $usingRed ? &#39;red&#39; : &#39;green&#39;}"></div>
}

func ExampleText() {
	fmt.Print(Div(ds.Text("$foo")))
	// Output: <div data-text="$foo"></div>
}

func TestDuration(t *testing.T) {
	t.Run("should panic on negative duration", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for negative duration")
			}
		}()
		ds.Duration(-1)
	})

	t.Run("should round zero duration to 0ms", func(t *testing.T) {
		n := Div(ds.OnInterval("$count++", ds.ModifierDuration, ds.Duration(0)))
		assert.Equal(t, `<div data-on-interval__duration.0ms="$count++"></div>`, n)
	})

	t.Run("should round small duration down to 0ms", func(t *testing.T) {
		n := Div(ds.OnInterval("$count++", ds.ModifierDuration, ds.Duration(100*time.Microsecond)))
		assert.Equal(t, `<div data-on-interval__duration.0ms="$count++"></div>`, n)
	})

	t.Run("should round 500us up to 1ms", func(t *testing.T) {
		n := Div(ds.OnInterval("$count++", ds.ModifierDuration, ds.Duration(500*time.Microsecond)))
		assert.Equal(t, `<div data-on-interval__duration.1ms="$count++"></div>`, n)
	})

	t.Run("should not round 1ms", func(t *testing.T) {
		n := Div(ds.OnInterval("$count++", ds.ModifierDuration, ds.Duration(time.Millisecond)))
		assert.Equal(t, `<div data-on-interval__duration.1ms="$count++"></div>`, n)
	})
}

func ExampleDuration_rounding() {
	fmt.Print(Div(ds.OnInterval("$count++", ds.ModifierDuration, ds.Duration(500*time.Microsecond))))
	// Output: <div data-on-interval__duration.1ms="$count++"></div>
}
