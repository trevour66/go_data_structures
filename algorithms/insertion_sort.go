package algorithms

import (
	"fmt"
	"go_data_structures/data_structures"
)


func Sort(sliceOfInt []int){
	var root *data_structures.Tree

	for _, value := range sliceOfInt {
		root = add(root, value)
	}

	fmt.Print(appendValues(sliceOfInt[:0], root))
}

func appendValues(values []int, root *data_structures.Tree) []int {
	if(root != nil){
		values = appendValues(values, root.Left)
		values = append(values, root.Value)
		values = appendValues(values, root.Right)
	}
	return values
}

func add(t *data_structures.Tree, v int) *data_structures.Tree  {
	if t == nil {
		t = new(data_structures.Tree)
		t.Value = v
		return t
	}

	if(v < t.Value){
		t.Left = add(t.Left, v)
	}else {
		t.Right = add(t.Right, v)
	}

	return t
}