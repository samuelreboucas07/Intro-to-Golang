package routes

import (
	"api-rest/controllers"
	"api-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentType)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/allPersonalities", controllers.AllPersonality).Methods("Get")
	r.HandleFunc("/api/getPersonalityById/{id}", controllers.GetPersonalityById).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.RemovePersolatity).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonlaty).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
