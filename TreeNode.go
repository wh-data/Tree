package Tree

type TreeNode struct {
	Children []*TreeNode
	Data     interface{}
}

func NewTreeNode(data interface{}) *TreeNode {
	if data == nil {
		return &TreeNode{}
	}
	return &TreeNode{
		Data: data,
	}
}

func (node *TreeNode) GetData() interface{} {
	if node != nil {
		return node.Data
	}
	return nil
}

func (node *TreeNode) GetChildren() []*TreeNode {
	if node != nil {
		return node.Children
	}
	return nil
}

func (node *TreeNode) GetChildByIndex(idx int) *TreeNode {
	if node != nil && idx < len(node.Children) && idx >= 0 {
		return node.Children[idx]
	}
	return nil
}
