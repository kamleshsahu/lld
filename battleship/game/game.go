package game

import (
	"errors"
	"fmt"
	"lld/battleship/entity"
	"lld/battleship/fireStrategy"
	"sync"
	"time"
)

type Game struct {
	Board        *entity.Board
	Players      []entity.Player
	FireStrategy fireStrategy.FireStrategy
	Turn         int
}

func (game *Game) InitGame(N int, players []entity.Player) error {
	game.Board = entity.NewBoard(N)
	game.Players = players
	game.mapPlayerToField()
	playerFields := make([]entity.Field, len(players))
	for i, player := range players {
		playerFields[i] = player.Field
	}
	game.FireStrategy.Init(playerFields)
	return nil
}

func (game *Game) AddShip(ship *entity.Ship, cell entity.Cell, Size int) error {
	if game.Board == nil {
		return errors.New("game has no board")
	}
	if game.Players == nil {
		return errors.New("game has no players")
	}

	if !game.Board.IsValidCell(cell) {
		return errors.New("invalid cell")
	}

	ship.Owner = &game.Players[game.getPlayer(cell)]
	ship.Owner.ShipCount++
	ship.Size = Size
	ship.Location = cell
	err := game.Board.AddShip(ship, cell, Size)
	fmt.Printf(game.Board.ViewBattleField())
	if err != nil {
		return err
	}
	return nil
}

func (game *Game) getPlayer(cell entity.Cell) int {
	size := len(game.Board.Cells) / len(game.Players)
	playerId := cell.X / size
	return playerId
}

func (game *Game) StartGame() error {
	fmt.Printf("Game started\n")
	fmt.Printf(game.Board.ViewBattleField())
	for {
		distroyerId := game.Turn
		targetPlayerId := distroyerId ^ 1
		cell, err := game.FireStrategy.GetFireLocation(targetPlayerId)
		if err != nil {
			return err
		}
		fmt.Printf("%s fired at %s\n", game.Players[distroyerId].Name, cell.ToString())
		_, err = game.Fire(*cell)
		if err != nil {
			return err
		}

		if !game.Players[targetPlayerId].HasShip() {
			fmt.Printf("game over, %s wins the game\n", game.Players[distroyerId].Name)
			break
		}
		time.Sleep(time.Second)
		game.Turn++
		game.Turn = game.Turn % len(game.Players)
	}

	return nil
}

func (g *Game) Fire(cell entity.Cell) (*entity.Ship, error) {
	if !g.Board.Cells[cell.X][cell.Y].HasShip(cell) {
		fmt.Printf("no ship found at cell :%s\n", cell.ToString())
		return nil, nil
	}

	ship, err := g.Board.RemoveShip(g.Board.Cells[cell.X][cell.Y])
	if ship != nil {
		ship.Owner.ShipCount -= 1
		fmt.Printf("%s's ship distroyed at cell :%s\n", ship.Owner.Name, ship.Location.ToString())
		fmt.Printf("updated battlefield:\n%s\n", g.Board.ViewBattleField())
	}
	return ship, err
}

func (game *Game) mapPlayerToField() {
	totalFieldSize := len(game.Board.Cells)
	size := totalFieldSize / len(game.Players)

	for x := 0; x < totalFieldSize; x++ {
		playerId := x / size
		for y := 0; y < totalFieldSize; y++ {
			game.Board.Cells[x][y].Owner = &game.Players[playerId]

			game.Players[playerId].Field.Cells = append(game.Players[playerId].Field.Cells, game.Board.Cells[x][y])
		}
	}
}

var singleton sync.Once
var gameInstance *Game

func NewGame(strategy fireStrategy.FireStrategy) IGame {
	singleton.Do(func() {
		gameInstance = &Game{FireStrategy: strategy}
	})
	return gameInstance
}
