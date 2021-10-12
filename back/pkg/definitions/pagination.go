package definitions

import "time"

// Pagination - paging settings
type Pagination struct {
	PagingKey time.Time `json:"pagingKey,omitempty" query:"pagingKey"` // paging key
	SortOrder string    `json:"sortOrder,omitempty" query:"sortOrder"` // sort key / default desc
	Limit     int       `json:"limit,omitempty"     query:"limit"`     // number of acquisitions
}
