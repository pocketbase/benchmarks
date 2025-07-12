_PocketBase v0.28.4_

_Railway Hobby Tier Configuration_

_CPU: 8 cores_
_Memory: 8GiBs_

_Actual usage:_

_CPU: Max 7.5 cores_

_Memory: Max 1.6GiBs_

_Dockerfile:_

```Dockerfile
FROM golang:1.23-alpine

RUN apk add --no-cache build-base


WORKDIR /pb

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=1 go build -o /pb/pocketbase

EXPOSE 8090

ENV GOMEMLIMIT=500MiB

ADD docker_cmd.sh /docker_cmd.sh

CMD /docker_cmd.sh

```

## Creating organizations (100)

#### Creating 50 organizations [reqs:50, conc:10, rule:`""`]

```
┌─ Best:      422.984µs
├─ Worst:     5.486995ms
├─ Completed: 10.899848ms
└─ Errors:    0
```

#### Creating 50 organizations [reqs:50, conc:10, rule:`"@request.body.name != ''"`]

```
┌─ Best:      403.488µs
├─ Worst:     15.106273ms
├─ Completed: 15.440765ms
└─ Errors:    0
```

## Creating permissions (50)

#### Creating 25 permissions [reqs:25, conc:5, rule:`""`]

```
┌─ Best:      440.548µs
├─ Worst:     2.024533ms
├─ Completed: 3.736119ms
└─ Errors:    0
```

#### Creating 25 permissions [reqs:25, conc:5, rule:`"@request.body.name != ''"`]

```
┌─ Best:      480.236µs
├─ Worst:     1.164276ms
├─ Completed: 3.877017ms
└─ Errors:    0
```

## Creating users (500 - expected to be slow due to passwordHash generation)

#### Creating 250 users [reqs:250, conc:50, rule:`""`]

```
┌─ Best:      438.289136ms
├─ Worst:     1.294221387s
├─ Completed: 4.229618672s
└─ Errors:    0
```

#### Creating 250 users [reqs:250, conc:50, rule:`"@request.body.email != '' && @request.body.permissions:length > 0"`]

```
┌─ Best:      439.747476ms
├─ Worst:     1.196151625s
├─ Completed: 4.310529732s
└─ Errors:    0
```

## Creating posts (10k, 25k, 50k, 100k)

#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`""`]

```
┌─ Best:      562.701µs
├─ Worst:     858.359431ms
├─ Completed: 1.055575348s
└─ Errors:    0
```

#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]

```
┌─ Best:      697.498µs
├─ Worst:     2.355779112s
├─ Completed: 3.392900399s
└─ Errors:    0
```

#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`""`]

```
┌─ Best:      449.978µs
├─ Worst:     789.604753ms
├─ Completed: 1.671651689s
└─ Errors:    0
```

#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]

```
┌─ Best:      698.793µs
├─ Worst:     2.101853397s
├─ Completed: 5.354824765s
└─ Errors:    0
```

#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`""`]

```
┌─ Best:      427.266µs
├─ Worst:     653.683309ms
├─ Completed: 3.490852589s
└─ Errors:    0
```

#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]

```
┌─ Best:      703.014µs
├─ Worst:     3.460234935s
├─ Completed: 10.372447133s
└─ Errors:    0
```

#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`""`]

```
┌─ Best:      409.641µs
├─ Worst:     652.967381ms
├─ Completed: 6.86895167s
└─ Errors:    0
```

#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]

```
┌─ Best:      667.055µs
├─ Worst:     2.652105104s
├─ Completed: 19.576743364s
└─ Errors:    0
```

## User auth with password (expected to be slow due to passwordHash verification)

#### users auth with email/pass - high concurrency [reqs:250, conc:250]

```
┌─ Best:      453.073033ms
├─ Worst:     2.072018336s
├─ Completed: 2.073125645s
└─ Errors:    0
```

#### users auth with email/pass - small concurrency [reqs:250, conc:10]

```
┌─ Best:      55.97836ms
├─ Worst:     95.306378ms
├─ Completed: 1.87558046s
└─ Errors:    0
```

## User auth refresh

#### users - auth refresh (high concurrency) [reqs:1000, conc:1000]

```
┌─ Best:      840.27µs
├─ Worst:     1.111711133s
├─ Completed: 1.12553792s
└─ Errors:    0
```

#### users - auth refresh (medium concurrency) [reqs:1000, conc:100]

