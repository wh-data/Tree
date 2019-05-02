package Tree

type TreeNode struct {
	Children []*TreeNode
	Data     *TreeNodeData
}

type TreeNodeData struct {
	Id   int
	Pid  int
	Data interface{}
}

func NewTreeNodeData(id, pid int, data interface{}) *TreeNodeData {
	return &TreeNodeData{
		Id:   id,
		Pid:  pid,
		Data: data,
	}
}

func NewTreeNode(data *TreeNodeData) *TreeNode {
	if data == nil {
		return &TreeNode{}
	}
	return &TreeNode{
		Data: data,
	}
}

func (data *TreeNodeData) GetId() int {
	if data != nil {
		return data.Id
	}
	return 0
}

func (node *TreeNode) GetId() int {
	if node != nil {
		return node.Data.GetId()
	}
	return 0
}

func (data *TreeNodeData) GetPid() int {
	if data != nil {
		return data.Pid
	}
	return 0
}

func (node *TreeNode) GetPid() int {
	if node != nil {
		return node.Data.GetPid()
	}
	return 0
}

func (node *TreeNode) GetData() *TreeNodeData {
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
