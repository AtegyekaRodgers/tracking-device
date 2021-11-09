package main


import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/afenet-demo/afenet-server/db"
    "github.com/julienschmidt/httprouter"
)

func uploadingNewPolicy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type","application/json")
	//TODO: grab the file uploaded
	//TODO: save the file in some directory
	//TODO: save a reference to the file in some variable
	
	var policy db.Policy
	_ = json.NewDecoder(r.Body).Decode(&policy)
	policy.DocumentLink = "some-link-will be generated"
    database.Create(&policy)
    json.NewEncoder(w).Encode(struct{Success string}{Success: "Policy uploaded successfully"})
}

func readPolicy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type","application/json")
	var policy db.Policy
    database.First(&policy, "name = ?", params["name"])
    
    json.NewEncoder(w).Encode(policy)
}

func agreeToPolicy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type","application/json")
	var action db.UserAction
	_ = json.NewDecoder(r.Body).Decode(&action)
	action.Description = "Agreed to a policy"
    database.Create(&action)
    json.NewEncoder(w).Encode(struct{Success string}{Success: "You have agreed to the policy."})
}

func disagreeToPolicy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type","application/json")
	var action db.UserAction
	_ = json.NewDecoder(r.Body).Decode(&action)
	action.Description = "Disagreed to a policy"
    database.Create(&action)
    json.NewEncoder(w).Encode(struct{Success string}{Success: "You have disagreed to the policy."})
}

func reportingPolicyViolation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type","application/json")
	var action db.UserAction
	_ = json.NewDecoder(r.Body).Decode(&action)
	action.Description = "Reported a policy violation"
    database.Create(&action)
    json.NewEncoder(w).Encode(struct{Success string}{Success: "You have reported a policy violation."})
}




