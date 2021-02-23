package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = *flag.String("name", "everyone", "The greeting object.")


	fmt.Printf(name)
}
