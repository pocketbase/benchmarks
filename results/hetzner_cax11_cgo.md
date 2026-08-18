_PocketBase v0.40.0-dev (Go 1.27rc3; CGO)_

## Creating organizations (100)
#### Creating 50 organizations [reqs:50, conc:10, rule:`""`]
```
┌─ Best:      421.85µs
├─ Worst:     9.378713ms
├─ Completed: 14.683884ms
└─ Errors:    0
```
#### Creating 50 organizations [reqs:50, conc:10, rule:`"@request.body.name != ''"`]
```
┌─ Best:      641.456µs
├─ Worst:     6.561603ms
├─ Completed: 12.50823ms
└─ Errors:    0
```

## Creating permissions (50)
#### Creating 25 permissions [reqs:25, conc:5, rule:`""`]
```
┌─ Best:      286.847µs
├─ Worst:     4.043181ms
├─ Completed: 7.546027ms
└─ Errors:    0
```
#### Creating 25 permissions [reqs:25, conc:5, rule:`"@request.body.name != ''"`]
```
┌─ Best:      424.411µs
├─ Worst:     3.985939ms
├─ Completed: 7.983118ms
└─ Errors:    0
```

## Creating users (500 - expected to be slow due to passwordHash generation)
#### Creating 250 users [reqs:250, conc:50, rule:`""`]
```
┌─ Best:      1.793600417s
├─ Worst:     4.219845063s
├─ Completed: 16.485505273s
└─ Errors:    0
```
#### Creating 250 users [reqs:250, conc:50, rule:`"@request.body.email != '' && @request.body.permissions:length > 0"`]
```
┌─ Best:      2.008897723s
├─ Worst:     4.387764192s
├─ Completed: 16.585044784s
└─ Errors:    0
```

## Creating posts (10k, 25k, 50k, 100k)
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`""`]
```
┌─ Best:      20.170701ms
├─ Worst:     493.531335ms
├─ Completed: 1.436489389s
└─ Errors:    0
```
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]
```
┌─ Best:      23.543624ms
├─ Worst:     546.985222ms
├─ Completed: 2.330642552s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`""`]
```
┌─ Best:      3.692892ms
├─ Worst:     687.514952ms
├─ Completed: 3.824895776s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]
```
┌─ Best:      25.219266ms
├─ Worst:     910.719814ms
├─ Completed: 6.340490241s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`""`]
```
┌─ Best:      6.089831ms
├─ Worst:     853.993926ms
├─ Completed: 7.709938646s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]
```
┌─ Best:      19.403002ms
├─ Worst:     882.603556ms
├─ Completed: 12.743354151s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`""`]
```
┌─ Best:      2.684666ms
├─ Worst:     944.509333ms
├─ Completed: 15.095029866s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]
```
┌─ Best:      12.556632ms
├─ Worst:     926.529966ms
├─ Completed: 23.522528723s
└─ Errors:    0
```

## User auth with password (expected to be slow due to passwordHash verification)
#### users auth with email/pass - high concurrency [reqs:250, conc:250]
```
┌─ Best:      7.746209968s
├─ Worst:     8.202944024s
├─ Completed: 8.204053693s
└─ Errors:    0
```
#### users auth with email/pass - small concurrency [reqs:250, conc:10]
```
┌─ Best:      120.326545ms
├─ Worst:     438.252694ms
├─ Completed: 8.254570293s
└─ Errors:    0
```

## User auth refresh
#### users - auth refresh (high concurrency) [reqs:1000, conc:1000]
```
┌─ Best:      128.686364ms
├─ Worst:     254.515851ms
├─ Completed: 254.929822ms
└─ Errors:    0
```
#### users - auth refresh (medium concurrency) [reqs:1000, conc:100]
```
┌─ Best:      360.929µs
├─ Worst:     91.958923ms
├─ Completed: 225.911104ms
└─ Errors:    0
```

