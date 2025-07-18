package main

import (
	"fmt"
	"os"

	"github.com/helson-lin/of/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}
}
