package datastar_test

import (
	"fmt"
	"testing"
	"time"

	. "maragu.dev/gomponents/html"

	data "maragu.dev/gomponents-datastar"
	"maragu.dev/gomponents-datastar/internal/assert"
)

func TestAttr(t *testing.T) {
	t.Run(`should output data-attr="{title: $title}"`, func(t *testing.T) {
		n := Div(data.Attr("title", "$title"))
		assert.Equal(t, `<div data-attr="{title: $title}"></div>`, n)
	})

	t.Run(`should output data-attr="{title: $title, id: $id}"`, func(t *testing.T) {
		n := Div(data.Attr("title", "$title", "id", "$id"))
		assert.Equal(t, `<div data-attr="{title: $title, id: $id}"></div>`, n)
	})
}

func TestBind(t *testing.T) {
	t.Run(`should output data-bind="foo"`, func(t *testing.T) {
		n := Input(data.Bind("foo"))
		assert.Equal(t, `<input data-bind="foo">`, n)
	})
}

func TestClass(t *testing.T) {
	t.Run(`should output data-class="{hidden: $hidden}"`, func(t *testing.T) {
		n := Div(data.Class("hidden", "$hidden"))
		assert.Equal(t, `<div data-class="{hidden: $hidden}"></div>`, n)
	})

	t.Run(`should output data-class="{hidden: $hidden, font-bold: $bold}"`, func(t *testing.T) {
		n := Div(data.Class("hidden", "$hidden", "font-bold", "$bold"))
		assert.Equal(t, `<div data-class="{hidden: $hidden, font-bold: $bold}"></div>`, n)
	})
}

func TestComputed(t *testing.T) {
	t.Run(`should output data-computed="{foo: () => $bar + $baz}"`, func(t *testing.T) {
		n := Div(data.Computed("foo", "$bar + $baz"))
		assert.Equal(t, `<div data-computed="{foo: () =&gt; $bar + $baz}"></div>`, n)
	})

	t.Run(`should output data-computed="{foo: () => $bar + $baz, total: () => $price * $quantity}"`, func(t *testing.T) {
		n := Div(data.Computed("foo", "$bar + $baz", "total", "$price * $quantity"))
		assert.Equal(t, `<div data-computed="{foo: () =&gt; $bar + $baz, total: () =&gt; $price * $quantity}"></div>`, n)
	})
}

func TestEffect(t *testing.T) {
	t.Run(`should output data-effect="$foo = $bar + $baz"`, func(t *testing.T) {
		n := Div(data.Effect("$foo = $bar + $baz"))
		assert.Equal(t, `<div data-effect="$foo = $bar + $baz"></div>`, n)
	})
}

func TestIgnore(t *testing.T) {
	t.Run(`should output data-ignore`, func(t *testing.T) {
		n := Div(data.Ignore())
		assert.Equal(t, `<div data-ignore></div>`, n)
	})

	t.Run(`should output data-ignore__self=""`, func(t *testing.T) {
		n := Div(data.Ignore(data.ModifierSelf))
		assert.Equal(t, `<div data-ignore__self></div>`, n)
	})
}

func TestIgnoreMorph(t *testing.T) {
	t.Run(`should output data-ignore-morph`, func(t *testing.T) {
		n := Div(data.IgnoreMorph())
		assert.Equal(t, `<div data-ignore-morph></div>`, n)
	})
}

