package helper

type Metadata struct {
	Page         int `json:"page"`
	Limit        int `json:"limit"`
	TotalPerPage int `json:"total_per_page"`
	TotalData    int `json:"total_data"`
}

func NewMetadata(page, limit, totalPerPage, totalData int) Metadata {
	return Metadata{
		Page:         page,
		Limit:        limit,
		TotalPerPage: totalPerPage,
		TotalData:    totalData,
	}
}
