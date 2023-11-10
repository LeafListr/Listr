package main

import "github.com/Linkinlog/LeafListr/internal/api"

func main() {
	err := api.ListenAndServe(":8080")
	if err != nil {
		panic(err)
	}
}
