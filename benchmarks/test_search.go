package benchmarks

import (
	"strings"

	"golang.org/x/sync/errgroup"
)

func (r *runner) searchRecords() error {
	r.write("## Search records")

	// revert any schema changes
	defer r.resetSchema(false)

	_, userToken, err := r.randomUserAuth()
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
		indexes     []string
		extraFunc   func() error
	}

	scenarios := []scenario{}

	collections := []string{
		colPosts10k,
		colPosts25k, colPosts50k, colPosts100k,
	}
	for _, col := range collections {
		scenarios = append(
			scenarios,
			scenario{"simpleA (many requests, no rules, no concurrent access)", 1000, 1, col, "?perPage=20", "", nil, nil},
			scenario{"simpleB (many requests, no rules, concurrent access)", 1000, 1000, col, "?perPage=20", "", nil, nil},
			scenario{"mixed read and write (simpleA list with 300 concurrent random " + col + " updates running in the background)", 1000, 1000, col, "?perPage=20", "", nil, func() error {
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
			scenario{"expand author", 100, 10, col, "?perPage=20&expand=author", "", nil, nil},
			scenario{"expand author (limited fields)", 100, 10, col, "?perPage=20&expand=author&fields=id,collectionId,expand.author.id", "", nil, nil},
			scenario{"expand author.permissions", 100, 10, col, "?perPage=20&expand=author.permissions", "", nil, nil},
			scenario{"expand author.permissions (limited fields)", 100, 10, col, "?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id", "", nil, nil},
			scenario{"simple auth rule", 100, 10, col, "?perPage=20", "@request.auth.id != ''", nil, nil},
			scenario{"author check (no index)", 100, 10, col, "?perPage=20", "author = @request.auth.id", nil, nil},
			scenario{"author check (with index)", 100, 10, col, "?perPage=20", "author = @request.auth.id", []string{
				"create index author_idx on " + col + " (author)",
			}, nil},
			scenario{"author.id (extra join) check (no index)", 100, 10, col, "?perPage=20", "author.id = @request.auth.id", nil, nil},
			scenario{"author.id (extra join) check (with index)", 100, 10, col, "?perPage=20", "author.id = @request.auth.id", []string{
				"create index author_idx on " + col + " (author)",
			}, nil},
			scenario{"loose large text search (no index)", 100, 10, col, "?perPage=20", "description ~ 'ipsum dolor'", nil, nil},
			scenario{"loose large text search (with index)", 100, 10, col, "?perPage=20", "description ~ 'ipsum dolor'", []string{
				"create index descriptions_idx on " + col + " (description)",
			}, nil},
			scenario{"multiple select :each (no index, match-all)", 100, 10, col, "?perPage=20", "type:each != 'c'", nil, nil},
			scenario{"multiple select :each (no index, at-least-one)", 100, 10, col, "?perPage=20", "type:each ?!= 'c'", nil, nil},
			scenario{"nested single relations lookup (no indexes)", 100, 10, col, "?perPage=20", "author.organization.name != 'test'", nil, nil},
			scenario{"nested multiple relations lookup (no indexes, match-all)", 100, 10, col, "?perPage=20", "author.permissions.active = true", nil, nil},
			scenario{"nested multiple relations lookup (no indexes, at-least-one)", 100, 10, col, "?perPage=20", "author.permissions.active ?= true", nil, nil},
			// @todo
			// the below takes too much time to complete
			// research how to optimize joins with large tables
			// scenario{"largest @collection.posts100k join on author (no indexes, match-all)", 100, 10, col, "?perPage=20", "@collection.posts100k.author != author", nil, nil},
			// scenario{"largest @collection.posts100k join on author (no indexes, at-least-one)", 100, 10, col, "?perPage=20", "@collection.posts100k.author ?!= author", nil, nil},
		)
	}

	for _, s := range scenarios {
		collectionData := map[string]any{
			"listRule": s.rule,
			"indexes":  s.indexes,
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
