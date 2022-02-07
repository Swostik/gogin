package main

import (
	"zenniz/controllers"
	"zenniz/services"

	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	r := gin.Default()

	db, err := leveldb.OpenFile("./test.db", nil)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	controller := controllers.NewUserController(services.NewuserService(db))

	r.GET("/users", controller.List)
	r.GET("/users/:personalCode", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:personalCode", controller.UpdateUser)
	r.DELETE("/users/:personalCode", controller.DeleteUser)
	r.Run()
}
