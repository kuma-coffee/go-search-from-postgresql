package entity

type Query struct {
	QueryParams string `form:"q"`
	SortQuery   string `form:"sort"`
}
