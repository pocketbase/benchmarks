package benchmarks

import (
	"strings"

	"golang.org/x/sync/errgroup"
)

func (r *runner) listRecords() error {
	r.write("## List records")

	// revert any schema changes
	defer r.resetSchema(false)

	user, userToken, err := r.randomUserAuth()
	if err != nil {
		return err
	}

	type scenario struct {
		comment     string
		iterations  int
		concurrency int
		collection  string
		query       string
		rule        string
		indexes     []string // set it to emtpy []string{} to reset the indexes
		extraFunc   func() error
	}

	scenarios := []scenario{
		{"getOne for auth refresh comparison (medium concurrency)", 1000, 100, colUsers, "/" + user.Id, "", nil, nil},
		{"getOne for auth refresh comparison (high concurrency)", 1000, 1000, colUsers, "/" + user.Id, "", nil, nil},
	}

	collections := []string{
		colPosts10k,
		colPosts25k,
		colPosts50k,
		colPosts100k,
	}
	for _, col := range collections {
		scenarios = append(
			scenarios,
			scenario{"simpleA (many requests, no rules, no concurrency)", 1000, 1, col, "?perPage=20", "", []string{}, nil},
			scenario{"simpleB (many requests, no rules, high concurrency)", 1000, 1000, col, "?perPage=20", "", []string{}, nil},
			scenario{"simpleC (many requests, no rules, high concurrency, skipTotal)", 1000, 1000, col, "?perPage=20&skipTotal=1", "", []string{}, nil},
			scenario{"mixed read and write (simpleA list with additional 300 concurrent random " + col + " updates running in the background)", 1000, 1000, col, "?perPage=20", "", []string{}, func() error {
				g := errgroup.Group{}
				g.SetLimit(-1)

				ids, err := r.randomRecordIds(col, 300)
				if err != nil {
					return err
				}

				for _, id := range ids {
					id := id
					g.Go(func() error {
						req := Request{
							Url:    r.baseUrl + "/api/collections/" + col + "/records/" + id,
							Method: "PATCH",
							Body:   strings.NewReader(`{"title": "update` + id + `"}`),
							Headers: map[string]string{
								"Authorization": userToken,
							},
						}

						return req.Send(nil)
					})
				}

				return g.Wait()
			}},
			scenario{"expand author", 100, 10, col, "?perPage=20&expand=author", "", []string{}, nil},
			scenario{"expand author (limited fields)", 100, 10, col, "?perPage=20&expand=author&fields=id,collectionId,expand.author.id", "", []string{}, nil},
			scenario{"expand author.permissions", 100, 10, col, "?perPage=20&expand=author.permissions", "", []string{}, nil},
			scenario{"expand author.permissions (limited fields)", 100, 10, col, "?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id", "", []string{}, nil},
			scenario{"simple auth rule", 100, 10, col, "?perPage=20", "@request.auth.id != ''", []string{}, nil},
			scenario{"author check (no index)", 100, 10, col, "?perPage=20", "author = @request.auth.id", []string{}, nil},
			scenario{"author check (with index)", 100, 10, col, "?perPage=20", "author = @request.auth.id", []string{
				"create index `idx_author_" + col + "` on " + col + " (author)",
			}, nil},
			scenario{"author check (with index and skipTotal)", 100, 10, col, "?perPage=20&skipTotal=1", "author = @request.auth.id", []string{
				"create index `idx_author_" + col + "` on " + col + " (author)",
			}, nil},
			scenario{"author.id (extra join) check (no index)", 100, 10, col, "?perPage=20", "author.id = @request.auth.id", []string{}, nil},
			scenario{"author.id (extra join) check (with index)", 100, 10, col, "?perPage=20", "author.id = @request.auth.id", []string{
				"create index `idx_author_" + col + "` on " + col + " (author)",
			}, nil},
			scenario{"author.id (extra join) check (with index and skipTotal)", 100, 10, col, "?perPage=20&skipTotal=1", "author.id = @request.auth.id", []string{
				"create index `idx_author_" + col + "` on " + col + " (author)",
			}, nil},
			scenario{"loose large text search (no index)", 100, 10, col, "?perPage=20", "description ~ 'ipsum dolor'", []string{}, nil},
			scenario{"loose large text search (with index)", 100, 10, col, "?perPage=20", "description ~ 'ipsum dolor'", []string{
				"create index `idx_descriptions_" + col + "` on " + col + " (description)",
			}, nil},
			scenario{"loose large text search (with index and skipTotal)", 100, 10, col, "?perPage=20&skipTotal=1", "description ~ 'ipsum dolor'", []string{
				"create index `idx_descriptions_" + col + "` on " + col + " (description)",
			}, nil},
			scenario{"multiple select :each (no index, match-all)", 100, 10, col, "?perPage=20", "type:each != 'c'", []string{}, nil},
			scenario{"multiple select :each (no index, match-all, skipTotal)", 100, 10, col, "?perPage=20&skipTotal=1", "type:each != 'c'", []string{}, nil},
			scenario{"multiple select :each (no index, at-least-one)", 100, 10, col, "?perPage=20", "type:each ?!= 'c'", []string{}, nil},
			scenario{"multiple select :each (no index, at-least-one, skipTotal)", 100, 10, col, "?perPage=20&skipTotal=1", "type:each ?!= 'c'", []string{}, nil},
			scenario{"nested single relations lookup (no indexes)", 100, 10, col, "?perPage=20", "author.organization.name != 'test'", []string{}, nil},
			scenario{"nested single relations lookup (no indexes, skipTotal)", 100, 10, col, "?perPage=20&skipTotal=1", "author.organization.name != 'test'", []string{}, nil},
			scenario{"nested multiple relations lookup (no indexes, match-all)", 100, 10, col, "?perPage=20", "author.permissions.active = true", []string{}, nil},
			scenario{"nested multiple relations lookup (no indexes, match-all, skipTotal)", 100, 10, col, "?perPage=20&skipTotal=1", "author.permissions.active = true", []string{}, nil},
			scenario{"nested multiple relations lookup (no indexes, at-least-one)", 100, 10, col, "?perPage=20", "author.permissions.active ?= true", []string{}, nil},
			scenario{"nested multiple relations lookup (no indexes, at-least-one, skipTotal)", 100, 10, col, "?perPage=20&skipTotal=1", "author.permissions.active ?= true", []string{}, nil},
		)
	}

	for _, s := range scenarios {
		if err := r.resetSchema(false); err != nil {
			return err
		}

		collectionData := map[string]any{
			"listRule": s.rule,
		}
		if s.indexes != nil {
			collectionData["indexes"] = s.indexes
		}

		if err := r.updateCollection(s.collection, collectionData); err != nil {
			return err
		}

		r.cooldown()

		r.write("#### %s - %s [reqs:%d, conc:%d, rule:`%q`, query:`%s`]",
			s.collection,
			s.comment,
			s.iterations,
			s.concurrency,
			s.rule,
			s.query,
		)

		g := errgroup.Group{}
		g.SetLimit(2)

		if s.extraFunc != nil {
			g.Go(func() error {
				return s.extraFunc()
			})
		}

		var result *BenchResult

		g.Go(func() (err error) {
			result, err = bench(func(i int) error {
				req := Request{
					Url:    r.baseUrl + "/api/collections/" + s.collection + "/records" + s.query,
					Method: "GET",
					Headers: map[string]string{
						"Authorization": userToken,
					},
				}

				return req.Send(nil)
			}, s.iterations, s.concurrency)

			return err
		})

		if err := g.Wait(); err != nil {
			return err
		}

		r.write(result.String())
	}

	r.write("")

	return nil
}
