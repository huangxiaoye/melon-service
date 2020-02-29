package interfaces

type IUserService interface {
	GetScores(player1Name string, player2Name string) (string, error)
}
