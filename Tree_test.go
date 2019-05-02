package Tree

import "testing"

func TestTree(t *testing.T) {
	data1 := NewTreeNodeData(1, 0, "0")
	data2 := NewTreeNodeData(2, 0, "x")
	data3 := NewTreeNodeData(3, 1, "y")
	dataArray := make([]*TreeNodeData, 0)
	dataArray = append(dataArray, data1)
	dataArray = append(dataArray, data2)
	dataArray = append(dataArray, data3)
	root := ConstructTree(dataArray)
	//can use protobuf,  then print can show its value
	t.Log(root)

}
