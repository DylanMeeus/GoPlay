package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	run()
}

func run() {
	w := js.Global.Get("innerWidth").Float()
	h := js.Global.Get("innerHeight").Float()

	fmt.Printf("%v %v\n", w, h)
}
