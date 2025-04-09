package dto

type Meta struct {
	TotalData int64 `json:"total_data"`
	TotalPage int64 `json:"total_page"`
	Page      int64 `json:"page"`
	PageSize  int64 `json:"page_size"`
}
