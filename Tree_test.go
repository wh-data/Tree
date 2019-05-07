package Tree

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	pb "github.com/wh-data/Tree/exampleData"
)

//this function is for test usage
func NewTreeNodeExampleData(id, pid *int32, extinfo []byte) *pb.TreeNodeExampleData {
	return &pb.TreeNodeExampleData{
		Id:      id,
		Pid:     pid,
		Extinfo: extinfo,
	}
}

func BuildTreeWithExampleData(data []*pb.TreeNodeExampleData) (tree *TreeNode) {
	root := NewTreeNode(nil)
	tempmap := make(map[int32]*TreeNode, 0)
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
		d, ok := e.Value.(*pb.TreeNodeExampleData)
		if !ok {
			fmt.Println("data type wrong! ")
			queue.PushBack(e)
			count++
			continue
		}
		if parent, exist := tempmap[d.GetPid()]; exist {
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

func TestTree(t *testing.T) {
	data1 := NewTreeNodeExampleData(proto.Int32(1), proto.Int32(0), []byte("hello tree"))
	data2 := NewTreeNodeExampleData(proto.Int32(2), proto.Int32(0), []byte("x"))
	data3 := NewTreeNodeExampleData(proto.Int32(3), proto.Int32(1), []byte("y"))
	dataArray := make([]*pb.TreeNodeExampleData, 0)
	dataArray = append(dataArray, data1)
	dataArray = append(dataArray, data2)
	dataArray = append(dataArray, data3)
	root := BuildTreeWithExampleData(dataArray)
	//can use protobuf,  then print can show its value
	if data, ok := root.GetChildByIndex(0).Data.(*pb.TreeNodeExampleData); ok {
		t.Logf("The id is: %d, pid is: %d, extinfo is: %v\n", data.GetId(), data.GetPid(), string(data.Extinfo))
	} else {
		t.Log("wrong data type")
	}

}
