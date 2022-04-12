package router_table

import (
	"context"
	"cst-lark-sdk/larksdk/protocol"
	"errors"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type RouterTable struct {
	client protocol.LarkProtocol
}

func NewRouterTable(client protocol.LarkProtocol) *RouterTable {
	return &RouterTable{
		client: client,
	}
}

/**
查找路由表，返回slot所在节点信息
*/
func (rt *RouterTable) LookupRouteTable(ctx context.Context, slot int) (int, error) {
	nodeId, ok := SlotRouterTable.Load(slot)
	if !ok || nodeId == nil {
		err := rt.GetRouteTable(ctx)
		if err != nil {
			return -1, err
		}
		nodeId, ok = SlotRouterTable.Load(slot)
		if !ok || nodeId == nil {
			return -1, err
		}
	}
	return nodeId.(int), nil
}

/**
GetRouteTable	拉取路由表信息
*/
var r = rate.Every(2 * time.Second)
var limit = rate.NewLimiter(r, 1)

func (rt *RouterTable) GetRouteTable(ctx context.Context) (err error) {
	// 调用larkseq服务的更新路由表接口
	var routeTable sync.Map
	SyncRouteTable := make(map[int]int)
	if limit.Allow() {
		SyncRouteTable, err = rt.client.SyncRouteTable(ctx)
		if err != nil {
			return err
		}
		for k, v := range SyncRouteTable {
			routeTable.Store(k, v)
		}
		SlotRouterTable = routeTable
		return nil
	}
	return errors.New("updating slot table")
	// 将获取的路由表更新
	//rt.UpdateSlotRouterTable(NewSlotTable().transport(syncRouteTableResp))
	return
}

//
///**
//UpdateSlotRouterTable	更新新的路由表
//*/
//func (rt *RouterTable) UpdateSlotRouterTable(routerTables []SlotWithNodeTable) {
//	var tmpRouterTable sync.Map
//	for _, routerTable := range routerTables {
//		for _, v := range routerTable.Slots {
//			for i := v.Start; i <= v.End; i++ {
//				tmpRouterTable.Store(i, routerTable.NodeId)
//			}
//		}
//	}
//	SlotRouterTable = tmpRouterTable
//	tmpRouterTable = sync.Map{}
//}
