_PocketBase v0.40.0-dev (Go 1.27rc3)_

## Creating organizations (100)
#### Creating 50 organizations [reqs:50, conc:10, rule:`""`]
```
┌─ Best:      749.7µs
├─ Worst:     7.963373ms
├─ Completed: 16.82421ms
└─ Errors:    0
```
#### Creating 50 organizations [reqs:50, conc:10, rule:`"@request.body.name != ''"`]
```
┌─ Best:      730.019µs
├─ Worst:     7.707326ms
├─ Completed: 17.92276ms
└─ Errors:    0
```

## Creating permissions (50)
#### Creating 25 permissions [reqs:25, conc:5, rule:`""`]
```
┌─ Best:      427.771µs
├─ Worst:     2.585389ms
├─ Completed: 6.453133ms
└─ Errors:    0
```
#### Creating 25 permissions [reqs:25, conc:5, rule:`"@request.body.name != ''"`]
```
┌─ Best:      544.095µs
├─ Worst:     3.651218ms
├─ Completed: 10.362997ms
└─ Errors:    0
```

## Creating users (500 - expected to be slow due to passwordHash generation)
#### Creating 250 users [reqs:250, conc:50, rule:`""`]
```
┌─ Best:      1.778036897s
├─ Worst:     4.104871666s
├─ Completed: 16.489465938s
└─ Errors:    0
```
#### Creating 250 users [reqs:250, conc:50, rule:`"@request.body.email != '' && @request.body.permissions:length > 0"`]
```
┌─ Best:      1.629831174s
├─ Worst:     4.282373492s
├─ Completed: 16.591378782s
└─ Errors:    0
```

## Creating posts (10k, 25k, 50k, 100k)
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`""`]
```
┌─ Best:      992.666µs
├─ Worst:     1.09630335s
├─ Completed: 1.762894492s
└─ Errors:    0
```
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]
```
┌─ Best:      1.859049ms
├─ Worst:     1.940439799s
├─ Completed: 2.777289413s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`""`]
```
┌─ Best:      1.193472ms
├─ Worst:     2.140505828s
├─ Completed: 4.574456462s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]
```
┌─ Best:      1.783888ms
├─ Worst:     2.386129355s
├─ Completed: 6.709106972s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`""`]
```
┌─ Best:      950.465µs
├─ Worst:     1.896914195s
├─ Completed: 8.706778222s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]
```
┌─ Best:      1.327155ms
├─ Worst:     3.240172389s
├─ Completed: 14.057668401s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`""`]
```
┌─ Best:      819.462µs
├─ Worst:     2.026562901s
├─ Completed: 17.542950343s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]
```
┌─ Best:      1.361037ms
├─ Worst:     3.018967354s
├─ Completed: 28.184594615s
└─ Errors:    0
```

## User auth with password (expected to be slow due to passwordHash verification)
#### users auth with email/pass - high concurrency [reqs:250, conc:250]
```
┌─ Best:      6.905382541s
├─ Worst:     8.237797323s
├─ Completed: 8.23804289s
└─ Errors:    0
```
#### users auth with email/pass - small concurrency [reqs:250, conc:10]
```
┌─ Best:      166.43021ms
├─ Worst:     453.259078ms
├─ Completed: 8.261187149s
└─ Errors:    0
```

## User auth refresh
#### users - auth refresh (high concurrency) [reqs:1000, conc:1000]
```
┌─ Best:      125.533436ms
├─ Worst:     356.767298ms
├─ Completed: 358.340901ms
└─ Errors:    0
```
#### users - auth refresh (medium concurrency) [reqs:1000, conc:100]
```
┌─ Best:      368.85µs
├─ Worst:     107.088703ms
├─ Completed: 223.051923ms
└─ Errors:    0
```

