package openShiftSetup

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type HTTPserver struct {
	ServeFiles string
	Dir        string
}

func HTTPserve() {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("/usr/local/http/"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
