package benchmarks

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type BenchResult struct {
	Best      time.Duration
	Worst     time.Duration
	Completed time.Duration
	Errors    []error
}

func (r BenchResult) String() string {
	return fmt.Sprintf(
		"```\n┌─ Best:      %s\n├─ Worst:     %s\n├─ Completed: %s\n└─ Errors:    %d\n```",
		r.Best,
		r.Worst,
		r.Completed,
		len(r.Errors),
	)
}

// A negative concurrency indicates no limit
// (aka. a go routine will be fired for each iteration).
func bench(action func(i int) error, iterations int, concurrency int) (*BenchResult, error) {
	if iterations < 1 {
		return nil, errors.New("iterations must be >= 1")
	}

	totalStart := time.Now()

	var errors []error

	times := make([]time.Duration, iterations)

	g := errgroup.Group{}
	g.SetLimit(concurrency)

	for i := 0; i < iterations; i++ {
		i := i
		g.Go(func() error {
			start := time.Now()

			if err := action(i); err != nil {
				errors = append(errors, err)
			}

			times[i] = time.Since(start)

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	result := &BenchResult{
		Best:      times[0],
		Worst:     times[0],
		Completed: time.Since(totalStart),
		Errors:    errors,
	}

	// find the best and worst time
	for _, t := range times {
		if t < result.Best {
			result.Best = t
		}
		if t > result.Worst {
			result.Worst = t
		}
	}

	return result, nil
}
