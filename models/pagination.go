package models

import "math"

//CustomerPage 分页计算方式
type CustomerPage struct {
	Per   int64
	Count int64
}

//SetPaginator 计算分页
func SetPaginator(count int64, per int64) int64 {
	countPage := int64(math.Ceil(float64(count) / float64(per)))
	return countPage
}
