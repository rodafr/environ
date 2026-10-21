package main

import (
	"fmt"
	"os"

	"github.com/rodafr/environ/pkg/environ"
)

func main() {
	load := environ.NewLoader(os.LookupEnv)
	fmt.Println(load.Str("EDITOR", "nvim"))
	load.Str("SOMETHINGTHATHISNOTFOUND")
	fmt.Println(load.Err().Error())
}
