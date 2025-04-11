package entities

type PagingRequest struct {
	Page uint32 `form:"page" json:"page" validate:"omitempty,gte=1,lte=30"`
	Size uint32 `form:"Size" json:"SizeSize" validate:"omitempty,gte=1"`
}

func (p *PagingRequest) SetDefautls() {
	//Set page value default = 1
	if p.Page < 1 {
		p.Page = 1
	}

	//Set pageSize value default = 20
	if p.Size < 1 || p.Size > 100 {
		p.Size = 20
	}
}

type PagingResponse[T any] struct {
	Data       []T   `json:"data"`
	TotalItems int64 `json:"totalItems"`
	Size       int64 `json:"size"`
}
