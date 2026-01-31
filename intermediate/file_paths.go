package intermediate

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	relativePath := "./data/trial.txt"
	absolutePath := "/home/data/trial.txt"

	// Join paths using filepath.join
	joinedPath := filepath.Join("downloads", "file.zip")
	fmt.Println("Joined Path: ", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path: ", normalizedPath)

	dir, file := filepath.Split("/home/noyan/trial.txt")
	fmt.Println("Directory: ", dir)
	fmt.Println("File: ", file)
	fmt.Println(filepath.Base("/home/noyan/trial.txt"))

	// Check if a path is absolute
	fmt.Println("Is Relative Path is Absolute", filepath.IsAbs(relativePath))
	fmt.Println("Is Absolute Path is Absolute", filepath.IsAbs(absolutePath))

	// Get the extension of file
	fmt.Println("Extension of file: ", filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	// Lets convert a path relative to another path
	rel, err := filepath.Rel("a/c", "a/b/text/test.txt")
	if err != nil {
		return
	}
	fmt.Println(rel)

	// Lets convert a path to absolute path
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		return
	}
	fmt.Println("Absolute Path: ", absPath)
}