## List records
#### users - getOne for auth refresh comparison (medium concurrency) [reqs:1000, conc:100, rule:`""`, query:`/gz9z57eia7u4ztx`]
```
┌─ Best:      430.211µs
├─ Worst:     117.758909ms
├─ Completed: 287.102636ms
└─ Errors:    0
```
#### users - getOne for auth refresh comparison (high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`/gz9z57eia7u4ztx`]
```
┌─ Best:      104.633238ms
├─ Worst:     276.955845ms
├─ Completed: 283.240132ms
└─ Errors:    0
```
#### posts10k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.282581ms
├─ Worst:     7.023508ms
├─ Completed: 3.050205429s
└─ Errors:    0
```
#### posts10k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      46.853973ms
├─ Worst:     1.606581273s
├─ Completed: 1.614474684s
└─ Errors:    0
```
#### posts10k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      189.61231ms
├─ Worst:     638.45195ms
├─ Completed: 655.765133ms
└─ Errors:    0
```
#### posts10k - mixed read and write (simpleA list with additional 300 concurrent random posts10k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      228.201941ms
├─ Worst:     1.851403539s
├─ Completed: 1.877787364s
└─ Errors:    0
```
#### posts10k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      3.532774ms
├─ Worst:     41.682194ms
├─ Completed: 186.152297ms
└─ Errors:    0
```
#### posts10k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      4.036628ms
├─ Worst:     71.845441ms
├─ Completed: 243.924682ms
└─ Errors:    0
```
#### posts10k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      6.048082ms
├─ Worst:     60.389815ms
├─ Completed: 275.521086ms
└─ Errors:    0
```
#### posts10k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      5.523748ms
├─ Worst:     71.505272ms
├─ Completed: 338.317845ms
└─ Errors:    0
```
#### posts10k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      2.60295ms
├─ Worst:     29.864798ms
├─ Completed: 142.803698ms
└─ Errors:    0
```
#### posts10k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      27.279009ms
├─ Worst:     333.510237ms
├─ Completed: 815.782091ms
└─ Errors:    0
```
#### posts10k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.444399ms
├─ Worst:     83.339748ms
├─ Completed: 134.020063ms
└─ Errors:    0
```
#### posts10k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.040828ms
├─ Worst:     27.603938ms
├─ Completed: 80.959725ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      32.1651ms
├─ Worst:     297.288428ms
├─ Completed: 785.618364ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.160711ms
├─ Worst:     68.487031ms
├─ Completed: 115.733174ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.059708ms
├─ Worst:     21.69442ms
├─ Completed: 62.417829ms
└─ Errors:    0
```
#### posts10k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      32.755515ms
├─ Worst:     279.965645ms
├─ Completed: 850.508859ms
└─ Errors:    0
```
#### posts10k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      23.506308ms
├─ Worst:     234.16486ms
├─ Completed: 803.380239ms
└─ Errors:    0
```
#### posts10k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.10779ms
├─ Worst:     57.577499ms
├─ Completed: 116.002541ms
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      366.822527ms
├─ Worst:     748.327807ms
├─ Completed: 5.108249493s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.478199ms
├─ Worst:     62.905202ms
├─ Completed: 123.825151ms
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      45.65686ms
├─ Worst:     357.841927ms
├─ Completed: 1.873170081s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.133111ms
├─ Worst:     56.897042ms
├─ Completed: 98.231306ms
└─ Errors:    0
```
#### posts10k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      243.730076ms
├─ Worst:     461.17509ms
├─ Completed: 3.522886146s
└─ Errors:    0
```
#### posts10k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.558842ms
├─ Worst:     41.303825ms
├─ Completed: 89.663397ms
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      1.97747627s
├─ Worst:     3.946486152s
├─ Completed: 30.051632132s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      18.32545ms
├─ Worst:     196.078643ms
├─ Completed: 568.265393ms
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      3.031455448s
├─ Worst:     4.936810749s
├─ Completed: 40.203436507s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      3.147484ms
├─ Worst:     95.715319ms
├─ Completed: 232.379613ms
└─ Errors:    0
```
#### posts25k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      4.619643ms
├─ Worst:     13.140992ms
├─ Completed: 6.203410413s
└─ Errors:    0
```
#### posts25k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      193.0118ms
├─ Worst:     3.075381543s
├─ Completed: 3.09586865s
└─ Errors:    0
```
#### posts25k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      86.868843ms
├─ Worst:     749.583361ms
├─ Completed: 757.107682ms
└─ Errors:    0
```
#### posts25k - mixed read and write (simpleA list with additional 300 concurrent random posts25k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      132.016649ms
├─ Worst:     3.249121107s
├─ Completed: 3.252980651s
└─ Errors:    0
```
#### posts25k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      7.511921ms
├─ Worst:     67.578767ms
├─ Completed: 336.568838ms
└─ Errors:    0
```
#### posts25k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      6.445612ms
├─ Worst:     85.638689ms
├─ Completed: 352.160455ms
└─ Errors:    0
```
#### posts25k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      9.391891ms
├─ Worst:     105.426099ms
├─ Completed: 393.229353ms
└─ Errors:    0
```
#### posts25k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      11.747954ms
├─ Worst:     111.141811ms
├─ Completed: 443.439615ms
└─ Errors:    0
```
#### posts25k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      4.930532ms
├─ Worst:     67.790373ms
├─ Completed: 258.07862ms
└─ Errors:    0
```
#### posts25k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      64.855414ms
├─ Worst:     632.034018ms
├─ Completed: 1.569201513s
└─ Errors:    0
```
#### posts25k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.172992ms
├─ Worst:     186.26762ms
├─ Completed: 232.537937ms
└─ Errors:    0
```
#### posts25k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.039748ms
├─ Worst:     23.305863ms
├─ Completed: 59.707957ms
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      52.892014ms
├─ Worst:     648.112608ms
├─ Completed: 1.551438439s
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.13439ms
├─ Worst:     146.233509ms
├─ Completed: 189.679511ms
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      936.145µs
├─ Worst:     25.572123ms
├─ Completed: 62.236664ms
└─ Errors:    0
```
#### posts25k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      78.096968ms
├─ Worst:     672.225453ms
├─ Completed: 2.116012653s
└─ Errors:    0
```
#### posts25k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      68.637875ms
├─ Worst:     559.542919ms
├─ Completed: 1.906508812s
└─ Errors:    0
```
#### posts25k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.11535ms
├─ Worst:     101.38199ms
├─ Completed: 133.745776ms
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      1.206009043s
├─ Worst:     1.886161068s
├─ Completed: 13.907335182s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.627123ms
├─ Worst:     111.196093ms
├─ Completed: 162.284699ms
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      249.260424ms
├─ Worst:     958.631829ms
├─ Completed: 4.651746328s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.144951ms
├─ Worst:     106.714853ms
├─ Completed: 156.186496ms
└─ Errors:    0
```
#### posts25k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      659.116582ms
├─ Worst:     1.409078712s
├─ Completed: 9.166266386s
└─ Errors:    0
```
#### posts25k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.212792ms
├─ Worst:     178.750419ms
├─ Completed: 224.014949ms
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      5.985163782s
├─ Worst:     8.896773141s
├─ Completed: 1m16.823787443s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      9.511134ms
├─ Worst:     243.956169ms
├─ Completed: 605.550927ms
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      8.818126602s
├─ Worst:     12.081556982s
├─ Completed: 1m40.105377463s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      3.167405ms
├─ Worst:     203.198798ms
├─ Completed: 314.105887ms
└─ Errors:    0
```
#### posts50k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      8.496467ms
├─ Worst:     19.03275ms
├─ Completed: 10.526114033s
└─ Errors:    0
```
#### posts50k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      189.186823ms
├─ Worst:     5.198116879s
├─ Completed: 5.205366992s
└─ Errors:    0
```
#### posts50k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      49.081393ms
├─ Worst:     717.738329ms
├─ Completed: 719.239889ms
└─ Errors:    0
```
#### posts50k - mixed read and write (simpleA list with additional 300 concurrent random posts50k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      127.77178ms
├─ Worst:     5.696303852s
├─ Completed: 5.718759453s
└─ Errors:    0
```
#### posts50k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      14.905559ms
├─ Worst:     129.721752ms
├─ Completed: 526.004558ms
└─ Errors:    0
```
#### posts50k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      9.966507ms
├─ Worst:     197.161796ms
├─ Completed: 599.561886ms
└─ Errors:    0
```
#### posts50k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      16.038269ms
├─ Worst:     194.124195ms
├─ Completed: 612.731119ms
└─ Errors:    0
```
#### posts50k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      11.923799ms
├─ Worst:     163.019723ms
├─ Completed: 671.584694ms
└─ Errors:    0
```
#### posts50k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      18.015842ms
├─ Worst:     161.77389ms
├─ Completed: 485.253547ms
└─ Errors:    0
```
#### posts50k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      106.451809ms
├─ Worst:     1.089433476s
├─ Completed: 2.913849464s
└─ Errors:    0
```
#### posts50k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.169431ms
├─ Worst:     299.974628ms
├─ Completed: 349.910725ms
└─ Errors:    0
```
#### posts50k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.091269ms
├─ Worst:     31.00231ms
├─ Completed: 90.843952ms
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      101.969849ms
├─ Worst:     1.034023354s
├─ Completed: 2.591261071s
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.373396ms
├─ Worst:     268.554468ms
├─ Completed: 324.422203ms
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      996.787µs
├─ Worst:     21.27077ms
├─ Completed: 60.323534ms
└─ Errors:    0
```
#### posts50k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      189.240025ms
├─ Worst:     1.043686492s
├─ Completed: 3.97149637s
└─ Errors:    0
```
#### posts50k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      192.172143ms
├─ Worst:     1.034511487s
├─ Completed: 3.806816722s
└─ Errors:    0
```
#### posts50k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      988.867µs
├─ Worst:     147.698833ms
├─ Completed: 179.410801ms
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      2.66635076s
├─ Worst:     3.587317488s
├─ Completed: 30.499895916s
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.669964ms
├─ Worst:     171.237103ms
├─ Completed: 236.896981ms
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      545.608722ms
├─ Worst:     1.925045481s
├─ Completed: 9.153275651s
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.282155ms
├─ Worst:     251.338486ms
├─ Completed: 303.844412ms
└─ Errors:    0
```
#### posts50k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      1.428694396s
├─ Worst:     2.631952799s
├─ Completed: 18.045069305s
└─ Errors:    0
```
#### posts50k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.47692ms
├─ Worst:     230.91546ms
├─ Completed: 284.079003ms
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      11.375946257s
├─ Worst:     16.630317121s
├─ Completed: 2m33.975203154s
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      7.066349ms
├─ Worst:     425.694873ms
├─ Completed: 720.930655ms
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      17.111877409s
├─ Worst:     23.072449333s
├─ Completed: 3m19.894552924s
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      3.225047ms
├─ Worst:     342.085016ms
├─ Completed: 441.253658ms
└─ Errors:    0
```
#### posts100k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      16.206399ms
├─ Worst:     28.257924ms
├─ Completed: 19.448259661s
└─ Errors:    0
```
#### posts100k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      204.658577ms
├─ Worst:     9.483271054s
├─ Completed: 9.503943213s
└─ Errors:    0
```
#### posts100k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      57.976088ms
├─ Worst:     736.552287ms
├─ Completed: 740.014622ms
└─ Errors:    0
```
#### posts100k - mixed read and write (simpleA list with additional 300 concurrent random posts100k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      139.373691ms
├─ Worst:     10.493344941s
├─ Completed: 10.499642152s
└─ Errors:    0
```
#### posts100k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      36.914719ms
├─ Worst:     247.552778ms
├─ Completed: 994.362903ms
└─ Errors:    0
```
#### posts100k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      28.133562ms
├─ Worst:     270.198791ms
├─ Completed: 991.499906ms
└─ Errors:    0
```
#### posts100k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      25.065438ms
├─ Worst:     267.799285ms
├─ Completed: 1.019962636s
└─ Errors:    0
```
#### posts100k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      34.822582ms
├─ Worst:     307.908331ms
├─ Completed: 1.075661702s
└─ Errors:    0
```
#### posts100k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      27.255177ms
├─ Worst:     198.510611ms
├─ Completed: 904.657595ms
└─ Errors:    0
```
#### posts100k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      207.69738ms
├─ Worst:     1.696324894s
├─ Completed: 5.640465285s
└─ Errors:    0
```
#### posts100k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.428186ms
├─ Worst:     290.284214ms
├─ Completed: 364.394819ms
└─ Errors:    0
```
#### posts100k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.009867ms
├─ Worst:     26.902888ms
├─ Completed: 61.903715ms
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      257.473286ms
├─ Worst:     1.01016081s
├─ Completed: 4.98312638s
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.419678ms
├─ Worst:     273.240712ms
├─ Completed: 326.226146ms
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.005707ms
├─ Worst:     34.042641ms
├─ Completed: 60.731563ms
└─ Errors:    0
```
#### posts100k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      396.106197ms
├─ Worst:     2.186438315s
├─ Completed: 7.986867128s
└─ Errors:    0
```
#### posts100k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      409.620723ms
├─ Worst:     1.141923895s
├─ Completed: 7.005699742s
└─ Errors:    0
```
#### posts100k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.036268ms
├─ Worst:     202.415276ms
├─ Completed: 261.030302ms
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      5.654287539s
├─ Worst:     6.936197982s
├─ Completed: 1m5.613344117s
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.747767ms
├─ Worst:     453.860679ms
├─ Completed: 532.21836ms
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      1.037624553s
├─ Worst:     3.053660058s
├─ Completed: 18.259423137s
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.204232ms
├─ Worst:     366.966928ms
├─ Completed: 431.523035ms
└─ Errors:    0
```
#### posts100k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      2.686788492s
├─ Worst:     4.284844129s
├─ Completed: 36.113672511s
└─ Errors:    0
```
#### posts100k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.024348ms
├─ Worst:     463.021728ms
├─ Completed: 511.187951ms
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      26.735136494s
├─ Worst:     33.297819544s
├─ Completed: 5m14.079495232s
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      17.1177ms
├─ Worst:     436.45061ms
├─ Completed: 891.604922ms
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      35.370157553s
├─ Worst:     43.853443902s
├─ Completed: 6m42.939435875s
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      3.162048ms
├─ Worst:     380.629587ms
├─ Completed: 502.08022ms
└─ Errors:    0
```

