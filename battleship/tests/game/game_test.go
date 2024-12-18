package game

import (
	"fmt"
	"lld/battleship/entity"
	game2 "lld/battleship/game"
	"lld/battleship/strategy/divideFieldStrategy"
	"lld/battleship/strategy/eliminationStrategy"
	"lld/battleship/strategy/fireStrategy"
	"lld/battleship/strategy/targetPlayerStrategy"
	"testing"
)

func getNewGameInstance() game2.IGame {
	fireStrategy := fireStrategy.NewRandomFireStrategy()
	eliminationStrategy := eliminationStrategy.NewDefaultEliminationStrategy()
	divideFieldStrategy := divideFieldStrategy.NewEqualDivideStrategy()
	targetPlayerStrategy := targetPlayerStrategy.NewDefaultTargetStrategy()

	gameInstance := game2.NewGame(fireStrategy, eliminationStrategy, divideFieldStrategy, targetPlayerStrategy)
	return gameInstance
}

// Test InitGame
func TestInitGame(t *testing.T) {

	gameInstance := getNewGameInstance()
	// check if n is odd
	err := gameInstance.InitGame(5)
	expected := entity.ErrEqualDivide(5, 2)
	if err.Error() != expected.Error() {
		t.Errorf("Expected error %v, but got %v", expected.Error(), err.Error())
	}

	//
	err = gameInstance.InitGame(6)
	if err != nil {
		t.Errorf("Game is not initialized")
	}

	// Check if the board and players are correctly initialized
	if gameInstance.GetBoard() == nil {
		t.Errorf("Board is not initialized")
	}
	if len(gameInstance.GetPlayers()) != 2 {
		t.Errorf("Expected 2 players, got %d", len(gameInstance.GetPlayers()))
	}
}

// Test AddShip
func TestAddShip(t *testing.T) {
	gameInstance := getNewGameInstance()

	err := gameInstance.InitGame(10)
	if err != nil {
		t.Fatalf("Failed to initialize game: %v", err)
	}

	// Add ships to both players
	err = gameInstance.AddShip("SHIP1", 2, 1, 1, 5, 2)
	if err != nil {
		t.Fatalf("Failed to add ship: %v", err)
	}

	// Verify that ships are added correctly by checking the player ship count
	if (gameInstance.GetPlayers())[0].ShipCount != 1 {
		t.Errorf("Player 1 should have 1 ship, got %d", (gameInstance.GetPlayers())[0].ShipCount)
	}
	if (gameInstance.GetPlayers())[1].ShipCount != 1 {
		t.Errorf("Player 2 should have 1 ship, got %d", (gameInstance.GetPlayers())[0].ShipCount)
	}

	// Add ships to both players
	err = gameInstance.AddShip("SHIP1", 2, 3, 3, 8, 4)
	if err != nil {
		t.Fatalf("Failed to add ship: %v", err)
	}
	// Verify that ships are added correctly by checking the player ship count
	if (gameInstance.GetPlayers())[0].ShipCount != 2 {
		t.Errorf("Player 1 should have 1 ship, got %d", (gameInstance.GetPlayers())[0].ShipCount)
	}
	if (gameInstance.GetPlayers())[1].ShipCount != 2 {
		t.Errorf("Player 2 should have 1 ship, got %d", (gameInstance.GetPlayers())[0].ShipCount)
	}

	// Invalid case, try adding ship before init game
	gameInstance.Reset()
	// Add ships to both players
	err = gameInstance.AddShip("SHIP1", 2, 1, 1, 5, 2)
	expected := entity.ERR_GAME_HAS_NO_BOARD
	if err.Error() != expected.Error() {
		t.Errorf("Expected error %v, but got %v", expected.Error(), err.Error())
	}

	// Invalid case, try adding ship to where ship is already present
	err = gameInstance.InitGame(10)
	// Add ships to both players
	err = gameInstance.AddShip("SHIP1", 2, 1, 1, 2, 2)
	expected = entity.ErrInvalidCellShip(entity.NewCell(2, 2), "SHIP1")
	if err.Error() != expected.Error() {
		t.Errorf("Expected error %v, but got %v", expected.Error(), err.Error())
	}

	// Invalid case, try adding ship in opponents place
	err = gameInstance.InitGame(10)
	// Add ships to both players
	err = gameInstance.AddShip("SHIP1", 2, 1, 1, 3, 5)
	expected = entity.ErrInvalidCellOwner(entity.NewCell(3, 5), entity.NewCell(3, 5), &entity.Player{Id: 0})
	if err.Error() != expected.Error() {
		t.Errorf("Expected error %v, but got %v", expected.Error(), err.Error())
	}

}

