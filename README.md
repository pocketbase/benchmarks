PocketBase benchmarks
======================================================================

This is a test PocketBase application with various benchmarks to serve
as a general idea of what you could expect from a basic PocketBase instance
without extra performance tuning or deployment on expensive server.

- [Results](#results)
- [Takeaways and things we'll have to improve](#takeaways-and-things-well-have-to-improve)
- [About the benchmarks](#about-the-benchmarks)


## Results

- **Hetzner CAX11 (2vCPU ARM64, 4GB RAM, €3.79)**
    - [Pure Go results (default)](results/hetzner_cax11.md)
    - [CGO results](results/hetzner_cax11_cgo.md)

- **Hetzner CAX41 (16vCPU ARM64, 32GB RAM, €24.49)**
    - [Pure Go results (default)](results/hetzner_cax41.md)
    - [CGO results](results/hetzner_cax41_cgo.md)

- **Hop.io (1vCPU, 512MB RAM, _free tier_)**
    - [Pure Go results (default)](results/hop_free_tier.md)
    _(fails the mixed `posts50k` tests most likely a result of some resources abuse protection)_

- **Fly.io (1vCPU, 256MB RAM, _free tier_)**
    - [Pure Go results (default)](results/fly_free_tier.md)
    _(it requires swap to be configured to prevent OOM errors and restarts; fails the heavier `posts100k` tests most likely a result of some resources abuse protection)_

_Keep in mind that the tests are performed on the same machine where the PocketBase instance is deployed so the app performance can be slightly affected by the benchmarks execution itself (most hosts providers have several protections in place and at the moment I don't have the time to create proper setup to run the tests from more than one IP._

_There are several optimizations that can be done and the benchmarks will change in the future so that the tests can run as part of the development workflow to track regressions, but for now improving the overall PocketBase dev experience remains with higher priority._


#### Takeaways and things we'll have to improve

- The Go and JS (_will be available with v0.17.0+_) custom routes and app hooks perform almost the same for low and medium concurrency (this is because by default we have 100 prewarmed `goja.Runtime`).

- At the moment there is no much difference in terms of query execution between the lower and higher spec Hetzner VMs (_probably because most of the operations are I/O bound_).

- We could experiment with different SQLite `PRAGMA` options to see if there will be any significant difference since we currently use mostly the defaults (_eg. increasing `page_size`, `cache_size`, heap limits, etc._).

- The default data DB connections limit (max:120, idle:20) could be changed to be dynamic based on the running hardware to improve the overall performance and reduce the memory usage.

- With higher concurrency individual query performance degrades (_probably because the runtime has to do more work, there is context switching involved, locks, etc._) but still the overall requests completion is better for most of the times.

- Basic API rules don't have significant impact on the performance. The performance starts to degrade when the rule condition requires a join with a large table (especially in case of a nested relation fields lookup). Joins over large datasets are slow even for simple queries like:

    ```sql
    SELECT COUNT(DISTINCT posts10k.rowid)
    FROM posts10k
    LEFT JOIN posts100k ON posts100k.author = posts10k.author
    ```
    Creating indexes for the relation fields can help speeding up the queries but further optimization techniques will have to be researched.

- To prevent unnecessary `JOIN` statements we can implement internally a special condition that will treat single `relation` field statements like `rel.id = @request.auth.id` the same as `rel = @request.auth.id`.

- The existing json column normalizations (eg. `CASE WHEN json_valid ...`) have some impact on the performance (~10%) and can be further optimized by removing them entirely and ensuring that the field values are always stored in the correct format before create/update persistence (either on app level or triggers).

- The CGO driver for some queries in _large datasets_ can be ~1.5x-4x times faster.
  Building with `CGO_ENABLED=1` is the easiest way to boost the query performance without changing your db structure or hardware but keep in mind that it complicates the cross-compilation.


## About the benchmarks

The application uses the `develop` branch of PocketBase.

The test database is ~180MB and has the following collections structure:

![db_erd](https://i.imgur.com/gYC8Qci.png)

In order to emulate real usage scenarios as close as possible, the tests are grouped in several categories:

- **create** - Tests the write (_insert_) HTTP API performance. It is also responsible for populating the test database.

- **search** - Tests the list/search (aka. `getList()`) HTTP API performance.
    It also contains scenarios with mixed list and update tests to observe how mixed read/write operations affects each other.

- **custom** - Tests for custom Go and JS code (routes, middlewares, hooks, etc.).

- **delete** - Test the delete HTTP API performance (including cascade delete).


## Run the benchmarks

To run the benchmarks locally or on your server, you can:

0. _[Install Go 1.18+](https://go.dev/doc/install) (if you haven't already)_
1. Clone/download the repo
2. Run `GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build` (_https://go.dev/doc/install/source#environment_)
3. Start the created executable by running `./app serve`.
4. Navigate to `http:localhost:8090/benchmarks`

_By default all tests are executed in the previously mentioned test categories order, but you can also
specify which exact tests to run using the `/benchmarks?run=custom,delete` query parameter._

The above will start the benchmarks in a new goroutine and once completed it will print its result to the terminal and in the `benchmarks` collection
(_note that some of the tests are slow and it may take some time to complete; we write the test result as a collection record to workaround various host providers DDoS protections and restrictions like persistent connections limit, read/write timeouts, etc._).
