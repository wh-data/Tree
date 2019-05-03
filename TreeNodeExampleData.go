//example data for tree node data
package Tree

type TreeNodeExampleData struct {
	Id      int
	Pid     int
	Extinfo []byte
}

func NewTreeNodeExampleData(id, pid int, extinfo []byte) *TreeNodeExampleData {
	return &TreeNodeExampleData{
		Id:      id,
		Pid:     pid,
		Extinfo: extinfo,
	}
}

func (data *TreeNodeExampleData) GetId() int {
	if data != nil {
		return data.Id
	}
	return 0
}

func (data *TreeNodeExampleData) GetPid() int {
	if data != nil {
		return data.Pid
	}
	return 0
}
