package game

import (
	"errors"
	"fmt"
	"lld/battleship/entity"
	"lld/battleship/strategy/divideFieldStrategy"
	"lld/battleship/strategy/eliminationStrategy"
	"lld/battleship/strategy/fireStrategy"
	"lld/battleship/strategy/targetPlayerStrategy"
	"lld/battleship/utils"
	"sync"
	"time"
)

type Game struct {
	board                     *entity.Board
	Players                   []*entity.Player
	fireStrategy              fireStrategy.FireStrategy
	eliminationStrategy       eliminationStrategy.IEliminationStrategy
	divideBattleFieldStrategy divideFieldStrategy.IDivideFieldStrategy
	targetPlayerStrategy      targetPlayerStrategy.ITargetPlayerStrategy
	turn                      int
}

func (game *Game) InitGame(N int) error {
	game.board = entity.NewBoard(N)
	game.Players = []*entity.Player{entity.NewPlayer(0), entity.NewPlayer(1)}

	err := game.divideBattleFieldStrategy.Divide(game.board, game.Players)
	if err != nil {
		return err
	}

	game.fireStrategy.Init(utils.ClonePlayerFields(game.Players))
	return nil
}

func (game *Game) AddShip(id string, size, player1ShipX, player1ShipY, player2ShipX, player2ShipY int) error {
	if game.board == nil {
		return errors.New("game has no board")
	}
	if game.Players == nil {
		return errors.New("game has no players")
	}

	cells := []entity.Cell{*entity.NewCell(player1ShipX, player1ShipY), *entity.NewCell(player2ShipX, player2ShipY)}

	for i := 0; i < len(cells); i++ {
		_, err := game.board.IsValidLocation(cells[i], size)
		if err != nil {
			return err
		}
	}

	for i := 0; i < len(cells); i++ {
		game.Players[i].ShipCount++
		ship := entity.NewShip(id, &cells[i], size, game.Players[i])
		err := game.board.AddShip(ship, cells[i], size)
		if err != nil {
			return err
		}
	}

	fmt.Printf(game.board.ViewBattleField())
	return nil
}

func (game *Game) StartGame() error {
	fmt.Printf("Game started\n")
	fmt.Printf(game.board.ViewBattleField())
	for {
		currentPlayer := game.Players[game.turn]
		targetPlayer, _ := game.targetPlayerStrategy.GetTargetPlayer(currentPlayer, game.Players)
		hitPosition, err := game.fireStrategy.GetFireLocation(targetPlayer.Id)
		if err != nil {
			return err
		}

		destroyedShip, err := game.fire(hitPosition)
		if err != nil {
			return err
		}

		fmt.Println(utils.TurnMessage(currentPlayer.GetName(), targetPlayer.GetName(), hitPosition.ToString(), destroyedShip))
		if destroyedShip != nil {
			fmt.Printf(game.board.ViewBattleField())
		}

		if !game.eliminationStrategy.IsEliminated(targetPlayer) {
			fmt.Printf("Game Over, %s wins the game\n", currentPlayer.GetName())
			break
		}
		time.Sleep(time.Second)
		game.nextturn()
	}

	return nil
}

func (game *Game) ViewBattleField() string {
	return game.board.ViewBattleField()
}

func (game *Game) nextturn() {
	game.turn++
	game.turn = game.turn % len(game.Players)
}

func (g *Game) fire(cell *entity.Cell) (*entity.Ship, error) {
	if !g.board.HasShip(cell) {
		return nil, nil
	}
	destroyedShip, err := g.board.RemoveShip(cell)
	if err != nil {
		return nil, err
	}
	destroyedShip.Owner.ShipCount -= 1
	return destroyedShip, nil
}

var singleton sync.Once
var gameInstance *Game

func NewGame(fireStrategy fireStrategy.FireStrategy,
	eliminationStrategy eliminationStrategy.IEliminationStrategy,
	divide divideFieldStrategy.IDivideFieldStrategy,
	targetPlayer targetPlayerStrategy.ITargetPlayerStrategy) IGame {
	singleton.Do(func() {
		gameInstance = &Game{
			fireStrategy:              fireStrategy,
			eliminationStrategy:       eliminationStrategy,
			divideBattleFieldStrategy: divide,
			targetPlayerStrategy:      targetPlayer,
		}
	})
	return gameInstance
}
