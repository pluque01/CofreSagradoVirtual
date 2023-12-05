package main

import (
	"fmt"

	"github.com/pluque01/CofreSagradoVirtual/internal/config"
	log "github.com/pluque01/CofreSagradoVirtual/internal/logger"
	"github.com/pluque01/CofreSagradoVirtual/internal/validcsv"
)

func main() {
	clientFile := validcsv.NewClientFile("test/data/validate_file_content.csv", ';')
	clientFile.Print()
	results := clientFile.ValidateFileContent()
	fmt.Println(results)

	fmt.Println(config.DefaultConfig)
	log.Close()
}
