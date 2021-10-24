package mik

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Is struct {
	Path     string
	FileName string
}

func (g *Is) Yield(input Is) (string, error) {
	g.createFolder(input)
	g.createFile(input)

	writes, err := g.writeFile(input)
	if err != nil {
		log.Fatal(err)
		return writes, err
	}

	g.readFile(input)

	return writes, nil

}

func (g *Is) createFolder(input Is) error {

	f := Is{}
	f.Path = filepath.Join("./", input.Path)
	f.FileName = input.FileName

	_, err := os.Stat(f.Path)

	if os.IsExist(err) {
		fmt.Println("your directory is already exist but that's ok")
		return err
	}

	err = os.Mkdir(f.Path, 0755)
	if err != nil {
		// log.Fatal(err)
		fmt.Println("your directory is already exist but that's ok")
		return err
	}

	return nil
}

func (g *Is) createFile(input Is) (string, error) {

	f := Is{}
	f.Path = filepath.Join("./", input.Path)
	f.FileName = input.FileName

	filepath, err := filepath.Abs(f.Path + "/" + f.FileName + ".go")
	if err != nil {
		log.Fatal("error")
		return filepath, err
	}

	filename, err := os.Create(filepath)

	if err != nil {
		log.Fatal("Cannot create a file please check your directory again")
		return filename.Name(), err
	}

	fmt.Printf("Create %s is successfully \n", filename.Name())
	return filepath, nil
}

func (g *Is) writeFile(input Is) (string, error) {

	f := Is{}
	f.Path = filepath.Join("./", input.Path)
	f.FileName = input.FileName
	path, _ := filepath.Abs(f.Path + "/" + f.FileName + ".go")

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return "", err
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(fmt.Sprintf("package %s", input.Path))
	if isError(err) {
		return "", err
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return "", err
	}

	fmt.Printf("Rewrite file is successfully \n")
	return f.FileName, nil
}

func (g *Is) readFile(input Is) {
	f := Is{}
	f.Path = filepath.Join("./", input.Path)
	f.FileName = input.FileName

	var file, err = os.OpenFile(f.Path+"/"+f.FileName+".go", os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Printf("Reading from file %s", f.FileName)
	fmt.Println()
	fmt.Println(string(text))
}
func isError(err error) bool {

	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
