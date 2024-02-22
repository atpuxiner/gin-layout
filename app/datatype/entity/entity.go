package entity

// 每个功能对应一个子文件夹，便于后期扩展（ent作为前缀）

type PageInfo struct {
	Page     int
	PageSize int
}

func (r PageInfo) Limit() int {
	return r.PageSize
}

func (r PageInfo) Offset() int {
	return (r.Page - 1) * r.PageSize
}
