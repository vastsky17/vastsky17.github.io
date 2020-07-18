package spiderhn

import (
	"github.com/jinzhu/gorm"
)

//PaginationQ gin handler query binding struct
type PaginationQ struct {
	Ok    bool        `json:"ok"`
	Size  uint        `form:"size" json:"size"`
	Page  uint        `form:"page" json:"page"`
	Data  interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list
	Total uint        `json:"total"`
}

//SearchAll optimized pagination method for gorm
func (p *PaginationQ) SearchAll(queryTx *gorm.DB) (data *PaginationQ, err error) {
	//99999 magic number for get all list without pagination
	if p.Size == 9999 || p.Size == 99999 {
		err = queryTx.Find(p.Data).Error
		p.Ok = err == nil
		return p, err
	}

	if p.Size < 1 {
		p.Size = 10
	}
	if p.Page < 1 {
		p.Page = 1
	}
	offset := p.Size * (p.Page - 1)
	err = queryTx.Count(&p.Total).Error
	if err != nil {
		return p, err
	}
	err = queryTx.Limit(p.Size).Offset(offset).Find(p.Data).Error
	p.Ok = err == nil
	return p, err
}
