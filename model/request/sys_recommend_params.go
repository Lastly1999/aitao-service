package request

type RecommendParams struct {
	Pager
	MaterialId string `json:"materialId"` // 官方的物料Id(详细物料id见：https://market.m.taobao.com/app/qn/toutiao-new/index-pc.html#/detail/10628875?_k=gpov9a)
}
