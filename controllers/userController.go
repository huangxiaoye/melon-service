package controllers

import (
	"fmt"
	"melon-service/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	interfaces.IUserService
}

func NewUserController(service interfaces.IUserService) UserController {
	return UserController{service}
}

func (controller UserController) Test(c *gin.Context) {

	player1Name := c.Query("player1")
	player2Name := c.Query("player2")
	fmt.Printf("player1Name is %s\n", player1Name)
	fmt.Printf("player2Name is %s\n", player2Name)
	scores, err := controller.GetScores(player1Name, player2Name)
	if err != nil {
		//
	}

	c.JSON(http.StatusOK, scores)
}
