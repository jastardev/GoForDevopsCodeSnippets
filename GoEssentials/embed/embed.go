package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"path/filepath"
)

//go:embed images/*
//go:embed index.html
var content embed.FS

func main() {
	err := fs.WalkDir(
		content,
		".",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() && filepath.Ext(path) == ".png" {
				fmt.Println("png file: ", path)
			}
			return nil
		},
	)
	if err != nil {
		panic(err)
	}
}
