_PocketBase v0.28.4_

_Railway Hobby Tier Configuration_

_CPU: 8 cores_
_Memory: 8GiBs_

_Actual usage:_

_CPU: Max 8 cores_

_Memory: Max 1GiBs_

## Creating organizations (100)

#### Creating 50 organizations [reqs:50, conc:10, rule:`""`]

```
┌─ Best:      482.635µs
├─ Worst:     4.920777ms
├─ Completed: 10.208344ms
└─ Errors:    0
```

#### Creating 50 organizations [reqs:50, conc:10, rule:`"@request.body.name != ''"`]

```
┌─ Best:      686.61µs
├─ Worst:     10.122601ms
├─ Completed: 14.7017ms
└─ Errors:    0
```

## Creating permissions (50)

#### Creating 25 permissions [reqs:25, conc:5, rule:`""`]

```
┌─ Best:      480.645µs
├─ Worst:     2.15987ms
├─ Completed: 5.558993ms
└─ Errors:    0
```

#### Creating 25 permissions [reqs:25, conc:5, rule:`"@request.body.name != ''"`]

```
┌─ Best:      670.999µs
├─ Worst:     1.564542ms
├─ Completed: 5.894937ms
└─ Errors:    0
```

## Creating users (500 - expected to be slow due to passwordHash generation)

#### Creating 250 users [reqs:250, conc:50, rule:`""`]

```
┌─ Best:      448.403162ms
├─ Worst:     1.299250477s
├─ Completed: 4.883065864s
└─ Errors:    0
```

#### Creating 250 users [reqs:250, conc:50, rule:`"@request.body.email != '' && @request.body.permissions:length > 0"`]

```
┌─ Best:      517.046841ms
├─ Worst:     1.41725161s
├─ Completed: 4.941040546s
└─ Errors:    0
```

## Creating posts (10k, 25k, 50k, 100k)

#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`""`]

```
┌─ Best:      839.874µs
├─ Worst:     933.395322ms
├─ Completed: 1.309496708s
└─ Errors:    0
```

#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]

```
┌─ Best:      1.063992ms
├─ Worst:     981.161023ms
├─ Completed: 1.376179141s
└─ Errors:    0
```

#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`""`]

```
┌─ Best:      732.911µs
├─ Worst:     1.006433793s
├─ Completed: 3.154084163s
└─ Errors:    0
```

#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]

```
┌─ Best:      1.044617ms
├─ Worst:     1.109628034s
├─ Completed: 3.41755219s
└─ Errors:    0
```

#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`""`]

```
┌─ Best:      731.733µs
├─ Worst:     1.116617024s
├─ Completed: 5.547888241s
└─ Errors:    0
```

#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]

```
┌─ Best:      1.048511ms
├─ Worst:     1.269509416s
├─ Completed: 6.793140112s
└─ Errors:    0
```

#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`""`]

```
┌─ Best:      701.618µs
├─ Worst:     1.209695668s
├─ Completed: 11.437265335s
└─ Errors:    0
```

#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`"@request.auth.id != '' && @request.body.public:isset = true"`]

```
┌─ Best:      969.405µs
├─ Worst:     1.269671715s
├─ Completed: 13.033361473s
└─ Errors:    0
```

## User auth with password (expected to be slow due to passwordHash verification)

#### users auth with email/pass - high concurrency [reqs:250, conc:250]

```
┌─ Best:      292.924476ms
├─ Worst:     2.428075099s
├─ Completed: 2.429228189s
└─ Errors:    0
```

#### users auth with email/pass - small concurrency [reqs:250, conc:10]

```
┌─ Best:      74.462132ms
├─ Worst:     115.671327ms
├─ Completed: 2.367791657s
└─ Errors:    0
```

## User auth refresh

#### users - auth refresh (high concurrency) [reqs:1000, conc:1000]

```
┌─ Best:      12.411337ms
├─ Worst:     267.32315ms
├─ Completed: 269.492168ms
└─ Errors:    0
```

#### users - auth refresh (medium concurrency) [reqs:1000, conc:100]

```
┌─ Best:      420.255µs
├─ Worst:     229.389828ms
├─ Completed: 230.600778ms
└─ Errors:    0
```

## List records

#### users - getOne for auth refresh comparison (high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`/ozl9aa9ro38s17h`]

