//go:build wasm && js

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	done := make(chan struct{})
	go foobar()
	<-done
}

func foobar() {
	fmt.Println("Hello from Web Assembly")

	document := js.Global().Get("document")
	appendToDoc := func(el js.Value) {
		document.Call("getElementById", "main").Call("appendChild", el)	
	}	

	document.Get("body").Get("classList").Call("add", "background") // add some css

    appendToDoc(createButton(&document))
    appendToDoc(document.Call("createElement", "br"))
    appendToDoc(createCanvas(&document))
}

func createButton(document *js.Value) js.Value {
	but := document.Call("createElement", "button")
    but.Set("innerHTML", "Hello WASM from Go!")
	
	// todo: required?
	// var f js.Func
    but.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        // defer f.Release()
		fmt.Println("click")
        return nil
    }))

	return but
}

func createCanvas(document *js.Value) js.Value {
	width := 200
	height := 100

	canv := document.Call("createElement", "canvas")
	canv.Set("width", width)
	canv.Set("height", height)
	canv.Set("style", "border:1px solid #000000;")
	canv.Set("id", "myCanvas")

	cxt := canv.Call("getContext", "2d")
	cxt.Call("moveTo", 0, 0)
	cxt.Call("lineTo", width, height)
	cxt.Call("stroke")

	return canv
}