## Go vs JS route execution
#### JS route (high concurrency) [reqs:500, conc:500]
```
┌─ Best:      1.075991543s
├─ Worst:     14.482336304s
├─ Completed: 14.499693944s
└─ Errors:    0
```
#### Go route (high concurrency) [reqs:500, conc:500]
```
┌─ Best:      1.242995913s
├─ Worst:     13.399614416s
├─ Completed: 13.40482748s
└─ Errors:    0
```
#### JS route (medium concurrency) [reqs:500, conc:50]
```
┌─ Best:      588.404083ms
├─ Worst:     1.494143126s
├─ Completed: 10.356135681s
└─ Errors:    0
```
#### Go route (medium concurrency) [reqs:500, conc:50]
```
┌─ Best:      600.753784ms
├─ Worst:     1.986653602s
├─ Completed: 10.522325549s
└─ Errors:    0
```
#### JS route (no concurrency) [reqs:500, conc:1]
```
┌─ Best:      30.552124ms
├─ Worst:     46.535685ms
├─ Completed: 17.37692557s
└─ Errors:    0
```
#### Go route (no concurrency) [reqs:500, conc:1]
```
┌─ Best:      30.144192ms
├─ Worst:     45.781304ms
├─ Completed: 17.565858145s
└─ Errors:    0
```

