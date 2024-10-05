package benchmarks

import (
	"math/rand"
	"strconv"
	"strings"
)

// pickRandom returns a random element from the provided list.
func pickRandom(list []string) string {
	return list[rand.Intn(len(list))]
}

func (r *runner) createOrganizations() error {
	r.write("## Creating organizations (100)")

	// revert any schema changes
	defer r.resetSchema(false)

	scenarios := []struct {
		iterations  int
		concurrency int
		collection  string
		rule        string
	}{
		{50, 10, colOrganizations, ""},
		{50, 10, colOrganizations, "@request.body.name != ''"},
	}

	for _, s := range scenarios {
		collectionData := map[string]any{
			"createRule": s.rule,
		}
		if err := r.updateCollection(s.collection, collectionData); err != nil {
			return err
		}

		total, err := r.app.CountRecords(s.collection)
		if err != nil {
			return err
		}

		r.cooldown()

		r.write(
			"#### Creating %d %s [reqs:%d, conc:%d, rule:`%q`]",
			s.iterations,
			s.collection,
			s.iterations,
			s.concurrency,
			s.rule,
		)

		b, err := bench(func(i int) error {
			name := s.collection + strconv.Itoa(i+int(total))

			req := Request{
				Url:    r.baseUrl + "/api/collections/" + s.collection + "/records",
				Method: "POST",
				Body:   strings.NewReader(`{"name":"` + name + `"}`),
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

func (r *runner) createPermissions() error {
	r.write("## Creating permissions (50)")

	// revert any schema changes
	defer r.resetSchema(false)

	scenarios := []struct {
		iterations  int
		concurrency int
		collection  string
		rule        string
	}{
		{25, 5, colPermissions, ""},
		{25, 5, colPermissions, "@request.body.name != ''"},
	}

	for _, s := range scenarios {
		collectionData := map[string]any{
			"createRule": s.rule,
		}
		if err := r.updateCollection(s.collection, collectionData); err != nil {
			return err
		}

		total, err := r.app.CountRecords(s.collection)
		if err != nil {
			return err
		}

		r.cooldown()

		r.write(
			"#### Creating %d %s [reqs:%d, conc:%d, rule:`%q`]",
			s.iterations,
			s.collection,
			s.iterations,
			s.concurrency,
			s.rule,
		)

		b, err := bench(func(i int) error {
			name := s.collection + strconv.Itoa(i+int(total))

			active := "false"
			if i%2 == 0 {
				active = "true"
			}

			req := Request{
				Url:    r.baseUrl + "/api/collections/" + s.collection + "/records",
				Method: "POST",
				Body:   strings.NewReader(`{"name":"` + name + `","active":` + active + `}`),
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

func (r *runner) createUsers() error {
	r.write("## Creating users (500 - expected to be slow due to passwordHash generation)")

	// revert any schema changes
	defer r.resetSchema(false)

	perms, err := r.randomRecordIds(colPermissions, 999)
	if err != nil {
		return err
	}

	orgs, err := r.randomRecordIds(colOrganizations, 999)
	if err != nil {
		return err
	}

	scenarios := []struct {
		iterations  int
		concurrency int
		collection  string
		rule        string
	}{
		{250, 50, colUsers, ""},
		{250, 50, colUsers, "@request.body.email != '' && @request.body.permissions:length > 0"},
	}

	for _, s := range scenarios {
		collectionData := map[string]any{
			"createRule": s.rule,
		}
		if err := r.updateCollection(s.collection, collectionData); err != nil {
			return err
		}

		total, err := r.app.CountRecords(s.collection)
		if err != nil {
			return err
		}

		r.cooldown()

		r.write(
			"#### Creating %d %s [reqs:%d, conc:%d, rule:`%q`]",
			s.iterations,
			s.collection,
			s.iterations,
			s.concurrency,
			s.rule,
		)

		b, err := bench(func(i int) error {
			name := s.collection + strconv.Itoa(i+int(total))

			req := Request{
				Url:    r.baseUrl + "/api/collections/" + s.collection + "/records",
				Method: "POST",
				Body: strings.NewReader(`{
					"email":"` + name + `@example.com",
					"username":"` + name + `",
					"name":"` + name + `",
					"organization":"` + pickRandom(orgs) + `",
					"permissions":["` + pickRandom(perms) + `", "` + pickRandom(perms) + `", "` + pickRandom(perms) + `"],
					"password":"1234567890",
					"passwordConfirm":"1234567890"
				}`),
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

func (r *runner) createPosts() error {
	r.write("## Creating posts (10k, 25k, 50k, 100k)")

	// revert any schema changes
	defer r.resetSchema(false)

	_, userToken, err := r.randomUserAuth()
	if err != nil {
		return err
	}

	types := []string{"a", "b", "c", "d"}

	users, err := r.randomRecordIds(colUsers, 999)
	if err != nil {
		return err
	}

	scenarios := []struct {
		iterations  int
		concurrency int
		collection  string
		rule        string
		token       string
	}{
		{5000, 500, colPosts10k, "", ""},
		{5000, 500, colPosts10k, "@request.auth.id != '' && @request.body.public:isset = true", userToken},
		{12500, 500, colPosts25k, "", ""},
		{12500, 500, colPosts25k, "@request.auth.id != '' && @request.body.public:isset = true", userToken},
		{25000, 500, colPosts50k, "", ""},
		{25000, 500, colPosts50k, "@request.auth.id != '' && @request.body.public:isset = true", userToken},
		{50000, 500, colPosts100k, "", ""},
		{50000, 500, colPosts100k, "@request.auth.id != '' && @request.body.public:isset = true", userToken},
	}

	for _, s := range scenarios {
		collectionData := map[string]any{
			"createRule": s.rule,
		}
		if err := r.updateCollection(s.collection, collectionData); err != nil {
			return err
		}

		total, err := r.app.CountRecords(s.collection)
		if err != nil {
			return err
		}

		r.cooldown()

		r.write(
			"#### Creating %d %s [reqs:%d, conc:%d, rule:`%q`]",
			s.iterations,
			s.collection,
			s.iterations,
			s.concurrency,
			s.rule,
		)
		b, err := bench(func(i int) error {
			name := s.collection + strconv.Itoa(i+int(total))

			public := "true"
			if i%2 == 0 {
				public = "false"
			}

			req := Request{
				Url:    r.baseUrl + "/api/collections/" + s.collection + "/records",
				Method: "POST",
				Body: strings.NewReader(`{
					"title":"` + name + `",
					"description":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec sit amet sodales nisl, quis pretium nunc. Suspendisse vel auctor velit, sed luctus lectus. Phasellus rhoncus imperdiet feugiat. Duis et laoreet felis, ut facilisis enim. Quisque aliquet aliquam magna eget eleifend. Duis sed tellus nibh. Nunc ac lacus auctor, scelerisque magna congue, euismod purus. Fusce sollicitudin pharetra egestas. Quisque pulvinar augue nec aliquam placerat. Suspendisse dapibus ornare sodales.",
					"public":` + public + `,
					"type":["` + pickRandom(types) + `", "` + pickRandom(types) + `"],
					"author":"` + pickRandom(users) + `"
				}`),
				Headers: map[string]string{},
			}
			if s.token != "" {
				req.Headers["Authorization"] = s.token
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
