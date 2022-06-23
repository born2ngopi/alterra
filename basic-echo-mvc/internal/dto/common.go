package dto

import "math"

type Pagination struct {
	Page     *int `query:"page" json:"page"`
	PageSize *int `query:"page_size" json:"page_size"`
}

type SearchGetRequest struct {
	Pagination
	Search   string   `query:"search"`
	AscField []string `query:"asc_field"`
	DscField []string `query:"dsc_field"`
}

type SearchGetResponse[T any] struct {
	Datas          []T `json:"data"`
	PaginationInfo PaginationInfo
}

type PaginationInfo struct {
	*Pagination
	Count       int  `json:"count"`
	MoreRecords bool `json:"more_records"`
	TotalPage   int  `json:"total_page"`
}

type ByIDRequest struct {
	ID uint `param:"id" validate:"required"`
}

func GetLimitOffset(p *Pagination) (limit, offset int) {

	if p.PageSize != nil {
		limit = *p.PageSize
	} else {
		limit = 10
		p.PageSize = &limit
	}

	if p.Page != nil {
		offset = (*p.Page - 1) * limit
	} else {
		offset = 0
	}

	return
}

func CheckInfoPagination(p *Pagination, count int64) *PaginationInfo {
	info := PaginationInfo{
		Pagination: p,
	}
	var page int
	if p.Page != nil {
		page = *p.Page
	} else {
		page = 1
	}
	info.Page = &page
	info.Count = int(count)
	info.MoreRecords = false

	info.TotalPage = int(math.Ceil(float64(count) / float64(*p.PageSize)))

	if info.TotalPage > *p.PageSize {
		info.MoreRecords = true
	}

	return &info
}
