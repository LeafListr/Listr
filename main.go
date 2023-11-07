package main

import "github.com/Linkinlog/LeafList/internal/api"

func main() {
	err := api.ListenAndServe("0.0.0.0:4200")
	if err != nil {
		panic(err)
	}
}