## List records
#### users - getOne for auth refresh comparison (medium concurrency) [reqs:1000, conc:100, rule:`""`, query:`/opnujn7rlcnj9en`]
```
┌─ Best:      1.381556ms
├─ Worst:     94.49919ms
├─ Completed: 242.611261ms
└─ Errors:    0
```
#### users - getOne for auth refresh comparison (high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`/opnujn7rlcnj9en`]
```
┌─ Best:      51.616469ms
├─ Worst:     260.104958ms
├─ Completed: 267.223944ms
└─ Errors:    0
```
#### posts10k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      980.985µs
├─ Worst:     5.156295ms
├─ Completed: 1.455219672s
└─ Errors:    0
```
#### posts10k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      211.524329ms
├─ Worst:     829.495919ms
├─ Completed: 841.729639ms
└─ Errors:    0
```
#### posts10k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      165.30424ms
├─ Worst:     587.201306ms
├─ Completed: 589.699772ms
└─ Errors:    0
```
#### posts10k - mixed read and write (simpleA list with additional 300 concurrent random posts10k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      149.209179ms
├─ Worst:     1.050019722s
├─ Completed: 1.069497992s
└─ Errors:    0
```
#### posts10k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      2.252179ms
├─ Worst:     45.254143ms
├─ Completed: 145.837292ms
└─ Errors:    0
```
#### posts10k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      5.006691ms
├─ Worst:     51.749073ms
├─ Completed: 160.469674ms
└─ Errors:    0
```
#### posts10k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      3.356328ms
├─ Worst:     66.160609ms
├─ Completed: 210.504421ms
└─ Errors:    0
```
#### posts10k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      7.477556ms
├─ Worst:     57.973355ms
├─ Completed: 270.662114ms
└─ Errors:    0
```
#### posts10k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      1.421517ms
├─ Worst:     19.147221ms
├─ Completed: 79.507438ms
└─ Errors:    0
```
#### posts10k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      3.525812ms
├─ Worst:     153.594934ms
├─ Completed: 308.900593ms
└─ Errors:    0
```
#### posts10k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.071948ms
├─ Worst:     21.675887ms
├─ Completed: 73.08955ms
└─ Errors:    0
```
#### posts10k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      961.305µs
├─ Worst:     11.322936ms
├─ Completed: 53.430556ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.898196ms
├─ Worst:     59.720801ms
├─ Completed: 201.877596ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.007066ms
├─ Worst:     19.403707ms
├─ Completed: 62.479113ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.325674ms
├─ Worst:     16.961523ms
├─ Completed: 80.867033ms
└─ Errors:    0
```
#### posts10k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      3.250365ms
├─ Worst:     51.228539ms
├─ Completed: 178.392903ms
└─ Errors:    0
```
#### posts10k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      3.05872ms
├─ Worst:     57.367339ms
├─ Completed: 209.745801ms
└─ Errors:    0
```
#### posts10k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.060428ms
├─ Worst:     13.611996ms
├─ Completed: 53.340354ms
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      40.603261ms
├─ Worst:     225.72918ms
├─ Completed: 1.216555635s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.345195ms
├─ Worst:     16.248265ms
├─ Completed: 67.465643ms
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      19.214023ms
├─ Worst:     177.719685ms
├─ Completed: 676.942532ms
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.194471ms
├─ Worst:     13.41279ms
├─ Completed: 57.240416ms
└─ Errors:    0
```
#### posts10k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      21.718528ms
├─ Worst:     143.638754ms
├─ Completed: 650.774408ms
└─ Errors:    0
```
#### posts10k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.232153ms
├─ Worst:     17.766585ms
├─ Completed: 55.774617ms
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      479.919983ms
├─ Worst:     762.083197ms
├─ Completed: 6.48021292s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      3.157002ms
├─ Worst:     36.901324ms
├─ Completed: 157.506076ms
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      717.577234ms
├─ Worst:     995.849026ms
├─ Completed: 8.811390446s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.777367ms
├─ Worst:     18.027631ms
├─ Completed: 79.493517ms
└─ Errors:    0
```
#### posts25k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      1.648803ms
├─ Worst:     5.281498ms
├─ Completed: 2.344780801s
└─ Errors:    0
```
#### posts25k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      231.610214ms
├─ Worst:     1.16701142s
├─ Completed: 1.181714285s
└─ Errors:    0
```
#### posts25k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      143.713356ms
├─ Worst:     573.16966ms
├─ Completed: 574.825143ms
└─ Errors:    0
```
#### posts25k - mixed read and write (simpleA list with additional 300 concurrent random posts25k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      286.750934ms
├─ Worst:     1.344469818s
├─ Completed: 1.347747543s
└─ Errors:    0
```
#### posts25k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      2.899356ms
├─ Worst:     36.34779ms
├─ Completed: 146.599072ms
└─ Errors:    0
```
#### posts25k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      3.243125ms
├─ Worst:     58.066598ms
├─ Completed: 196.178608ms
└─ Errors:    0
```
#### posts25k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      4.20223ms
├─ Worst:     71.769835ms
├─ Completed: 235.137186ms
└─ Errors:    0
```
#### posts25k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      4.991691ms
├─ Worst:     66.982671ms
├─ Completed: 302.954558ms
└─ Errors:    0
```
#### posts25k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      2.042213ms
├─ Worst:     23.938466ms
├─ Completed: 110.321363ms
└─ Errors:    0
```
#### posts25k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      5.544345ms
├─ Worst:     119.996896ms
├─ Completed: 324.980374ms
└─ Errors:    0
```
#### posts25k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.040907ms
├─ Worst:     15.513725ms
├─ Completed: 62.047702ms
└─ Errors:    0
```
#### posts25k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      886.504µs
├─ Worst:     13.539993ms
├─ Completed: 55.760418ms
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      6.685694ms
├─ Worst:     200.682045ms
├─ Completed: 487.740588ms
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.301394ms
├─ Worst:     19.88404ms
├─ Completed: 63.104969ms
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.008586ms
├─ Worst:     16.476191ms
├─ Completed: 63.774306ms
└─ Errors:    0
```
#### posts25k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      7.386553ms
├─ Worst:     115.231371ms
├─ Completed: 432.78219ms
└─ Errors:    0
```
#### posts25k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      6.621173ms
├─ Worst:     123.18662ms
├─ Completed: 431.025584ms
└─ Errors:    0
```
#### posts25k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      993.586µs
├─ Worst:     26.507773ms
├─ Completed: 58.886139ms
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      151.011466ms
├─ Worst:     490.586661ms
├─ Completed: 3.107878864s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.185511ms
├─ Worst:     13.377629ms
├─ Completed: 62.37899ms
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      64.850935ms
├─ Worst:     279.976398ms
├─ Completed: 1.469370442s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      2.166176ms
├─ Worst:     13.38619ms
├─ Completed: 56.691442ms
└─ Errors:    0
```
#### posts25k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      49.867703ms
├─ Worst:     292.272358ms
├─ Completed: 1.440758894s
└─ Errors:    0
```
#### posts25k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.310954ms
├─ Worst:     12.764213ms
├─ Completed: 57.517704ms
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      1.308953929s
├─ Worst:     1.766705213s
├─ Completed: 16.142675968s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      2.927117ms
├─ Worst:     37.543821ms
├─ Completed: 150.802061ms
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      1.915900752s
├─ Worst:     2.416770523s
├─ Completed: 21.760979763s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.486518ms
├─ Worst:     35.832737ms
├─ Completed: 84.58361ms
└─ Errors:    0
```
#### posts50k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.858755ms
├─ Worst:     12.587889ms
├─ Completed: 4.007554858s
└─ Errors:    0
```
#### posts50k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      379.477438ms
├─ Worst:     1.707720551s
├─ Completed: 1.722854226s
└─ Errors:    0
```
#### posts50k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      147.29773ms
├─ Worst:     563.162598ms
├─ Completed: 566.029953ms
└─ Errors:    0
```
#### posts50k - mixed read and write (simpleA list with additional 300 concurrent random posts50k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      415.336614ms
├─ Worst:     1.891779042s
├─ Completed: 1.895262853s
└─ Errors:    0
```
#### posts50k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      4.001104ms
├─ Worst:     58.827737ms
├─ Completed: 232.572278ms
└─ Errors:    0
```
#### posts50k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      4.351994ms
├─ Worst:     60.579663ms
├─ Completed: 252.652483ms
└─ Errors:    0
```
#### posts50k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      5.070733ms
├─ Worst:     69.554258ms
├─ Completed: 280.692456ms
└─ Errors:    0
```
#### posts50k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      5.602667ms
├─ Worst:     92.420936ms
├─ Completed: 349.684779ms
└─ Errors:    0
```
#### posts50k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      3.248685ms
├─ Worst:     33.962208ms
├─ Completed: 169.86988ms
└─ Errors:    0
```
#### posts50k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      41.471083ms
├─ Worst:     287.396831ms
├─ Completed: 1.21637163s
└─ Errors:    0
```
#### posts50k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      999.386µs
├─ Worst:     65.563713ms
├─ Completed: 124.148245ms
└─ Errors:    0
```
#### posts50k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      904.304µs
├─ Worst:     11.939672ms
├─ Completed: 49.650218ms
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      46.798503ms
├─ Worst:     345.183061ms
├─ Completed: 1.209917101s
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.210232ms
├─ Worst:     48.896478ms
├─ Completed: 90.076234ms
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      961.105µs
├─ Worst:     24.158792ms
├─ Completed: 77.751752ms
└─ Errors:    0
```
#### posts50k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      56.096466ms
├─ Worst:     267.317186ms
├─ Completed: 1.292633583s
└─ Errors:    0
```
#### posts50k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      24.718246ms
├─ Worst:     254.147162ms
├─ Completed: 760.555317ms
└─ Errors:    0
```
#### posts50k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      998.706µs
├─ Worst:     26.882942ms
├─ Completed: 62.627157ms
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      407.272044ms
├─ Worst:     1.040034981s
├─ Completed: 7.1023335s
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.15955ms
├─ Worst:     21.637165ms
├─ Completed: 72.30993ms
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      154.174349ms
├─ Worst:     542.634661ms
├─ Completed: 3.166379513s
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      993.266µs
├─ Worst:     11.389698ms
├─ Completed: 58.776656ms
└─ Errors:    0
```
#### posts50k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      219.530137ms
├─ Worst:     458.652787ms
├─ Completed: 3.254749383s
└─ Errors:    0
```
#### posts50k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.032507ms
├─ Worst:     15.121515ms
├─ Completed: 55.844699ms
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      2.914177842s
├─ Worst:     3.547793321s
├─ Completed: 33.093810187s
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      3.06448ms
├─ Worst:     47.107111ms
├─ Completed: 155.35426ms
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      4.257713356s
├─ Worst:     4.587313209s
├─ Completed: 44.040574041s
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.845408ms
├─ Worst:     30.58996ms
├─ Completed: 90.518206ms
└─ Errors:    0
```
#### posts100k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      5.061332ms
├─ Worst:     13.729919ms
├─ Completed: 6.81384368s
└─ Errors:    0
```
#### posts100k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      206.483916ms
├─ Worst:     3.265234176s
├─ Completed: 3.275893456s
└─ Errors:    0
```
#### posts100k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      174.052029ms
├─ Worst:     596.974762ms
├─ Completed: 602.520947ms
└─ Errors:    0
```
#### posts100k - mixed read and write (simpleA list with additional 300 concurrent random posts100k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      444.420015ms
├─ Worst:     3.097574475s
├─ Completed: 3.110184885s
└─ Errors:    0
```
#### posts100k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      6.423888ms
├─ Worst:     73.226994ms
├─ Completed: 346.958507ms
└─ Errors:    0
```
#### posts100k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      6.600092ms
├─ Worst:     85.027262ms
├─ Completed: 365.222985ms
└─ Errors:    0
```
#### posts100k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      7.413874ms
├─ Worst:     109.397419ms
├─ Completed: 414.800401ms
└─ Errors:    0
```
#### posts100k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      15.768052ms
├─ Worst:     116.048593ms
├─ Completed: 474.932253ms
└─ Errors:    0
```
#### posts100k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      5.36286ms
├─ Worst:     79.164149ms
├─ Completed: 273.663152ms
└─ Errors:    0
```
#### posts100k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      89.130489ms
├─ Worst:     356.935129ms
├─ Completed: 2.112138321s
└─ Errors:    0
```
#### posts100k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.100548ms
├─ Worst:     55.177242ms
├─ Completed: 98.805983ms
└─ Errors:    0
```
#### posts100k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.296473ms
├─ Worst:     13.36523ms
├─ Completed: 55.233483ms
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      132.335459ms
├─ Worst:     296.084858ms
├─ Completed: 2.109467211s
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.110029ms
├─ Worst:     51.397103ms
├─ Completed: 90.106675ms
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      934.345µs
├─ Worst:     40.508858ms
├─ Completed: 74.314702ms
└─ Errors:    0
```
#### posts100k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      107.559171ms
├─ Worst:     371.040537ms
├─ Completed: 2.378354279s
└─ Errors:    0
```
#### posts100k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      138.088889ms
├─ Worst:     283.027237ms
├─ Completed: 2.02948756s
└─ Errors:    0
```
#### posts100k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      959.065µs
├─ Worst:     46.023123ms
├─ Completed: 78.353967ms
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      1.279204872s
├─ Worst:     1.795144316s
├─ Completed: 15.465665275s
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.471599ms
├─ Worst:     23.934706ms
├─ Completed: 80.079293ms
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      472.944401ms
├─ Worst:     715.950791ms
├─ Completed: 6.129388472s
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      993.266µs
├─ Worst:     14.351215ms
├─ Completed: 56.805644ms
└─ Errors:    0
```
#### posts100k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      380.554146ms
├─ Worst:     754.481678ms
├─ Completed: 6.521493279s
└─ Errors:    0
```
#### posts100k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.252953ms
├─ Worst:     16.516471ms
├─ Completed: 57.246016ms
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      6.367730578s
├─ Worst:     6.999949384s
├─ Completed: 1m8.401830439s
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      9.191326ms
├─ Worst:     71.93321ms
├─ Completed: 249.888625ms
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      8.444086291s
├─ Worst:     8.93637826s
├─ Completed: 1m27.932536234s
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      4.319996ms
├─ Worst:     21.044205ms
├─ Completed: 102.637674ms
└─ Errors:    0
```