## Go vs JS hooks execution
#### JS OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]
```
┌─ Best:      1.290476ms
├─ Worst:     129.525375ms
├─ Completed: 168.749258ms
└─ Errors:    0
```
#### Go OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]
```
┌─ Best:      1.033748ms
├─ Worst:     21.095142ms
├─ Completed: 54.151415ms
└─ Errors:    0
```

## Deleting records
#### deleting 100 posts10k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      750.261µs
├─ Worst:     24.911168ms
├─ Completed: 45.358892ms
└─ Errors:    0
```
#### deleting 100 posts10k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      1.222153ms
├─ Worst:     28.787195ms
├─ Completed: 64.254934ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      838.423µs
├─ Worst:     10.792418ms
├─ Completed: 39.684216ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      967.387µs
├─ Worst:     14.537801ms
├─ Completed: 40.334474ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      1.053189ms
├─ Worst:     8.266508ms
├─ Completed: 36.425046ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      1.037149ms
├─ Worst:     13.424411ms
├─ Completed: 39.328406ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      821.143µs
├─ Worst:     42.145404ms
├─ Completed: 78.763454ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      898.985µs
├─ Worst:     9.809631ms
├─ Completed: 35.264334ms
└─ Errors:    0
```
#### deleting 100 users - with cascade deleting all associated posts [conc:10, rule:`""`]
```
┌─ Best:      165.889899ms
├─ Worst:     8.853374717s
├─ Completed: 18.379476005s
└─ Errors:    0
```
#### deleting 100 organizations - with cascade deleting all users and associated posts [conc:10, rule:`""`]
```
┌─ Best:      216.971109ms
├─ Worst:     33.515866445s
├─ Completed: 44.077648523s
└─ Errors:    0
```
