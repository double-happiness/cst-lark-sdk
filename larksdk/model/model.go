package model

type SegmentResp struct {
	BizTag string `json:"biz_tag"`
	Start  int64  `json:"start"`	// 开始seq
	End    int64  `json:"end"`		// 结束seq
}
