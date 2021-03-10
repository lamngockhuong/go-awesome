package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("India", "Indiana"))
	fmt.Println(strings.Compare("Indiana", "India"))
	fmt.Println(strings.Compare("India", "India"))
}