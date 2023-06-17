package models

import "gorm.io/gorm"

type Attendance struct {
	ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
	UserID    uint    `json:"userId"`
	Degree    *string `json:"degree"`
	Year      *string `json:"year"`
	Block     *string `json:"block"`
	Subject   *string `json:"subject"`
	Date      *string `json:"date"`
	StartTime *string `json:"startTime"`
	EndTime   *string `json:"endTime"`
}

func MigrateAttendace(db *gorm.DB) error {
	err := db.AutoMigrate(&Attendance{})
	return err
}
