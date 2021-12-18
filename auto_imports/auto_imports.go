package main

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/emirpasic/gods/sets/treeset"
)

func main() {
	dirs := treeset.NewWithStringComparator()
	fset := token.NewFileSet()
	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			p, err := parser.ParseDir(fset, path, func(info os.FileInfo) bool {
				return true
			}, parser.AllErrors)
			if err != nil {
				panic(err)
			}
			for _, v := range p {
				for _, v2 := range v.Files {
					for _, v3 := range v2.Imports {
						if v3.Path.Value == "\"git.xios.club/xios/gc\"" {
							dirs.Add(path)
						}
					}
				}
			}
		}
		return nil
	})

	imports := ""
	for _, v := range dirs.Values() {
		path := v.(string)
		if strings.HasPrefix(path, "./") {
			continue
		}
		if runtime.GOOS == "windows" {
			imports += "\t_ \"admin/" + strings.ReplaceAll(path, "\\", "/") + "\"\n"
		} else {
			imports += "\t_ \"admin/" + path + "\"\n"
		}
	}

	f, err := os.OpenFile("./imports.go", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString("package main\n\nimport (\n" + imports + ")")
	if err != nil {
		panic(err)
	}
}
