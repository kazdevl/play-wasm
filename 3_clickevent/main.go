//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	window := js.Global()
	message := window.Get("document").Call("getElementById", "message")

	var on bool = false
	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		on = !on
		if on {
			message.Set("innerHTML", "Clicked!!")
		} else {
			message.Set("innerHTML", "tern off")
		}
		// cb.Release()
		fmt.Printf("on: %v\n", on)
		return nil
	})

	message.Call("addEventListener", "click", cb)

	select {}
}
