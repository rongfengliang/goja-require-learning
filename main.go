package main

import (
	"fmt"
	"log"
)

func main() {
	value, _ := myvm.Exec()
	fmt.Println(value)
	m, err := myrequire.Require("app.js")
	if err != nil {
		log.Panic(err)
	} else {
		ob := m.ToObject(myvm.jsRuntime)
		fmt.Print(ob.Get("filteruser").String())
	}
}
