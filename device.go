package main


import (
    "fmt"
    "sync"
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
var mu sync.Mutex

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
    rows, err := database.Table("devices").Select("id", "unique_label", "ownername", "ownerphone", "owneremail", "type", "state", "current_holder", "latitude","longitude", "obsvr_latitude","obsvr_longitude", "missing").Rows()   
    if err != nil {
        fmt.Println("!!Failed to load devices from db: ", err)
    }else{
        //for each row, 
        for rows.Next() {
            var hdv HotDevice
            rows.Scan(
                &hdv.ID,
                &hdv.UniqueLabel, &hdv.Ownername, &hdv.Ownerphone, &hdv.Owneremail, &hdv.Type, &hdv.State, &hdv.CurrentHolder, 
                &hdv.Latitude, &hdv.Longitude, &hdv.ObsvrLatitude, &hdv.ObsvrLongitude, &hdv.Missing )
            //create a key for the target device in the devices map
            key := fmt.Sprintf("%s_%s", hdv.Type, hdv.UniqueLabel)
            //add hdv to the devices map
            mu.Lock()
            devices[key] = hdv
            mu.Unlock()
        }
    }
}

func registerDevice(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registering ...")
	w.Header().Set("Content-Type","application/json")
	var device db.Device
	_ = json.NewDecoder(r.Body).Decode(&device)
	database.Create(&device)
	json.NewEncoder(w).Encode(struct{Success string}{Success: "Device registered."})
}

func updateDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var dv db.Device
	_ = json.NewDecoder(r.Body).Decode(&dv)
	
    uniquelabel := mux.Vars(r)["uniquelabel"]
    
    var device db.Device
    database.Model(&device).Where("unique_label = ?", uniquelabel).Updates(dv)
    
    database.Save(&device)
    json.NewEncoder(w).Encode(struct{Success string}{Success: "Device updated."})
}


func getMyDevices(w http.ResponseWriter, r *http.Request) {
    type Tmp struct {
        Owneremail string 
    }
    var tmp Tmp
	_ = json.NewDecoder(r.Body).Decode(&tmp)
    
    var dvces []HotDevice
    err := database.Table("devices").Where("owneremail = ?", tmp.Owneremail).Find(&dvces).Error
    if err!=nil {
        json.NewEncoder(w).Encode(struct{Err string;}{Err:err.Error()})
    }
    json.NewEncoder(w).Encode(struct{Success string; Devices []HotDevice}{Success: "devices found", Devices:dvces})
}

func readUpdates(w http.ResponseWriter, r *http.Request) {
    //TODO: read which device's updates are needed
    //retrieve from db and respond
    json.NewEncoder(w).Encode(struct{Success string;}{Success: "resource under development"})
}

func distanceBtnGPSCoordinates(lat1, lng1, lat2, lng2 float32) float32 {
	x := lng2-lng1
	y := lat2-lat1
	distance := math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
	return float32(distance)
}








