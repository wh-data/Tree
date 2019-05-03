package Tree

import "testing"

func TestTree(t *testing.T) {
	data1 := NewTreeNodeExampleData(1, 0, []byte("hello tree"))
	data2 := NewTreeNodeExampleData(2, 0, []byte("x"))
	data3 := NewTreeNodeExampleData(3, 1, []byte("y"))
	dataArray := make([]*TreeNodeExampleData, 0)
	dataArray = append(dataArray, data1)
	dataArray = append(dataArray, data2)
	dataArray = append(dataArray, data3)
	root := ConstructTreeWithExampleData(dataArray)
	//can use protobuf,  then print can show its value
	if data, ok := root.GetChildByIndex(0).Data.(*TreeNodeExampleData); ok {
		t.Logf("The id is: %d, pid is: %d, extinfo is: %v\n", data.GetId(), data.GetPid(), string(data.Extinfo))
	} else {
		t.Log("wrong data type")
	}

}
