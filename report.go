package main


import (
    //"strconv"
    "net/http"
    "encoding/json"
    //"github.com/gorilla/mux"
    "github.com/AtegyekaRodgers/tracking-device/db"
)

func reportDeviceLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var dv db.Device
	_ = json.NewDecoder(r.Body).Decode(&dv)
	//r.Body = `{"type": dv.Type, "uniquelabel": dv.UniqueLabel, "longitude": dv.Longitude, "latitude": dv.Latitude}`
	
	//TODO: index device in the devices map
	//if device is in the map, then
	    //determine distance btn boserver and the device
	    //compare the reported device location with previous location
	    //if device has moved, 
	        //update the location details: longitude and latitude
	        //update the persisted location in db (deviceobject.updateLocation)
	    //end if
	//end if
    
    json.NewEncoder(w).Encode(dv)
}

func reportObserverLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var dv db.Device
	_ = json.NewDecoder(r.Body).Decode(&dv)
	//r.Body = `{"type": dv.Type, "uniquelabel": dv.UniqueLabel, "obsvrlongitude": dv.ObsvrLongitude, "obsvrlatitude": dv.ObsvrLatitude}`
	//TODO: update observer location in the hot devices map
    json.NewEncoder(w).Encode(dv)
}






