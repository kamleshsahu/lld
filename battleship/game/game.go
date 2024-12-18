package game

import (
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
	players                   []*entity.Player
	fireStrategy              fireStrategy.FireStrategy
	eliminationStrategy       eliminationStrategy.IEliminationStrategy
	divideBattleFieldStrategy divideFieldStrategy.IDivideFieldStrategy
	targetPlayerStrategy      targetPlayerStrategy.ITargetPlayerStrategy
	isStarted                 bool
	turn                      int
}

func (game *Game) InitGame(N int) error {
	if game.isStarted {
		return entity.ERR_GAME_ALREADY_STARTED
	}
	game.board = entity.NewBoard(N)
	game.players = []*entity.Player{entity.NewPlayer(0), entity.NewPlayer(1)}

	err := game.divideBattleFieldStrategy.Divide(game.board, game.players)
	if err != nil {
		return err
	}

	game.fireStrategy.Init(utils.ClonePlayerFields(game.players))
	return nil
}

func (game *Game) AddShip(id string, size, player1ShipX, player1ShipY, player2ShipX, player2ShipY int) error {
	if game.board == nil {
		return entity.ERR_GAME_HAS_NO_BOARD
	}
	if game.players == nil {
		return entity.ERR_GAME_HAS_NO_PLAYERS
	}

	if game.isStarted {
		return entity.ERR_GAME_ALREADY_STARTED
	}

	cells := []*entity.Cell{entity.NewCell(player1ShipX, player1ShipY), entity.NewCell(player2ShipX, player2ShipY)}

	for i := 0; i < len(cells); i++ {
		_, err := game.board.CanPlaceShip(cells[i], size, game.players[i])
		if err != nil {
			return err
		}
		ship := entity.NewShip(id, cells[i], size, game.players[i])
		err = game.board.AddShip(ship)
		if err != nil {
			return err
		}
		game.players[i].ShipCount++
	}

	fmt.Printf(game.board.ViewBattleField())
	return nil
}

func (game *Game) StartGame() error {
	if game.isStarted {
		return entity.ERR_GAME_ALREADY_STARTED
	}
	if game.board == nil {
		return entity.ERR_GAME_HAS_NO_BOARD
	}

	game.isStarted = true
	fmt.Printf("Game started\n")
	fmt.Printf(game.board.ViewBattleField())
	for {
		currentPlayer := game.players[game.turn]
		targetPlayer, _ := game.targetPlayerStrategy.GetTargetPlayer(currentPlayer, game.players)
		hitPosition, err := game.fireStrategy.GetFireLocation(targetPlayer.Id)
		if err != nil {
			return err
		}

		destroyedShip, err := game.Fire(hitPosition)
		if err != nil {
			return err
		}

		fmt.Println(utils.TurnMessage(currentPlayer.GetName(), targetPlayer.GetName(), hitPosition.ToString(), destroyedShip))
		if destroyedShip != nil {
			fmt.Printf(game.board.ViewBattleField())
		}

		if !game.eliminationStrategy.IsEliminated(targetPlayer) {
			fmt.Printf("Game Over, Player %s won the game\n", currentPlayer.GetName())
			break
		}
		time.Sleep(time.Second)
		game.nextturn()
	}
	game.Reset()
	return nil
}

func (game *Game) Reset() {
	game.isStarted = false
	game.turn = 0
	game.board = nil
	game.players = []*entity.Player{}
}

func (game *Game) ViewBattleField() string {
	return game.board.ViewBattleField()
}

func (game *Game) nextturn() {
	game.turn++
	game.turn = game.turn % len(game.players)
}

func (g *Game) Fire(cell *entity.Cell) (*entity.Ship, error) {
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

func (game *Game) GetBoard() *entity.Board {
	return game.board
}

func (game *Game) GetPlayers() []*entity.Player {
	return game.players
}
