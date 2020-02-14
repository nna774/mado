package debug

import (
	"syscall/js"

	"github.com/nna774/mado/internal"
)

// Debugger is debug class
type Debugger struct {
	jsValue js.Value
}

// NewDebugger creates debugger
func NewDebugger(id string) Debugger {
	debug := internal.GetElementByID(id)
	ul := internal.CreateElement("ul")
	debug.Call("appendChild", ul)
	return Debugger{jsValue: ul}
}

// Log makes log
func (d *Debugger) Log(msg string) {
	li := internal.CreateElement("li")
	d.jsValue.Call("appendChild", li)
	li.Set("innerText", msg)
}
