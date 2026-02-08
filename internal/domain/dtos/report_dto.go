package dtos

type TodayReportDto struct {
	TotalRevenue       int                    `json:"total_revenue"`
	TotalTransactions  int                    `json:"total_transactions"`
	BestSellingProduct *BestSellingProductDto `json:"best_selling_product"`
}

type BestSellingProductDto struct {
	Name     string `json:"name"`
	QtySold  int    `json:"qty_sold"`
}
