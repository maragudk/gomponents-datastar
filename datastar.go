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
	ModifierDocument       Modifier = "__document" // Attaches the event listener to the document element
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

// Prop outputs the __prop modifier for [Bind], for example "__prop.checked".
// It binds to a specific property instead of the default binding. The property must not be read-only.
//
// Unlike [Duration] and [Threshold], which are reused across several modifiers and so take one as a separate argument,
// __prop takes only a property name, so it is a single modifier.
//
// Datastar reads the property name from the attribute key, which HTML lowercases, and converts it from kebab case.
// Give multi-word property names in kebab case, so "inner-html" rather than "innerHTML".
//
// See https://data-star.dev/reference/attributes#data-bind
func Prop(name string) Modifier {
	return Modifier("__prop." + name)
}

// Event outputs the __event modifier for [Bind], for example "__event.input.change".
// It defines which events sync the element property back to the signal.
//
// Unlike [Duration] and [Threshold], which are reused across several modifiers and so take one as a separate argument,
// __event takes only event names, so it is a single modifier.
//
// Datastar reads the event names from the attribute key, which HTML lowercases.
// Give multi-word event names in kebab case, so "my-event" rather than "myEvent".
//
// Datastar adds no listeners if no event names are given.
//
// See https://data-star.dev/reference/attributes#data-bind
func Event(names ...string) Modifier {
	v := "__event"
	for _, name := range names {
		v += "." + name
	}
	return Modifier(v)
}

// Attr sets the value of any HTML attribute to an expression, and keeps it in sync.
//
// <div data-attr:title="$foo"></div>
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

// Bind creates a signal (if one doesn’t already exist) and sets up two-way data binding between it and an element’s current bound state.
// When the signal changes, Datastar writes that value to the element.
// When one of the bind events fires, Datastar reads the element’s current bound property/value and writes that back to the signal.
//
// The `data-bind` attribute can be placed on any HTML element on which data can be input or choices selected (`input`, `select`, `textarea` elements, and web components).
// Native elements use their built-in bind semantics automatically.
// Generic custom elements default to binding through value and listening on change.
//
// `data-bind` does not inspect the event payload. It only uses the configured event as a signal to re-read the element’s current bound property/value.
// If you need to pull data from event itself, use `data-on:*` instead.
//
// <input data-bind:foo />
//
// The signal name can be specified in the key (as above), or in the value (as below). This can be useful depending on the templating language you are using.
//
// <input data-bind="foo" />
//
// Attribute casing rules apply to the signal name. Both of these create the signal $fooBar:
//
// <input data-bind:foo-bar />
// <input data-bind="fooBar" />
//
// The initial value of the signal is set to the value of the element, unless a signal has already been defined. So in the example below, $fooBar is set to baz.
//
// <input data-bind:foo-bar value="baz" />
//
// Whereas in the example below, $fooBar inherits the value fizz of the predefined signal.
//
// <div data-signals:foo-bar="'fizz'">
// <input data-bind:foo-bar value="baz" />
// </div>
//
// Input fields of type file will automatically encode file contents in base64. This means that a form is not required.
//
// <input type="file" data-bind:files multiple />
//
// Modifiers allow you to modify behavior when binding signals using a key, so passing any modifier emits the key form of the attribute.
// Datastar then reads the signal name from the key, which HTML lowercases, and converts it to camel case,
// so give the name in kebab case: "my-signal" becomes $mySignal.
//
// Native form controls use their built-in binding semantics automatically. Generic custom elements default to value and change.
// Use [Prop] and [Event] when a custom element’s live state is stored somewhere else.
//
// <my-toggle data-bind:is-checked__prop.checked__event.change></my-toggle>
//
// [Prop] and [Event] are available since Datastar v1.0.0, and can be used independently of each other since v1.0.1.
//
// See https://data-star.dev/reference/attributes#data-bind
func Bind(name string, modifiers ...Modifier) g.Node {
	if len(modifiers) == 0 {
		return data("bind", name)
	}
	nameWithModifiers := name
	for _, modifier := range modifiers {
		nameWithModifiers += string(modifier)
	}
	return data("bind:" + nameWithModifiers)
}

