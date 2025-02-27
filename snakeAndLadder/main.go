package main

import (
	"lld/snakeAndLadder/game"
)

func main() {

	g := game.NewGame(20)
	g.AddPlayer(game.Player{Name: "Kamlesh", Cell: 0})
	g.AddPlayer(game.Player{Name: "Tikesh", Cell: 0})

	objects := []game.IObject{game.NewSnake(14, 7, 1), game.NewLadder(10, 5)}
	g.SetObjects(objects)

	g.Start()
}