```
┌─ Best:      354.408µs
├─ Worst:     992.761864ms
├─ Completed: 993.490211ms
└─ Errors:    0
```

## List records

#### users - getOne for auth refresh comparison (high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`/27ish09p35d950k`]

```
┌─ Best:      6.887284ms
├─ Worst:     1.416295338s
├─ Completed: 1.416975351s
└─ Errors:    0
```

#### users - getOne for auth refresh comparison (medium concurrency) [reqs:1000, conc:100, rule:`""`, query:`/27ish09p35d950k`]

```
┌─ Best:      240.785µs
├─ Worst:     1.100421791s
├─ Completed: 1.101196728s
└─ Errors:    0
```

#### posts10k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      703.418µs
├─ Worst:     4.919004ms
├─ Completed: 1.025149259s
└─ Errors:    0
```

#### posts10k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      14.552531ms
├─ Worst:     1.398906259s
├─ Completed: 1.400046076s
└─ Errors:    0
```

#### posts10k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      15.911166ms
├─ Worst:     1.213084264s
├─ Completed: 1.214741395s
└─ Errors:    0
```

#### posts10k - mixed read and write (simpleA list with additional 300 concurrent random posts10k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      11.545727ms
├─ Worst:     1.705058095s
├─ Completed: 1.706994761s
└─ Errors:    0
```

#### posts10k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]

```
┌─ Best:      1.521831ms
├─ Worst:     8.520169ms
├─ Completed: 39.951861ms
└─ Errors:    0
```

#### posts10k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]

```
┌─ Best:      1.895759ms
├─ Worst:     11.938343ms
├─ Completed: 50.93242ms
└─ Errors:    0
```

#### posts10k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]

```
┌─ Best:      2.422357ms
├─ Worst:     11.049687ms
├─ Completed: 60.76289ms
└─ Errors:    0
```

#### posts10k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]

```
┌─ Best:      3.170763ms
├─ Worst:     14.152437ms
├─ Completed: 84.149659ms
└─ Errors:    0
```

#### posts10k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]

```
┌─ Best:      842.579µs
├─ Worst:     12.557214ms
├─ Completed: 29.925622ms
└─ Errors:    0
```

#### posts10k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      2.08466ms
├─ Worst:     145.200031ms
├─ Completed: 161.577669ms
└─ Errors:    0
```

#### posts10k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      798.291µs
├─ Worst:     59.232638ms
├─ Completed: 71.766985ms
└─ Errors:    0
```

#### posts10k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      722.153µs
├─ Worst:     20.480344ms
├─ Completed: 33.617527ms
└─ Errors:    0
```

#### posts10k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      1.887245ms
├─ Worst:     148.736212ms
├─ Completed: 171.960024ms
└─ Errors:    0
```

#### posts10k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      738.542µs
├─ Worst:     45.631664ms
├─ Completed: 58.594793ms
└─ Errors:    0
```

#### posts10k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      766.846µs
├─ Worst:     25.138496ms
├─ Completed: 38.865818ms
└─ Errors:    0
```

#### posts10k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      2.318793ms
├─ Worst:     91.986325ms
├─ Completed: 116.791591ms
└─ Errors:    0
```

#### posts10k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      2.752604ms
├─ Worst:     106.175127ms
├─ Completed: 127.179643ms
└─ Errors:    0
```

#### posts10k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      739.65µs
├─ Worst:     40.534974ms
├─ Completed: 51.427493ms
└─ Errors:    0
```

#### posts10k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]

```
┌─ Best:      108.16627ms
├─ Worst:     298.371926ms
├─ Completed: 1.684894796s
└─ Errors:    0
```

#### posts10k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.354485ms
├─ Worst:     26.604943ms
├─ Completed: 49.278605ms
└─ Errors:    0
```

#### posts10k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]

```
┌─ Best:      92.825944ms
├─ Worst:     278.683404ms
├─ Completed: 1.586094876s
└─ Errors:    0
```

#### posts10k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      994.547µs
├─ Worst:     36.472975ms
├─ Completed: 47.892778ms
└─ Errors:    0
```

#### posts10k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]

```
┌─ Best:      9.115476ms
├─ Worst:     125.19944ms
├─ Completed: 233.437459ms
└─ Errors:    0
```

#### posts10k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.006026ms
├─ Worst:     42.221649ms
├─ Completed: 48.358876ms
└─ Errors:    0
```

