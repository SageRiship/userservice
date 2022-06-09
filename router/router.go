package router

import (
	"github.com/SageRiship/userservice/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", controller.GetUserById).Methods("GET")
	router.HandleFunc("/api/user/name/{name}", controller.GetUserByName).Methods("GET")
	router.HandleFunc("/api/user/userid/{userid}", controller.GetUserByUserId).Methods("GET")
	router.HandleFunc("/api/user", controller.AddUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", controller.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/api/user/userid/{userid}", controller.DeleteUserByUserId).Methods("DELETE")
	router.HandleFunc("/api/user", controller.DeleteAllUser).Methods("DELETE")
	return router
}
