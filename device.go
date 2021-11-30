package main


import (
    "fmt"
    "math"
    //"strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/AtegyekaRodgers/tracking-device/db"
)

type HotDevice struct {
    db.Device
    DistanceFromObserver float32 `json:"distance"`
    Units string `json:"units" gorm:"default:m"`
}

var devices map[string]HotDevice

func (dv *HotDevice) updateDeviceLocation() {
    var device db.Device
    database.First(&device, "unique_label = ?", dv.UniqueLabel)
    device.Longitude = dv.Longitude
    device.Latitude = dv.Latitude
    database.Save(&device)
}

func (dv *HotDevice) updateObserverLocation() {
    var device db.Device
    database.First(&device, "unique_label = ?", dv.UniqueLabel)
    device.ObsvrLongitude = dv.ObsvrLongitude
    device.ObsvrLatitude  = dv.ObsvrLatitude
    database.Save(&device)
}

//function to load all devices data from db into devices map
func loadDevices() {
    //read all records in the devices table 
    rows, err := database.Table("devices").Select("id", "unique_label", "ownername", "ownerphone", "owneremail", "type", "state", "current_holder", "latitude","longitude", "obsvrlatitude","obsvrlongitude").Rows()   
    //for each row, 
    for rows.Next() {
        var hdv HotDevice
        rows.Scan(
            &hdv.ID,
            &hdv.UniqueLabel, &hdv.Ownername, &hdv.Ownerphone, &hdv.Owneremail, &hdv.Type, &hdv.State, &hdv.CurrentHolder, 
            &hdv.Latitude, &hdv.Longitude, &hdv.ObsvrLatitude, &hdv.ObsvrLongitude )
        //create a key for the target device in the devices map
        key := fmt.Sprintf("%s_%s", hdv.Type, hdv.UniqueLabel)
        //add hdv to the devices map
        devices[key] = hdv
    }
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


func getMyDevices(w http.ResponseWriter, r *http.Request) {
    type Tmp struct {
        Owneremail string 
    }
    var tmp Tmp
	_ = json.NewDecoder(r.Body).Decode(&tmp)
    
    var dvces []db.Device
    err := database.Where("owneremail = ?", tmp.Owneremail).Find(&dvces).Error
    
    json.NewEncoder(w).Encode(dvces)
}

func readUpdates(w http.ResponseWriter, r *http.Request) {
    //TODO: read which device's updates are needed 
    //retrieve from db and respond
    json.NewEncoder(w).Encode(struct{Success string}{Success: "resource under development"})
}

func reportDeviceLocation(w http.ResponseWriter, r *http.Request) {
    fmt.Println("reported location: ")
    var device db.Device
	_ = json.NewDecoder(r.Body).Decode(&device)
    //create a key to find the target device in the devices map
    key := fmt.Sprintf("%s_%s", device.Type, device.UniqueLabel)
    //HotDevice
    //index the map and update the obsvrlatitude and obsvrlongitude
    devices[key].Longitude = device.Longitude
    devices[key].Latitude  = device.Latitude
    //update the db with the new changes
    devices[key].updateDeviceLocation()
    //respond with the device, including the calculated distance
    json.NewEncoder(w).Encode(struct{Success string; Device HotDevice}{Success: "reported location", Device: devices[key]})
}

func reportObserverLocation(w http.ResponseWriter, r *http.Request) {
    fmt.Println("reported location: ")
    var device db.Device
	_ = json.NewDecoder(r.Body).Decode(&device)
    //create a key to find the target device in the devices map
    key := fmt.Sprintf("%s_%s", device.Type, device.UniqueLabel)
    //HotDevice
    //index the map and update the obsvrlatitude and obsvrlongitude
    devices[key].ObsvrLongitude = device.ObsvrLongitude
    devices[key].ObsvrLatitude  = device.ObsvrLatitude
    //read the device longitude and latitude
    deviceLongitude := devices[key].Longitude
    deviceLatitude := devices[key].Latitude
    //determine distance between the two: device and observer, multiply by 1.60934 to convert to kilometres from miles
    distanceInKm := 1.60934 * float32(distanceBtnGPSCoordinates(float64(deviceLatitude), float64(deviceLongitude), float64(device.ObsvrLatitude), float64(device.ObsvrLongitude)))
    distanceInMetres := distanceInKm * 1000
    devices[key].DistanceFromObserver = distanceInMetres
    //update the db with the new changes
    devices[key].updateObserverLocation()
    //respond with the device, including the calculated distance
    json.NewEncoder(w).Encode(struct{Success string; Device HotDevice}{Success: "reported location", Device: devices[key]})
}


func distanceBtnGPSCoordinates(lat1, lng1, lat2, lng2 float64) float64 {
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)
	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)
	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}
	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515
	return dist
}








