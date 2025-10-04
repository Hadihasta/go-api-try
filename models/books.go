// declare package models agar dapat di import di main.go nantinya
package models

import "gorm.io/gorm"

// jika di monggo DB kita bisa dengan mudah connect database tanpa harus declare databasenya namun
// jika di postgreSql maka harus di create connect seperti ini jika mau integrate
type Books struct {
	// dengan gorm dapat kita declare primary key seperti ini maka dari itu lihat dokumentasinya
	ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}


func MigrateBooks(db *gorm.DB) error{
	err := db.AutoMigrate(&Books{})
	return err
}