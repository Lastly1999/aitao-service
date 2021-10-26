package response

type CommoditParams struct {
	Keywords string `json:"keywords"`
	PageSize int    `json:"pageSize"`
	PageOn   int    `json:"pageOn"`
}