## Go vs JS route execution
#### JS route (high concurrency) [reqs:500, conc:500]
```
┌─ Best:      213.851418ms
├─ Worst:     5.684440604s
├─ Completed: 5.687974779s
└─ Errors:    0
```
#### Go route (high concurrency) [reqs:500, conc:500]
```
┌─ Best:      319.787301ms
├─ Worst:     5.140661734s
├─ Completed: 5.147648401s
└─ Errors:    0
```
#### JS route (medium concurrency) [reqs:500, conc:50]
```
┌─ Best:      79.263366ms
├─ Worst:     1.072515177s
├─ Completed: 5.39920595s
└─ Errors:    0
```
#### Go route (medium concurrency) [reqs:500, conc:50]
```
┌─ Best:      77.454278ms
├─ Worst:     866.811258ms
├─ Completed: 4.982427607s
└─ Errors:    0
```
#### JS route (no concurrency) [reqs:500, conc:1]
```
┌─ Best:      10.462601ms
├─ Worst:     18.401973ms
├─ Completed: 6.14978181s
└─ Errors:    0
```
#### Go route (no concurrency) [reqs:500, conc:1]
```
┌─ Best:      10.323677ms
├─ Worst:     17.14022ms
├─ Completed: 6.074233903s
└─ Errors:    0
```

