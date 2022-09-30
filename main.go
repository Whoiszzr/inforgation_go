package main

import (
	"fmt"
	"inforgation_go/modules"
)

func main() {
	err := modules.Fofa("114.114.114.114", "**********", "**********@*****.com")
	if err != nil {
		fmt.Println("FOFA查询失败！" + err.Error())
	} else {
		fmt.Println("OK")
	}
}