// Test StartGame
func TestStartGame(t *testing.T) {
	gameInstance := getNewGameInstance()
	gameInstance.Reset()

	err := gameInstance.InitGame(10)
	if err != nil {
		t.Fatalf("Failed to initialize game: %v", err)
	}

	err = gameInstance.AddShip("SHIP1", 2, 1, 1, 5, 2)
	if err != nil {
		t.Fatalf("Failed to add ship: %v", err)
	}

	err = gameInstance.AddShip("SHIP1", 2, 3, 3, 8, 4)
	if err != nil {
		t.Fatalf("Failed to add ship: %v", err)
	}

	// Start the game
	err = gameInstance.StartGame()
	if err != nil {
		t.Fatalf("Game failed to start: %v", err)
	}

}

// Test StartGame
func TestInvalidStartGame(t *testing.T) {
	gameInstance := getNewGameInstance()
	gameInstance.Reset()

	// Start the game without init
	err := gameInstance.StartGame()
	expected := entity.ERR_GAME_HAS_NO_BOARD
	if err.Error() != expected.Error() {
		t.Errorf("Expected error %v, but got %v", expected.Error(), err.Error())
	}

	// Start the game without adding ship
	gameInstance.InitGame(10)
	err = gameInstance.StartGame()
	if err != nil {
		t.Fatalf("Game failed to start: %v", err)
	}

}

// Test Fire
func TestFire(t *testing.T) {
	gameInstance := getNewGameInstance()

	err := gameInstance.InitGame(10)
	if err != nil {
		t.Fatalf("Failed to initialize game: %v", err)
	}
	// Add ships to both players
	err = gameInstance.AddShip("SHIP1", 2, 1, 1, 5, 5)
	if err != nil {
		t.Errorf("Failed to add ship: %v", err)
	}

	// Fire at a ship's position
	cell := entity.NewCell(1, 1)
	destroyedShip, err := gameInstance.Fire(cell)
	if err != nil {
		t.Errorf("Failed to fire: %v", err)
	}

	// Verify that a ship is destroyed
	if destroyedShip == nil {
		t.Errorf("Expected ship to be destroyed, but got nil")
	}
	startLocation := entity.NewCell(1, 1)
	size := 2
	for i := startLocation.X; i < startLocation.X+size; i++ {
		for j := startLocation.Y; j < startLocation.Y+size; j++ {
			if gameInstance.GetBoard().GetCells()[j][i].Ship != nil {
				t.Errorf("Expected ship to be removed from battlefield, but its present")
			}
		}
	}
	fmt.Println(gameInstance.ViewBattleField())
}

// Test ViewBattleField
func TestViewBattleField(t *testing.T) {
	gameInstance := getNewGameInstance()

	err := gameInstance.InitGame(10)
	if err != nil {
		t.Fatalf("Failed to initialize game: %v", err)
	}

	// Add ships to both players
	err = gameInstance.AddShip("SHIP1", 2, 0, 0, 5, 1)
	if err != nil {
		t.Fatalf("Failed to add ship: %v", err)
	}

	// Add ships to both players
	err = gameInstance.AddShip("SHIP1", 3, 2, 2, 7, 7)
	if err != nil {
		t.Fatalf("Failed to add ship: %v", err)
	}

	// Get the battlefield view
	view := gameInstance.ViewBattleField()
	if view == "" {
		t.Errorf("Expected non-empty battlefield view")
	}
}
