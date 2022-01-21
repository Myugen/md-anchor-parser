package main

import (
	"github.com/myugen/md-anchor-parser/cmd"
	"log"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
