package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"learn-go-api/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)	//	turn on
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)		//	handler function which will set up the router, ie add endpoints

	fmt.Println("Starting GO API service....")

	//listening

	err := http.ListenAndServe("localhost:8000",r)
	if err != nil{
		log.Error(err)
	}
}