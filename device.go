package main


import (
    //"strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/AtegyekaRodgers/tracking-device/db"
)

type HotDevice struct {
    db.Device
}

var devices map[string]HotDevice

func (dv *HotDevice) updateLocation() {
    var device db.Device
    database.First(&device, "unique_label = ?", dv.UniqueLabel)
    device.Longitude = dv.Longitude
    device.Latitude = dv.Latitude
    database.Save(&device)
}

//TODO: function to load all devices data from db into devices map
func loadDevices() {
    //read all records in the devices table 
    //for each row in result, 
        //create an instance of HotDevice with record data
        //push it to the devices map
    //end loop
}

func registerDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var device db.Device
	_ = json.NewDecoder(r.Body).Decode(&device)
	database.Create(&device)
	json.NewEncoder(w).Encode(device)
}

func updateDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var dv db.Device
	_ = json.NewDecoder(r.Body).Decode(&dv)
	
    uniquelabel := mux.Vars(r)["uniquelabel"]
    
    var device db.Device
    database.Model(&device).Where("unique_label = ?", uniquelabel).Updates(dv)
    
    //database.First(&device, "unique_label = ?", uniquelabel)
    
    database.Save(&device)
    json.NewEncoder(w).Encode(device)
}

func readUpdates(w http.ResponseWriter, r *http.Request) {
    //TODO: read which device's updates are needed 
    //retrieve from db and respond
    json.NewEncoder(w).Encode(struct{Success string}{Success: "resource under development"})
}








