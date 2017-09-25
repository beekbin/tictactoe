package main

import (
	"math/rand"
	"time"
	"flag"
	"github.com/songbinliu/tictac/pkg/game"
	"github.com/songbinliu/tictac/pkg/mcts"
	"fmt"
)

func setFlag() {
	flag.Set("logtostderr", "true")
	flag.Set("v", "3")

	flag.Parse()
}


func testSingleMove() {
	mc := mcts.NewMCTS()

	player := 1
	opponent := 2

	size := 3
	board := game.NewTicTac(size)
	board.Board[0][0] = opponent

	for i := 0; i < 1; i ++ {
		//1. player1 move
		p := mc.GetNextMove(player, opponent, board)
		fmt.Printf("%v\n", p)

		board.Move(player, p)
		if board.GetStatus() != -1 {
			break
		}

		////2. opponent move
		////board.RandomPlay(opponent)
		//p = mc.GetNextMove(opponent, player, board)
		//board.Move(opponent, p)
		//if board.GetStatus() != -1 {
		//	break
		//}
	}

	board.PrintBoard()
	board.PrintStatus()

	return
}

func testDraw() (int, []*game.Position) {
	mc := mcts.NewMCTS()

	player := 1
	opponent := 2

	size := 3
	board := game.NewTicTac(size)
	board.Board[0][0] = opponent
	//board.Board[1][1] = opponent
	//board.Board[2][0] = player

	moves := []*game.Position{}

	for i := 0; i < size*size; i ++ {
		//1. player1 move
		p := mc.GetNextMove(player, opponent, board)
		board.Move(player, p)
		moves = append(moves, p)

		if board.GetStatus() != -1 {
			break
		}

		//2. opponent move
		//board.RandomPlay(opponent)
		p = mc.GetNextMove(opponent, player, board)
		board.Move(opponent, p)
		moves = append(moves, p)
		if board.GetStatus() != -1 {
			break
		}
	}

	//board.PrintBoard()
	board.PrintStatus()

	return board.GetStatus(), moves
}

func testWin() {
	mc := mcts.NewMCTS()

	player := 1
	opponent := 2

	size := 3
	board := game.NewTicTac(size)

	for i := 0; i < size*size; i ++ {
		//1. player1 move
		p := mc.GetNextMove(player, opponent, board)
		board.Move(player, p)
		if board.GetStatus() != -1 {
			break
		}

		//2. opponent move
		board.RandomPlay(opponent)
		if board.GetStatus() != -1 {
			break
		}
	}

	board.PrintBoard()
	board.PrintStatus()
}


func testDraw2() {
	errNum := 0
	total := 100
	for i := 0; i < total; i ++ {
		status, moves := testDraw()
		if status != 0 {
			errNum += 1
			fmt.Printf("errNum %d/%d %d\n", errNum, i, total)
			fmt.Printf("%v \n", moves)
		}
	}

	fmt.Printf("errNum %d/%d\n", errNum, total)
}


func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	setFlag()
	//testDraw()
	testDraw2()
	//testSingleMove()
	//testWin()
}