package db

import (
	"gorm.io/gorm"
)

type Device struct {
  gorm.Model
  UniqueLabel string  `json:"uniquelabel"`
  Ownername string `json:"ownername"`
  Ownerphone string `json:"ownerphone"`
  Owneremail  string  `json:"owneremail"`
  Type string `json:"type" gorm:"default:car"`
  State string `json:"state" gorm:"default:stationary"`
  CurrentHolder string `json:"currentHolder" gorm:"default:owner"`
  Longitude float32 `json:"longitude"`
  Latitude float32 `json:"latitude"`
  ObsvrLongitude float32 `json:"obsvrlongitude"`
  ObsvrLatitude float32 `json:"obsvrlatitude"`
  Missing bool `json:"missing" gorm:"default:false"`
}

type DeviceHistory struct {
  gorm.Model
  UniqueLabel string  `json:"uniquelabel"`
  Type string `json:"type" gorm:"default:car"`
  State string `json:"state" gorm:"default:stationary"`
  CurrentHolder string `json:"currentHolder" gorm:"default:owner"`
  Longitude float32 `json:"longitude"`
  Latitude float32 `json:"latitude"`
  Missing bool `json:"missing" gorm:"default:false"`
  Notice string `json:"notice"`
}



