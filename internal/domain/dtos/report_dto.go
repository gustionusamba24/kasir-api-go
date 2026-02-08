package dtos

type TodayReportDto struct {
	TotalRevenue       int                    `json:"total_revenue"`
	TotalTransactions  int                    `json:"total_transactions"`
	BestSellingProduct *BestSellingProductDto `json:"best_selling_product"`
}
type DateRangeReportDto struct {
	StartDate          string                 `json:"start_date"`
	EndDate            string                 `json:"end_date"`
	TotalRevenue       int                    `json:"total_revenue"`
	TotalTransactions  int                    `json:"total_transactions"`
	BestSellingProduct *BestSellingProductDto `json:"best_selling_product"`
}
type BestSellingProductDto struct {
	Name     string `json:"name"`
	QtySold  int    `json:"qty_sold"`
}
