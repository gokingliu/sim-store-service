package main

import (
	"fmt"
	"net/http"
	"sim/service"
)

func main() {
	service.Init()

	err := http.ListenAndServe(":8688", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v", err)
	}
}
