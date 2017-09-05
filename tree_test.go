package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"strings"
)

type goodNode struct {
	left, right *goodNode
	val         int64
}

func TestNilDoesNotPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()
	var x *goodNode = nil
	ToTree(x)
}

func TestGettingInOrderListOfTree(t *testing.T) {
	x := &goodNode{val: 1, left: &goodNode{val: 2}, right: &goodNode{val: 3}}
	i := InOrder(x)
	assert.EqualValues(t, strings.Split("2,1,3", ","), i)
}
