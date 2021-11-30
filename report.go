package main


import (
    "fmt"
    //"strconv"
    "net/http"
    "encoding/json"
    //"github.com/gorilla/mux"
    "github.com/AtegyekaRodgers/tracking-device/db"
)

func reportDeviceLocation(w http.ResponseWriter, r *http.Request) {
    fmt.Println("reported location: ")
    var device db.Device
	_ = json.NewDecoder(r.Body).Decode(&device)
    //create a key to find the target device in the devices map
    key := fmt.Sprintf("%s_%s", device.Type, device.UniqueLabel)
    tmphdv := devices[key]
    //index the map and update the obsvrlatitude and obsvrlongitude
    tmphdv.Longitude = device.Longitude
    tmphdv.Latitude  = device.Latitude
    //update the db with the new changes
    tmphdv.updateDeviceLocation()
    devices[key] = tmphdv
    //respond with the device, including the calculated distance
    json.NewEncoder(w).Encode(struct{Success string; Device HotDevice}{Success: "device reported location", Device: tmphdv})
}

func reportObserverLocation(w http.ResponseWriter, r *http.Request) {
    fmt.Println("reported location: ")
    var device db.Device
	_ = json.NewDecoder(r.Body).Decode(&device)
    //create a key to find the target device in the devices map
    key := fmt.Sprintf("%s_%s", device.Type, device.UniqueLabel)
    tmphdv := devices[key]
    //index the map and update the obsvrlatitude and obsvrlongitude
    tmphdv.ObsvrLongitude = device.ObsvrLongitude
    tmphdv.ObsvrLatitude  = device.ObsvrLatitude
    //read the device longitude and latitude
    deviceLongitude := devices[key].Longitude
    deviceLatitude := devices[key].Latitude
    //determine distance between the two: device and observer, multiply by 1.60934 to convert to kilometres from miles
    distanceInKm := 1.60934 * float32(distanceBtnGPSCoordinates(deviceLatitude, deviceLongitude, device.ObsvrLatitude, device.ObsvrLongitude))
    distanceInMetres := distanceInKm * 1000
    tmphdv.DistanceFromObserver = distanceInMetres
    //update the db with the new changes
    tmphdv.updateObserverLocation()
    devices[key] = tmphdv
    //respond with the device, including the calculated distance
    json.NewEncoder(w).Encode(struct{Success string; Device HotDevice}{Success: "observer reported location", Device: tmphdv})
}






