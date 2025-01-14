package main

import (
	"fmt"

	v8 "rogchap.com/v8go"
)

func main() {
	ctx := v8.NewContext()
	ctx.RunScript("const add = (a, b) => a + b", "math.js")
	ctx.RunScript("const result = add(3, 4)", "main.js")
	val, _ := ctx.RunScript("result", "value.js")
	fmt.Printf("addition result: %s\n", val)
}
