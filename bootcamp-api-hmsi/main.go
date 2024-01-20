package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kisahtegar/bootcamp-api-hmsi/connectDB"
	"github.com/kisahtegar/bootcamp-api-hmsi/modules/customers/customerHandler"
	"github.com/kisahtegar/bootcamp-api-hmsi/modules/customers/customerRepository"
	"github.com/kisahtegar/bootcamp-api-hmsi/modules/customers/customerUsecase"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_DRIVER := os.Getenv("DB_DRIVER")
	PORT := os.Getenv("PORT")

	log.Info().Msg(DB_HOST)
	log.Info().Msg(DB_NAME)
	log.Info().Msg(DB_PORT)
	log.Info().Msg(DB_USER)
	log.Info().Msg(DB_PASSWORD)
	log.Info().Msg(DB_DRIVER)
	log.Info().Msg(PORT)

	db, errConn := connectDB.GetConnPostgres(DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_DRIVER)
	if errConn != nil {
		log.Error().Msg(errConn.Error())
		os.Exit(1)
	}
	fmt.Println("Successfully connected!")

	// // DB struct initialize
	// DB := query.DB{Conn: db}

	// // // Create customer
	// err = DB.Create(&query.Customers{
	// 	Name:  "Kisah Tegar",
	// 	Phone: "0822736133",
	// 	Email: "kisah@mail.com",
	// 	Age:   18,
	// })
	// fmt.Println("err", err)
	// if err != nil {
	// 	log.Error().Msg(errConn.Error())
	// 	os.Exit(1)
	// }
	// fmt.Println("Insert Data Berhasil")

	// // Read customer
	// result, err := DB.Read()
	// if err != nil {
	// 	log.Error().Msg(errConn.Error())
	// 	os.Exit(1)
	// }
	// fmt.Println(result)

	// // Update customer
	// err = DB.Update(&query.Customers{
	// 	Id:    1,
	// 	Name:  "Budi",
	// 	Phone: "012345564",
	// 	Email: "budi@example.com",
	// 	Age:   31,
	// })
	// if err != nil {
	// 	log.Error().Msg(errConn.Error())
	// 	os.Exit(1)
	// }
	// fmt.Println("Update Berhasil")

	// // Delete customer
	// err = DB.Delete(1)
	// if err != nil {
	// 	log.Error().Msg(errConn.Error())
	// 	os.Exit(1)
	// }
	// fmt.Println("Delet Berhasil")

	// initialize router
	router := gin.Default()

	// initialize module customers
	customerRepo := customerRepository.NewCustomerRepository(db)
	customerUC := customerUsecase.NewCustomerUsecase(customerRepo)
	customerHandler.NewCustomerHandler(router, customerUC)

	log.Debug().Msg("Service started on Port: " + PORT)
	router.Run(":" + PORT)
}
