package main

import "github.com/Linkinlog/LeafListr/internal/api"

func main() {
	addr := ":8080"
	timeout := 3
	err := api.ListenAndServe(addr, timeout)
	if err != nil {
		panic(err)
	}
}
