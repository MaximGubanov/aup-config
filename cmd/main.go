package main

import (
	"fmt"
	"os"

	"github.com/MaximGubanov/aup-config"
)

func main() {
	execDir, err := os.Getwd()
	fmt.Printf("%s\n", execDir)
	cfg, err := aup_config.NewConfig(execDir, "../sgs.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", cfg)
}
