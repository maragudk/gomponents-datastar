// Package datastar provides Datastar attributes and helpers for gomponents.
// See https://data-star.dev
package datastar

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type Modifier string

type Filter struct {
	Include string
	Exclude string
}

const (
	ModifierCapture        Modifier = "__capture"
	ModifierCase           Modifier = "__case"
	ModifierDebounce       Modifier = "__debounce"
	ModifierDelay          Modifier = "__delay"
	ModifierDuration       Modifier = "__duration"
	ModifierExit           Modifier = "__exit"
	ModifierFull           Modifier = "__full"
	ModifierHalf           Modifier = "__half"
	ModifierIfMissing      Modifier = "__ifmissing"
	ModifierOnce           Modifier = "__once"
	ModifierOutside        Modifier = "__outside"
	ModifierPassive        Modifier = "__passive"
	ModifierPrevent        Modifier = "__prevent"
	ModifierSelf           Modifier = "__self"
	ModifierStop           Modifier = "__stop"
	ModifierTerse          Modifier = "__terse"
	ModifierThreshold      Modifier = "__threshold"
	ModifierThrottle       Modifier = "__throttle"
	ModifierViewTransition Modifier = "__viewtransition"
	ModifierWindow         Modifier = "__window"
)

const (
	ModifierCamel      Modifier = ".camel" // Camel case: myEvent
	ModifierKebab      Modifier = ".kebab" // Kebab case: my-event
	ModifierLeading    Modifier = ".leading"
	ModifierNoLeading  Modifier = ".noleading"
	ModifierNoTrailing Modifier = ".notrailing"
	ModifierPascal     Modifier = ".pascal" // Pascal case: MyEvent
	ModifierSnake      Modifier = ".snake"  // Snake case: my_event
	ModifierTrailing   Modifier = ".trailing"
)

// Duration outputs millisecond values for durations, rounded to the nearest millisecond.
// Panics if the duration is negative.
func Duration(d time.Duration) Modifier {
	if d < 0 {
		panic(fmt.Sprintf("duration must not be negative, but is: %v", d))
	}
	return Modifier(fmt.Sprintf(".%vms", d.Round(time.Millisecond).Milliseconds()))
}

// Threshold outputs a visibility percentage threshold for the __threshold modifier.
// The value must be between 0.0 (exclusive) and 1.0 (inclusive).
// For values less than 1.0, the value is rounded to two decimal places (e.g., 0.25 for 25% visibility).
// For the value 1.0, it is formatted as ".100" representing 100% visibility.
// Panics if the threshold is outside the valid range.
func Threshold(threshold float64) Modifier {
	if threshold <= 0 || threshold > 1 {
		panic(fmt.Sprintf("threshold must be between 0.0 (exclusive) and 1.0 (inclusive), but is: %v", threshold))
	}
	// Special case: 1 represents 100% visibility
	if threshold == 1 {
		return Modifier(".100")
	}
	// Round to 2 decimal places and remove leading "0"
	return Modifier(strings.TrimPrefix(fmt.Sprintf("%.2f", threshold), "0"))
}

// Attr sets the value of any HTML attribute to an expression, and keeps it in sync.
//
// <div data-attr-title="$foo"></div>
//
// The `data-attr` attribute can also be used to set the values of multiple attributes on an element using a set of key-value pairs,
// where the keys represent attribute names and the values represent expressions.
//
// <div data-attr="{title: $foo, disabled: $bar}"></div>
//
// See https://data-star.dev/reference/attributes#data-attr
func Attr(pairs ...string) g.Node {
	if len(pairs)%2 == 1 {
		panic("each attribute name must have a value")
	}
	return data("attr", toObject(pairs))
}

// Bind creates a signal (if one doesn’t already exist) and sets up two-way data binding between it and an element’s value.
// This means that the value of the element is updated when the signal changes, and the signal value is updated when the value of the element changes.
//
// The `data-bind` attribute can be placed on any HTML element on which data can be input or choices selected (`input`, `select`, `textarea` elements, and web components).
// Event listeners are added for change and input events.
//
// <input data-bind-foo />
//
// The signal name can be specified in the key (as above), or in the value (as below). This can be useful depending on the templating language you are using.
//
// <input data-bind="foo" />
//
// The initial value of the signal is set to the value of the element, unless a signal has already been defined. So in the example below, $foo is set to bar.
//
// <input data-bind-foo value="bar" />
//
// Whereas in the example below, $foo inherits the value baz of the predefined signal.
//
// <div data-signals-foo="baz">
// <input data-bind-foo value="bar" />
// </div>
//
// See https://data-star.dev/reference/attributes#data-bind
func Bind(name string) g.Node {
	return data("bind", name)
}

