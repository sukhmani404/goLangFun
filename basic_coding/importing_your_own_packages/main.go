package main

import (
	"fmt"

	"github.com/sukhmani404/goLangFun/basic_coding/importing_your_own_packages/strutil"
)

func main() {
	fmt.Println("Hello there my friend this is how you do stuff~!!")
	var name = "Sukhmani"
	fmt.Println(name)
	fmt.Println(strutil.Reverse("sukhmani"))
}
