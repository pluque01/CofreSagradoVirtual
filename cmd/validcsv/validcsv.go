package main

import (
	"net/http"

	"github.com/pluque01/CofreSagradoVirtual/internal/api"
	log "github.com/pluque01/CofreSagradoVirtual/internal/logger"
)

func main() {
	http.ListenAndServe(":8080", api.GetServeMux())
	log.Close()
}
