package intermediate

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// err := os.Mkdir("subdir", 0755)
	// checkError(err)
	// // defer os.RemoveAll("./subdir")
	// os.WriteFile("subdir/trial.txt", []byte("Hello World!"), 0755)

	// checkError(os.MkdirAll("subdir/parent/child1", 0755))
	// checkError(os.MkdirAll("subdir/parent/child2", 0755))
	// checkError(os.MkdirAll("subdir/parent/child3", 0755))

	// os.WriteFile("subdir/parent/file1", []byte("File1 let see what happens without txt!"), 0755)
	// os.WriteFile("subdir/parent/child3/file_child_3", []byte("File child 3 let see what happens without txt!"), 0755)

	result, err := os.ReadDir("subdir/parent")
	checkError(err)
	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}
	checkError(os.Chdir("subdir/parent/child3"))
	result, err = os.ReadDir(".")
	checkError(err)

	for _, entry := range result {
		fmt.Println(entry.Name())
	}
	checkError(os.Chdir("../../.."))
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)

	// We got 2 different variations of filepath
	// filepath.Walk and filepath.Walkdir
	pathFile := "subdir"
	err = filepath.WalkDir(pathFile, func(path string, d os.DirEntry, err error) error {
		checkError(err)
		fmt.Println(path)
		return nil
	})
	checkError(err)
	checkError(os.RemoveAll("./subdir"))
}
