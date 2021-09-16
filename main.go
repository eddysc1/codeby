package main

import (
	"fmt"
	"os"
)

func main() {
	if (len(os.Args) == 1 || len(os.Args) > 2) {
		fmt.Println("usage: codeby \"full name\"")
	} else {
		fmt.Println(addedRemovedLoc(os.Args[1]))
	}
}
