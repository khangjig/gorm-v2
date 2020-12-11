package model

type Paginator struct {
	Page  int `json:"page,omitempty" query:"page"`
	Limit int `json:"limit,omitempty" query:"limit"`
}

type Condition struct {
	Type    string
	Pattern string
	Values  []interface{}
}

type GetAllRequest struct {
	Paginator Paginator `json:"paginator"`
	SortBy    SortType  `json:"sort_by,omitempty" query:"sort_by"`
	OrderBy   string    `json:"order_by,omitempty" query:"order_by"`
}