func TestJSONSignals(t *testing.T) {
	t.Run(`should output data-json-signals`, func(t *testing.T) {
		n := Pre(data.JSONSignals(data.Filter{}))
		assert.Equal(t, `<pre data-json-signals></pre>`, n)
	})

	t.Run(`should output data-json-signals__terse`, func(t *testing.T) {
		n := Pre(data.JSONSignals(data.Filter{}, data.ModifierTerse))
		assert.Equal(t, `<pre data-json-signals__terse></pre>`, n)
	})

	t.Run(`should output data-json-signals="{include: /user/}"`, func(t *testing.T) {
		n := Pre(data.JSONSignals(data.Filter{Include: "/user/"}))
		assert.Equal(t, `<pre data-json-signals="{include: /user/}"></pre>`, n)
	})

	t.Run(`should output data-json-signals="{exclude: /temp$/}"`, func(t *testing.T) {
		n := Pre(data.JSONSignals(data.Filter{Exclude: "/temp$/"}))
		assert.Equal(t, `<pre data-json-signals="{exclude: /temp$/}"></pre>`, n)
	})

	t.Run(`should output data-json-signals="{include: /^app/, exclude: /password/}"`, func(t *testing.T) {
		n := Pre(data.JSONSignals(data.Filter{Include: "/^app/", Exclude: "/password/"}))
		assert.Equal(t, `<pre data-json-signals="{include: /^app/, exclude: /password/}"></pre>`, n)
	})

	t.Run(`should output data-json-signals__terse="{include: /counter/}"`, func(t *testing.T) {
		n := Pre(data.JSONSignals(data.Filter{Include: "/counter/"}, data.ModifierTerse))
		assert.Equal(t, `<pre data-json-signals__terse="{include: /counter/}"></pre>`, n)
	})
}

func TestIndicator(t *testing.T) {
	t.Run(`should output data-indicator="fetching"`, func(t *testing.T) {
		n := Button(data.Indicator("fetching"))
		assert.Equal(t, `<button data-indicator="fetching"></button>`, n)
	})

	tests := []struct {
		name     string
		modifier data.Modifier
		expected string
	}{
		{name: `should output data-indicator__case.camel="fetching"`, modifier: data.ModifierCamel, expected: `<button data-indicator__case.camel="fetching"></button>`},
		{name: `should output data-indicator__case.kebab="fetching"`, modifier: data.ModifierKebab, expected: `<button data-indicator__case.kebab="fetching"></button>`},
		{name: `should output data-indicator__case.snake="fetching"`, modifier: data.ModifierSnake, expected: `<button data-indicator__case.snake="fetching"></button>`},
		{name: `should output data-indicator__case.pascal="fetching"`, modifier: data.ModifierPascal, expected: `<button data-indicator__case.pascal="fetching"></button>`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := Button(data.Indicator("fetching", data.ModifierCase, test.modifier))
			assert.Equal(t, test.expected, n)
		})
	}
}

func TestOn(t *testing.T) {
	t.Run(`should output data-on:click="$foo = ''"`, func(t *testing.T) {
		n := Button(data.On("click", "$foo = ''"))
		assert.Equal(t, `<button data-on:click="$foo = &#39;&#39;"></button>`, n)
	})

	t.Run(`should output data-on:click__window__debounce.500ms.leading="$foo = ''"`, func(t *testing.T) {
		n := Button(data.On("click", "$foo = ''", data.ModifierWindow, data.ModifierDebounce, data.Duration(500*time.Millisecond), data.ModifierLeading))
		assert.Equal(t, `<button data-on:click__window__debounce.500ms.leading="$foo = &#39;&#39;"></button>`, n)
	})
}

func TestOnIntersect(t *testing.T) {
	t.Run(`should output data-on-intersect="$intersected = true"`, func(t *testing.T) {
		n := Div(data.OnIntersect("$intersected = true"))
		assert.Equal(t, `<div data-on-intersect="$intersected = true"></div>`, n)
	})

	t.Run(`should output data-on-intersect__once__full="$fullyIntersected = true"`, func(t *testing.T) {
		n := Div(data.OnIntersect("$fullyIntersected = true", data.ModifierOnce, data.ModifierFull))
		assert.Equal(t, `<div data-on-intersect__once__full="$fullyIntersected = true"></div>`, n)
	})

	t.Run(`should output data-on-intersect__half="$halfIntersected = true"`, func(t *testing.T) {
		n := Div(data.OnIntersect("$halfIntersected = true", data.ModifierHalf))
		assert.Equal(t, `<div data-on-intersect__half="$halfIntersected = true"></div>`, n)
	})

	t.Run(`should output data-on-intersect__exit="$exited = true"`, func(t *testing.T) {
		n := Div(data.OnIntersect("$exited = true", data.ModifierExit))
		assert.Equal(t, `<div data-on-intersect__exit="$exited = true"></div>`, n)
	})

	t.Run(`should output data-on-intersect__threshold.25="$visible = true"`, func(t *testing.T) {
		n := Div(data.OnIntersect("$visible = true", data.ModifierThreshold, data.Threshold(0.25)))
		assert.Equal(t, `<div data-on-intersect__threshold.25="$visible = true"></div>`, n)
	})
}

