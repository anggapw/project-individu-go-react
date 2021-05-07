package main

import (
	"doctorsAppointment/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/doctors", handler.GetAllDoctor)
	r.GET("/patients", handler.GetAllPatient)
	r.GET("/appointments", handler.GetAllAppointment)
	r.Run(":4444")
}
