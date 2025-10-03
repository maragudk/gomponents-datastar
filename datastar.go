// Package datastar provides Datastar attributes and helpers for gomponents.
// See https://data-star.dev
package datastar

import (
	"fmt"
	"time"

	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type Modifier string

const (
	ModifierCapture        Modifier = "__capture"
	ModifierCase           Modifier = "__case"
	ModifierDebounce       Modifier = "__debounce"
	ModifierDelay          Modifier = "__delay"
	ModifierOnce           Modifier = "__once"
	ModifierOutside        Modifier = "__outside"
	ModifierPassive        Modifier = "__passive"
	ModifierPrevent        Modifier = "__prevent"
	ModifierSelf           Modifier = "__self"
	ModifierStop           Modifier = "__stop"
	ModifierThrottle       Modifier = "__throttle"
	ModifierViewTransition Modifier = "__viewtransition"
	ModifierWindow         Modifier = "__window"
)

const (
	ModifierCamel     Modifier = ".camel" // Camel case: myEvent
	ModifierKebab     Modifier = ".kebab" // Kebab case: my-event
	ModifierLeading   Modifier = ".leading"
	ModifierNoLeading Modifier = ".noleading"
	ModifierNoTrail   Modifier = ".notrail"
	ModifierPascal    Modifier = ".pascal" // Pascal case: MyEvent
	ModifierSnake     Modifier = ".snake"  // Snake case: my_event
	ModifierTrail     Modifier = ".trail"
)

// ModifierDuration outputs millisecond values for durations under 1 second, otherwise second values.
func ModifierDuration(d time.Duration) Modifier {
	if d.Seconds() < 1 {
		return Modifier(fmt.Sprintf(".%vms", d.Milliseconds()))
	}
	return Modifier(fmt.Sprintf(".%vs", int(d.Seconds())))
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
// <div data-computed-foo="$bar + $baz"></div>
//
// Computed signals are useful for memoizing expressions containing other signals. Their values can be used in other expressions.
//
// <div data-computed-foo="$bar + $baz"></div>
// <div data-text="$foo"></div>
//
// Computed signal expressions must not be used for performing actions (changing other signals, actions, JavaScript functions, etc.).
// If you need to perform an action in response to a signal change, use the data-effect attribute.
//
// See https://data-star.dev/reference/attributes#data-computed
func Computed(name, expression string, modifiers ...Modifier) g.Node {
	nameWithModifiers := name
	for _, modifier := range modifiers {
		nameWithModifiers += string(modifier)
	}
	return data("computed-"+nameWithModifiers, expression)
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
// When using data-indicator with a fetch request initiated in a data-on-load attribute, you should ensure that the indicator signal is created before the fetch request is initialized.
//
// <div data-indicator-fetching data-on-load="@get('/endpoint')"></div>
//
// See https://data-star.dev/reference/attributes#data-indicator
func Indicator(name string, modifiers ...Modifier) g.Node {
	nameWithModifiers := ""
	for _, modifier := range modifiers {
		nameWithModifiers += string(modifier)
	}
	return data("indicator"+nameWithModifiers, name)
}

// Text binds the text content of an element to an expression.
//
// <div data-text="$foo"></div>
//
// See https://data-star.dev/reference/attributes#data-text
func Text(v string) g.Node {
	return data("text", v)
}

// On attaches an event listener to an element, executing an expression whenever the event is triggered.
//
// <button data-on-click="$foo = ”">Reset</button>
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
	return data("on-"+eventWithModifiers, expression)
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

func data(name string, value ...string) g.Node {
	if len(value) > 0 {
		return html.Data(name, value[0])
	}
	return g.Attr("data-" + name)
}
