package pagination

type PageOptions struct {
	Page    int    `json:"page"`
	Take    int    `json:"take"`
	OrderBy string `json:"order_by"`
}

type PageMetadata struct {
	Page            int  `json:"page"`
	Take            int  `json:"take"`
	ItemCount       int  `json:"item_count"`
	PageCount       int  `json:"page_count"`
	HasPreviousPage bool `json:"has_previous_page"`
	HasNextPage     bool `json:"has_next_page"`
}

type PageDto struct {
	Metadata PageMetadata  `json:"metadata"`
	Data     []interface{} `json:"data"`
}

func (p PageOptions) Skip() int {
	return (p.Page - 1) * p.Take
}

func NewPageMetadata(pageOptions PageOptions, itemCount int) PageMetadata {
	pageCount := itemCount / pageOptions.Take
	if itemCount%pageOptions.Take != 0 {
		pageCount++
	}

	return PageMetadata{
		Page:            pageOptions.Page,
		Take:            pageOptions.Take,
		ItemCount:       itemCount,
		PageCount:       pageCount,
		HasPreviousPage: pageOptions.Page > 1,
		HasNextPage:     pageOptions.Page < pageCount,
	}
}
func NewPageDto(data []interface{}, metadata PageMetadata) PageDto {
	return PageDto{
		Metadata: metadata,
		Data:     data,
	}
}
