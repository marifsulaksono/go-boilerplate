package model

type Pagination struct {
	Page  int `json:"page" query:"page" validate:"gte=1"`
	Limit int `json:"pageSize" query:"pageSize" validate:"gte=1,lte=100"`
}

func (p *Pagination) SetDefault() {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit < 1 {
		p.Limit = 25
	}
}

type PaginationResponse struct {
	List         interface{} `json:"lists"`
	Page         int         `json:"page"`
	Limit        int         `json:"limit"`
	TotalPerPage int         `json:"total_per_page"`
	TotalData    int         `json:"total_data"`
}
