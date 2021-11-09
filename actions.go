package main


import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/afenet-demo/afenet-server/db"
    "github.com/julienschmidt/httprouter"
)

func monitoringActions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type","application/json")
	
    var actions []db.UserAction
    
    database.Find(&actions)
    if rows := result.RowsAffected; rows < 1 {
        json.NewEncoder(w).Encode(struct{Error string}{ Empty: "There are no actions taken by any user." })
    }else{
        json.NewEncoder(w).Encode(actions)
    }
}



