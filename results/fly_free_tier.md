## Creating organizations (100)
#### Creating 50 organizations [reqs:50, conc:10, rule:`""`]
```
┌─ Best:      5.30762ms
├─ Worst:     24.090968ms
├─ Completed: 64.147246ms
└─ Errors:    0
```
#### Creating 50 organizations [reqs:50, conc:10, rule:`"@request.data.name != ''"`]
```
┌─ Best:      3.2291ms
├─ Worst:     16.942609ms
├─ Completed: 51.122717ms
└─ Errors:    0
```

## Creating permissions (50)
#### Creating 25 permissions [reqs:25, conc:5, rule:`""`]
```
┌─ Best:      1.97615ms
├─ Worst:     10.90429ms
├─ Completed: 33.531238ms
└─ Errors:    0
```
#### Creating 25 permissions [reqs:25, conc:5, rule:`"@request.data.name != ''"`]
```
┌─ Best:      3.91094ms
├─ Worst:     8.703899ms
├─ Completed: 27.065068ms
└─ Errors:    0
```

## Creating users (500 - expected to be slow due to passwordHash generation)
#### Creating 250 users [reqs:250, conc:50, rule:`""`]
```
┌─ Best:      22.729618694s
├─ Worst:     29.315346221s
├─ Completed: 2m13.25037132s
└─ Errors:    0
```
#### Creating 250 users [reqs:250, conc:50, rule:`"@request.data.email != '' && @request.data.permissions:length > 0"`]
```
┌─ Best:      20.196515766s
├─ Worst:     30.183112059s
├─ Completed: 2m13.190505244s
└─ Errors:    0
```

## Creating posts (10k, 25k, 50k, 100k)
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`""`]
```
┌─ Best:      146.505441ms
├─ Worst:     2.0124111s
├─ Completed: 7.575088748s
└─ Errors:    0
```
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      103.805414ms
├─ Worst:     2.467506043s
├─ Completed: 9.370156682s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`""`]
```
┌─ Best:      104.038504ms
├─ Worst:     2.079795815s
├─ Completed: 18.674421056s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      122.697172ms
├─ Worst:     2.660540161s
├─ Completed: 23.756208043s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`""`]
```
┌─ Best:      96.359624ms
├─ Worst:     2.804224873s
├─ Completed: 37.301102315s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      133.709512ms
├─ Worst:     3.332385971s
├─ Completed: 47.151979057s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`""`]
```
┌─ Best:      25.885878ms
├─ Worst:     2.554985667s
├─ Completed: 1m18.367738785s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      102.149374ms
├─ Worst:     3.248287306s
├─ Completed: 1m36.457011126s
└─ Errors:    0
```

