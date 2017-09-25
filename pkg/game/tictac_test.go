package game

import (
	"testing"
)

func TestTicTac(t *testing.T) {
	size := 5
	player := 1

	game := NewTicTac(size)

	p := NewPosition(1, 1)
	game.Move(player, p)

	b := game.GetBoard()

	if b[1][1] != player {
		t.Errorf("failed [%d Vs. %d]", b[1][1], player)
	}
	game.PrintBoard()
}

func TestTicTac_GetStatus(t *testing.T) {
	size := 3
	player := 1

	game := NewTicTac(size)

	p := NewPosition(0, 0)
	game.Move(player, p)
	p.Set(1, 1)
	game.Move(player, p)
	p.Set(2, 2)

	if result := game.GetStatus(); result != -1 {
		t.Errorf("Wrong status: %d Vs. %d", result, -1)
	}

	game.Move(player, p)

	result := game.GetStatus()
	if result != player {
		t.Errorf("Wrong status: %d Vs. %d", result, player)
	}

	game.PrintBoard()
}

func TestTicTac_GetStatus2(t *testing.T) {
	size := 3
	player := 1

	game := NewTicTac(size)

	p := NewPosition(0, 1)
	game.Move(player, p)
	p.Set(1, 1)
	game.Move(player, p)
	p.Set(2, 1)

	if result := game.GetStatus(); result != -1 {
		t.Errorf("Wrong status: %d Vs. %d", result, -1)
	}

	game.Move(player, p)
	if result := game.GetStatus(); result != player {
		t.Errorf("Wrong status: %d Vs. %d", result, player)
	}

	game.PrintBoard()
}


func TestTicTac_GetStatus3(t *testing.T) {
	size := 3
	player := 1

	game := NewTicTac(size)

	p := NewPosition(1, 1)
	game.Move(player, p)
	p.Set(1, 0)
	game.Move(player, p)
	p.Set(1, 2)

	if result := game.GetStatus(); result != -1 {
		t.Errorf("Wrong status: %d Vs. %d", result, -1)
	}

	game.Move(player, p)
	if result := game.GetStatus(); result != player {
		t.Errorf("Wrong status: %d Vs. %d", result, player)
	}

	game.PrintBoard()
	game.PrintStatus()
}
