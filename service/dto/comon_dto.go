package dto

// CommonIDDTO 通用ID DTO
type CommonIDDTO struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

// 分页TDO
type Paginate struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"Limit,omitempty" form:"Limit"`
}

func (p *Paginate) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Paginate) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	return p.Limit
}
