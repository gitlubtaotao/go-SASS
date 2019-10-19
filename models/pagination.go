package models

import "math"

//CustomerPage 分页计算方式
type CustomerPage struct {
	Per   int64
	Count int64
}

//SetPaginator 计算分页
func SetPaginator(count int64) int64 {
	per := UserPerPage()
	countPage := int64(math.Ceil(float64(count) / float64(per)))
	return countPage
}

//对应的分页次数，计算limit的值
func UserPerPage() int64 {
	return 10
}
//计算offset对应的值
func GetOffsetPage(page int64) int64  {
	result :=UserPerPage() * (page - 1)
	return result
}
