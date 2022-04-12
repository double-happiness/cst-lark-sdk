package router_table

import "sync"

/**
map[int]int
key	slot
value	nodeId
*/
var SlotRouterTable sync.Map