## Search records
#### posts10k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      1.83847ms
├─ Worst:     8.846719ms
├─ Completed: 2.298487912s
└─ Errors:    0
```
#### posts10k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      182.884819ms
├─ Worst:     2.358694609s
├─ Completed: 2.361265939s
└─ Errors:    0
```
#### posts10k - mixed read and write (simpleA list with 300 concurrent random posts10k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      292.430872ms
├─ Worst:     6.594547697s
├─ Completed: 6.622903395s
└─ Errors:    0
```
#### posts10k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      14.422999ms
├─ Worst:     346.2293ms
├─ Completed: 1.308052442s
└─ Errors:    0
```
#### posts10k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      15.111699ms
├─ Worst:     96.412025ms
├─ Completed: 601.099714ms
└─ Errors:    0
```
#### posts10k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      16.681369ms
├─ Worst:     139.932242ms
├─ Completed: 607.761254ms
└─ Errors:    0
```
#### posts10k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      21.559038ms
├─ Worst:     121.668252ms
├─ Completed: 657.59907ms
└─ Errors:    0
```
#### posts10k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      10.847879ms
├─ Worst:     39.353428ms
├─ Completed: 208.580958ms
└─ Errors:    0
```
#### posts10k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      10.11558ms
├─ Worst:     306.797361ms
├─ Completed: 747.317815ms
└─ Errors:    0
```
#### posts10k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      4.08488ms
├─ Worst:     45.673787ms
├─ Completed: 176.627829ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      34.299798ms
├─ Worst:     408.157576ms
├─ Completed: 2.158887751s
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      5.51203ms
├─ Worst:     54.389027ms
├─ Completed: 217.276837ms
└─ Errors:    0
```
#### posts10k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      17.980039ms
├─ Worst:     151.890411ms
├─ Completed: 920.254255ms
└─ Errors:    0
```
#### posts10k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      41.630307ms
├─ Worst:     143.144341ms
├─ Completed: 748.222904ms
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      206.905128ms
├─ Worst:     919.851085ms
├─ Completed: 5.02183333s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      36.512028ms
├─ Worst:     319.094491ms
├─ Completed: 1.68642262s
└─ Errors:    0
```
#### posts10k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      110.973444ms
├─ Worst:     485.600011ms
├─ Completed: 2.813985362s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      621.909403ms
├─ Worst:     2.003569311s
├─ Completed: 13.432806909s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      2.947073514s
├─ Worst:     4.208830449s
├─ Completed: 36.651123964s
└─ Errors:    0
```
#### posts25k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.87965ms
├─ Worst:     19.372179ms
├─ Completed: 3.651000421s
└─ Errors:    0
```
#### posts25k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      1.027673149s
├─ Worst:     9.191014202s
├─ Completed: 9.206168801s
└─ Errors:    0
```
#### posts25k - mixed read and write (simpleA list with 300 concurrent random posts25k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      1.097943263s
├─ Worst:     13.607030378s
├─ Completed: 13.676249294s
└─ Errors:    0
```
#### posts25k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      71.248495ms
├─ Worst:     465.984283ms
├─ Completed: 1.399847806s
└─ Errors:    0
```
#### posts25k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      22.137899ms
├─ Worst:     124.031603ms
├─ Completed: 737.049266ms
└─ Errors:    0
```
#### posts25k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      17.811899ms
├─ Worst:     147.843311ms
├─ Completed: 742.925145ms
└─ Errors:    0
```
#### posts25k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      13.890559ms
├─ Worst:     173.047389ms
├─ Completed: 823.872331ms
└─ Errors:    0
```
#### posts25k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      10.055729ms
├─ Worst:     57.465596ms
├─ Completed: 324.08067ms
└─ Errors:    0
```
#### posts25k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      72.774016ms
├─ Worst:     1.392553966s
├─ Completed: 3.116473303s
└─ Errors:    0
```
#### posts25k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      6.87322ms
├─ Worst:     65.353636ms
├─ Completed: 283.314183ms
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      89.614045ms
├─ Worst:     1.266820455s
├─ Completed: 4.69782101s
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      6.624649ms
├─ Worst:     68.788316ms
├─ Completed: 263.135224ms
└─ Errors:    0
```
#### posts25k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      103.842103ms
├─ Worst:     640.641101ms
├─ Completed: 2.631580283s
└─ Errors:    0
```
#### posts25k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      65.980126ms
├─ Worst:     333.28954ms
├─ Completed: 2.121103423s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      806.962802ms
├─ Worst:     2.900251528s
├─ Completed: 15.414402741s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      159.025311ms
├─ Worst:     737.815146ms
├─ Completed: 4.251034696s
└─ Errors:    0
```
#### posts25k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      384.399827ms
├─ Worst:     2.3502315s
├─ Completed: 11.480474086s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      1.548625137s
├─ Worst:     4.432553055s
├─ Completed: 34.591187616s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      6.976880434s
├─ Worst:     10.391837179s
├─ Completed: 1m30.468687342s
└─ Errors:    0
```
#### posts50k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      5.01712ms
├─ Worst:     43.792448ms
├─ Completed: 5.996207022s
└─ Errors:    0
```
#### posts50k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      1.3446655s
├─ Worst:     27.227335205s
├─ Completed: 27.242292464s
└─ Errors:    0
```
#### posts50k - mixed read and write (simpleA list with 300 concurrent random posts50k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      1.888016838s
├─ Worst:     30.763229245s
├─ Completed: 30.885333338s
└─ Errors:    0
```
#### posts50k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      43.674907ms
├─ Worst:     704.539468ms
├─ Completed: 1.905625926s
└─ Errors:    0
```
#### posts50k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      21.295408ms
├─ Worst:     173.090559ms
├─ Completed: 941.731613ms
└─ Errors:    0
```
#### posts50k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      18.299289ms
├─ Worst:     167.1327ms
├─ Completed: 940.208224ms
└─ Errors:    0
```
#### posts50k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      44.637977ms
├─ Worst:     176.13865ms
├─ Completed: 1.057369517s
└─ Errors:    0
```
#### posts50k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      7.35769ms
├─ Worst:     95.455154ms
├─ Completed: 593.537354ms
└─ Errors:    0
```
#### posts50k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      210.902738ms
├─ Worst:     2.312385782s
├─ Completed: 7.252535648s
└─ Errors:    0
```
#### posts50k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      7.97178ms
├─ Worst:     84.217085ms
├─ Completed: 294.094583ms
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      347.297569ms
├─ Worst:     2.075561126s
├─ Completed: 8.772255467s
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      9.533719ms
├─ Worst:     126.930602ms
├─ Completed: 347.154048ms
└─ Errors:    0
```
#### posts50k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      437.467644ms
├─ Worst:     1.478749542s
├─ Completed: 8.071212709s
└─ Errors:    0
```
#### posts50k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      200.255378ms
├─ Worst:     972.996332ms
├─ Completed: 5.346640661s
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      1.50287929s
├─ Worst:     6.375945279s
├─ Completed: 36.615077166s
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      236.396436ms
├─ Worst:     2.3483813s
├─ Completed: 11.997042354s
└─ Errors:    0
```
#### posts50k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      563.422946ms
├─ Worst:     4.512212851s
├─ Completed: 28.467320842s
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      5.822087773s
├─ Worst:     9.032644712s
├─ Completed: 1m14.509751034s
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      17.115728279s
├─ Worst:     20.135512079s
├─ Completed: 3m5.346883342s
└─ Errors:    0
```
#### posts100k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      8.811189ms
├─ Worst:     51.595766ms
├─ Completed: 10.604332207s
└─ Errors:    0
```
#### posts100k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      1.159538891s
├─ Worst:     1m1.778535355s
├─ Completed: 1m1.792103434s
└─ Errors:    186
```
#### posts100k - mixed read and write (simpleA list with 300 concurrent random posts100k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

---

ERROR: Benchmark run interrupted...

---

## Deleting records
#### deleting 100 posts10k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      4.348709ms
├─ Worst:     39.956648ms
├─ Completed: 221.240197ms
└─ Errors:    0
```
#### deleting 100 posts10k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      4.70196ms
├─ Worst:     60.355286ms
├─ Completed: 237.585546ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      3.998119ms
├─ Worst:     24.407438ms
├─ Completed: 167.35884ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      12.482419ms
├─ Worst:     89.654015ms
├─ Completed: 318.216391ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      7.23102ms
├─ Worst:     37.857007ms
├─ Completed: 207.286667ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      7.07949ms
├─ Worst:     46.066757ms
├─ Completed: 219.886357ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      4.7814ms
├─ Worst:     31.215768ms
├─ Completed: 154.099911ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      4.85479ms
├─ Worst:     28.638939ms
├─ Completed: 159.967491ms
└─ Errors:    0
```
#### deleting 100 users - with cascade deleting all associated posts [conc:10, rule:`""`]
```
┌─ Best:      1.528596928s
├─ Worst:     1m9.309482194s
├─ Completed: 2m14.479974616s
└─ Errors:    0
```
#### deleting 100 organizations - with cascade deleting all users and associated posts [conc:10, rule:`""`]
```
┌─ Best:      1.34219032s
├─ Worst:     2m26.064052325s
├─ Completed: 4m23.945224022s
└─ Errors:    0
```
