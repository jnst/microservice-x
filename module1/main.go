package main

import (
	"fmt"

	"github.com/jnst/microservice-x/module1/db"
	"github.com/samber/lo"
)

func Uniq() []string {
	return lo.Uniq[string]([]string{"Samuel", "John", "Samuel"})
}

func main() {
	//names := Uniq()
	//fmt.Println(names)
	//
	//fmt.Println(reverse.String("Hello, world!"))
	fmt.Println(db.Now())
}
