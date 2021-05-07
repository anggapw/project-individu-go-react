package entity

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	FirstName         string
	LastName          string
	MedicalSpecialist []MedicalSpecialist `gorm:"foreignKey:DoctorID"`
	Appointment       []Appointment       `gorm:"foreignKey:DoctorID"`
}

type MedicalSpecialist struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string
}

type Patient struct {
	gorm.Model
	FirstName   string
	LastName    string
	PhoneNumber int
	Address     string
	Appointment []Appointment `gorm:"foreignKey:PatientID"`
}

type Appointment struct {
	ID          int `gorm:"primaryKey"`
	Time        time.Time
	Date        time.Time
	IsCompleted bool
}
