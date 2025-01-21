package main

import (
	"lld/battleship/game"
	"lld/battleship/strategy/divideFieldStrategy"
	"lld/battleship/strategy/eliminationStrategy"
	"lld/battleship/strategy/fireStrategy"
	"lld/battleship/strategy/targetPlayerStrategy"
	"sort"
)

func main() {
	n := 10
	fireStrategy := fireStrategy.NewRandomFireStrategy()
	eliminationStrategy := eliminationStrategy.NewDefaultEliminationStrategy()
	targetPlayerStrategy := targetPlayerStrategy.NewDefaultTargetStrategy()
	divideField := divideFieldStrategy.NewEqualDivideStrategy()
	gameService := game.NewGame(fireStrategy, eliminationStrategy, divideField, targetPlayerStrategy)

	err := gameService.InitGame(n)
	if err != nil {
		panic(err)
	}

	sort.SearchInts()

	err = gameService.AddShip("SH1", 2, 2, 0, 5, 5)
	if err != nil {
		panic(err)
	}

	err = gameService.AddShip("SH2", 3, 1, 4, 7, 2)
	if err != nil {
		panic(err)
	}
	err = gameService.StartGame()
	if err != nil {
		panic(err)
	}
}
