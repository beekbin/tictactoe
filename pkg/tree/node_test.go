package tree


import (
	"testing"
	"github.com/beekbin/tictactoe/pkg/game"
	"fmt"
)

func TestNode_Simulation(t *testing.T) {
	player := 1
	opponent := 2

	num := 50
	board := game.NewTicTac(4)
	allPositions := board.GetEmptyPositions()
	scores := make([]int, len(allPositions))

	for i, p := range allPositions {
		node := NewNode(opponent, player, board)
		node.board.Move(player, p)

		scores[i] = 0
		for j := 0; j < num; j ++ {
			result := node.Simulation(opponent)
			if result == player {
				scores[i] += 1
			}
		}

		//fmt.Printf("%v = %d\n", p, scores[i])
	}
}

func TestNode_Simulation2(t *testing.T) {
	player := 1
	opponent := 2

	num := 50
	board := game.NewTicTac(3)
	board.Board[0][0] = opponent
	board.Board[1][1] = opponent
	board.Board[2][1] = opponent
	board.Board[2][0] = player
	board.Board[2][2] = player

	allPositions := board.GetEmptyPositions()
	scores := make([]int, len(allPositions))

	for i, p := range allPositions {
		node := NewNode(opponent, player, board)
		node.board.Move(player, p)

		scores[i] = 0
		for j := 0; j < num; j ++ {
			result := node.Simulation(opponent)
			if result == player {
				scores[i] += 1
			} else if result == opponent {
				scores[i] -= 1
			}
		}

		//fmt.Printf("V2--%v = %d\n", p, scores[i])
	}
}


func TestNode_Simulation3(t *testing.T) {
	player := 1
	opponent := 2

	num := 500000
	board := game.NewTicTac(3)
	board.Board[2][2] = opponent
	board.Board[1][1] = opponent
	board.Board[2][0] = player

	allPositions := board.GetEmptyPositions()
	scores := make([]int, len(allPositions))

	for i, p := range allPositions {
		node := NewNode(opponent, player, board)
		node.board.Move(player, p)

		scores[i] = 0
		for j := 0; j < num; j ++ {
			result := node.Simulation(opponent)
			if result == player {
				scores[i] += 1
			} else if result == opponent {
				scores[i] -= 1
			}
		}

		fmt.Printf("V3--%v = %d\n", p, scores[i])
	}
}