```
┌─ Best:      4.945779ms
├─ Worst:     319.854206ms
├─ Completed: 320.841954ms
└─ Errors:    0
```

#### users - getOne for auth refresh comparison (medium concurrency) [reqs:1000, conc:100, rule:`""`, query:`/ozl9aa9ro38s17h`]

```
┌─ Best:      516.761µs
├─ Worst:     264.957545ms
├─ Completed: 266.558382ms
└─ Errors:    0
```

#### posts10k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      1.307563ms
├─ Worst:     6.633886ms
├─ Completed: 1.956517408s
└─ Errors:    0
```

#### posts10k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      10.713846ms
├─ Worst:     594.576243ms
├─ Completed: 596.491433ms
└─ Errors:    0
```

#### posts10k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      76.582674ms
├─ Worst:     417.309646ms
├─ Completed: 418.406414ms
└─ Errors:    0
```

#### posts10k - mixed read and write (simpleA list with additional 300 concurrent random posts10k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      12.597968ms
├─ Worst:     566.205689ms
├─ Completed: 567.303759ms
└─ Errors:    0
```

#### posts10k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]

```
┌─ Best:      3.171799ms
├─ Worst:     13.65483ms
├─ Completed: 77.217729ms
└─ Errors:    0
```

#### posts10k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]

```
┌─ Best:      2.899438ms
├─ Worst:     23.43073ms
├─ Completed: 68.151445ms
└─ Errors:    0
```

#### posts10k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]

```
┌─ Best:      3.479886ms
├─ Worst:     20.864367ms
├─ Completed: 86.226367ms
└─ Errors:    0
```

#### posts10k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]

```
┌─ Best:      4.383076ms
├─ Worst:     30.430483ms
├─ Completed: 122.375482ms
└─ Errors:    0
```

#### posts10k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]

```
┌─ Best:      1.629373ms
├─ Worst:     15.607162ms
├─ Completed: 41.753166ms
└─ Errors:    0
```

#### posts10k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      31.616188ms
├─ Worst:     278.402426ms
├─ Completed: 635.766447ms
└─ Errors:    0
```

#### posts10k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      1.049482ms
├─ Worst:     78.980791ms
├─ Completed: 97.430825ms
└─ Errors:    0
```

#### posts10k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.049036ms
├─ Worst:     15.018288ms
├─ Completed: 44.624682ms
└─ Errors:    0
```

#### posts10k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      30.489972ms
├─ Worst:     240.534734ms
├─ Completed: 642.618131ms
└─ Errors:    0
```

#### posts10k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      1.128708ms
├─ Worst:     101.076993ms
├─ Completed: 121.570381ms
└─ Errors:    0
```

#### posts10k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.100987ms
├─ Worst:     14.315547ms
├─ Completed: 37.731487ms
└─ Errors:    0
```

#### posts10k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      12.511405ms
├─ Worst:     161.676227ms
├─ Completed: 420.850318ms
└─ Errors:    0
```

#### posts10k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      12.179948ms
├─ Worst:     131.055956ms
├─ Completed: 338.796601ms
└─ Errors:    0
```

#### posts10k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.01627ms
├─ Worst:     30.498971ms
├─ Completed: 54.633376ms
└─ Errors:    0
```

#### posts10k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]

```
┌─ Best:      210.553456ms
├─ Worst:     581.001521ms
├─ Completed: 3.719270645s
└─ Errors:    0
```

#### posts10k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.586302ms
├─ Worst:     51.623269ms
├─ Completed: 86.32977ms
└─ Errors:    0
```

#### posts10k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]

```
┌─ Best:      67.148862ms
├─ Worst:     325.472748ms
├─ Completed: 1.217460064s
└─ Errors:    0
```

#### posts10k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.277762ms
├─ Worst:     46.906166ms
├─ Completed: 74.87334ms
└─ Errors:    0
```

#### posts10k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]

```
┌─ Best:      187.640783ms
├─ Worst:     396.719439ms
├─ Completed: 3.009229472s
└─ Errors:    0
```

#### posts10k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.211941ms
├─ Worst:     51.463381ms
├─ Completed: 77.709568ms
└─ Errors:    0
```

#### posts10k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]

```
┌─ Best:      749.750414ms
├─ Worst:     1.159591322s
├─ Completed: 8.865796614s
└─ Errors:    0
```

