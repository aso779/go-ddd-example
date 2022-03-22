package infrastructure

type Page struct {
	Size   int `json:"size"`
	Number int `json:"number"`
}

func (r *Page) GetSize() int {
	return r.Size
}

func (r *Page) GetNumber() int {
	return r.Number
}

func (r *Page) GetOffset() int {
	return r.GetSize() * (r.GetNumber() - 1)
}

func (r *Page) IsEmpty() bool {
	return nil == r
}

type PageInfo struct {
	Size       int `json:"size"`
	Number     int `json:"number"`
	TotalCount int `json:"totalCount"`
}
