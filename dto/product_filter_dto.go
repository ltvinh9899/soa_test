package dto

type ProductFilter struct {
    SearchQuery string   `form:"search"`
	Type 	  string   `form:"type"`
    Status      string   `form:"status"`
}
