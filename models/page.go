package models

type Page struct {
	ShowCount     int //每页显示记录数
	TotalPage     int //总页数
	TotalResult   int //总记录数
	CurrentPage   int //当前页
	CurrentResult int //当前记录起始索引
}

func DefaultPerson() *Page {
	return &Page{ShowCount: 10}
}
func NewPage(currentPage int,showCount int) *Page{
	return &Page{CurrentPage:currentPage,ShowCount: showCount}
}

func (m *Page) GetTotalPage() *Page {
	if d := (m.TotalResult % m.ShowCount); d == 0 {
		m.TotalPage = m.TotalResult / m.ShowCount
	} else {
		m.TotalPage = (m.TotalResult / m.ShowCount) + 1
	}
	return m
}
