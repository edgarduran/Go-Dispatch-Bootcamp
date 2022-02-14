package router

import (
	"net/http"

	"github.com/edgarduran/Go-Dispatch-Bootcamp/controllers"
	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/vehicles", controllers.GetAllVehicles).Methods(http.MethodGet)

	return rtr
}
