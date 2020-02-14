package internal

import "syscall/js"

// GetElementByID is javascript getElementById
func GetElementByID(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

// CreateElement javascript createElement
func CreateElement(name string) js.Value {
	return js.Global().Get("document").Call("createElement", name)
}
