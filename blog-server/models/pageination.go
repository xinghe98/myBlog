package models

// Pageination 分页结构体
type Pageination struct {
	Limit  int    `json:"limit" form:"limit" uri:"limit"`
	Page   int    `json:"page" form:"page" uri:"page"`
	Sort   string `json:"sort" form:"sort" uri:"sort"`
	Status int    `json:"status" form:"status" uri:"status"`
}
