package larksdk

import (
	"context"
	"cst-lark-sdk/hash_slot"
	"cst-lark-sdk/larksdk/model"
	"cst-lark-sdk/larksdk/protocol"
	"cst-lark-sdk/router_table"
	"strings"
)

type LarkSDK struct {
	opt         *Options
	ctx         context.Context
	client      protocol.LarkProtocol
	routerTable *router_table.RouterTable
}

func NewLarkSDK(context context.Context, client protocol.LarkProtocol, opt *Options) *LarkSDK {
	opt.init()
	routerTable := router_table.NewRouterTable(client)
	return &LarkSDK{
		opt:         opt,
		ctx:         context,
		client:      client,
		routerTable: routerTable,
	}
}

func (larkSDK *LarkSDK) NextSeq(ctx context.Context, bizTag string) (seqId int64, err error) {
	// bizTag为空
	if len(strings.Trim(bizTag, " ")) == 0 {
		return
	}
	// 计算槽
	slot := hash_slot.HashSlot(bizTag)
	// 获取槽对应的节点信息
	val, err := router_table.NewRouterTable(larkSDK.client).LookupRouteTable(ctx, slot)
	if err != nil || val == -1 {
		return
	}
	maxRetry := larkSDK.opt.MaxRedirects
RETRY:
	for {
		// TODO：获取seq
		seqId, err = larkSDK.client.NextSeq(ctx, bizTag)
		// todo:出错
		if maxRetry > 0 {
			maxRetry--
			//同步路由表
			goto RETRY
		}
		return
	}
	return
}

func (larkSDK *LarkSDK) Segment(ctx context.Context, bizTag string, count int32) (resp *model.SegmentResp, err error) {
	// bizTag为空
	if len(strings.Trim(bizTag, " ")) == 0 {
		return
	}
	resp.BizTag = bizTag
	if count == 0 {
		id, nextSeqErr := larkSDK.NextSeq(ctx, bizTag)
		if nextSeqErr != nil {
			return nil, nextSeqErr
		}
		resp.Start = id
		resp.End = id
	} else {
		slot := hash_slot.HashSlot(bizTag)
		_, getRouterTableErr := router_table.NewRouterTable(larkSDK.client).LookupRouteTable(ctx, slot)
		if getRouterTableErr != nil {
			return
		}
		maxRetry := larkSDK.opt.MaxRedirects
	RETRY:
		for {
			// TODO：获取Segment
			resp, err = larkSDK.client.Segment(ctx, bizTag, count)
			// todo:出错
			if maxRetry > 0 {
				maxRetry--
				goto RETRY
			}
		}
	}
	return
}
