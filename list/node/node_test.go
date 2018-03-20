package node_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zdishou/goexpirements/list/node"
)

func TestNew(t *testing.T) {
	assert.Equal(t, "[1]", node.New(1).String())
	assert.Equal(t, "[1, 2]", node.New(1, 2).String())
	assert.Equal(t, "[1, 2, 3]", node.New(1, 2, 3).String())
}

func TestGetTail(t *testing.T) {
	list1 := node.New(1)
	assert.Nil(t, list1.GetTail())

	list2 := node.New(1, 2)
	assert.Equal(t, "[2]", list2.GetTail().String())

	list3 := node.New(1, 2, 3)
	assert.Equal(t, "[2, 3]", list3.GetTail().String())
}

func TestReverseList(t *testing.T) {
	list1 := node.New(1)
	assert.Equal(t, "[1]", list1.Reverse().String())

	list2 := node.New(1, 2)
	assert.Equal(t, "[2, 1]", list2.Reverse().String())
	assert.Equal(t, "[2, 1]", list2.Reverse().String())

	list3 := node.New(1, 2, 3)
	assert.Equal(t, "[3, 2, 1]", list3.Reverse().String())
	assert.Equal(t, "[3, 2, 1]", list3.Reverse().String())
}
