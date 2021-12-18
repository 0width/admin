package tests

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"testing"
)

func TestParser(t *testing.T) {
	fset := token.NewFileSet()
	p, err := parser.ParseDir(fset, "deep/tests", func(info os.FileInfo) bool {
		return true
	}, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}

func TestWalk(t *testing.T) {
	os.Chdir("../")
	ps := []string{}
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
						if v3.Path.Value == "\"golang.org/x/crypto/bcrypt\"" {
							ps = append(ps, path)
						}
					}
				}
			}
		}
		return nil
	})
	print("a")
}
