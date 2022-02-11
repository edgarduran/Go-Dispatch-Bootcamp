package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", func(w http.ResponseWriter, rtr *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test rtr resp"))
	})

	return rtr
}
