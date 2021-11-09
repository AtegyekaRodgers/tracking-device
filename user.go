package main


import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/afenet-demo/afenet-server/db"
    "github.com/julienschmidt/httprouter"
)

type User struct {
  Firstname  string  `json:"firstname"`
  Lastname string  `json:"lastname"`
  Phone string `json:"unique"`
  Email string `json:"email"`
  EmployeeID string `json:"employeeid"`
  Title string `json:"title"`
  Country string `json:"country"`
  Branch string `json:"branch"`
  AccessRights int `json:"accessRights"`
}

func registeringUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type","application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	
	database.Create(db.User{
	      Username: user.Username,
          Password: user.Password,
	      Firstname: user.Firstname,
          Lastname: user.Lastname,
          Phone: user.Phone,
          Email: user.Email,
          EmployeeID: user.EmployeeID,
          Title: user.Title,
          Country: user.Country,
          Branch: user.Branch,
          AccessRights: user.AccessRights })
	json.NewEncoder(w).Encode(user)
}

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"pxwd"`
}

func userLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type","application/json")
	var credentials Credentials
	_ = json.NewDecoder(r.Body).Decode(&credentials)
	
    var users []User
    
    database.Where("username = ? AND Password >= ?", user.Username, user.Password).Find(&users)
    if rows := result.RowsAffected; rows < 1 {
        json.NewEncoder(w).Encode(struct{Error string}{ Error: "Access denied." })
    }else{
        json.NewEncoder(w).Encode(users)
    }
}

type UserRights struct {
    Username string  `json:"username"`
    Rights int `json:"rights"`
}

func assignUserRights(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type","application/json")
	var rights UserRights
	_ = json.NewDecoder(r.Body).Decode(&rights)
	
    var user db.User
    database.First(&user, "username = ?", rights.Username)
    user.AccessRights = rights.Rights
    database.Save(&user)
    json.NewEncoder(w).Encode(user)
}


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



