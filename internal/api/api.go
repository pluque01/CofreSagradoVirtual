package api

import (
	"log"
	"net/http"
	"os"

	"github.com/pluque01/CofreSagradoVirtual/internal/api/handlers"
)

var mux *http.ServeMux

func newServeMux() *http.ServeMux {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	sm := http.NewServeMux()
	sm.Handle("/clientdata/", handlers.NewClientData(l))
	return sm
}

func GetServeMux() *http.ServeMux {
	if mux == nil {
		mux = newServeMux()
	}
	return mux
}
