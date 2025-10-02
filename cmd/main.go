package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MaximGubanov/aup-config"
)

func main() {
	execPath, _ := os.Executable()
	execDir := filepath.Dir(execPath)
	fmt.Printf("%s\n", execDir)
	cfg, err := aup_config.NewConfig("../sgs.json", execDir)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", cfg)
}
