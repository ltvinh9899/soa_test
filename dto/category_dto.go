package dto

type CategoryStat struct {
    CategoryName string `json:"categoryName"`
    ProductCount int    `json:"productCount"`
}

type DashboardResponse struct {
    CategoryStats []CategoryStat `json:"categoryDistribution"`
}