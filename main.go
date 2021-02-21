package main

import (
	"fmt"
	"time"

	bl "github.com/xcsk/breakline"
)

func main() {
	fmt.Printf("⏰: %s\n", time.Now().Format("2006-01-02T15:04:05-0700"))
	fmt.Println(bl.BreakLine(28))
}
