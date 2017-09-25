package mcts

import (
	"time"
	_ "github.com/golang/glog"

	"github.com/songbinliu/tictac/pkg/tree"
	"github.com/songbinliu/tictac/pkg/game"
	_ "fmt"
)

const (

	defaultWinScore = float64(10.0)
	defaultMaxDepth = 10
	defaultMaxDuration = time.Millisecond * 820
)

type MCTS struct {
	WinScore float64
	MaxDepth int
	MaxDuration time.Duration

	step int
}

func NewMCTS() *MCTS {
	return &MCTS {
		WinScore: defaultWinScore,
		MaxDepth: defaultMaxDepth,
		MaxDuration: defaultMaxDuration,
		step : 0,
	}
}

func (mcts *MCTS) GetNextMove(player, opponent int, board *game.TicTac) *game.Position {

	node := tree.NewNode(player, opponent, board)
	mytree := tree.NewTree(node)
	deadline := time.Now().Add(mcts.MaxDuration)

	i := 0
	for time.Now().Before(deadline) {
		//1. select by UCT
		snode := mytree.SelectLeafNode()

		//2. expand
		if snode.GetStatus() == -1 {
			snode.Expand()
			snode = snode.RandomChild()
		}

		//3. simulation
		result := snode.Simulation(opponent)

		//4. update by back propagation
		reward := mcts.WinScore
		if result == 0 {
			reward = 0
		}
		snode.BackPropagate(result, reward)
		i ++
	}

	//fmt.Printf("i = %d\n", i)
	mcts.step ++
	//if mcts.step == 1 {
	//	glog.V(2).Infof("begin to print tree L2")
	//	mytree.PrintTree()
	//	glog.V(2).Infof("end of printing tree L2")
	//}

	child := mytree.GetRoot().SelectMostVisitedChild()
	//fmt.Printf("move=%v\n", child.Move)
	return child.GetMove()
}
