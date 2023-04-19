package request

// Paging common input parameter structure
type PageInfo struct {
	PageNo   int    `json:"pageNo" form:"pageNo"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	TokenId  string `json:"tokenId" form:"tokenId"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// Get role by id structure
type GetAuthorityId struct {
	AuthorityId string
}

type Empty struct{}
