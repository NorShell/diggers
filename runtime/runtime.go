package runtime

import (
	"fmt"
	"log"
	"os"

	v8 "rogchap.com/v8go"
)

func Execute() {
	iso := v8.NewIsolate()
	defer iso.Dispose()

	printfn := v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		fmt.Printf("%v\n", info.Args())
		return nil
	})

	global := v8.NewObjectTemplate(iso)
	global.Set("print", printfn)

	ctx := v8.NewContext(iso, global)
	defer ctx.Close()

	jsCode, err := os.ReadFile("script.js")
	if err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
		return
	}

	_, err = ctx.RunScript(string(jsCode), "script.js")
	if err != nil {
		log.Fatal(err)
	}
}
