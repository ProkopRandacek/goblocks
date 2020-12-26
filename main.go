package main

import (
	"fmt"
)

func main() {
	blocks := Load()
	cont := content{}
	cont.blocks = &blocks
		
	for i := range blocks {
		go run(blocks[i], &cont, i)
	}
	fmt.Scanln()
}
