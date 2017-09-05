package main

import (
	"fmt"
	"reflect"
)

// Tree is the interface that allows for traversal
type Tree interface {
	fmt.Stringer
	Left() Tree
	Right() Tree
}

func InOrder(root interface{}) []string {
	v := reflect.ValueOf(root)
	if v.Kind() != reflect.Ptr {
		return nil
	}
	var res []string
	t := ToTree(root)
	if t == nil {
		return nil
	}
	inOrderHelper(t, &res)
	return res
}

func ToTree(x interface{}) Tree {
	y := toTree(reflect.ValueOf(x))
	return y
}

// Takes a value and tries to walk it like a tree
func toTree(vv reflect.Value) Tree {
	for vv.Kind() == reflect.Ptr || vv.Type().Kind() == reflect.Interface {
		if vv.IsNil() {
			return nil
		}
		vv = vv.Elem()
	}
	var l, r, v, e reflect.Value
	for _, i := range []string{"left", "Left"} {
		l = vv.FieldByName(i)
		if l != e {
			break
		}
	}
	for _, i := range []string{"right", "Right"} {
		r = vv.FieldByName(i)
		if r != e {
			break
		}
	}
	if l == e || r == e {
		return nil
	}
	for _, i := range []string{"val", "Val", "payload", "Payload"} {
		v = vv.FieldByName(i)
		if v != e {
			// todo: print nicer things if this is an interface
			break
		}
	}
	if v == e {
		v = reflect.ValueOf("cannot determine value of node")
	}
	return &tree{
		left:  toTree(l),
		right: toTree(r),
		val:   v,
	}
}

type tree struct {
	left  Tree
	right Tree
	val   interface{}
}

func (l *tree) Left() Tree {
	return l.left
}
func (l *tree) Right() Tree {
	return l.right
}
func (l *tree) String() string {
	return fmt.Sprintf("%v", l.val)
}

func inOrderHelper(root Tree, res *[]string) {
	if root == nil {
		return
	}
	inOrderHelper(root.Left(), res)
	*res = append(*res, root.String())
	inOrderHelper(root.Right(), res)
}