// Class adds or removes a class to or from an element based on an expression.
//
// <div data-class-hidden="$foo"></div>
//
// If the expression evaluates to true, the `hidden` class is added to the element; otherwise, it is removed.
//
// The `data-class` attribute can also be used to add or remove multiple classes from an element using a set of key-value pairs,
// where the keys represent class names and the values represent expressions.
//
// <div data-class="{hidden: $foo, 'font-bold': $bar}"></div>
//
// See https://data-star.dev/reference/attributes#data-class
func Class(pairs ...string) g.Node {
	if len(pairs)%2 == 1 {
		panic("each class name must have a value")
	}
	return data("class", toObject(pairs))
}

// Computed creates a signal that is computed based on an expression. The computed signal is read-only,
// and its value is automatically updated when any signals in the expression are updated.
//
// <div data-computed="{foo: () => $bar + $baz}"></div>
//
// The `data-computed` attribute can also be used to create multiple computed signals using a set of key-value pairs,
// where the keys represent signal names and the values represent expressions.
//
// <div data-computed="{foo: () => $bar + $baz, total: () => $price * $quantity}"></div>
//
// Computed signals are useful for memoizing expressions containing other signals. Their values can be used in other expressions.
//
// <div data-computed="{foo: () => $bar + $baz}"></div>
// <div data-text="$foo"></div>
//
// Computed signal expressions must not be used for performing actions (changing other signals, actions, JavaScript functions, etc.).
// If you need to perform an action in response to a signal change, use the data-effect attribute.
//
// See https://data-star.dev/reference/attributes#data-computed
func Computed(pairs ...string) g.Node {
	if len(pairs)%2 == 1 {
		panic("each computed signal name must have an expression")
	}
	return data("computed", toComputed(pairs))
}

// Effect executes an expression on page load and whenever any signals in the expression change.
// This is useful for performing side effects, such as updating other signals, making requests to the backend, or manipulating the DOM.
//
// <div data-effect="$foo = $bar + $baz"></div>
//
// See https://data-star.dev/reference/attributes#data-effect
func Effect(expression string) g.Node {
	return data("effect", expression)
}

// Ignore tells Datastar to ignore an element and its descendants.
// Datastar walks the entire DOM and applies plugins to each element it encounters.
// It's possible to tell Datastar to ignore an element and its descendants by placing a `data-ignore` attribute on it.
// This is useful for preventing naming conflicts with third-party libraries or avoiding processing elements with potentially unsafe user input.
//
// <div data-ignore data-show-thirdpartylib="">
//
//	<div>
//	    Datastar will not process this element.
//	</div>
//
// </div>
//
// See https://data-star.dev/reference/attributes#data-ignore
func Ignore(modifiers ...Modifier) g.Node {
	eventWithModifiers := ""
	for _, modifier := range modifiers {
		eventWithModifiers += string(modifier)
	}
	return data("ignore" + eventWithModifiers)
}

// IgnoreMorph tells the `PatchElements` watcher to skip processing an element and its children when morphing elements.
//
// <div data-ignore-morph>
//
//	This element will not be morphed.
//
// </div>
//
// See https://data-star.dev/reference/attributes#data-ignore-morph
func IgnoreMorph() g.Node {
	return data("ignore-morph")
}

// Indicator creates a signal and sets its value to `true` while a fetch request is in flight, otherwise `false`. The signal can be used to show a loading indicator.
//
// <button data-on-click="@get('/endpoint')" data-indicator="fetching" data-attr-disabled="$fetching"></button>
// <div data-show="$fetching">Loading...</div>
//
// When using data-indicator with a fetch request initiated in a data-init attribute, you should ensure that the indicator signal is created before the fetch request is initialized.
//
// <div data-indicator-fetching data-init="@get('/endpoint')"></div>
//
// See https://data-star.dev/reference/attributes#data-indicator
func Indicator(name string, modifiers ...Modifier) g.Node {
	nameWithModifiers := ""
	for _, modifier := range modifiers {
		nameWithModifiers += string(modifier)
	}
	return data("indicator"+nameWithModifiers, name)
}

