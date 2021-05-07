package config

import (
	"doctorsAppointment/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := "root@tcp(localhost)/doctors_appointment?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&entity.Doctor{})
	db.AutoMigrate(&entity.MedicalSpecialist{})
	db.AutoMigrate(&entity.Patient{})
	db.AutoMigrate(&entity.Appointment{})

	return db
}
