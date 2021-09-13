package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func remove(path string, destDir string) error {
	return os.Rename(path, filepath.Join(destDir, filepath.Base(path)))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "move files/directories into %v, like mv unix command\n", os.TempDir())
		fmt.Fprintf(os.Stderr, "usage: %v file1|dir1 [file2|dir2 ...]\n", os.Args[0])
	}

	dir, err := ioutil.TempDir("", "removed")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create tempory directory\n")
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		if err := remove(arg, dir); err != nil {
			fmt.Fprintf(os.Stderr, "could not move %v into %v: %v\n", arg, dir, err)
		}
	}
}
