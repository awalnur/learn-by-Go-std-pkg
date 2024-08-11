package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "Input Database username")
	flag.Parse()
	fmt.Println(*username)
}
