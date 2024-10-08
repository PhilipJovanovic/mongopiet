package opts

import (
	"testing"
)

func TestMain(t *testing.T) {
	testSkipLimitS(t)
	testRequest(t)
}

func testSkipLimitS(t *testing.T) {
	skip, limit, err := SkipLimit("", "")
	if err != nil {
		t.Fatal(err)
	}

	// default limit
	if skip != 0 || limit != 16 {
		t.Fatal(skip, limit)
	}

	_, _, err = SkipLimit("test", "")
	// invalid page
	if err == nil {
		t.Fatal(err)
	}

	_, _, err = SkipLimit("1", "invalid")
	// invalid limit == default limit
	if err != nil {
		t.Fatal(err)
	}

	skip, limit, err = SkipLimit("1", "10")
	if err != nil {
		t.Fatal(err)
	}

	// first page with custom limit (10)
	if skip != 0 || limit != 10 {
		t.Fatal(skip, limit)
	}

	// second page with custom limit (10)
	skip, limit, err = SkipLimit("2", "10")
	if err != nil {
		t.Fatal(err)
	}

	if skip != 10 || limit != 10 {
		t.Fatal(skip, limit)
	}
}

type Req struct {
	Form map[string]string
}

func (r *Req) FormValue(s string) string {
	return r.Form[s]
}

func testRequest(t *testing.T) {
	req := &Req{
		Form: map[string]string{
			"page":  "2",
			"limit": "10",
		},
	}

	skip, limit, err := SkipLimit(req.FormValue("page"), req.FormValue("limit"))
	if err != nil {
		t.Fatal(err)
	}

	if skip != 10 || limit != 10 {
		t.Fatal(skip, limit)
	}

	req.Form["page"] = "1"
	req.Form["limit"] = "20"

	skip, limit, err = SkipLimit(req.FormValue("page"), req.FormValue("limit"))
	if err != nil {
		t.Fatal(err)
	}

	if skip != 0 || limit != 20 {
		t.Fatal(skip, limit)
	}

	empty := &Req{
		Form: map[string]string{},
	}

	skip, limit, err = SkipLimit(empty.FormValue("page"), empty.FormValue("limit"))
	if err != nil {
		t.Fatal(err)
	}

	if skip != 0 || limit != 16 {
		t.Fatal(skip, limit)
	}
}
