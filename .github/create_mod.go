// Command create_mod builds a Go module zip in proxy.golang.org layout using
// golang.org/x/mod/zip.CreateFromDir.
//
// Usage: create_mod <module-path> <version> <dir> <output-zip>
package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "usage: %s <module-path> <version> <dir> <output-zip>\n", os.Args[0])
		os.Exit(2)
	}
	modPath := os.Args[1]
	version := os.Args[2]
	dir := os.Args[3]
	out := os.Args[4]

	f, err := os.Create(out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create output: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	m := module.Version{Path: modPath, Version: version}
	if err := zip.CreateFromDir(f, m, dir); err != nil {
		fmt.Fprintf(os.Stderr, "CreateFromDir: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("wrote %s\n", out)
}
