package main

import (
	"fmt"
	"os"

	"github.com/NikitaKissa/cvgen/internal/cli"
	"github.com/NikitaKissa/cvgen/internal/template"
)

func main() {
	// template import to binary
	_ = template.DefaultTemplateHTML
	_ = template.DefaultStyleCSS
	// ^^^ DON'T touch this part ^^^

	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
