package benchmarks

import (
	"strings"
	"sync"
)

func (r *runner) customRoute() error {
	r.write("## Go vs JS route execution")

	_, adminToken, err := r.randomAdmin()
	if err != nil {
		return err
	}

	scenarios := []struct {
		comment     string
		iterations  int
		concurrency int
		path        string
	}{
		{"JS route (high concurrency)", 500, 500, "/js"},
		{"Go route (high concurrency)", 500, 500, "/go"},
		{"JS route (medium concurrency)", 500, 50, "/js"},
		{"Go route (medium concurrency)", 500, 50, "/go"},
		{"JS route (no concurrency)", 500, 1, "/js"},
		{"Go route (no concurrency)", 500, 1, "/go"},
	}

	for _, s := range scenarios {
		r.cooldown()

		r.write("#### %s [reqs:%d, conc:%d]",
			s.comment,
			s.iterations,
			s.concurrency,
		)

		b, err := bench(func(i int) error {
			req := Request{
				Url:    r.baseUrl + s.path,
				Method: "GET",
				Headers: map[string]string{
					"Authorization": adminToken,
				},
			}

			return req.Send(nil)
		}, s.iterations, s.concurrency)
		if err != nil {
			return err
		}

		r.write(b.String())
	}

	r.write("")

	return nil
}

func (r *runner) customHook() error {
	r.write("## Go vs JS hooks execution")

	_, adminToken, err := r.randomAdmin()
	if err != nil {
		return err
	}

	scenarios := []struct {
		comment     string
		updateCount int
		concurrency int
		collection  string
		tempName    string
	}{
		{"JS OnRecordBeforeUpdateRequest hook handler", 100, 10, colPosts10k, "js"},
		{"Go OnRecordBeforeUpdateRequest hook handler", 100, 10, colPosts10k, "go"},
	}

	for _, s := range scenarios {
		// temp rename the collection to ensure that the hook is triggered only for the specific case
		if err := r.updateCollection(s.collection, map[string]any{"name": s.tempName}); err != nil {
			return err
		}

		idsToUpdate, err := r.randomRecordIds(s.tempName, s.updateCount)
		if err != nil {
			return err
		}

		idsMap := sync.Map{}
		for i, id := range idsToUpdate {
			idsMap.Store(i, id)
		}

		r.cooldown()

		r.write("#### %s - [reqs:%d, conc:%d]",
			s.comment,
			len(idsToUpdate),
			s.concurrency,
		)
		b, benchErr := bench(func(i int) error {
			raw, _ := idsMap.Load(i)
			id := raw.(string)

			req := Request{
				Url:    r.baseUrl + "/api/collections/" + s.tempName + "/records/" + id,
				Method: "PATCH",
				Body:   strings.NewReader(`{"title": "hook_update"}`),
				Headers: map[string]string{
					"Authorization": adminToken,
				},
			}

			return req.Send(nil)
		}, len(idsToUpdate), s.concurrency)

		// revert temp schema changes
		if err := r.resetSchema(false); err != nil {
			return err
		}

		if benchErr != nil {
			return benchErr
		}

		r.write(b.String())
	}

	r.write("")

	return nil
}
