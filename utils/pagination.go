package utils

import "fmt"

type PaginationMeta struct {
	Total           int64   `json:"total"`
	CurrentPage     int     `json:"current_page"`
	PerPage         int     `json:"per_page"`
	TotalPages      int64   `json:"total_pages"`
	First           *int    `json:"first,omitempty"`
	Previous        *int    `json:"previous,omitempty"`
	Next            *int    `json:"next,omitempty"`
	Last            *int    `json:"last,omitempty"`
	HasMore         bool    `json:"has_more"`
	HasPrevious     bool    `json:"has_previous"`
	IsEmpty         bool    `json:"is_empty"`
	FirstPageURL    string  `json:"first_page_url"`
	NextPageURL     *string `json:"next_page_url,omitempty"`
	LastPageURL     string  `json:"last_page_url"`
	PreviousPageURL *string `json:"previous_page_url,omitempty"`
}

func GeneratePaginationMeta(total int64, page int, pageSize int) PaginationMeta {
	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)

	intPtr := func(i int) *int { return &i }
	strPtr := func(s string) *string { return &s }

	meta := PaginationMeta{
		Total:        total,
		CurrentPage:  page,
		PerPage:      pageSize,
		TotalPages:   totalPages,
		HasMore:      page < int(totalPages),
		HasPrevious:  page > 1,
		IsEmpty:      total == 0,
		FirstPageURL: fmt.Sprintf("?page=1&page_size=%d", pageSize),
		LastPageURL:  fmt.Sprintf("?page=%d&page_size=%d", totalPages, pageSize),
	}

	if page > 1 {
		meta.First = intPtr(1)
		meta.Previous = intPtr(page - 1)
		meta.PreviousPageURL = strPtr(fmt.Sprintf("?page=%d&page_size=%d", page-1, pageSize))
	}

	if page < int(totalPages) {
		meta.Next = intPtr(page + 1)
		meta.Last = intPtr(int(totalPages))
		meta.NextPageURL = strPtr(fmt.Sprintf("?page=%d&page_size=%d", page+1, pageSize))
	}

	return meta
}
