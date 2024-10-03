package main

import (
	"fmt"

	"github.com/mshirdel/sandbox/brokers"
)

func main() {
	if err := brokers.Run("hello from golang"); err != nil {
		fmt.Println(err)
		return 
	}
	
	fmt.Println("OK")
}
