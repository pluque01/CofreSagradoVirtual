package main

import (
	"fmt"

	"github.com/pluque01/CofreSagradoVirtual/internal/validcsv"
)

func main() {
	clientFile := validcsv.NewClientFile("test/data/validate_file_content.csv", ';')
	clientFile.Print()
	results := clientFile.ValidateFileContent()
	fmt.Println(results)
}