// Class adds or removes a class to or from an element based on an expression.
//
// <div data-class:hidden="$foo"></div>
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
// <button data-on:click="@get('/endpoint')" data-indicator="fetching" data-attr:disabled="$fetching"></button>
// <div data-show="$fetching">Loading...</div>
//
// When using data-indicator with a fetch request initiated in a data-init attribute, you should ensure that the indicator signal is created before the fetch request is initialized.
//
// <div data-indicator:fetching data-init="@get('/endpoint')"></div>
//
// Modifiers allow you to modify behavior when defining indicator signals using a key,
// so passing any modifier emits the key form of the attribute.
// Datastar then reads the signal name from the key, which HTML lowercases, and converts it to camel case,
// so give the name in kebab case: "my-signal" becomes $mySignal.
//
// See https://data-star.dev/reference/attributes#data-indicator
func Indicator(name string, modifiers ...Modifier) g.Node {
	if len(modifiers) == 0 {
		return data("indicator", name)
	}
	nameWithModifiers := name
	for _, modifier := range modifiers {
		nameWithModifiers += string(modifier)
	}
	return data("indicator:" + nameWithModifiers)
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

// Nonce enables CSP mode. Place it on the `html` element.
// Its value must match the nonce in your CSP’s `script-src` directive.
// Generate a new cryptographically secure random nonce on the server for every full-page response.
//
//	<html data-nonce="{page-nonce}">
//	    <head>
//	        <meta http-equiv="Content-Security-Policy"
//	            content="script-src 'self' 'nonce-{page-nonce}';"
//	        >
//	        <script type="module" src="/datastar.js"></script>
//	    </head>
//	    <body>
//	        <button data-on:click="$count++">Increment</button>
//	    </body>
//	</html>
//
// Without it, Datastar uses the `Function()` constructor to evaluate expressions,
// and the Content Security Policy for that mode must include `unsafe-eval`.
//
// Datastar reads the nonce and removes the `data-nonce` attribute.
// It applies the nonce when compiling expressions and when executing scripts received in element patches or JavaScript responses.
// Element patch responses do not need to include the nonce.
//
// CSP mode does not make Datastar expressions safe to use with untrusted content.
//
// Available since Datastar v1.0.3.
// Panics if the value is empty, because Datastar throws on an empty `data-nonce` and then never initializes.
//
// See https://data-star.dev/reference/security#csp-mode
func Nonce(value string) g.Node {
	if value == "" {
		panic("nonce must not be empty")
	}
	return data("nonce", value)
}

// On attaches an event listener to an element, executing an expression whenever the event is triggered.
//
// <button data-on:click="$foo = ' '">Reset</button>
//
// An evt variable that represents the event object is available in the expression.
//
// <div data-on:my-event="$foo = evt.detail"></div>
//
// The `data-on` attribute works with events and custom events. The `data-on:submit` event listener prevents the default submission behavior of forms.
//
// [ModifierWindow] attaches the event listener to the window element, and [ModifierDocument] attaches it to the document element.
// The latter is useful for events that are only available on document and that do not bubble.
//
// <button data-on:fullscreenchange__document="$fullscreen = !$fullscreen"></button>
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
// <div data-ref:foo></div>
//
// The signal name can be specified in the key (as above), or in the value (as below). This can be useful depending on the templating language you are using.
//
// <div data-ref="foo"></div>
//
// The signal value can then be used to reference the element.
//
// $foo is a reference to a <span data-text="$foo.tagName"></span> element
//
// Modifiers allow you to modify behavior when defining references using a key,
// so passing any modifier emits the key form of the attribute.
// Datastar then reads the signal name from the key, which HTML lowercases, and converts it to camel case,
// so give the name in kebab case: "my-signal" becomes $mySignal.
//
// <div data-ref:my-signal__case.kebab></div>
//
// See https://data-star.dev/reference/attributes#data-ref
func Ref(name string, modifiers ...Modifier) g.Node {
	if len(modifiers) == 0 {
		return data("ref", name)
	}
	nameWithModifiers := name
	for _, modifier := range modifiers {
		nameWithModifiers += string(modifier)
	}
	return data("ref:" + nameWithModifiers)
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

// Signal patches (adds, updates or removes) a single signal into the existing signals, using the key form of the attribute.
// Values defined later in the DOM tree override those defined earlier.
//
// <div data-signals:foo="1"></div>
//
// Signals can be nested using dot-notation.
//
// <div data-signals:foo.bar="1"></div>
//
// Setting a signal's value to null or undefined removes the signal, so a nil value removes it,
// unless [ModifierIfMissing] is given, which suppresses the removal.
//
// <div data-signals:foo="null"></div>
//
// The value is JSON encoded. Datastar does not parse it as JSON: it evaluates the attribute value as a Datastar expression,
// which JSON survives because it is a subset of the JavaScript object notation that syntax is built on.
// Datastar rewrites its own action syntax first, though, without regard for string literals,
// so a string value containing an @ followed by a name and an opening parenthesis does not survive.
//
// Keys used in data-signals:* are converted to camel case, so the signal name mySignal must be written as data-signals:my-signal.
// HTML lowercases attribute names, so give the name in kebab case.
//
// Signals beginning with an underscore are not included in requests to the backend by default.
// You can opt to include them by modifying the value of the filterSignals option.
//
// Signal names cannot begin with nor contain a double underscore (__), due to its use as a modifier delimiter.
//
// Modifiers allow you to modify behavior when patching signals using a key, so both [ModifierCase] and [ModifierIfMissing] apply here.
// [ModifierCase] replaces the camel case conversion rather than adding to it, and Datastar has no kebab conversion,
// so [ModifierKebab] leaves the name exactly as written.
//
//	<div data-signals:my-signal__case.kebab="1"
//	     data-signals:foo__ifmissing="1"
//	></div>
//
// Panics if the value cannot be JSON encoded, such as a NaN or infinite float, a channel, or a function.
//
// See https://data-star.dev/reference/attributes#data-signals
func Signal(name string, value any, modifiers ...Modifier) g.Node {
	nameWithModifiers := name
	for _, modifier := range modifiers {
		nameWithModifiers += string(modifier)
	}
	return data("signals:"+nameWithModifiers, toJSON(value))
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
// Datastar converts the casing of a signal name it reads from the attribute key, and this form puts no name there,
// so [ModifierCase] and the casing modifiers are still emitted but never applied. Use [Signal] to apply them to a single signal.
// [ModifierIfMissing] applies to both forms.
//
// Panics if the signals cannot be JSON encoded, such as a NaN or infinite float, a channel, or a function.
//
// See https://data-star.dev/reference/attributes#data-signals
func Signals(signals map[string]any, modifiers ...Modifier) g.Node {
	nameWithModifiers := ""
	for _, modifier := range modifiers {
		nameWithModifiers += string(modifier)
	}
	return data("signals"+nameWithModifiers, toJSON(signals))
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
// <div style="color: red;" data-style:color="$x && 'green'"></div>
//
// <!-- When $hiding is true, display becomes none; when false, reverts to flex from inline style -->
// <div style="display: flex;" data-style:display="$hiding && 'none'"></div>
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

func toJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal value: %v", err))
	}
	return string(b)
}

func data(name string, value ...string) g.Node {
	if len(value) > 0 {
		return html.Data(name, value[0])
	}
	return g.Attr("data-" + name)
}