func TestOnInterval(t *testing.T) {
	t.Run(`should output data-on-interval="$count++"`, func(t *testing.T) {
		n := Div(data.OnInterval("$count++"))
		assert.Equal(t, `<div data-on-interval="$count++"></div>`, n)
	})

	t.Run(`should output data-on-interval__duration.500ms="$count++"`, func(t *testing.T) {
		n := Div(data.OnInterval("$count++", data.ModifierDuration, data.Duration(500*time.Millisecond)))
		assert.Equal(t, `<div data-on-interval__duration.500ms="$count++"></div>`, n)
	})
}

func TestInit(t *testing.T) {
	t.Run(`should output data-init="$count = 1"`, func(t *testing.T) {
		n := Div(data.Init("$count = 1"))
		assert.Equal(t, `<div data-init="$count = 1"></div>`, n)
	})

	t.Run(`should output data-init__delay.500ms="$count = 1"`, func(t *testing.T) {
		n := Div(data.Init("$count = 1", data.ModifierDelay, data.Duration(500*time.Millisecond)))
		assert.Equal(t, `<div data-init__delay.500ms="$count = 1"></div>`, n)
	})
}

func TestOnSignalPatch(t *testing.T) {
	t.Run(`should output data-on-signal-patch="console.log('Signal patch:', patch)"`, func(t *testing.T) {
		n := Div(data.OnSignalPatch("console.log('Signal patch:', patch)"))
		assert.Equal(t, `<div data-on-signal-patch="console.log(&#39;Signal patch:&#39;, patch)"></div>`, n)
	})

	t.Run(`should output data-on-signal-patch__debounce.500ms="doSomething()"`, func(t *testing.T) {
		n := Div(data.OnSignalPatch("doSomething()", data.ModifierDebounce, data.Duration(500*time.Millisecond)))
		assert.Equal(t, `<div data-on-signal-patch__debounce.500ms="doSomething()"></div>`, n)
	})
}

func TestOnSignalPatchFilter(t *testing.T) {
	t.Run(`should output data-on-signal-patch-filter="{include: /^counter$/}"`, func(t *testing.T) {
		n := Div(data.OnSignalPatchFilter(data.Filter{Include: "/^counter$/"}))
		assert.Equal(t, `<div data-on-signal-patch-filter="{include: /^counter$/}"></div>`, n)
	})

	t.Run(`should output data-on-signal-patch-filter="{exclude: /changes$/}"`, func(t *testing.T) {
		n := Div(data.OnSignalPatchFilter(data.Filter{Exclude: "/changes$/"}))
		assert.Equal(t, `<div data-on-signal-patch-filter="{exclude: /changes$/}"></div>`, n)
	})

	t.Run(`should output data-on-signal-patch-filter="{include: /user/, exclude: /password/}"`, func(t *testing.T) {
		n := Div(data.OnSignalPatchFilter(data.Filter{Include: "/user/", Exclude: "/password/"}))
		assert.Equal(t, `<div data-on-signal-patch-filter="{include: /user/, exclude: /password/}"></div>`, n)
	})
}

func TestPreserveAttr(t *testing.T) {
	t.Run(`should output data-preserve-attr="open"`, func(t *testing.T) {
		n := Details(data.PreserveAttr("open"))
		assert.Equal(t, `<details data-preserve-attr="open"></details>`, n)
	})

	t.Run(`should output data-preserve-attr="open class"`, func(t *testing.T) {
		n := Details(data.PreserveAttr("open", "class"))
		assert.Equal(t, `<details data-preserve-attr="open class"></details>`, n)
	})
}