#### posts10k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]

```
┌─ Best:      215.086329ms
├─ Worst:     348.632487ms
├─ Completed: 2.612608266s
└─ Errors:    0
```

#### posts10k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      4.497201ms
├─ Worst:     57.247158ms
├─ Completed: 123.000747ms
└─ Errors:    0
```

#### posts10k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]

```
┌─ Best:      256.527857ms
├─ Worst:     406.39793ms
├─ Completed: 3.290275198s
└─ Errors:    0
```

#### posts10k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.558196ms
├─ Worst:     45.347552ms
├─ Completed: 61.872213ms
└─ Errors:    0
```

#### posts25k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      978.415µs
├─ Worst:     6.720651ms
├─ Completed: 1.36529583s
└─ Errors:    0
```

#### posts25k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      88.552546ms
├─ Worst:     1.437806648s
├─ Completed: 1.438551852s
└─ Errors:    0
```

#### posts25k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      9.955083ms
├─ Worst:     1.07109135s
├─ Completed: 1.073258647s
└─ Errors:    0
```

#### posts25k - mixed read and write (simpleA list with additional 300 concurrent random posts25k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      6.044054ms
├─ Worst:     2.718650448s
├─ Completed: 2.721424578s
└─ Errors:    0
```

#### posts25k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]

```
┌─ Best:      1.649533ms
├─ Worst:     12.010725ms
├─ Completed: 44.217772ms
└─ Errors:    0
```

#### posts25k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]

```
┌─ Best:      2.120108ms
├─ Worst:     19.555342ms
├─ Completed: 57.358642ms
└─ Errors:    0
```

#### posts25k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]

```
┌─ Best:      2.71148ms
├─ Worst:     11.350985ms
├─ Completed: 59.498058ms
└─ Errors:    0
```

#### posts25k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]

```
┌─ Best:      3.352577ms
├─ Worst:     23.831295ms
├─ Completed: 87.00575ms
└─ Errors:    0
```

#### posts25k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]

```
┌─ Best:      1.126163ms
├─ Worst:     15.36862ms
├─ Completed: 33.747852ms
└─ Errors:    0
```

#### posts25k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      7.703304ms
├─ Worst:     291.243724ms
├─ Completed: 420.501302ms
└─ Errors:    0
```

#### posts25k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      796.598µs
├─ Worst:     67.381495ms
├─ Completed: 82.72452ms
└─ Errors:    0
```

#### posts25k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      714.159µs
├─ Worst:     19.783234ms
├─ Completed: 31.858652ms
└─ Errors:    0
```

#### posts25k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      7.522806ms
├─ Worst:     299.625605ms
├─ Completed: 418.002146ms
└─ Errors:    0
```

#### posts25k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      840.348µs
├─ Worst:     79.387875ms
├─ Completed: 94.98958ms
└─ Errors:    0
```

#### posts25k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      771.2µs
├─ Worst:     21.047726ms
├─ Completed: 30.91323ms
└─ Errors:    0
```

#### posts25k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      10.334168ms
├─ Worst:     104.319247ms
├─ Completed: 217.575206ms
└─ Errors:    0
```

#### posts25k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      7.902227ms
├─ Worst:     136.133214ms
├─ Completed: 236.162929ms
└─ Errors:    0
```

#### posts25k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      700.287µs
├─ Worst:     48.903109ms
├─ Completed: 56.258594ms
└─ Errors:    0
```

#### posts25k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]

```
┌─ Best:      239.832229ms
├─ Worst:     636.627959ms
├─ Completed: 3.887768899s
└─ Errors:    0
```

#### posts25k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.590943ms
├─ Worst:     77.17992ms
├─ Completed: 99.36677ms
└─ Errors:    0
```

#### posts25k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]

```
┌─ Best:      262.56567ms
├─ Worst:     526.949124ms
├─ Completed: 3.687527896s
└─ Errors:    0
```

#### posts25k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.063194ms
├─ Worst:     76.339497ms
├─ Completed: 80.299284ms
└─ Errors:    0
```

#### posts25k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]

```
┌─ Best:      24.598987ms
├─ Worst:     232.497956ms
├─ Completed: 506.435887ms
└─ Errors:    0
```

#### posts25k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      941.519µs
├─ Worst:     73.095704ms
├─ Completed: 73.121192ms
└─ Errors:    0
```

#### posts25k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]

