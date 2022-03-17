package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
)

func remove(path string, destDir string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "file/directory %v does not exist.\n", path)
			return nil
		}
		return err
	}

	return os.Rename(path, filepath.Join(destDir, filepath.Base(path)))
}

const prefix = "removed"

func cleanCmd(arg string) error {
	return os.RemoveAll(arg)
}

func main() {
	clean := pflag.Bool("clean", false, "definitively remove args")
	pflag.Parse()
	args := pflag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "move files/directories into %v, like mv unix command\n", os.TempDir())
		fmt.Fprintf(os.Stderr, "usage: %v file1|dir1 [file2|dir2 ...]\n", os.Args[0])
	}

	var dir string

	if !(*clean) {
		var err error
		dir, err = ioutil.TempDir("", prefix)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not create tempory directory\n")
			os.Exit(1)
		}
	}

	for _, arg := range args {
		var err error
		var callback func()

		switch *clean {
		case true:
			err = cleanCmd(arg)
			callback = func() {
				fmt.Println("delete file or dir:", arg)
			}
		case false:
			err = remove(arg, dir)
			callback = func() {
				fmt.Println("move file or dir:", arg, "to", dir)
			}
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not remove %v : %v\n", arg, err)
		} else {
			callback()
		}

	}
}
