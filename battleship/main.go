package main

import (
	"lld/battleship/entity"
	"lld/battleship/fireStrategy"
	"lld/battleship/game"
)

func main() {
	n := 10
	fireStrategy := fireStrategy.NewRandomFireStrategy()
	gameService := game.NewGame(fireStrategy)

	players := []entity.Player{{Name: "Kamlesh"}, {Name: "Tikesh"}}

	gameService.InitGame(n, players)

	ship1 := &entity.Ship{Name: "SH1"}
	err := gameService.AddShip(ship1, entity.Cell{X: 0, Y: 0}, 3)
	if err != nil {
		panic(err)
	}
	ship2 := &entity.Ship{Name: "SH2"}
	err = gameService.AddShip(ship2, entity.Cell{X: 2, Y: 5}, 2)
	if err != nil {
		panic(err)
	}
	ship3 := &entity.Ship{Name: "SH3"}
	err = gameService.AddShip(ship3, entity.Cell{X: 5, Y: 2}, 3)
	if err != nil {
		panic(err)
	}
	ship4 := &entity.Ship{Name: "SH4"}
	err = gameService.AddShip(ship4, entity.Cell{X: 6, Y: 6}, 3)
	if err != nil {
		panic(err)
	}
	err = gameService.StartGame()
	if err != nil {
		panic(err)
	}
}
