package routers

import (
	"melon-service/databases"

	"github.com/gin-gonic/gin"
)

// Engine return Engine
func Engine() *gin.Engine {
	//设置生产环境
	//gin.SetMode(config.RunMode)
	r := gin.Default()
	db := databases.GetDB()
	// userRepositoy := repositories.UserRepository{DB: db}
	// userService := services.UserService{userRepositoy}
	// userController := controllers.UserController{userService}
	userController := InitializeUserController(db)

	r.GET("/", userController.Test)

	return r
}
