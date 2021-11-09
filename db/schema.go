package db

import (
	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  Username string `json:"username"`
  Password string `json:"pxwd"`
  Firstname  string  `json:"firstname"`
  Lastname string  `json:"lastname"`
  Phone string `json:"unique"`
  Email string `json:"email"`
  EmployeeID string `json:"employeeid"`
  Title string `json:"title"`
  Country string `json:"country"`
  Branch string `json:"branch"`
  AccessRights int `json:"accessRights"`
  Alerts []Alert `json:"alerts"`
  Notifications []Notification `json:"notifications"`
}

type Policy struct {
  gorm.Model
  Name string `josn:"name"`
  Version uint `josn:"version"`
  DocumentLink string `josn:"documentLink"`
}

type UserAction struct {
  gorm.Model
  ActionID string `json:"actionid"`
  UserID string `json:"userid"`
  Description string `json:"description"`
  OnEntity string `json:"onEntity"`
  EntityOption string `json:"entityOption"`
  ViolationBy string `json:"violationBy"`
  Country string `json:"country"`
  Latitude float `json:"latitude"`
  Longitude float `json:"longitude"`
}

type Alert struct {
  gorm.Model
  Message string
  Read bool
}

type Notification struct {
  gorm.Model
  Message string
  Read bool
}



