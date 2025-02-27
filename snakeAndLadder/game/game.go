package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Id   int
	Cell int
	Name string
}

type Board struct {
	Cells []Cell
}

type IObject interface {
	JumpPosition() int
	GetStart() int
}

type Snake struct {
	Start int
	End   int
	Bite  int
}

func (s *Snake) GetStart() int {
	return s.Start
}

func (s *Snake) JumpPosition() int {
	if s.Bite == 0 {
		return s.Start
	}
	s.Bite--
	return s.End
}

type Ladder struct {
	Bottom int
	Top    int
}

func (l *Ladder) GetStart() int {
	return l.Bottom
}

func (l *Ladder) JumpPosition() int {
	return l.Top
}

func NewLadder(Top, Bottom int) IObject {
	return &Ladder{Top: Top, Bottom: Bottom}
}

func NewSnake(Head, Tail int, Bite int) IObject {
	return &Snake{Start: Head, End: Tail, Bite: Bite}
}

type Cell struct {
	Position int
	Object   *IObject
}

type IDice interface {
	SetMaxValue(val int)
	Roll() int
}

type DefaultDice struct {
	MaxValue        int
	randomGenerator *rand.Rand
}

func NewDefaultDice(maxValue int) IDice {
	return &DefaultDice{MaxValue: maxValue, randomGenerator: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (d *DefaultDice) SetMaxValue(val int) {
	d.MaxValue = val
}

func (d *DefaultDice) Roll() int {
	return d.randomGenerator.Intn(d.MaxValue) + 1
}

type IGame interface {
	AddPlayer(player Player)
	Start()
	IsStarted() bool
	SetObjects([]IObject)
	IsOver() bool
	GetWinner() int
}

type Game struct {
	players   []Player
	isStarted bool
	cells     []Cell
	Dice      IDice
}

func (g *Game) AddPlayer(player Player) {
	g.players = append(g.players, player)
}

func (g *Game) Start() {
	if g.players == nil || len(g.players) == 1 {
		fmt.Println("less than 2 players")
		return
	}
	if g.isStarted {
		return
	}
	g.isStarted = true

	for {
		player := g.players[0]
		diceValue := (g.Dice).Roll()
		position := player.Cell + diceValue
		if position > len(g.cells) {
			fmt.Println("invalid move")
			g.players = g.players[1:]
			g.players = append(g.players, player)
			continue
		}
		xposition := position

		if position != len(g.cells) && g.cells[position].Object != nil {
			position = (*g.cells[position].Object).JumpPosition()
			fmt.Println("Player ", player.Name, " dicevalue :", diceValue, " got jump from ", xposition, " to ", position)
		}

		player.Cell = position
		g.players = g.players[1:]
		if position == len(g.cells) {
			fmt.Println("Player ", player.Name, " finished the game")
		} else {
			g.players = append(g.players, player)
		}
		fmt.Println("Player ", player.Name, "player position ", player.Cell)
		if len(g.players) == 0 {
			break
		}
	}

	g.isStarted = false
}

func (g *Game) IsStarted() bool {
	return g.isStarted
}

func (g *Game) SetObjects(objects []IObject) {
	for _, obj := range objects {
		g.cells[obj.GetStart()].Object = &obj
	}
}

func (g *Game) IsOver() bool {
	return g.isStarted == false
}

func (g *Game) GetWinner() int {
	return 0
}

func NewGame(boardSize int) IGame {
	dice := NewDefaultDice(6)
	return &Game{players: make([]Player, 0), isStarted: false, cells: make([]Cell, boardSize), Dice: dice}
}
