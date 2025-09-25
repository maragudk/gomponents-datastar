// Package datastar provides Datastar attributes and helpers for gomponents.
// See https://data-star.dev
package datastar

import (
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
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
func Attr(name, v string) g.Node {
	return Data("attr-"+name, v)
}