#### posts10k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      6.634253ms
├─ Worst:     80.267505ms
├─ Completed: 178.108563ms
└─ Errors:    0
```

#### posts10k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]

```
┌─ Best:      531.618308ms
├─ Worst:     1.206275829s
├─ Completed: 6.858195792s
└─ Errors:    0
```

#### posts10k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      2.365851ms
├─ Worst:     50.919761ms
├─ Completed: 86.534524ms
└─ Errors:    0
```

#### posts25k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      2.019829ms
├─ Worst:     7.705829ms
├─ Completed: 3.001926691s
└─ Errors:    0
```

#### posts25k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      41.983644ms
├─ Worst:     928.428011ms
├─ Completed: 931.12081ms
└─ Errors:    0
```

#### posts25k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      3.277785ms
├─ Worst:     420.499343ms
├─ Completed: 421.753406ms
└─ Errors:    0
```

#### posts25k - mixed read and write (simpleA list with additional 300 concurrent random posts25k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      28.3576ms
├─ Worst:     1.130164069s
├─ Completed: 1.131643385s
└─ Errors:    0
```

#### posts25k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]

```
┌─ Best:      5.020172ms
├─ Worst:     38.114671ms
├─ Completed: 127.487342ms
└─ Errors:    0
```

#### posts25k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]

```
┌─ Best:      4.197231ms
├─ Worst:     36.619301ms
├─ Completed: 106.570457ms
└─ Errors:    0
```

#### posts25k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]

```
┌─ Best:      4.395492ms
├─ Worst:     34.602015ms
├─ Completed: 120.290591ms
└─ Errors:    0
```

#### posts25k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]

```
┌─ Best:      6.069353ms
├─ Worst:     54.674898ms
├─ Completed: 166.356836ms
└─ Errors:    0
```

#### posts25k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]

```
┌─ Best:      2.843573ms
├─ Worst:     35.543764ms
├─ Completed: 83.436298ms
└─ Errors:    0
```

#### posts25k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      81.018207ms
├─ Worst:     559.229198ms
├─ Completed: 1.617921679s
└─ Errors:    0
```

#### posts25k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      1.081608ms
├─ Worst:     95.788884ms
├─ Completed: 136.741341ms
└─ Errors:    0
```

#### posts25k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.054226ms
├─ Worst:     22.705444ms
├─ Completed: 37.003163ms
└─ Errors:    0
```

#### posts25k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      60.651584ms
├─ Worst:     510.558727ms
├─ Completed: 1.530540449s
└─ Errors:    0
```

#### posts25k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      1.043462ms
├─ Worst:     129.170133ms
├─ Completed: 189.008938ms
└─ Errors:    0
```

#### posts25k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.060053ms
├─ Worst:     11.823944ms
├─ Completed: 36.437729ms
└─ Errors:    0
```

#### posts25k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      58.486214ms
├─ Worst:     344.200665ms
├─ Completed: 1.101118184s
└─ Errors:    0
```

#### posts25k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      42.340667ms
├─ Worst:     336.179987ms
├─ Completed: 891.942573ms
└─ Errors:    0
```

#### posts25k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.039042ms
├─ Worst:     78.365848ms
├─ Completed: 96.528911ms
└─ Errors:    0
```

#### posts25k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]

```
┌─ Best:      1.050692681s
├─ Worst:     1.621532207s
├─ Completed: 12.441948061s
└─ Errors:    0
```

#### posts25k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.640652ms
├─ Worst:     90.633049ms
├─ Completed: 120.903432ms
└─ Errors:    0
```

#### posts25k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]

```
┌─ Best:      139.083509ms
├─ Worst:     849.894695ms
├─ Completed: 3.161926931s
└─ Errors:    0
```

#### posts25k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.200951ms
├─ Worst:     101.488666ms
├─ Completed: 132.284126ms
└─ Errors:    0
```

#### posts25k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]

```
┌─ Best:      603.862371ms
├─ Worst:     1.122642253s
├─ Completed: 8.237433895s
└─ Errors:    0
```

#### posts25k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.239705ms
├─ Worst:     93.085754ms
├─ Completed: 114.816919ms
└─ Errors:    0
```

#### posts25k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]

```
┌─ Best:      2.139542994s
├─ Worst:     3.19858901s
├─ Completed: 25.322508509s
└─ Errors:    0
```

