package main

import (
	"flag"
	"fmt"
)

func main() {

	filename := flag.String("f", "filecreate.txt", "flag for creating file with given filename")
	path := flag.String("p", "./", "path for searching files")
	filesToSearch := flag.String("ext", "*", "extension to searching files")
	flag.Parse()

	*filesToSearch = fmt.Sprintf("*.%s", *filesToSearch)
	paths, err := WalkMatch(*path, *filesToSearch)
	check(err)

	createFile(paths, *filename)
}