```
┌─ Best:      550.457712ms
├─ Worst:     734.663353ms
├─ Completed: 6.069286882s
└─ Errors:    0
```

#### posts25k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      4.748819ms
├─ Worst:     90.380175ms
├─ Completed: 144.74999ms
└─ Errors:    0
```

#### posts25k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]

```
┌─ Best:      635.111008ms
├─ Worst:     907.743674ms
├─ Completed: 7.750270858s
└─ Errors:    0
```

#### posts25k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.625027ms
├─ Worst:     76.858189ms
├─ Completed: 93.035139ms
└─ Errors:    0
```

#### posts50k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      1.814909ms
├─ Worst:     6.476353ms
├─ Completed: 2.189062206s
└─ Errors:    0
```

#### posts50k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      77.449224ms
├─ Worst:     1.789289278s
├─ Completed: 1.791112697s
└─ Errors:    0
```

#### posts50k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      18.495214ms
├─ Worst:     1.02146381s
├─ Completed: 1.023512807s
└─ Errors:    0
```

#### posts50k - mixed read and write (simpleA list with additional 300 concurrent random posts50k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      69.581916ms
├─ Worst:     4.696047221s
├─ Completed: 4.697991617s
└─ Errors:    0
```

#### posts50k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]

```
┌─ Best:      2.803475ms
├─ Worst:     32.960195ms
├─ Completed: 82.486996ms
└─ Errors:    0
```

#### posts50k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]

```
┌─ Best:      3.393085ms
├─ Worst:     20.865782ms
├─ Completed: 67.917379ms
└─ Errors:    0
```

#### posts50k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]

```
┌─ Best:      3.579773ms
├─ Worst:     21.233903ms
├─ Completed: 81.105103ms
└─ Errors:    0
```

#### posts50k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]

```
┌─ Best:      4.147578ms
├─ Worst:     21.651642ms
├─ Completed: 98.293172ms
└─ Errors:    0
```

#### posts50k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]

```
┌─ Best:      2.293667ms
├─ Worst:     16.787059ms
├─ Completed: 46.44091ms
└─ Errors:    0
```

#### posts50k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      18.293087ms
├─ Worst:     279.949378ms
├─ Completed: 546.887711ms
└─ Errors:    0
```

#### posts50k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      802.315µs
├─ Worst:     99.819408ms
├─ Completed: 105.12653ms
└─ Errors:    0
```

#### posts50k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      732.703µs
├─ Worst:     20.171592ms
├─ Completed: 30.851659ms
└─ Errors:    0
```

#### posts50k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      17.131003ms
├─ Worst:     294.053157ms
├─ Completed: 532.523346ms
└─ Errors:    0
```

#### posts50k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      758.27µs
├─ Worst:     96.268158ms
├─ Completed: 127.561319ms
└─ Errors:    0
```

#### posts50k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      843.418µs
├─ Worst:     23.780857ms
├─ Completed: 36.189492ms
└─ Errors:    0
```

#### posts50k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      20.94289ms
├─ Worst:     189.933061ms
├─ Completed: 445.510535ms
└─ Errors:    0
```

#### posts50k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      19.315848ms
├─ Worst:     195.47659ms
├─ Completed: 406.235897ms
└─ Errors:    0
```

#### posts50k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      673.275µs
├─ Worst:     58.113265ms
├─ Completed: 59.307219ms
└─ Errors:    0
```

#### posts50k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]

```
┌─ Best:      510.706035ms
├─ Worst:     1.014801867s
├─ Completed: 7.453630116s
└─ Errors:    0
```

#### posts50k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.310106ms
├─ Worst:     52.030682ms
├─ Completed: 77.090753ms
└─ Errors:    0
```

#### posts50k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]

```
┌─ Best:      616.169017ms
├─ Worst:     966.413664ms
├─ Completed: 7.633839868s
└─ Errors:    0
```

#### posts50k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.221638ms
├─ Worst:     67.929077ms
├─ Completed: 68.21867ms
└─ Errors:    0
```

#### posts50k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]

```
┌─ Best:      50.059595ms
├─ Worst:     233.524181ms
├─ Completed: 836.003182ms
└─ Errors:    0
```

#### posts50k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      826.703µs
├─ Worst:     64.255357ms
├─ Completed: 64.427103ms
└─ Errors:    0
```

#### posts50k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]

