package dto

type Paginate struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (p *Paginate) GetLimit() int {
	if p.Limit == 0 {
		return 10
	}
	return p.Limit
}

func (p *Paginate) GetPage() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

func (p *Paginate) Offset() int {
	return (p.GetPage() - 1) * p.Limit
}
