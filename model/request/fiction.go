package request

// User register structure
type FictionApiParam struct {
	PageInfo
	SortId    int64 	`json:"sortId" form:"sortId"`
}
