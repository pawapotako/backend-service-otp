package model

type Paging struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Sort       string `json:"sort,omitempty"`
	TotalRows  int64  `json:"totalRows"`
	TotalPages int    `json:"totalPages"`
}

func (p *PagingPayload) GetOffset() int {

	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PagingPayload) GetLimit() int {

	switch {
	case p.Paging.Limit > 100:
		p.Paging.Limit = 100
	case p.Paging.Limit <= 0:
		p.Paging.Limit = 10
	}

	return p.Paging.Limit
}

func (p *PagingPayload) GetPage() int {

	if p.Paging.Page == 0 {
		p.Paging.Page = 1
	}

	return p.Paging.Page
}

func (p *PagingPayload) GetSort() string {

	if p.Paging.Sort == "" {
		p.Paging.Sort = "Id asc"
	}

	return p.Paging.Sort
}
