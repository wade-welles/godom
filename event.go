// This file implements Event interface
// https://developer.mozilla.org/en-US/docs/Web/API/Event
package godom

import (
	"github.com/gopherjs/gopherjs/js"
)

type Event struct {
	*js.Object
}

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/Event/target
func (e Event) Target() *Object {
	return &Object{e.Get("target")}
}

// Methods