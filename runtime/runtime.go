package runtime

import (
	"log"

	"github.com/NorShell/diggers/utils"
	v8 "rogchap.com/v8go"
)

func Execute(file string) {
	// create an isolate
	iso := v8.NewIsolate()
	defer iso.Dispose()

	// collecting the callblacks
	printfn := Print(iso)

	// binding the callback
	global := v8.NewObjectTemplate(iso)
	global.Set("print", printfn)

	// creating a new context, the actual runtime
	ctx := v8.NewContext(iso, global)
	defer ctx.Close()

	jsCode, err := utils.ReadFile(file)

	_, err = ctx.RunScript(string(jsCode), "script.js")
	if err != nil {
		log.Fatal(err)
	}
}
