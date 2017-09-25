package tree

import (
	"github.com/beekbin/tictactoe/pkg/game"

	"github.com/golang/glog"
	"math"
	"math/rand"
	"fmt"
)

const (
	uctEmptyValue = float64(1e10)
	defaultLoseBigScore = float64(-1e12)
)

type Node struct {
	player int
	opponent int
	board *game.TicTac

	winCount int64
	visitCount float64
	winScore float64
	UCTValue float64

	Parent   *Node
	Children []*Node
	Move *game.Position
}

func NewNode(player, opponent int, board *game.TicTac) *Node {
	return &Node{
		player: player,
		opponent: opponent,
		board: board.Clone(),
		visitCount: 0.0,
		winCount: 0,
		winScore: 0.0,
		UCTValue: uctEmptyValue,

		Children: []*Node{},
		Parent: nil,
	}
}

func (n *Node) String() string {
	//return "aa"
	p := ""
	if n.Move != nil {
		p = fmt.Sprintf("p=(%d, %d)", n.Move.X, n.Move.Y)
	}
	return fmt.Sprintf("player=%d, opponent=%d, visitCount=%.1f, winScore=%.1f, uct=%.1f, %s",
	n.player, n.opponent, n.visitCount, n.winScore, n.UCTValue, p)
}

func (n *Node) SetWinScore(s float64) {
	n.winScore = s;
}

func (n *Node) CalcUCT() {
	if n.Parent == nil {
		return
	}

	if n.visitCount < 1 {
		//glog.Error("potential bug.")
		n.UCTValue = uctEmptyValue
		return
	}

	avg := n.winScore/n.visitCount
	explore := math.Sqrt(math.Log(n.Parent.visitCount)/n.visitCount)

	n.UCTValue = avg + 1.41 * explore
}

func (n *Node) GetStatus() int {
	return n.board.GetStatus()
}

func (n *Node) GetMove() *game.Position {
	return n.Move
}

func (n *Node) Expand() {
	positions := n.board.GetEmptyPositions()
	children := make([]*Node, len(positions))

	for i, position := range positions {
		aboard := n.board.Clone()
		aboard.Move(n.player, position)

		//switch players
		anode := NewNode(n.opponent, n.player, aboard)
		anode.Parent = n
		anode.Move = position
		children[i] = anode
	}

	n.Children = children
}

func (n *Node) RandomChild() *Node {
	if n.Children == nil || len(n.Children) < 1 {
		fmt.Println("Potential bug.")
		return n
	}

	num := len(n.Children)
	rs := rand.Int31n(int32(num))
	node := n.Children[rs]

	return node
}

func (n *Node) SelectMostUCTChild_heavy() *Node {
	if len(n.Children) < 1 {
		glog.Errorf("Potential bug.")
		return n
	}

	node := n.Children[0]
	node.CalcUCT()

	for i := 1; i < len(n.Children); i ++ {
		n.Children[i].CalcUCT()
		if n.Children[i].UCTValue > node.UCTValue {
			node = n.Children[i]
		}
	}

	return node
}

func (n *Node) SelectMostUCTChild() *Node {
	if len(n.Children) < 1 {
		glog.Errorf("Potential bug.")
		return n
	}

	node := n.Children[0]
	for i := 1; i < len(n.Children); i ++ {
		if n.Children[i].UCTValue > node.UCTValue {
			node = n.Children[i]
		}
	}

	return node
}

func (n *Node) SelectMostVisitedChild() *Node {
	if n.Children == nil || len(n.Children) < 1 {
		glog.Errorf("Potential bug.")
		return n
	}

	node := n.Children[0]
	for i := 1; i < len(n.Children); i ++ {
		if n.Children[i].visitCount > node.visitCount {
		//if n.Children[i].winCount > node.winCount {
			node = n.Children[i]
		}
	}

	return node
}

func (n *Node) setBigValue(opponent, winner int ) {
	if winner == 0 {
		return
	}

	p := n.Parent
	if p.opponent == winner {
		p.SetWinScore(1e8)
	} else {
		p.SetWinScore(defaultLoseBigScore)
	}
}

func (n *Node) Simulation(overall_opponent int) int {
	result := n.board.GetStatus()
	if result != -1 {
		//n.setBigValue(overall_opponent, result)
		if result == overall_opponent {
			glog.V(4).Infof("begin to set big loss. why?")
			n.Parent.SetWinScore(defaultLoseBigScore)
		}
		return result
	}

	board := n.board.Clone()

	player := n.player
	opponent := n.opponent

	for result == -1 {
		board.RandomPlay(player)
		result = board.GetStatus()

		player, opponent = opponent, player
	}

	return result
}

func (n *Node) BackPropagate(winner int, reward float64) {
	node := n

	for node != nil {
		node.visitCount ++
		if node.opponent == winner {
			node.winScore += reward
			node.winCount += 1
		}

		node.CalcUCT()
		node = node.Parent
	}
}