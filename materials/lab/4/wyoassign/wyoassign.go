package wyoassign

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

)

type Response struct{
	Assignments []Assignment `json:"assignments"`
}

type Assignment struct {
	Id string `json:"id"`
	Title string `json:"title`
	Description string `json:"desc"`
	Points int `json:"points"`
}

var Assignments []Assignment
const Valkey string = "FooKey"

func InitAssignments(){
	var assignmnet Assignment
	assignmnet.Id = "Mike1A"
	assignmnet.Title = "Lab 4 "
	assignmnet.Description = "Some lab this guy made yesteday?"
	assignmnet.Points = 20
	Assignments = append(Assignments, assignmnet)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}


func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assignments = Assignments

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	var found = false
	
	for _, assignment := range Assignments {
		if assignment.Id == params["id"]{
			json.NewEncoder(w).Encode(assignment)
			found = true;
			break
		}
	}
	if !found{
		response := make(map[string]string)
		response["status"] = "No Such ID to Retrieve"
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return
		}
		w.Write(jsonResponse)
	}
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, assignment := range Assignments {
			if assignment.Id == params["id"]{
				Assignments = append(Assignments[:index], Assignments[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	r.ParseForm()
	var duplicate = false;
	response := make(map[string]string)
	response["status"] = "No Such ID to Update"
	
	if params["id"]!=r.FormValue("id"){
		for _, assignment := range Assignments {
			if assignment.Id == r.FormValue("id"){
				duplicate = true;
				response["status"] = "New ID Already In Use: Assignment Not Updated"
				break
			}
		}
	}
	if !duplicate{
		for index, assignment := range Assignments {
			if assignment.Id == params["id"]{
				if(r.FormValue("id")!=""){
					assignment.Id =  r.FormValue("id")
				}
				if(r.FormValue("title")!=""){
					assignment.Title =  r.FormValue("title")
				}
				if(r.FormValue("desc")!=""){
					assignment.Description =  r.FormValue("desc")
				}
				if(r.FormValue("points")!=""){
					assignment.Points, _ =  strconv.Atoi(r.FormValue("points"))
				}
				Assignments[index]=assignment
				response["status"] = "Success"
				break
			}
		}
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var assignment Assignment
	r.ParseForm()
	var duplicate = false;
	response := make(map[string]string)
	response["status"] = "Assignment Creation Failed"

	for _, assignment := range Assignments {
		if assignment.Id == r.FormValue("id"){
			duplicate = true;
			response["status"] = "ID Already In Use: Assignment Not Created"
			break
		}
	}
	
	if(r.FormValue("id") != "")&&(!duplicate){
		assignment.Id =  r.FormValue("id")
		if(r.FormValue("title")!=""){
			assignment.Title =  r.FormValue("title")
		}else{
			assignment.Title = "Title"
		}
		if(r.FormValue("desc")!=""){
			assignment.Description =  r.FormValue("desc")
		}else{
			assignment.Description = "Description"
		}
		if(r.FormValue("points")!=""){
			assignment.Points, _ =  strconv.Atoi(r.FormValue("points"))
		}else{
			assignment.Points =  0
		}
		Assignments = append(Assignments, assignment)
		response["status"] = "Success"
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)

}