// JSONSignals sets the text content of an element to a reactive JSON stringified version of signals.
// Useful when troubleshooting an issue.
//
// You can optionally provide a filter object to include or exclude specific signals using regular expressions.
//
// <!-- Only show signals that include "user" in their path -->
// <pre data-json-signals="{include: /user/}"></pre>
//
// <!-- Show all signals except those ending with "temp" -->
// <pre data-json-signals="{exclude: /temp$/}"></pre>
//
// <!-- Combine include and exclude filters -->
// <pre data-json-signals="{include: /^app/, exclude: /password/}"></pre>
//
// <pre data-json-signals></pre>
//
// See https://data-star.dev/reference/attributes#data-json-signals
func JSONSignals(filter Filter, modifiers ...Modifier) g.Node {
	nameWithModifiers := ""
	for _, modifier := range modifiers {
		nameWithModifiers += string(modifier)
	}
	if filter.Include == "" && filter.Exclude == "" {
		return data("json-signals" + nameWithModifiers)
	}
	return data("json-signals"+nameWithModifiers, toFilter(filter))
}

// On attaches an event listener to an element, executing an expression whenever the event is triggered.
//
// <button data-on-click="$foo = ' '">Reset</button>
//
// An evt variable that represents the event object is available in the expression.
//
// <div data-on-myevent="$foo = evt.detail"></div>
//
// The `data-on` attribute works with events and custom events. The `data-on-submit` event listener prevents the default submission behavior of forms.
//
// See https://data-star.dev/reference/attributes#data-on
func On(event, expression string, modifiers ...Modifier) g.Node {
	eventWithModifiers := event
	for _, modifier := range modifiers {
		eventWithModifiers += string(modifier)
	}
	return data("on:"+eventWithModifiers, expression)
}

// OnIntersect runs an expression when the element intersects with the viewport.
//
// <div data-on-intersect="$intersected = true"></div>
//
// See https://data-star.dev/reference/attributes#data-on-intersect
func OnIntersect(expression string, modifiers ...Modifier) g.Node {
	eventWithModifiers := ""
	for _, modifier := range modifiers {
		eventWithModifiers += string(modifier)
	}
	return data("on-intersect"+eventWithModifiers, expression)
}

// OnInterval runs an expression at a regular interval. The interval duration defaults to one second and can be modified using the __duration modifier.
//
// <div data-on-interval="$count++"></div>
//
// See https://data-star.dev/reference/attributes#data-on-interval
func OnInterval(expression string, modifiers ...Modifier) g.Node {
	eventWithModifiers := ""
	for _, modifier := range modifiers {
		eventWithModifiers += string(modifier)
	}
	return data("on-interval"+eventWithModifiers, expression)
}

// Init runs an expression when an element is loaded into the DOM.
//
// The expression contained in the data-init attribute is executed when the element attribute is loaded into the DOM.
// This can happen on page load, when an element is patched into the DOM, and any time the attribute is modified (via a backend action or otherwise).
//
// <div data-init="$count = 1"></div>
//
// See https://data-star.dev/reference/attributes#data-init
func Init(expression string, modifiers ...Modifier) g.Node {
	eventWithModifiers := ""
	for _, modifier := range modifiers {
		eventWithModifiers += string(modifier)
	}
	return data("init"+eventWithModifiers, expression)
}

// OnSignalPatch runs an expression whenever one or more signals are patched.
// This is useful for tracking changes, updating computed values, or triggering side effects when data updates.
//
// <div data-on-signal-patch="console.log('Signal patch:', patch)"></div>
//
// The patch variable is available in the expression and contains the signal patch details.
//
// <div data-on-signal-patch="console.log('Signal patch:', patch)"></div>
//
// You can filter which signals to watch using the data-on-signal-patch-filter attribute.
//
// See https://data-star.dev/reference/attributes#data-on-signal-patch
func OnSignalPatch(expression string, modifiers ...Modifier) g.Node {
	eventWithModifiers := ""
	for _, modifier := range modifiers {
		eventWithModifiers += string(modifier)
	}
	return data("on-signal-patch"+eventWithModifiers, expression)
}

// OnSignalPatchFilter filters which signals to watch when using the `data-on-signal-patch` attribute.
//
// <div data-on-signal-patch-filter="{include: /^counter$/}"></div>
//
// See https://data-star.dev/reference/attributes#data-on-signal-patch-filter
func OnSignalPatchFilter(filter Filter) g.Node {
	return data("on-signal-patch-filter", toFilter(filter))
}

// PreserveAttr preserves the value of an attribute when morphing DOM elements.
//
// <details open data-preserve-attr="open">
//
//	<summary>Title</summary>
//	Content
//
// </details>
//
// You can preserve multiple attributes.
//
// See https://data-star.dev/reference/attributes#data-preserve-attr
func PreserveAttr(attrs ...string) g.Node {
	return data("preserve-attr", strings.Join(attrs, " "))
}