```
┌─ Best:      1.041009603s
├─ Worst:     1.360762984s
├─ Completed: 11.890172812s
└─ Errors:    0
```

#### posts50k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      4.08872ms
├─ Worst:     84.963837ms
├─ Completed: 140.864032ms
└─ Errors:    0
```

#### posts50k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]

```
┌─ Best:      1.374840153s
├─ Worst:     1.678229815s
├─ Completed: 15.499317918s
└─ Errors:    0
```

#### posts50k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.622757ms
├─ Worst:     66.458018ms
├─ Completed: 85.729319ms
└─ Errors:    0
```

#### posts100k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      2.926102ms
├─ Worst:     10.336858ms
├─ Completed: 3.780975324s
└─ Errors:    0
```

#### posts100k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      145.71105ms
├─ Worst:     2.536017142s
├─ Completed: 2.536555461s
└─ Errors:    0
```

#### posts100k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      23.331337ms
├─ Worst:     1.582810539s
├─ Completed: 1.583643337s
└─ Errors:    0
```

#### posts100k - mixed read and write (simpleA list with additional 300 concurrent random posts100k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      217.652626ms
├─ Worst:     8.469911426s
├─ Completed: 8.472220219s
└─ Errors:    0
```

#### posts100k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]

```
┌─ Best:      4.513142ms
├─ Worst:     39.873784ms
├─ Completed: 108.788166ms
└─ Errors:    0
```

#### posts100k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]

```
┌─ Best:      4.946103ms
├─ Worst:     49.995496ms
├─ Completed: 131.272438ms
└─ Errors:    0
```

#### posts100k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]

```
┌─ Best:      5.476714ms
├─ Worst:     32.223553ms
├─ Completed: 113.557697ms
└─ Errors:    0
```

#### posts100k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]

```
┌─ Best:      6.015487ms
├─ Worst:     43.828422ms
├─ Completed: 164.297142ms
└─ Errors:    0
```

#### posts100k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]

```
┌─ Best:      3.307646ms
├─ Worst:     40.452603ms
├─ Completed: 87.641832ms
└─ Errors:    0
```

#### posts100k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      37.58286ms
├─ Worst:     270.050515ms
├─ Completed: 796.012205ms
└─ Errors:    0
```

#### posts100k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      769.8µs
├─ Worst:     85.278258ms
├─ Completed: 104.636363ms
└─ Errors:    0
```

#### posts100k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      739.611µs
├─ Worst:     23.498438ms
├─ Completed: 36.691889ms
└─ Errors:    0
```

#### posts100k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      35.646002ms
├─ Worst:     296.335468ms
├─ Completed: 788.660282ms
└─ Errors:    0
```

#### posts100k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      777.743µs
├─ Worst:     94.469123ms
├─ Completed: 109.295718ms
└─ Errors:    0
```

#### posts100k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      782.431µs
├─ Worst:     18.260006ms
├─ Completed: 31.81687ms
└─ Errors:    0
```

#### posts100k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      42.14636ms
├─ Worst:     171.325805ms
├─ Completed: 695.920765ms
└─ Errors:    0
```

#### posts100k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      36.377141ms
├─ Worst:     217.345989ms
├─ Completed: 646.500447ms
└─ Errors:    0
```

#### posts100k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      710.242µs
├─ Worst:     61.206384ms
├─ Completed: 62.555793ms
└─ Errors:    0
```

#### posts100k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]

```
┌─ Best:      1.295308757s
├─ Worst:     1.795331336s
├─ Completed: 14.768416541s
└─ Errors:    0
```

#### posts100k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.613594ms
├─ Worst:     57.19545ms
├─ Completed: 87.908655ms
└─ Errors:    0
```

#### posts100k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]

```
┌─ Best:      1.140773751s
├─ Worst:     1.831045808s
├─ Completed: 14.746897029s
└─ Errors:    0
```

#### posts100k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.222771ms
├─ Worst:     69.78761ms
├─ Completed: 70.929503ms
└─ Errors:    0
```

#### posts100k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]

```
┌─ Best:      110.576535ms
├─ Worst:     304.112192ms
├─ Completed: 1.617739341s
└─ Errors:    0
```

#### posts100k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      901.252µs
├─ Worst:     49.858528ms
├─ Completed: 54.890892ms
└─ Errors:    0
```

#### posts100k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]

```
┌─ Best:      2.255989708s
├─ Worst:     2.594283664s
├─ Completed: 23.898465059s
└─ Errors:    0
```