#### posts25k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      9.650555ms
├─ Worst:     197.865473ms
├─ Completed: 331.746978ms
└─ Errors:    0
```

#### posts25k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]

```
┌─ Best:      1.373190425s
├─ Worst:     2.642801447s
├─ Completed: 16.767806309s
└─ Errors:    0
```

#### posts25k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      2.425109ms
├─ Worst:     112.838857ms
├─ Completed: 151.541601ms
└─ Errors:    0
```

#### posts50k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      3.854452ms
├─ Worst:     15.541027ms
├─ Completed: 5.306787598s
└─ Errors:    0
```

#### posts50k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      38.089919ms
├─ Worst:     1.586631799s
├─ Completed: 1.588288982s
└─ Errors:    0
```

#### posts50k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      85.270701ms
├─ Worst:     441.024272ms
├─ Completed: 442.72215ms
└─ Errors:    0
```

#### posts50k - mixed read and write (simpleA list with additional 300 concurrent random posts50k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      106.30899ms
├─ Worst:     2.07055083s
├─ Completed: 2.07797134s
└─ Errors:    0
```

#### posts50k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]

```
┌─ Best:      5.941513ms
├─ Worst:     55.778715ms
├─ Completed: 179.844672ms
└─ Errors:    0
```

#### posts50k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]

```
┌─ Best:      7.550725ms
├─ Worst:     33.612389ms
├─ Completed: 152.415417ms
└─ Errors:    0
```

#### posts50k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]

```
┌─ Best:      7.313401ms
├─ Worst:     38.409684ms
├─ Completed: 178.103805ms
└─ Errors:    0
```

#### posts50k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]

```
┌─ Best:      6.883634ms
├─ Worst:     35.850896ms
├─ Completed: 178.30548ms
└─ Errors:    0
```

#### posts50k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]

```
┌─ Best:      5.826706ms
├─ Worst:     31.555393ms
├─ Completed: 116.306035ms
└─ Errors:    0
```

#### posts50k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      121.489051ms
├─ Worst:     645.037132ms
├─ Completed: 2.614913992s
└─ Errors:    0
```

#### posts50k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      1.033949ms
├─ Worst:     119.51638ms
├─ Completed: 141.253693ms
└─ Errors:    0
```

#### posts50k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.038784ms
├─ Worst:     15.659457ms
├─ Completed: 33.162733ms
└─ Errors:    0
```

#### posts50k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      141.939118ms
├─ Worst:     396.454929ms
├─ Completed: 2.417309999s
└─ Errors:    0
```

#### posts50k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      1.185761ms
├─ Worst:     133.206853ms
├─ Completed: 153.001501ms
└─ Errors:    0
```

#### posts50k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.07018ms
├─ Worst:     10.85953ms
├─ Completed: 29.45794ms
└─ Errors:    0
```

#### posts50k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      158.370533ms
├─ Worst:     289.845705ms
├─ Completed: 2.322445394s
└─ Errors:    0
```

#### posts50k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      99.309154ms
├─ Worst:     336.234731ms
├─ Completed: 1.719054594s
└─ Errors:    0
```

#### posts50k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      955.215µs
├─ Worst:     47.071093ms
├─ Completed: 63.606421ms
└─ Errors:    0
```

#### posts50k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]

```
┌─ Best:      2.401003137s
├─ Worst:     3.039045378s
├─ Completed: 27.347116055s
└─ Errors:    0
```

#### posts50k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.716658ms
├─ Worst:     180.504279ms
├─ Completed: 217.763389ms
└─ Errors:    0
```

#### posts50k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]

```
┌─ Best:      286.158707ms
├─ Worst:     1.026873699s
├─ Completed: 5.976731172s
└─ Errors:    0
```

#### posts50k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.304083ms
├─ Worst:     142.566653ms
├─ Completed: 176.52963ms
└─ Errors:    0
```

#### posts50k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]

```
┌─ Best:      1.445297129s
├─ Worst:     2.27333964s
├─ Completed: 19.082172161s
└─ Errors:    0
```

#### posts50k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.308207ms
├─ Worst:     176.705377ms
├─ Completed: 206.950798ms
└─ Errors:    0
```

#### posts50k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]

```
┌─ Best:      4.925791404s
├─ Worst:     6.380875143s
├─ Completed: 57.666650645s
└─ Errors:    0
```

