package game

type IGame interface {
	InitGame(N int) error
	AddShip(id string, size, player1ShipX, player1ShipY, player2ShipX, player2ShipY int) error
	StartGame() error
	ViewBattleField() string
}
