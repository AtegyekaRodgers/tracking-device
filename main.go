package main


import (
    "fmt"
    "log"
	"flag"
	"embed"
	"io/fs"
	"strings"
	"net/http"
	"gorm.io/gorm"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/AtegyekaRodgers/tracking-device/db"
)

//go:embed client/observer/web/*
var obsvrweb embed.FS

//go:embed client/device/web/*
var devceweb embed.FS

var database *gorm.DB

func index(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(struct{Success string}{Success: "Welcome !"})
}

func resourceNotFound(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(struct{Success string}{Success: "Resource not found !"})
}

func getRouter() *mux.Router {
	clientObserverWeb, _ := fs.Sub(obsvrweb, "client/observer/web")
	client_device_web, _ := fs.Sub(devceweb, "client/device/web")
	router := mux.NewRouter()
	router.HandleFunc("/ajax/updates", readUpdates).Methods("GET")
	router.HandleFunc("/report/observer", reportObserverLocation).Methods("POST")
	router.HandleFunc("/report/device", reportDeviceLocation).Methods("POST")
	router.HandleFunc("/device/register", registerDevice).Methods("POST")
	router.HandleFunc("/device/update/{uniquelabel}", updateDevice).Methods("PUT")
	router.PathPrefix("/client/observer/web").Handler( http.FileServer(http.FS(clientObserverWeb)) ).Methods("GET")
	router.PathPrefix("/client/device/web").Handler( http.FileServer(http.FS(client_device_web)) ).Methods("GET")
	router.PathPrefix("/").Handler( http.FileServer(http.FS(clientObserverWeb)) ).Methods("GET")
	router.HandleFunc("/", index).Methods("POST")
	//router.NotFoundHandler = http.HandlerFunc(resourceNotFound)
	return router
}

func main() {
    //++++| os.Args |+++++
    httpEndPoint := ":7000" 
    addr := flag.String("addr", httpEndPoint, "API service address") 
    flag.Parse()
    //++++++++++++++++++++
    database, _ = db.Connect()
    
    loadDevices()
    
    fmt.Println("Server listening on port: "+(strings.Split(httpEndPoint,":")[1])) 
    log.Fatal(http.ListenAndServe(*addr, getRouter()))
}