func TestRef(t *testing.T) {
	t.Run(`should output data-ref="foo"`, func(t *testing.T) {
		n := Div(data.Ref("foo"))
		assert.Equal(t, `<div data-ref="foo"></div>`, n)
	})

	tests := []struct {
		name     string
		modifier data.Modifier
		expected string
	}{
		{name: `should output data-ref__case.camel="foo"`, modifier: data.ModifierCamel, expected: `<div data-ref__case.camel="foo"></div>`},
		{name: `should output data-ref__case.kebab="foo"`, modifier: data.ModifierKebab, expected: `<div data-ref__case.kebab="foo"></div>`},
		{name: `should output data-ref__case.snake="foo"`, modifier: data.ModifierSnake, expected: `<div data-ref__case.snake="foo"></div>`},
		{name: `should output data-ref__case.pascal="foo"`, modifier: data.ModifierPascal, expected: `<div data-ref__case.pascal="foo"></div>`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := Div(data.Ref("foo", data.ModifierCase, test.modifier))
			assert.Equal(t, test.expected, n)
		})
	}
}

func TestShow(t *testing.T) {
	t.Run(`should output data-show="$foo"`, func(t *testing.T) {
		n := Div(data.Show("$foo"))
		assert.Equal(t, `<div data-show="$foo"></div>`, n)
	})
}

func TestSignals(t *testing.T) {
	t.Run(`should output data-signals="{"foo":1}"`, func(t *testing.T) {
		n := Div(data.Signals(map[string]any{"foo": 1}))
		assert.Equal(t, `<div data-signals="{&#34;foo&#34;:1}"></div>`, n)
	})

	t.Run(`should output data-signals="{"bar":2,"foo":1}"`, func(t *testing.T) {
		n := Div(data.Signals(map[string]any{"foo": 1, "bar": 2}))
		assert.Equal(t, `<div data-signals="{&#34;bar&#34;:2,&#34;foo&#34;:1}"></div>`, n)
	})

	t.Run(`should output data-signals with nested objects`, func(t *testing.T) {
		n := Div(data.Signals(map[string]any{"foo": map[string]any{"bar": 1, "baz": 2}}))
		assert.Equal(t, `<div data-signals="{&#34;foo&#34;:{&#34;bar&#34;:1,&#34;baz&#34;:2}}"></div>`, n)
	})

	t.Run(`should output data-signals__ifmissing="{"foo":1}"`, func(t *testing.T) {
		n := Div(data.Signals(map[string]any{"foo": 1}, data.ModifierIfMissing))
		assert.Equal(t, `<div data-signals__ifmissing="{&#34;foo&#34;:1}"></div>`, n)
	})

	t.Run(`should output data-signals__case.kebab="{"foo":1}"`, func(t *testing.T) {
		n := Div(data.Signals(map[string]any{"foo": 1}, data.ModifierCase, data.ModifierKebab))
		assert.Equal(t, `<div data-signals__case.kebab="{&#34;foo&#34;:1}"></div>`, n)
	})
}

func TestStyle(t *testing.T) {
	t.Run(`should output data-style="{display: $hiding ? 'none' : 'flex'}"`, func(t *testing.T) {
		n := Div(data.Style("display", "$hiding ? 'none' : 'flex'"))
		assert.Equal(t, `<div data-style="{display: $hiding ? &#39;none&#39; : &#39;flex&#39;}"></div>`, n)
	})

	t.Run(`should output data-style="{display: $hiding ? 'none' : 'flex', color: $usingRed ? 'red' : 'green'}"`, func(t *testing.T) {
		n := Div(data.Style("display", "$hiding ? 'none' : 'flex'", "color", "$usingRed ? 'red' : 'green'"))
		assert.Equal(t, `<div data-style="{display: $hiding ? &#39;none&#39; : &#39;flex&#39;, color: $usingRed ? &#39;red&#39; : &#39;green&#39;}"></div>`, n)
	})
}

