package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      []Role `gorm:"foreignKey:UserID"`
}

type Role struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string
}

type MedicalSpecialist struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string
}

type Appointment struct {
	ID          int `gorm:"primaryKey"`
	Time        time.Time
	Date        time.Time
	IsCompleted bool
}

type Doctor struct {
	ID                int                 `gorm:"primaryKey"`
	User              []User              `gorm:"foreignKey:DoctorID"`
	MedicalSpecialist []MedicalSpecialist `gorm:"foreignKey:DoctorID"`
	Appointment       []Appointment       `gorm:"foreignKey:DoctorID"`
}

type Patient struct {
	ID          int           `gorm:"primaryKey"`
	User        []User        `gorm:"foreignKey:DoctorID"`
	Appointment []Appointment `gorm:"foreignKey:DoctorID"`
}
