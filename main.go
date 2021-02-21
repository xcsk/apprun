package main

import (
	"fmt"
	"time"

	bl "github.com/xcsk/breakline"
	tl "github.com/xcsk/breakline/title"
)

func main() {
	fmt.Printf("⏰: %s\n", time.Now().Format("2006-01-02T15:04:05-0700"))
	fmt.Println(bl.BreakLine(28))

	for i := 0; i < 3; i++ {
		fmt.Println(tl.Nao())
		fmt.Println(tl.Aika())
	}
}
