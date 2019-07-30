package node_test

import (
	"testing"

	"github.com/a6cexz/goexperiments/list/node"
	"github.com/stretchr/testify/assert"
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

func TestGetValue(t *testing.T) {
	list1 := node.New(1)
	assert.Equal(t, 1, list1.GetValue())

	list2 := node.New(1, 2)
	assert.Equal(t, 2, list2.GetTail().GetValue())

	list3 := node.New(1, 2, 3)
	assert.Equal(t, 3, list3.GetTail().GetTail().GetValue())
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

func TestSum(t *testing.T) {
	list1 := node.New(1)
	assert.Equal(t, 1, list1.Sum())

	list2 := node.New(1, 2)
	assert.Equal(t, 3, list2.Sum())

	list3 := node.New(1, 2, 3)
	assert.Equal(t, 6, list3.Sum())
}

func TestMin(t *testing.T) {
	list1 := node.New(1)
	assert.Equal(t, 1, list1.Min())

	list2 := node.New(1, 2)
	assert.Equal(t, 1, list2.Min())

	list3 := node.New(1, 2, 3)
	assert.Equal(t, 1, list3.Min())

	list4 := node.New(3, 2, 1)
	assert.Equal(t, 1, list4.Min())
}

func TestMax(t *testing.T) {
	list1 := node.New(1)
	assert.Equal(t, 1, list1.Max())

	list2 := node.New(1, 2)
	assert.Equal(t, 2, list2.Max())

	list3 := node.New(1, 2, 3)
	assert.Equal(t, 3, list3.Max())
}

func TestMinNegative(t *testing.T) {
	list1 := node.New(-10, -20, 0, 100, -100)
	assert.Equal(t, -100, list1.Min())
}

func TestMaxNegative(t *testing.T) {
	list1 := node.New(-10, -20, 0, 100, -100)
	assert.Equal(t, 100, list1.Max())
}
