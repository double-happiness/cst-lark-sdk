package protocol

import (
	"context"
	"cst-lark-sdk/larksdk/model"
)

type LarkProtocol interface {
	SyncRouteTable(ctx context.Context) (map[int]int, error)
	NextSeq(ctx context.Context, bizTag string) (int64, error)
	Segment(ctx context.Context, bizTag string, count int32) (segmentResp *model.SegmentResp, err error)
}
