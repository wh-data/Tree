package Tree

import (
	"container/list"
	"fmt"
)

type Tree struct {
	Root   *TreeNode
	Height int
	Size   int
}

func (tree *Tree) GetHeight() int {
	if tree != nil {
		return tree.Height
	}
	return 0
}

func (tree *Tree) GetSize() int {
	if tree != nil {
		return tree.Size
	}
	return 0
}

func (tree *Tree) SetHeight() {
	if tree == nil && tree.Root != nil {
		tree.Height = tree.Root.GetHeight()
		return
	}
	return
}

func (tree *Tree) SetSize() {
	if tree == nil && tree.Root != nil {
		tree.Size = tree.Root.GetSize()
		return
	}
	return
}

func ConstructTreeWithExampleData(data []*TreeNodeExampleData) (tree *TreeNode) {
	root := NewTreeNode(nil)
	tempmap := make(map[int]*TreeNode, 0)
	queue := list.New()
	for _, d := range data {
		e := queue.PushBack(d)
		if d.GetPid() == 0 {
			node := NewTreeNode(d)
			root.Children = append(root.Children, node)
			queue.Remove(e)
			tempmap[d.GetId()] = node
		}

	}
	count := 0
	for {
		if queue.Len() == 0 || count == queue.Len() {
			break
		}
		e := queue.Front()
		queue.Remove(e)
		if d, ok := e.Value.(*TreeNodeExampleData); !ok {
			fmt.Println("wrong type")
		} else if parent, ok := tempmap[d.GetPid()]; ok {
			node := NewTreeNode(d)
			parent.Children = append(parent.Children, node)
			tempmap[d.GetId()] = node
			count = 0
		} else {
			queue.PushBack(e)
			count++
		}

	}
	return root
}
