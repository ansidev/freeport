package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ansidev/freeport"
)

func main() {
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strconv.Itoa(port))
}
