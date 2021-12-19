package common

type Page struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}
