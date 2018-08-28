package main

import (
	"flag"
	"os"
	"path/filepath"

	pb "gopkg.in/cheggaaa/pb.v1"
)

func main() {
	directories := []string{}
	files := []string{}

	root := os.Args[1]

	flag.Parse()

	if root == "" {
		flag.Usage()
		return
	}
	root = filepath.Clean(root)

	stats, err := os.Lstat(root)
	if err != nil {
		panic(err)
	}

	if stats.IsDir() {
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				directories = append(directories, path)
			} else {
				files = append(files, path)
			}
			return nil
		})
	} else {
		os.Remove(root)
	}

	bar := pb.StartNew(len(files) + len(directories))
	for _, file := range files {
		if err := os.Remove(file); err != nil {
			panic(err)
		}
		bar.Increment()
	}
	for i := len(directories) - 1; i >= 0; i-- {
		if err := os.Remove(directories[i]); err != nil {

			panic(err)
		}
		bar.Increment()
	}
	bar.Finish()

}
