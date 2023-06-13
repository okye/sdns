//go:build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// middleware list order very important, handlers call via this order.
var middlewareList = []string{
	"recovery",
	"loop",
	"metrics",
	"accesslist",
	"ratelimit",
	"edns",
	"accesslog",
	"chaos",
	"hostsfile",
	"blocklist",
	"as112",
	"cache",
	"failover",
	"resolver",
	"forwarder",
}

func main() {
	var pathlist []string
	for _, name := range middlewareList {
		stat, err := os.Stat(filepath.Join(middlewareDir, name))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if !stat.IsDir() {
			fmt.Println("path is not directory")
			os.Exit(1)
		}
		pathlist = append(pathlist, filepath.Join(prefixDir, middlewareDir, name))
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	file.WriteString("// Code generated by gen.go DO NOT EDIT.\n")

	file.WriteString("\npackage main\n\nimport (\n")

	for _, path := range pathlist {
		file.WriteString("\t_ \"" + path + "\"\n")
	}

	file.WriteString(")")
}

const (
	filename      = "generated.go"
	prefixDir     = "github.com/semihalev/sdns"
	middlewareDir = "middleware"
)
