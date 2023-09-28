package main

import (
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/example/hello/reverse"
)

func main() {
	names := lo.Uniq[string]([]string{"Samuel", "John", "Samuel"})
	fmt.Println(names)

	fmt.Println(reverse.String("Hello, world!"))
}