#### posts50k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      11.529863ms
├─ Worst:     217.893936ms
├─ Completed: 363.152434ms
└─ Errors:    0
```

#### posts50k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]

```
┌─ Best:      2.749889714s
├─ Worst:     4.531582238s
├─ Completed: 33.218313856s
└─ Errors:    0
```

#### posts50k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      2.25933ms
├─ Worst:     164.232984ms
├─ Completed: 201.9622ms
└─ Errors:    0
```

#### posts100k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      6.757815ms
├─ Worst:     19.790926ms
├─ Completed: 9.826061295s
└─ Errors:    0
```

#### posts100k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      61.318572ms
├─ Worst:     3.562504521s
├─ Completed: 3.566206024s
└─ Errors:    0
```

#### posts100k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      120.055941ms
├─ Worst:     469.364203ms
├─ Completed: 470.00468ms
└─ Errors:    0
```

#### posts100k - mixed read and write (simpleA list with additional 300 concurrent random posts100k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

```
┌─ Best:      250.781838ms
├─ Worst:     3.714347338s
├─ Completed: 3.717081769s
└─ Errors:    0
```

#### posts100k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]

```
┌─ Best:      11.875381ms
├─ Worst:     95.139171ms
├─ Completed: 259.155089ms
└─ Errors:    0
```

#### posts100k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]

```
┌─ Best:      11.127901ms
├─ Worst:     89.376822ms
├─ Completed: 312.368512ms
└─ Errors:    0
```

#### posts100k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]

```
┌─ Best:      15.671355ms
├─ Worst:     89.694635ms
├─ Completed: 313.388546ms
└─ Errors:    0
```

#### posts100k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]

```
┌─ Best:      16.505175ms
├─ Worst:     90.849571ms
├─ Completed: 333.590443ms
└─ Errors:    0
```

#### posts100k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]

```
┌─ Best:      9.411387ms
├─ Worst:     95.018721ms
├─ Completed: 242.940137ms
└─ Errors:    0
```

#### posts100k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      298.2215ms
├─ Worst:     1.200754474s
├─ Completed: 5.049755045s
└─ Errors:    0
```

#### posts100k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      1.166512ms
├─ Worst:     53.054005ms
├─ Completed: 73.336832ms
└─ Errors:    0
```

#### posts100k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      990.373µs
├─ Worst:     21.207495ms
├─ Completed: 37.027781ms
└─ Errors:    0
```

#### posts100k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      361.592355ms
├─ Worst:     729.102688ms
├─ Completed: 5.08962521s
└─ Errors:    0
```

#### posts100k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]

```
┌─ Best:      1.08294ms
├─ Worst:     48.940865ms
├─ Completed: 73.471917ms
└─ Errors:    0
```

#### posts100k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.083766ms
├─ Worst:     19.062841ms
├─ Completed: 40.896821ms
└─ Errors:    0
```

#### posts100k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      350.443843ms
├─ Worst:     599.285746ms
├─ Completed: 4.485079348s
└─ Errors:    0
```

#### posts100k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]

```
┌─ Best:      222.633828ms
├─ Worst:     428.206026ms
├─ Completed: 3.3120999s
└─ Errors:    0
```

#### posts100k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      846.437µs
├─ Worst:     26.290382ms
├─ Completed: 54.077362ms
└─ Errors:    0
```

#### posts100k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]

```
┌─ Best:      5.253579815s
├─ Worst:     6.234286021s
├─ Completed: 58.297950279s
└─ Errors:    0
```

#### posts100k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.706805ms
├─ Worst:     220.536155ms
├─ Completed: 257.620323ms
└─ Errors:    0
```

#### posts100k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]

```
┌─ Best:      763.54904ms
├─ Worst:     1.397437626s
├─ Completed: 10.659772672s
└─ Errors:    0
```

#### posts100k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.270796ms
├─ Worst:     127.601333ms
├─ Completed: 151.68123ms
└─ Errors:    0
```

#### posts100k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]

```
┌─ Best:      2.66274049s
├─ Worst:     4.225892606s
├─ Completed: 36.244929912s
└─ Errors:    0
```

#### posts100k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      1.376921ms
├─ Worst:     175.03562ms
├─ Completed: 206.120167ms
└─ Errors:    0
```

#### posts100k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]

```
┌─ Best:      10.723482455s
├─ Worst:     13.671957537s
├─ Completed: 2m0.73805432s
└─ Errors:    0
```

#### posts100k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      10.048743ms
├─ Worst:     262.686627ms
├─ Completed: 405.751292ms
└─ Errors:    0
```

