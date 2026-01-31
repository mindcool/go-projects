package intermediate

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed sample.txt
var content string

//go:embed basics
var basicsFolder embed.FS

func main() {
	fmt.Println("Embeded Content: ", content)
	content, err := basicsFolder.ReadFile("basics/hello.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	err = fs.WalkDir(basicsFolder, "basics", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