## Go vs JS hooks execution
#### JS OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]
```
┌─ Best:      785.221µs
├─ Worst:     12.069483ms
├─ Completed: 42.441179ms
└─ Errors:    0
```
#### Go OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]
```
┌─ Best:      783.821µs
├─ Worst:     11.576951ms
├─ Completed: 38.807921ms
└─ Errors:    0
```

## Deleting records
#### deleting 100 posts10k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      477.293µs
├─ Worst:     7.814849ms
├─ Completed: 29.43847ms
└─ Errors:    0
```
#### deleting 100 posts10k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      610.536µs
├─ Worst:     24.890467ms
├─ Completed: 51.314857ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      663.498µs
├─ Worst:     9.592977ms
├─ Completed: 40.326522ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      610.577µs
├─ Worst:     5.59627ms
├─ Completed: 29.82004ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      589.056µs
├─ Worst:     6.263888ms
├─ Completed: 27.313253ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      747.541µs
├─ Worst:     9.770582ms
├─ Completed: 33.90663ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      539.055µs
├─ Worst:     39.16757ms
├─ Completed: 65.059386ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      809.982µs
├─ Worst:     6.447613ms
├─ Completed: 29.945324ms
└─ Errors:    0
```
#### deleting 100 users - with cascade deleting all associated posts [conc:10, rule:`""`]
```
┌─ Best:      102.904321ms
├─ Worst:     4.306383508s
├─ Completed: 11.756722294s
└─ Errors:    0
```
#### deleting 100 organizations - with cascade deleting all users and associated posts [conc:10, rule:`""`]
```
┌─ Best:      294.308337ms
├─ Worst:     18.202586809s
├─ Completed: 29.539645163s
└─ Errors:    0
```
