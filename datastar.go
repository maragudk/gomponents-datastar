// Package datastar provides Datastar attributes and helpers for gomponents.
// See https://data-star.dev
package datastar

import (
	"fmt"

	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

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
	return html.Data("attr", toObject(pairs))
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
	return html.Data("bind", name)
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
	return html.Data("class", toObject(pairs))
}

// Text binds the text content of an element to an expression.
//
// <div data-text="$foo"></div>
//
// See https://data-star.dev/reference/attributes#data-text
func Text(v string) g.Node {
	return html.Data("text", v)
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