func TestText(t *testing.T) {
	t.Run(`should output data-text="$foo"`, func(t *testing.T) {
		n := Div(data.Text("$foo"))
		assert.Equal(t, `<div data-text="$foo"></div>`, n)
	})
}

func ExampleAttr() {
	fmt.Print(Div(data.Attr("title", "$title")))
	// Output: <div data-attr="{title: $title}"></div>
}

func ExampleAttr_multiple() {
	fmt.Print(Div(data.Attr("title", "$title", "id", "$id")))
	// Output: <div data-attr="{title: $title, id: $id}"></div>
}

func ExampleBind() {
	fmt.Print(Input(data.Bind("foo")))
	// Output: <input data-bind="foo">
}

func ExampleClass() {
	fmt.Print(Div(data.Class("hidden", "$hidden")))
	// Output: <div data-class="{hidden: $hidden}"></div>
}

func ExampleClass_multiple() {
	fmt.Print(Div(data.Class("hidden", "$hidden", "font-bold", "$bold")))
	// Output: <div data-class="{hidden: $hidden, font-bold: $bold}"></div>
}

func ExampleComputed() {
	fmt.Print(Div(data.Computed("foo", "$bar + $baz")))
	// Output: <div data-computed="{foo: () =&gt; $bar + $baz}"></div>
}

func ExampleComputed_multiple() {
	fmt.Print(Div(data.Computed("foo", "$bar + $baz", "total", "$price * $quantity")))
	// Output: <div data-computed="{foo: () =&gt; $bar + $baz, total: () =&gt; $price * $quantity}"></div>
}

func ExampleEffect() {
	fmt.Print(Div(data.Effect("$foo = $bar + $baz")))
	// Output: <div data-effect="$foo = $bar + $baz"></div>
}

func ExampleIgnore() {
	fmt.Print(Div(data.Ignore()))
	// Output: <div data-ignore></div>
}

func ExampleIgnore_withModifierSelf() {
	fmt.Print(Div(data.Ignore(data.ModifierSelf)))
	// Output: <div data-ignore__self></div>
}

func ExampleIgnoreMorph() {
	fmt.Print(Div(data.IgnoreMorph()))
	// Output: <div data-ignore-morph></div>
}

func ExampleIndicator() {
	fmt.Print(Button(data.Indicator("fetching")))
	// Output: <button data-indicator="fetching"></button>
}

func ExampleIndicator_withModifierCase() {
	fmt.Print(Button(data.Indicator("fetching", data.ModifierCase, data.ModifierKebab)))
	// Output: <button data-indicator__case.kebab="fetching"></button>
}

func ExampleJSONSignals() {
	fmt.Print(Pre(data.JSONSignals(data.Filter{})))
	// Output: <pre data-json-signals></pre>
}

func ExampleJSONSignals_withFilter() {
	fmt.Print(Pre(data.JSONSignals(data.Filter{Include: "/user/"})))
	// Output: <pre data-json-signals="{include: /user/}"></pre>
}

func ExampleJSONSignals_withModifier() {
	fmt.Print(Pre(data.JSONSignals(data.Filter{}, data.ModifierTerse)))
	// Output: <pre data-json-signals__terse></pre>
}

func ExampleOn_click() {
	fmt.Print(Button(data.On("click", "$foo = ''")))
	// Output: <button data-on:click="$foo = &#39;&#39;"></button>
}

func ExampleOn_withModifiers() {
	fmt.Print(Button(data.On("click", "$foo = ''", data.ModifierWindow, data.ModifierDebounce, data.Duration(500*time.Millisecond), data.ModifierLeading)))
	// Output: <button data-on:click__window__debounce.500ms.leading="$foo = &#39;&#39;"></button>
}

func ExampleOnIntersect() {
	fmt.Print(Div(data.OnIntersect("$intersected = true")))
	// Output: <div data-on-intersect="$intersected = true"></div>
}

