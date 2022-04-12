package router_table

type SlotWithNodeTable struct {
	NodeId int32   `json:"node_id"` // 节点
	Slots  []Slots `json:"slots"`   // 节点对应插槽
}

/**
Slots对应存储关系
1.单个槽：start=end=slot
2.连续槽：[startslot, endslot],满足左闭右闭
*/
type Slots struct {
	Start int32 `json:"start"` // 开始
	End   int32 `json:"end"`   // 结束
}

func NewSlotTable() *SlotWithNodeTable {
	return &SlotWithNodeTable{}
}

// 将调用方返回的节点插槽关系进行映射
func (st *SlotWithNodeTable) transport() []SlotWithNodeTable {
	ret := make([]SlotWithNodeTable, 0)
	return ret
}
