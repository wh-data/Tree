package Tree

type Tree struct {
	Root   *TreeNode
	Height int
	Size   int
}

func NewTree(root *TreeNode) *Tree {
	if root == nil {
		return &Tree{}
	}
	return &Tree{
		Root:   root,
		Height: root.GetHeight(),
		Size:   root.GetSize(),
	}
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
