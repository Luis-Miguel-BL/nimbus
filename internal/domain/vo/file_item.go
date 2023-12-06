package vo

type FileItem struct {
	index int
	data  map[string]any
}

func (v FileItem) Index() int {
	return v.index
}
