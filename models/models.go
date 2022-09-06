package models

type User struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT" json:"ID"`
	Username string `gorm:"size(20)" json:"Username"`
	Password string `gorm:"size(100); not null" json:"Password"`
}
