package main

import (
	"fmt"
	"log"
	"os"

	v8 "rogchap.com/v8go"
)

func main() {
	iso := v8.NewIsolate()
	defer iso.Dispose()

	ctx := v8.NewContext(iso)
	defer ctx.Close()

	jsCode, err := os.ReadFile("calculator.js")
	if err != nil {
		log.Fatal(err)
	}

	_, err = ctx.RunScript(string(jsCode), "calculator.js")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := ctx.RunScript("globalResults.sum", "get_sum.js")
	if err != nil {
		_ = fmt.Errorf("Error getting sum: %v\n", err)
	}

	difference, err := ctx.RunScript("globalResults.difference", "get_difference.js")
	if err != nil {
		_ = fmt.Errorf("Errot getting the difference: %v\n", err)
	}

	product, err := ctx.RunScript("globalResults.product", "get_product.js")
	if err != nil {
		_ = fmt.Errorf("Errot getting the product: %v\n", err)
	}

	fmt.Printf("Sum: %v\n", sum)
	fmt.Printf("Difference: %v\n", difference)
	fmt.Printf("Product: %v\n", product)
}
