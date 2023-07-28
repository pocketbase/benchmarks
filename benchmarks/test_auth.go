package benchmarks

import "strings"

func (r *runner) authWithPassword() error {
	r.write("## User auth with password (expected to be slow due to passwordHash verification)")

	scenarios := []struct {
		comment     string
		iterations  int
		concurrency int
	}{
		{"high concurrency", 250, 250},
		{"small concurrency", 250, 10},
	}

	for _, s := range scenarios {
		r.cooldown()

		r.write("#### %s auth with email/pass - %s [reqs:%d, conc:%d]",
			colUsers,
			s.comment,
			s.iterations,
			s.concurrency,
		)

		b, err := bench(func(i int) error {
			req := Request{
				Url:    r.baseUrl + "/api/collections/" + colUsers + "/auth-with-password",
				Method: "POST",
				Body:   strings.NewReader(`{"identity":"` + colUsers + `0@example.com","password":"1234567890"}`),
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

func (r *runner) authRefresh() error {
	r.write("## User auth refresh")

	_, userToken, err := r.randomUserAuth()
	if err != nil {
		return err
	}

	scenarios := []struct {
		comment     string
		iterations  int
		concurrency int
	}{
		{"auth refresh (high concurrency)", 1000, 1000},
		{"auth refresh (medium concurrency)", 1000, 100},
	}

	for _, s := range scenarios {
		r.cooldown()

		r.write("#### %s - %s [reqs:%d, conc:%d]",
			colUsers,
			s.comment,
			s.iterations,
			s.concurrency,
		)

		b, err := bench(func(i int) error {
			req := Request{
				Url:    r.baseUrl + "/api/collections/" + colUsers + "/auth-refresh",
				Method: "POST",
				Headers: map[string]string{
					"Authorization": userToken,
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