// Ref creates a new signal that is a reference to the element on which the data attribute is placed.
//
// <div data-ref="foo"></div>
//
// The signal value can then be used to reference the element.
//
// $foo is a reference to a <span data-text="$foo.tagName"></span> element
//
// See https://data-star.dev/reference/attributes#data-ref
func Ref(name string, modifiers ...Modifier) g.Node {
	nameWithModifiers := ""
	for _, modifier := range modifiers {
		nameWithModifiers += string(modifier)
	}
	return data("ref"+nameWithModifiers, name)
}

// Show or hide an element based on whether an expression evaluates to true or false.
// For anything with custom requirements, use data-class instead.
//
// <div data-show="$foo"></div>
//
// To prevent flickering of the element before Datastar has processed the DOM, you can add a display: none style to the element to hide it initially.
//
// <div data-show="$foo" style="display: none"></div>
//
// See https://data-star.dev/reference/attributes#data-show
func Show(expression string) g.Node {
	return data("show", expression)
}

// Signals patches (adds, updates or removes) one or more signals into the existing signals. Values defined later in the DOM tree override those defined earlier.
//
// <div data-signals="{foo: {bar: 1, baz: 2}}"></div>
//
// Setting a signal's value to null will remove the signal.
//
// <div data-signals="{foo: null}"></div>
//
// Signals beginning with an underscore are not included in requests to the backend by default.
// You can opt to include them by modifying the value of the filterSignals option.
//
// Signal names cannot begin with nor contain a double underscore (__), due to its use as a modifier delimiter.
//
// See https://data-star.dev/reference/attributes#data-signals
func Signals(signals map[string]any, modifiers ...Modifier) g.Node {
	nameWithModifiers := ""
	for _, modifier := range modifiers {
		nameWithModifiers += string(modifier)
	}
	return data("signals"+nameWithModifiers, toSignals(signals))
}

// Style sets the value of inline CSS styles on an element based on an expression, and keeps them in sync.
//
// The data-style attribute can be used to set multiple style properties on an element using a set of key-value pairs,
// where the keys represent CSS property names and the values represent expressions.
//
//	<div data-style="{
//	  display: $hiding ? 'none' : 'flex',
//	  flexDirection: 'column',
//	  color: $usingRed ? 'red' : 'green'
//	}"></div>
//
// Style properties can be specified in either camelCase (e.g., backgroundColor) or kebab-case (e.g., background-color).
// They will be automatically converted to the appropriate format.
//
// Empty string, null, undefined, or false values will restore the original inline style value if one existed,
// or remove the style property if there was no initial value. This allows you to use the logical AND operator (&&)
// for conditional styles: $condition && 'value' will apply the style when the condition is true and restore the original value when false.
//
// <!-- When $x is false, color remains red from inline style -->
// <div style="color: red;" data-style-color="$x && 'green'"></div>
//
// <!-- When $hiding is true, display becomes none; when false, reverts to flex from inline style -->
// <div style="display: flex;" data-style-display="$hiding && 'none'"></div>
//
// The plugin tracks initial inline style values and restores them when data-style expressions become falsy or during cleanup.
// This ensures existing inline styles are preserved and only the dynamic changes are managed by Datastar.
//
// See https://data-star.dev/reference/attributes#data-style
func Style(pairs ...string) g.Node {
	if len(pairs)%2 == 1 {
		panic("each style property must have a value")
	}
	return data("style", toObject(pairs))
}

// Text binds the text content of an element to an expression.
//
// <div data-text="$foo"></div>
//
// See https://data-star.dev/reference/attributes#data-text
func Text(v string) g.Node {
	return data("text", v)
}

func toObject(pairs []string) string {
	v := "{"
	for i := 0; i < len(pairs); i += 2 {
		v += fmt.Sprintf(`%s: %s`, pairs[i], pairs[i+1])
		if i < len(pairs)-2 {
			v += ", "
		}
	}
	v += "}"
	return v
}

func toComputed(pairs []string) string {
	v := "{"
	for i := 0; i < len(pairs); i += 2 {
		v += fmt.Sprintf(`%s: () => %s`, pairs[i], pairs[i+1])
		if i < len(pairs)-2 {
			v += ", "
		}
	}
	v += "}"
	return v
}

func toFilter(filter Filter) string {
	v := "{"
	if filter.Include != "" {
		v += fmt.Sprintf("include: %s", filter.Include)
		if filter.Exclude != "" {
			v += ", "
		}
	}
	if filter.Exclude != "" {
		v += fmt.Sprintf("exclude: %s", filter.Exclude)
	}
	v += "}"
	return v
}

func toSignals(signals map[string]any) string {
	b, err := json.Marshal(signals)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal signals: %v", err))
	}
	return string(b)
}

func data(name string, value ...string) g.Node {
	if len(value) > 0 {
		return html.Data(name, value[0])
	}
	return g.Attr("data-" + name)
}
