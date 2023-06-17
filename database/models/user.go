package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID         uint       `gorm:"primary key; autoIncrement" json:"id"`
	Username   *string    `json:"username"`
	Email      *string    `json:"email"`
	Password   *string    `json:"password"`
	Attendance Attendance `gorm:"foreignKey:UserID"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Attendance{})
	if err != nil {

		return err
	}

	err = db.Exec("ALTER TABLE attendances ADD CONSTRAINT fk_attendances_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE").Error
	return err
}