#### posts100k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]

```
┌─ Best:      5.793550303s
├─ Worst:     6.970289776s
├─ Completed: 1m3.959203816s
└─ Errors:    0
```

#### posts100k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]

```
┌─ Best:      2.38684ms
├─ Worst:     180.4251ms
├─ Completed: 225.614944ms
└─ Errors:    0
```

## Go vs JS route execution

#### JS route (high concurrency) [reqs:500, conc:500]

```
┌─ Best:      2.326803305s
├─ Worst:     8.968689451s
├─ Completed: 8.969707613s
└─ Errors:    0
```

#### Go route (high concurrency) [reqs:500, conc:500]

```
┌─ Best:      1.231658508s
├─ Worst:     8.001293332s
├─ Completed: 8.008423123s
└─ Errors:    0
```

#### JS route (medium concurrency) [reqs:500, conc:50]

```
┌─ Best:      313.553601ms
├─ Worst:     725.508399ms
├─ Completed: 5.249022786s
└─ Errors:    0
```

#### Go route (medium concurrency) [reqs:500, conc:50]

```
┌─ Best:      328.489169ms
├─ Worst:     741.248261ms
├─ Completed: 5.401618658s
└─ Errors:    0
```

#### JS route (no concurrency) [reqs:500, conc:1]

```
┌─ Best:      11.804224ms
├─ Worst:     30.890173ms
├─ Completed: 7.903498721s
└─ Errors:    0
```

#### Go route (no concurrency) [reqs:500, conc:1]

```
┌─ Best:      11.881209ms
├─ Worst:     27.510127ms
├─ Completed: 7.798411011s
└─ Errors:    0
```

## Go vs JS hooks execution

#### JS OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]

```
┌─ Best:      1.362342ms
├─ Worst:     119.963279ms
├─ Completed: 137.282978ms
└─ Errors:    0
```

#### Go OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]

```
┌─ Best:      1.345835ms
├─ Worst:     18.551225ms
├─ Completed: 44.628002ms
└─ Errors:    0
```

## Deleting records

#### deleting 100 posts10k - simple (no cascade, no rule) [conc:10, rule:`""`]

```
┌─ Best:      1.198982ms
├─ Worst:     19.097674ms
├─ Completed: 43.584159ms
└─ Errors:    0
```

#### deleting 100 posts10k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]

```
┌─ Best:      1.067387ms
├─ Worst:     15.890757ms
├─ Completed: 34.147516ms
└─ Errors:    0
```

#### deleting 100 posts25k - simple (no cascade, no rule) [conc:10, rule:`""`]

```
┌─ Best:      812.828µs
├─ Worst:     7.689085ms
├─ Completed: 24.573058ms
└─ Errors:    0
```

#### deleting 100 posts25k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]

```
┌─ Best:      647.688µs
├─ Worst:     5.948505ms
├─ Completed: 22.702355ms
└─ Errors:    0
```

#### deleting 100 posts50k - simple (no cascade, no rule) [conc:10, rule:`""`]

```
┌─ Best:      991.294µs
├─ Worst:     12.809638ms
├─ Completed: 30.111334ms
└─ Errors:    0
```

#### deleting 100 posts50k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]

```
┌─ Best:      898.355µs
├─ Worst:     6.34087ms
├─ Completed: 20.771464ms
└─ Errors:    0
```

#### deleting 100 posts100k - simple (no cascade, no rule) [conc:10, rule:`""`]

```
┌─ Best:      1.065754ms
├─ Worst:     20.2072ms
├─ Completed: 41.352171ms
└─ Errors:    0
```

#### deleting 100 posts100k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]

```
┌─ Best:      845.671µs
├─ Worst:     6.195983ms
├─ Completed: 23.988194ms
└─ Errors:    0
```

#### deleting 100 users - with cascade deleting all associated posts [conc:10, rule:`""`]

```
┌─ Best:      124.296395ms
├─ Worst:     6.503176765s
├─ Completed: 13.497309818s
└─ Errors:    0
```

#### deleting 100 organizations - with cascade deleting all users and associated posts [conc:10, rule:`""`]

```
┌─ Best:      151.476058ms
├─ Worst:     15.196593059s
├─ Completed: 31.825532901s
└─ Errors:    0
```

---

Completed!