func ExampleOnIntersect_withModifiers() {
	fmt.Print(Div(data.OnIntersect("$fullyIntersected = true", data.ModifierOnce, data.ModifierFull)))
	// Output: <div data-on-intersect__once__full="$fullyIntersected = true"></div>
}

func ExampleOnInterval() {
	fmt.Print(Div(data.OnInterval("$count++")))
	// Output: <div data-on-interval="$count++"></div>
}

func ExampleOnInterval_withDuration() {
	fmt.Print(Div(data.OnInterval("$count++", data.ModifierDuration, data.Duration(500*time.Millisecond))))
	// Output: <div data-on-interval__duration.500ms="$count++"></div>
}

func ExampleInit() {
	fmt.Print(Div(data.Init("$count = 1")))
	// Output: <div data-init="$count = 1"></div>
}

func ExampleInit_withDelay() {
	fmt.Print(Div(data.Init("$count = 1", data.ModifierDelay, data.Duration(500*time.Millisecond))))
	// Output: <div data-init__delay.500ms="$count = 1"></div>
}

func ExampleOnSignalPatch() {
	fmt.Print(Div(data.OnSignalPatch("console.log('Signal patch:', patch)")))
	// Output: <div data-on-signal-patch="console.log(&#39;Signal patch:&#39;, patch)"></div>
}

func ExampleOnSignalPatch_withModifiers() {
	fmt.Print(Div(data.OnSignalPatch("doSomething()", data.ModifierDebounce, data.Duration(500*time.Millisecond))))
	// Output: <div data-on-signal-patch__debounce.500ms="doSomething()"></div>
}

func ExampleOnSignalPatchFilter() {
	fmt.Print(Div(data.OnSignalPatchFilter(data.Filter{Include: "/^counter$/"})))
	// Output: <div data-on-signal-patch-filter="{include: /^counter$/}"></div>
}

func ExampleOnSignalPatchFilter_withExclude() {
	fmt.Print(Div(data.OnSignalPatchFilter(data.Filter{Include: "/user/", Exclude: "/password/"})))
	// Output: <div data-on-signal-patch-filter="{include: /user/, exclude: /password/}"></div>
}

func ExamplePreserveAttr() {
	fmt.Print(Details(data.PreserveAttr("open")))
	// Output: <details data-preserve-attr="open"></details>
}

func ExamplePreserveAttr_multiple() {
	fmt.Print(Details(data.PreserveAttr("open", "class")))
	// Output: <details data-preserve-attr="open class"></details>
}

func ExampleRef() {
	fmt.Print(Div(data.Ref("foo")))
	// Output: <div data-ref="foo"></div>
}

func ExampleRef_withModifierCase() {
	fmt.Print(Div(data.Ref("foo", data.ModifierCase, data.ModifierKebab)))
	// Output: <div data-ref__case.kebab="foo"></div>
}

func ExampleShow() {
	fmt.Print(Div(data.Show("$foo")))
	// Output: <div data-show="$foo"></div>
}

func ExampleSignals() {
	fmt.Print(Div(data.Signals(map[string]any{"foo": 1})))
	// Output: <div data-signals="{&#34;foo&#34;:1}"></div>
}

func ExampleSignals_withModifierIfMissing() {
	fmt.Print(Div(data.Signals(map[string]any{"foo": 1}, data.ModifierIfMissing)))
	// Output: <div data-signals__ifmissing="{&#34;foo&#34;:1}"></div>
}

func ExampleStyle() {
	fmt.Print(Div(data.Style("display", "$hiding ? 'none' : 'flex'")))
	// Output: <div data-style="{display: $hiding ? &#39;none&#39; : &#39;flex&#39;}"></div>
}

func ExampleStyle_multiple() {
	fmt.Print(Div(data.Style("display", "$hiding ? 'none' : 'flex'", "color", "$usingRed ? 'red' : 'green'")))
	// Output: <div data-style="{display: $hiding ? &#39;none&#39; : &#39;flex&#39;, color: $usingRed ? &#39;red&#39; : &#39;green&#39;}"></div>
}

