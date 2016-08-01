package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	fmt.Println("Test if go get works on this repo..")
	fmt.Println("Simple BSON: %+v", bson.M{"a": 1, "b": true})
}
