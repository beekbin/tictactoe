package game

import (
	"fmt"
	"math/rand"
	//"github.com/golang/glog"
)

type Position struct {
	X int
	Y int
}

func NewPosition(x, y int) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

func (p *Position) Set(x, y int) {
	p.X = x
	p.Y = y
}

func (p *Position) String() string {
	return fmt.Sprintf("<%d, %d>", p.X, p.Y)
}

type TicTac struct {
	Board [][]int
	size  int
	totalMove int
}

func NewTicTac(size int) *TicTac {
	t := &TicTac{
		size: size,
		totalMove: 0,
	}

	a := make([][]int, size)
	for i := range a {
		a[i] = make([]int, size)
	}

	t.Board = a
	return t
}

func (t *TicTac) Clone() *TicTac {
	clone := &TicTac{
		size: t.size,
		totalMove: t.totalMove,
	}

	a := make([][]int, t.size)
	for i := range a {
		a[i] = make([]int, t.size)
	}

	for i := 0; i < t.size; i ++ {
		for j := 0; j < t.size; j ++ {
			a[i][j] = t.Board[i][j]
		}
	}
	clone.Board = a

	return clone
}

func (t *TicTac) GetBoard() [][]int {
	return t.Board
}

func (t *TicTac) Move(player int, p *Position) {
	t.Board[p.X][p.Y] = player
	t.totalMove ++
}

// check the status of the Game
// -1: in progress, 0: draw, 1: player1, 2: player2
func (t *TicTac) GetStatus() int {
	//glog.V(2).Infof("status to be done")
	result := t.checkRowStatus()
	if result != 0 {
		return result
	}

	result = t.checkColumnStatus()
	if result != 0 {
		return result
	}

	result = t.checkDialogStatus()
	if result != 0 {
		return result
	}

	if tmp := t.GetEmptyPositions(); len(tmp) == 0 {
		return 0
	}

	return -1
}

func (t *TicTac) checkDialogStatus() int {
	size := t.size

	row := make([]int, size)
	for i := 0; i < size; i ++ {
		row[i] = t.Board[i][i]
	}
	if result := checkRow(row); result != 0 {
		return result
	}

	for i := 0; i < size; i ++ {
		row[i] = t.Board[i][size-i-1]
	}
	if result := checkRow(row); result != 0 {
		return result
	}
	return 0
}

func (t *TicTac) checkColumnStatus() int {
	size := t.size

	row := make([]int, size)
	for i := 0; i < size; i ++ {
		for j := 0; j < size; j ++ {
			row[j] = t.Board[j][i]
		}

		if result := checkRow(row); result != 0 {
			return result
		}
	}

	return 0
}

func checkRow(row []int) int {
	if row[0] == 0 {
		return 0
	}

	for _, n := range row {
		if n != row[0] {
			return 0
		}
	}

	return row[0]
}

func (t *TicTac) checkRowStatus() int {
	size := t.size
	for i := 0; i < size; i++ {
		if result := checkRow(t.Board[i]); result != 0 {
			return result
		}
	}
	return 0
}


func (t *TicTac) GetEmptyPositions() []*Position {
	var result []*Position

	size := t.size

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if t.Board[i][j] == 0 {
				p := NewPosition(i, j)
				result = append(result, p)
			}
		}
	}

	return result
}

func (t *TicTac) RandomPlay(player int) error {
	empty := t.GetEmptyPositions()
	num := int32(len(empty))
	if num < 1 {
		return fmt.Errorf("Game over, no empty position.")
	}

	choice := rand.Int31n(num)
	p := empty[choice]

	t.Move(player, p)

	return nil
}

func (t *TicTac) PrintBoard() {

	size := t.size
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%d ", t.Board[i][j])
		}
		fmt.Println("")
	}
}

func (t *TicTac) PrintStatus() {
	status := t.GetStatus()
	switch (status) {
	case 1:
		fmt.Println("Player 1 win.")
	case 2:
		fmt.Println("Player 2 win.")
	case 0:
		fmt.Println("Game draw.")
	case -1:
		fmt.Println("Game in progress.")
	default:
		fmt.Printf("Wrong status [%d]!", status)
	}

}