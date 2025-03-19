package store

import (
	"net/http"
	"strconv"
)

type FeedPagination struct {
	Limit  int    `json:"limit" validate:"required,gte=1,lte=20"`
	Offset int    `json:"offset" validate:"gte=0"`
	Sort   string `json:"sort" validate:"oneof=ASC DESC"`
}

func FeedPaginationParse(r *http.Request) (*FeedPagination, error) {
	var (
		limitKey  = "limit"
		offsetKey = "offset"
		sortKey   = "sort"
	)

	var fp FeedPagination

	// get query string
	qs := r.URL.Query()

	// set limit
	limit := qs.Get(limitKey)
	if limit == "" {
		fp.Limit = 10
	} else {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
		fp.Limit = limitInt
	}

	// set offset
	offset := qs.Get(offsetKey)
	if offset == "" {
		fp.Offset = 0
	} else {
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			return nil, err
		}
		fp.Offset = offsetInt
	}

	// set sort
	sort := qs.Get(sortKey)
	if sort == "" {
		fp.Sort = "DESC"
	} else {
		fp.Sort = sort
	}

	return &fp, nil
}