func ExampleText() {
	fmt.Print(Div(data.Text("$foo")))
	// Output: <div data-text="$foo"></div>
}

func TestDuration(t *testing.T) {
	t.Run("should panic on negative duration", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for negative duration")
			}
		}()
		data.Duration(-1)
	})

	t.Run("should round zero duration to 0ms", func(t *testing.T) {
		n := Div(data.OnInterval("$count++", data.ModifierDuration, data.Duration(0)))
		assert.Equal(t, `<div data-on-interval__duration.0ms="$count++"></div>`, n)
	})

	t.Run("should round small duration down to 0ms", func(t *testing.T) {
		n := Div(data.OnInterval("$count++", data.ModifierDuration, data.Duration(100*time.Microsecond)))
		assert.Equal(t, `<div data-on-interval__duration.0ms="$count++"></div>`, n)
	})

	t.Run("should round 500us up to 1ms", func(t *testing.T) {
		n := Div(data.OnInterval("$count++", data.ModifierDuration, data.Duration(500*time.Microsecond)))
		assert.Equal(t, `<div data-on-interval__duration.1ms="$count++"></div>`, n)
	})

	t.Run("should not round 1ms", func(t *testing.T) {
		n := Div(data.OnInterval("$count++", data.ModifierDuration, data.Duration(time.Millisecond)))
		assert.Equal(t, `<div data-on-interval__duration.1ms="$count++"></div>`, n)
	})
}

func ExampleDuration_rounding() {
	fmt.Print(Div(data.OnInterval("$count++", data.ModifierDuration, data.Duration(500*time.Microsecond))))
	// Output: <div data-on-interval__duration.1ms="$count++"></div>
}

func TestThreshold(t *testing.T) {
	t.Run("should panic on negative threshold", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for negative threshold")
			}
		}()
		data.Threshold(-0.1)
	})

	t.Run("should panic on threshold equal to 0", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for threshold equal to 0")
			}
		}()
		data.Threshold(0)
	})

	t.Run("should panic on threshold greater than 1", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for threshold greater than 1")
			}
		}()
		data.Threshold(1.1)
	})

	t.Run("should format 0.25 as .25", func(t *testing.T) {
		n := Div(data.OnIntersect("$visible = true", data.ModifierThreshold, data.Threshold(0.25)))
		assert.Equal(t, `<div data-on-intersect__threshold.25="$visible = true"></div>`, n)
	})

	t.Run("should format 0.5 as .50", func(t *testing.T) {
		n := Div(data.OnIntersect("$visible = true", data.ModifierThreshold, data.Threshold(0.5)))
		assert.Equal(t, `<div data-on-intersect__threshold.50="$visible = true"></div>`, n)
	})

	t.Run("should format 1 as 100", func(t *testing.T) {
		n := Div(data.OnIntersect("$visible = true", data.ModifierThreshold, data.Threshold(1)))
		assert.Equal(t, `<div data-on-intersect__threshold.100="$visible = true"></div>`, n)
	})

	t.Run("should round 0.333 to .33", func(t *testing.T) {
		n := Div(data.OnIntersect("$visible = true", data.ModifierThreshold, data.Threshold(0.333)))
		assert.Equal(t, `<div data-on-intersect__threshold.33="$visible = true"></div>`, n)
	})

	t.Run("should round 0.335 to .34", func(t *testing.T) {
		n := Div(data.OnIntersect("$visible = true", data.ModifierThreshold, data.Threshold(0.335)))
		assert.Equal(t, `<div data-on-intersect__threshold.34="$visible = true"></div>`, n)
	})
}

func ExampleOnIntersect_withExit() {
	fmt.Print(Div(data.OnIntersect("$exited = true", data.ModifierExit)))
	// Output: <div data-on-intersect__exit="$exited = true"></div>
}

func ExampleOnIntersect_withThreshold() {
	fmt.Print(Div(data.OnIntersect("$visible = true", data.ModifierThreshold, data.Threshold(0.25))))
	// Output: <div data-on-intersect__threshold.25="$visible = true"></div>
}
