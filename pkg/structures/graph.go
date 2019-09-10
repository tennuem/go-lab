package structures

import "github.com/pkg/errors"

type Node struct {
	Value int
	Lines []*Node
}

type Graph interface {
	AddNode(value int)
	FindNode(value int) *Node
	AddLine(startValue, endValue int) error
}

func NewGraph() Graph {
	return &graph{}
}

type graph struct {
	nodes []*Node
}

func (g *graph) AddNode(value int) {
	g.nodes = append(g.nodes, &Node{Value: value})
}

func (g *graph) FindNode(value int) *Node {
	var result *Node
	for _, v := range g.nodes {
		if value == v.Value {
			result = v
		}
	}
	return result
}

func (g *graph) AddLine(startValue, endValue int) error {
	startNode := g.FindNode(startValue)
	endNode := g.FindNode(endValue)
	if startNode == nil || endNode == nil {
		return errors.New("Node not found")
	}
	startNode.Lines = append(startNode.Lines, endNode)
	return nil
}
