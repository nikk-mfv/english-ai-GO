package entities

import "gorm.io/gorm"

var defaultPageSize uint32 = 30

type PagingRequest struct {
	Size uint32 `form:"size" json:"size" validate:"omitempty,gte=1"`
	Page uint32 `form:"page" json:"page" validate:"omitempty,gte=1"`
}

func (p PagingRequest) GormPaging(db *gorm.DB) *gorm.DB {
	return db.Limit(int(p.Size)).Offset(int(p.Size) * (int(p.Page) - 1))
}
