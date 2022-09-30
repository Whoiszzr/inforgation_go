package main

import (
	"fmt"
	"inforgation_go/modules"
)

func main() {
	err := modules.Fofa("114.114.114.114", "**********", "**********@*****.com")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}
}
