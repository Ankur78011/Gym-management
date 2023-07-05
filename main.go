package main

import (
	"database/sql"
	"fmt"
	"gym/ankur/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "gym1"
)

func main() {
	psqlCon := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlCon)
	if err != nil {
		fmt.Println("Connection to database failed")
		return
	}
	defer db.Close()
	fmt.Println("DATABASE CONNECTED SUCCESFULLY")
	router := gin.Default()
	apiHandler := handlers.NewApiHandler(db)

	//////login//
	router.POST("/login", apiHandler.Login())
	////////////----GYM OWNERS----/////////
	router.GET("/owner", apiHandler.GetOwner())
	router.POST("/owner", apiHandler.CreateOwner())
	router.DELETE("/owner/:id", apiHandler.DeleteOwner())
	router.PUT("/owner/:id", apiHandler.UpdateOwner())
	/////-----GYMS----//////
	router.GET("/gyms", apiHandler.AuthOwner(), apiHandler.GetGyms())
	router.POST("gyms", apiHandler.AuthOwner(), apiHandler.CreateGyms())
	router.DELETE("/gyms/:id", apiHandler.AuthOwner(), apiHandler.DeleteGyms())
	router.PUT("/gyms/:id", apiHandler.UpdateGyms())
	// /////////----CUSTOMERS----/////
	router.GET("/customer", apiHandler.AuthOwner(), apiHandler.GetCustomer())
	router.POST("customer", apiHandler.AuthOwner(), apiHandler.CreateCustomer())
	router.DELETE("/customer/:id", apiHandler.AuthOwner(), apiHandler.DeleteCustomer())
	router.PUT("/customer/:id", apiHandler.UpdateCustomer())
	//////---PERSONAL TRAINERS---/////
	router.GET("/trainers", apiHandler.GetTrainer())
	router.POST("trainers", apiHandler.CreateTrainer())
	router.DELETE("/trainers/:id", apiHandler.DeleteTrainer())
	router.PUT("/trainers/:id", apiHandler.UpdateTrainer())
	// ////----CUSTOMER STATS---//////////
	router.GET("/report/:id", apiHandler.GetReports())
	router.POST("/report/:id", apiHandler.CreateReport())
	router.Run(":8001")
}
