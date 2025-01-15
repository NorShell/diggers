package runtime

import (
	"fmt"

	v8 "rogchap.com/v8go"
)

func Print(iso *v8.Isolate) *v8.FunctionTemplate {
	printFunc := v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		fmt.Printf(" %v\n", info.Args())
		return nil
	})
	return printFunc
}
