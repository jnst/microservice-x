package main

import (
	"fmt"
	"github.com/jnst/microservice-x/module1/db"
	"github.com/samber/lo"
	"golang.org/x/example/hello/reverse"
)

func Uniq() []string {
	return lo.Uniq[string]([]string{"Samuel", "John", "Samuel"})
}

func main() {
	//names := Uniq()
	//fmt.Println(names)

	fmt.Println(reverse.String("Hello, world!"), reverse.Int(1234567890))
	fmt.Println(db.Now())
}
