package models

import "math"

//CustomerPage 分页计算方式
type CustomerPage struct {
	Page int
	Per int
	Count int
}
//SetPaginator 计算分页
func (p *CustomerPage) SetPaginator() map[string] int  {
	
	offset := p.Per * (p.Page - 1)
	countPage := int(math.Ceil(float64(p.Count) / float64(p.Per)))
	mapValue := map[string]int{
		"Offset": offset,
		"CountPage": countPage,
	}
	return mapValue
}