package main

// main.go HAS FOUR TODOS - TODO_1 - TODO_4

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"scrape/scrape"
)
var Log_Level = 2;
//There is another Log_Level variable in scrapeapi.go

//TODONE_1: Logging right now just happens, create a global constant integer LOG_LEVEL 
//TODONE_1: When LOG_LEVEL = 0 DO NOT LOG anything
//TODONE_1: When LOG_LEVEL = 1 LOG API details only 
//TODONE_1: When LOG_LEVEL = 2 LOG API details and file matches (e.g., everything)

func main() {
	if(Log_Level != 0){
		log.Println("starting API server")
	}
	//create a new router
	router := mux.NewRouter()
	if(Log_Level != 0){
		log.Println("creating routes")
	}
	//specify endpoints
	router.HandleFunc("/", scrape.MainPage).Methods("GET")

	router.HandleFunc("/api-status", scrape.APISTATUS).Methods("GET")

	router.HandleFunc("/indexer", scrape.IndexFiles).Methods("GET")
	router.HandleFunc("/search", scrape.FindFile).Methods("GET")		
	
    router.HandleFunc("/addsearch/{regex}", scrape.AddRegEx).Methods("GET")
	
    router.HandleFunc("/clear", scrape.ClearRegEx).Methods("GET")
	
    router.HandleFunc("/reset", scrape.ResetRegEx).Methods("GET")
	

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}