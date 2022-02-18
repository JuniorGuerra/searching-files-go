package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func createFile(paths []string, fileName string) {
	ext := filepath.Ext(fileName)

	var commentType string
	switch ext {
	case ".sql":
		commentType = "-- "

	case ".go":
		commentType = "// "

	case ".bash":
		commentType = "# "

	case ".vim":
		commentType = "# "

	case ".init":
		commentType = "# "

	default:
		commentType = "// "

	}

	file, err := os.Create(fileName)
	check(err)

	writer := bufio.NewWriter(file)
	for _, path := range paths {
		data, err := os.ReadFile(path)
		check(err)

		comment := fmt.Sprintf(`
%s #############################################
%s Start: %s
%s #############################################

`, commentType, commentType, path, commentType)

		text := fmt.Sprintf("%s\n", data)

		writer.WriteString(comment)
		writer.WriteString(text)

		defer file.Close()

	}
	writer.Flush()
}
