package Tree

type TreeNode struct {
	Parent   *TreeNode
	Children []*TreeNode
	Data     interface{}
	Height   int
	Size     int
}

func NewTreeNode(data interface{}) *TreeNode {
	if data == nil {
		return &TreeNode{}
	}
	return &TreeNode{
		Data:   data,
		Height: 0,
		Size:   1,
	}
}

func (node *TreeNode) GetData() interface{} {
	if node != nil {
		return node.Data
	}
	return nil
}

func (node *TreeNode) GetParent() *TreeNode {
	if node != nil {
		return node.Parent
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

func (node *TreeNode) GetHeight() int {
	if node != nil {
		return node.Height
	}
	return 0
}

func (node *TreeNode) GetSize() int {
	if node != nil {
		return node.Size
	}
	return 0
}

func (node *TreeNode) SetHeight() {
	if node == nil {
		return
	}
	height := 0
	for _, ch := range node.GetChildren() {
		if ch.GetHeight()+1 > height {
			height = ch.GetHeight() + 1
		}
	}
	if node.GetHeight() == height {
		//no need regression
		return
	}
	node.Height = height
	if node.GetParent() == nil {
		//the end of regression
		return
	}
	node.GetParent().SetHeight()

}

func (node *TreeNode) SetSize() {
	if node == nil {
		return
	}
	size := 1
	for _, ch := range node.GetChildren() {
		size += ch.GetSize()
	}
	if node.GetSize() == size {
		//no need regression
		return
	}
	node.Size = size
	if node.GetParent() == nil {
		//the end of regression
		return
	}
	node.GetParent().SetSize()
}

//todo
//set left child
//set right child
//cut off from parent
//go protobuf
//map lock, unlock