#### posts100k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      5.672143ms
├─ Worst:     89.778761ms
├─ Completed: 155.057025ms
└─ Errors:    0
```

#### posts100k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]

```
┌─ Best:      2.858024638s
├─ Worst:     3.32676391s
├─ Completed: 30.753608721s
└─ Errors:    0
```

#### posts100k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.900096ms
├─ Worst:     77.219114ms
├─ Completed: 92.091564ms
└─ Errors:    0
```

## Go vs JS route execution

#### JS route (high concurrency) [reqs:500, conc:500]

```
┌─ Best:      4.896988378s
├─ Worst:     26.828036004s
├─ Completed: 26.831431272s
└─ Errors:    0
```

#### Go route (high concurrency) [reqs:500, conc:500]

```
┌─ Best:      2.975519546s
├─ Worst:     24.815333852s
├─ Completed: 24.817102387s
└─ Errors:    0
```

#### JS route (medium concurrency) [reqs:500, conc:50]

```
┌─ Best:      717.539529ms
├─ Worst:     5.25811619s
├─ Completed: 21.493397372s
└─ Errors:    0
```

#### Go route (medium concurrency) [reqs:500, conc:50]

```
┌─ Best:      907.162164ms
├─ Worst:     4.768224094s
├─ Completed: 23.430452125s
└─ Errors:    0
```

#### JS route (no concurrency) [reqs:500, conc:1]

```
┌─ Best:      9.135043ms
├─ Worst:     17.695781ms
├─ Completed: 5.391244922s
└─ Errors:    0
```

#### Go route (no concurrency) [reqs:500, conc:1]

```
┌─ Best:      8.92006ms
├─ Worst:     19.481095ms
├─ Completed: 5.330896469s
└─ Errors:    0
```

## Go vs JS hooks execution

#### JS OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]

```
┌─ Best:      608.247µs
├─ Worst:     92.57722ms
├─ Completed: 94.906467ms
└─ Errors:    0
```

#### Go OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]

```
┌─ Best:      585.244µs
├─ Worst:     17.459427ms
├─ Completed: 24.61267ms
└─ Errors:    0
```

## Deleting records

#### deleting 100 posts10k - simple (no cascade, no rule) [conc:10, rule:`""`]

```
┌─ Best:      407.679µs
├─ Worst:     19.221512ms
├─ Completed: 24.60406ms
└─ Errors:    0
```

#### deleting 100 posts10k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]

```
┌─ Best:      511.026µs
├─ Worst:     10.552196ms
├─ Completed: 21.22611ms
└─ Errors:    0
```

#### deleting 100 posts25k - simple (no cascade, no rule) [conc:10, rule:`""`]

```
┌─ Best:      459.919µs
├─ Worst:     1.829055ms
├─ Completed: 9.055699ms
└─ Errors:    0
```

#### deleting 100 posts25k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]

```
┌─ Best:      444.101µs
├─ Worst:     1.580376ms
├─ Completed: 9.206707ms
└─ Errors:    0
```

#### deleting 100 posts50k - simple (no cascade, no rule) [conc:10, rule:`""`]

```
┌─ Best:      429.301µs
├─ Worst:     6.208246ms
├─ Completed: 14.900668ms
└─ Errors:    0
```

#### deleting 100 posts50k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]

```
┌─ Best:      481.998µs
├─ Worst:     2.187118ms
├─ Completed: 9.417857ms
└─ Errors:    0
```

#### deleting 100 posts100k - simple (no cascade, no rule) [conc:10, rule:`""`]

```
┌─ Best:      447.181µs
├─ Worst:     12.066901ms
├─ Completed: 20.477985ms
└─ Errors:    0
```

#### deleting 100 posts100k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]

```
┌─ Best:      598.746µs
├─ Worst:     2.95285ms
├─ Completed: 11.97499ms
└─ Errors:    0
```

#### deleting 100 users - with cascade deleting all associated posts [conc:10, rule:`""`]

```
┌─ Best:      84.124054ms
├─ Worst:     4.080352648s
├─ Completed: 8.200931897s
└─ Errors:    0
```

#### deleting 100 organizations - with cascade deleting all users and associated posts [conc:10, rule:`""`]

```
┌─ Best:      108.025796ms
├─ Worst:     9.706040142s
├─ Completed: 19.308718112s
└─ Errors:    0
```

---

Completed!
