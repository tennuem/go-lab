package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	g := NewGraph()

	g.AddNode(1)
	g.AddNode(2)
	err := g.AddLine(1, 2)

	firstNode := g.FindNode(1)
	endNode := g.FindNode(2)

	assert.NoError(t, err)
	assert.Equal(t, endNode, firstNode.Lines[0])
}
