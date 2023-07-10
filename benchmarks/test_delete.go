package benchmarks

import (
	"sync"
)

func (r *runner) deleteRecords() error {
	r.write("## Deleting records")

	// revert any schema changes
	defer r.resetSchema(false)

	_, userToken, err := r.randomUserAuth()
	if err != nil {
		return err
	}

	scenarios := []struct {
		comment     string
		deleteCount int
		concurrency int
		collection  string
		rule        string
	}{
		// posts10k
		{"simple (no cascade, no rule)", 100, 10, colPosts10k, ""},
		{"simple (no cascade, with rule)", 100, 10, colPosts10k, "@request.auth.id != ''"},
		// posts25k
		{"simple (no cascade, no rule)", 100, 10, colPosts25k, ""},
		{"simple (no cascade, with rule)", 100, 10, colPosts25k, "@request.auth.id != ''"},
		// posts50k
		{"simple (no cascade, no rule)", 100, 10, colPosts50k, ""},
		{"simple (no cascade, with rule)", 100, 10, colPosts50k, "@request.auth.id != ''"},
		// posts100k
		{"simple (no cascade, no rule)", 100, 10, colPosts100k, ""},
		{"simple (no cascade, with rule)", 100, 10, colPosts100k, "@request.auth.id != ''"},
		// users
		{"with cascade deleting all associated posts", 100, 10, colUsers, ""},
		// organizations
		{"with cascade deleting all users and associated posts", 100, 10, colOrganizations, ""},
	}

	for _, s := range scenarios {
		collectionData := map[string]any{
			"deleteRule": s.rule,
		}
		if err := r.updateCollection(s.collection, collectionData); err != nil {
			return err
		}

		idsToDelete, err := r.randomRecordIds(s.collection, s.deleteCount)
		if err != nil {
			return err
		}

		idsMap := sync.Map{}
		for i, id := range idsToDelete {
			idsMap.Store(i, id)
		}

		r.cooldown()

		r.write("#### deleting %d %s - %s [conc:%d, rule:`%q`]",
			len(idsToDelete),
			s.collection,
			s.comment,
			s.concurrency,
			s.rule,
		)

		b, err := bench(func(i int) error {
			raw, _ := idsMap.Load(i)
			id := raw.(string)

			req := Request{
				Url:    r.baseUrl + "/api/collections/" + s.collection + "/records/" + id,
				Method: "DELETE",
				Headers: map[string]string{
					"Authorization": userToken,
				},
			}

			return req.Send(nil)
		}, len(idsToDelete), s.concurrency)
		if err != nil {
			return err
		}

		r.write(b.String())
	}

	r.write("")

	return nil
}
