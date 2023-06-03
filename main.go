package main

import (
	"flag" //https://gobyexample.com/command-line-flags

	"github.com/fatih/color"
)

func main() {

	var lang string
	flag.StringVar(&lang, "language", "error", "language for project (rust, go, javascript)")
	flag.StringVar(&lang, "lang", "error", "language for project (shorthand)")
	flag.StringVar(&lang, "l", "error", "language for project (shorthand)")

	var name string
	flag.StringVar(&name, "name", "hello_world", "name for project")
	flag.StringVar(&name, "n", "hello_world", "name for project (shorthand)")

	var code bool
	flag.BoolVar(&code, "code", false, "open project in vscode")
	flag.BoolVar(&code, "c", false, "open project in vscode (shorthand)")

	var mkdir bool
	flag.BoolVar(&mkdir, "mkdir", false, "make directory for project (not for rust)")
	flag.BoolVar(&mkdir, "m", false, "make directory for project (shorthand) (not for rust)")
	flag.BoolVar(&mkdir, "d", false, "make directory for project (shorthand) (not for rust)")

	var overwrite bool
	flag.BoolVar(&overwrite, "overwrite", false, "overwrite existing project")
	flag.BoolVar(&overwrite, "o", false, "overwrite existing project (shorthand)")

	flag.Parse()

	if lang == "error" {
		color.Red("Error: Language not specified")
	}

	switch lang {
	case "go":
		goProject(name, code, mkdir, overwrite)
	case "rust":
		rustProject(name, code, overwrite)
	case "javascript":
		npmProject(name, code, mkdir, overwrite)
	case "js":
		npmProject(name, code, mkdir, overwrite)

	}

}
