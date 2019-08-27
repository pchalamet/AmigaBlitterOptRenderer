package main

// "go.inferGopath": false,
// "go.useLanguageServer": true,
// "go.formatTool": "goreturns",

import (
	"fmt"
	"github.com/pchalamet/AmigaBlitterOptRenderer/foundation"
)

func main() {
	vec := foundation.NewVector3(1, 2, 3)
	fmt.Println(vec.X)
	fmt.Println("Hello")
}
