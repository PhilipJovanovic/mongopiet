package opts

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SkipLimitI returns skip and limit for pagination with raw integers
//
// page: page number
//
// limit: limit per page (default 16)
func SkipLimitI(page int, limit int) (int64, int64, error) {
	rLimit := 16

	if limit > 0 {
		rLimit = limit
	}

	if page < 1 {
		return 0, int64(rLimit), nil
	}

	return int64((page - 1) * rLimit), int64(rLimit), nil
}

// used like SkipLimit but with string input (f.e. from URL query)
//
// page 0 == page 1
//
// page: page number (returns error if invalid)
//
// limit: limit per page (default 16) (default if invalid)
//
// Example:
//
//	page := r.URL.Query().Get("page")
//	limit := r.URL.Query().Get("limit")
//
//	skip, mongolimit, err := SkipLimit(page, limit)
//	if err != nil {
//		return err
//	}
func SkipLimit(page string, limit string) (int64, int64, error) {
	defaultLimit := 16

	if li, err := strconv.Atoi(limit); err == nil && li > 0 {
		defaultLimit = li
	}

	if page == "" || page == "0" {
		return 0, int64(defaultLimit), nil
	}

	pa, err := strconv.Atoi(page)
	if err != nil {
		return 0, 0, err
	}

	return SkipLimitI(pa, defaultLimit)
}

type Request interface {
	FormValue(key string) string
}

// QueryParam can be used to get pagination and sorting options from a request.
//
// r: Request interface
//
// sortField: field to sort by
//
// Example:
//
//	opt := opts.QueryParam(r, "createdAt")
//	if opt != nil {
//		// do something with opt
//	}
func QueryParam(r Request, sortField string) *options.FindOptions {
	opt := options.Find()

	page := r.FormValue("page")
	limit := r.FormValue("limit")

	skip, mongolimit, err := SkipLimit(page, limit)
	if err != nil {
		return nil
	}

	opt.SetSkip(skip)
	opt.SetLimit(mongolimit)

	// add sorting based on sortfield
	if sort := r.FormValue("sort"); sort != "" {
		if sort == "asc" {
			opt.SetSort(bson.M{sortField: 1})
		} else {
			opt.SetSort(bson.M{sortField: -1})
		}
	}

	return opt
}
