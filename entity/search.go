package entity

type Query struct {
	SearchQuery string `form:"search"`
	SortQuery   string `form:"sort"`
}
