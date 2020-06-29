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
		fmt.Println(ob.Get("filteruser").String())
		fmt.Println(ob.Get("id").String())
		fmt.Println(ob.Get("id2").String())
	}
}
