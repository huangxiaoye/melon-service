package services

import (
	"fmt"
	"melon-service/interfaces"
)

type UserService struct {
	interfaces.IUserRepository
}

func NewUserService(repo interfaces.IUserRepository) UserService {
	return UserService{repo}
}

func (service UserService) GetScores(player1Name string, player2Name string) (string, error) {

	baseScore := [4]string{"Love", "Fifteen", "Thirty", "Forty"}
	var result string
	fmt.Printf("user1 name is %s\n", player1Name)
	player1, err := service.GetUserByName(player1Name)
	fmt.Printf("user1 is %v\n", player1)
	if err != nil {
		//Handle error
		return "", err
	}
	fmt.Printf("user2 name is %s\n", player2Name)
	player2, err := service.GetUserByName(player2Name)
	if err != nil {
		//Handle error
		return "", err
	}
	fmt.Printf("user2 is %v\n", player2)
	if player1.Score < 4 && player2.Score < 4 && !(player1.Score+player2.Score == 6) {

		s := baseScore[player1.Score]

		if player1.Score == player2.Score {
			result = s + "-All"
		} else {
			result = s + "-" + baseScore[player2.Score]
		}
	}

	if player1.Score == player2.Score {
		result = "Deuce"
	}

	return result, nil
}
