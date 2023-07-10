## Creating organizations (100)
#### Creating 50 organizations [reqs:50, conc:10, rule:`""`]
```
┌─ Best:      2.441289ms
├─ Worst:     33.707965ms
├─ Completed: 63.274674ms
└─ Errors:    0
```
#### Creating 50 organizations [reqs:50, conc:10, rule:`"@request.data.name != ''"`]
```
┌─ Best:      2.425329ms
├─ Worst:     23.945129ms
├─ Completed: 46.096771ms
└─ Errors:    0
```

## Creating permissions (50)
#### Creating 25 permissions [reqs:25, conc:5, rule:`""`]
```
┌─ Best:      2.124008ms
├─ Worst:     9.610076ms
├─ Completed: 21.044718ms
└─ Errors:    0
```
#### Creating 25 permissions [reqs:25, conc:5, rule:`"@request.data.name != ''"`]
```
┌─ Best:      2.097288ms
├─ Worst:     10.367119ms
├─ Completed: 27.707423ms
└─ Errors:    0
```

## Creating users (500 - expected to be slow due to passwordHash generation)
#### Creating 250 users [reqs:250, conc:50, rule:`""`]
```
┌─ Best:      957.122636ms
├─ Worst:     2.730076203s
├─ Completed: 8.815766798s
└─ Errors:    0
```
#### Creating 250 users [reqs:250, conc:50, rule:`"@request.data.email != '' && @request.data.permissions:length > 0"`]
```
┌─ Best:      858.53644ms
├─ Worst:     2.760179508s
├─ Completed: 8.763106467s
└─ Errors:    0
```

## Creating posts (10k, 25k, 50k, 100k)
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`""`]
```
┌─ Best:      5.766382ms
├─ Worst:     4.829458427s
├─ Completed: 4.963676768s
└─ Errors:    0
```
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      5.631581ms
├─ Worst:     5.456205487s
├─ Completed: 5.662862701s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`""`]
```
┌─ Best:      3.872255ms
├─ Worst:     9.509186329s
├─ Completed: 13.137581927s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      3.909695ms
├─ Worst:     11.217687027s
├─ Completed: 14.490130417s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`""`]
```
┌─ Best:      3.427413ms
├─ Worst:     13.861141529s
├─ Completed: 26.099528651s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      3.873215ms
├─ Worst:     17.648307064s
├─ Completed: 29.028796979s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`""`]
```
┌─ Best:      3.117252ms
├─ Worst:     26.51586884s
├─ Completed: 55.232834935s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      3.948216ms
├─ Worst:     27.919594309s
├─ Completed: 58.135969276s
└─ Errors:    0
```

## Search records
#### posts10k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.118689ms
├─ Worst:     23.894333ms
├─ Completed: 3.444245151s
└─ Errors:    0
```
#### posts10k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      65.004973ms
├─ Worst:     417.733104ms
├─ Completed: 421.787281ms
└─ Errors:    0
```
#### posts10k - mixed read and write (simpleA list with 300 concurrent random posts10k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      43.54441ms
├─ Worst:     354.606259ms
├─ Completed: 358.226514ms
└─ Errors:    0
```
#### posts10k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      3.293773ms
├─ Worst:     8.174272ms
├─ Completed: 51.927802ms
└─ Errors:    0
```
#### posts10k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      3.717454ms
├─ Worst:     26.113182ms
├─ Completed: 73.391966ms
└─ Errors:    0
```
#### posts10k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      4.719538ms
├─ Worst:     16.137943ms
├─ Completed: 75.597774ms
└─ Errors:    0
```
#### posts10k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      5.814223ms
├─ Worst:     20.3878ms
├─ Completed: 94.901969ms
└─ Errors:    0
```
#### posts10k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      1.983048ms
├─ Worst:     6.429945ms
├─ Completed: 35.139217ms
└─ Errors:    0
```
#### posts10k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      6.634066ms
├─ Worst:     35.7981ms
├─ Completed: 110.967392ms
└─ Errors:    0
```
#### posts10k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.517446ms
├─ Worst:     18.645953ms
├─ Completed: 35.770619ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      16.666185ms
├─ Worst:     53.618969ms
├─ Completed: 254.490391ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.875647ms
├─ Worst:     18.723593ms
├─ Completed: 52.619005ms
└─ Errors:    0
```
#### posts10k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      5.590662ms
├─ Worst:     34.931377ms
├─ Completed: 110.28355ms
└─ Errors:    0
```
#### posts10k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      5.09078ms
├─ Worst:     32.395527ms
├─ Completed: 92.636601ms
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      160.149105ms
├─ Worst:     522.376317ms
├─ Completed: 3.976814997s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      171.731431ms
├─ Worst:     288.967408ms
├─ Completed: 2.345825238s
└─ Errors:    0
```
#### posts10k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      13.376852ms
├─ Worst:     38.707111ms
├─ Completed: 194.080518ms
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      521.484078ms
├─ Worst:     1.090837303s
├─ Completed: 8.498351885s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      2.412446418s
├─ Worst:     4.265549094s
├─ Completed: 34.852307659s
└─ Errors:    0
```
#### posts25k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.71541ms
├─ Worst:     22.313368ms
├─ Completed: 4.388361452s
└─ Errors:    0
```
#### posts25k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      9.954879ms
├─ Worst:     399.250367ms
├─ Completed: 401.577496ms
└─ Errors:    0
```
#### posts25k - mixed read and write (simpleA list with 300 concurrent random posts25k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      28.657952ms
├─ Worst:     456.403271ms
├─ Completed: 458.084638ms
└─ Errors:    0
```
#### posts25k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      3.734175ms
├─ Worst:     11.073844ms
├─ Completed: 59.086152ms
└─ Errors:    0
```
#### posts25k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      4.285017ms
├─ Worst:     10.628002ms
├─ Completed: 60.312477ms
└─ Errors:    0
```
#### posts25k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      5.209461ms
├─ Worst:     43.523211ms
├─ Completed: 137.44478ms
└─ Errors:    0
```
#### posts25k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      6.079504ms
├─ Worst:     17.395708ms
├─ Completed: 101.031077ms
└─ Errors:    0
```
#### posts25k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      2.41777ms
├─ Worst:     11.958287ms
├─ Completed: 40.54072ms
└─ Errors:    0
```
#### posts25k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      11.834847ms
├─ Worst:     32.260167ms
├─ Completed: 177.460617ms
└─ Errors:    0
```
#### posts25k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.621927ms
├─ Worst:     17.220268ms
├─ Completed: 49.241754ms
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      27.602549ms
├─ Worst:     76.989343ms
├─ Completed: 361.44398ms
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.812688ms
├─ Worst:     19.777237ms
├─ Completed: 46.479063ms
└─ Errors:    0
```
#### posts25k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      11.109923ms
├─ Worst:     38.023429ms
├─ Completed: 170.185709ms
└─ Errors:    0
```
#### posts25k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      9.872599ms
├─ Worst:     58.941752ms
├─ Completed: 167.154657ms
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      581.847689ms
├─ Worst:     1.305399813s
├─ Completed: 9.807906853s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      189.266104ms
├─ Worst:     812.377956ms
├─ Completed: 5.759887948s
└─ Errors:    0
```
#### posts25k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      29.982478ms
├─ Worst:     65.358378ms
├─ Completed: 401.3935ms
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      688.348313ms
├─ Worst:     2.820182315s
├─ Completed: 20.874546719s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      6.550659174s
├─ Worst:     11.474612253s
├─ Completed: 1m29.083013613s
└─ Errors:    0
```
#### posts50k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      3.786815ms
├─ Worst:     20.400161ms
├─ Completed: 6.780669974s
└─ Errors:    0
```
#### posts50k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      49.153875ms
├─ Worst:     521.340705ms
├─ Completed: 524.735759ms
└─ Errors:    0
```
#### posts50k - mixed read and write (simpleA list with 300 concurrent random posts50k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      87.909308ms
├─ Worst:     607.711327ms
├─ Completed: 612.140825ms
└─ Errors:    0
```
#### posts50k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      4.815099ms
├─ Worst:     13.942335ms
├─ Completed: 72.609127ms
└─ Errors:    0
```
#### posts50k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      5.160381ms
├─ Worst:     17.866311ms
├─ Completed: 82.343966ms
└─ Errors:    0
```
#### posts50k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      6.387585ms
├─ Worst:     28.467232ms
├─ Completed: 114.882495ms
└─ Errors:    0
```
#### posts50k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      7.279989ms
├─ Worst:     26.299824ms
├─ Completed: 123.173128ms
└─ Errors:    0
```
#### posts50k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      3.472334ms
├─ Worst:     13.944135ms
├─ Completed: 55.875102ms
└─ Errors:    0
```
#### posts50k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      19.903038ms
├─ Worst:     35.714421ms
├─ Completed: 258.824666ms
└─ Errors:    0
```
#### posts50k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.551326ms
├─ Worst:     17.745511ms
├─ Completed: 39.318436ms
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      32.629649ms
├─ Worst:     62.214926ms
├─ Completed: 419.192542ms
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.756687ms
├─ Worst:     19.850359ms
├─ Completed: 55.54606ms
└─ Errors:    0
```
#### posts50k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      20.34492ms
├─ Worst:     64.083055ms
├─ Completed: 317.94046ms
└─ Errors:    0
```
#### posts50k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      18.340793ms
├─ Worst:     70.491639ms
├─ Completed: 267.398421ms
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      1.353012569s
├─ Worst:     2.582801607s
├─ Completed: 19.577102055s
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      198.788709ms
├─ Worst:     1.481791361s
├─ Completed: 11.367870197s
└─ Errors:    0
```
#### posts50k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      58.159271ms
├─ Worst:     128.44091ms
├─ Completed: 699.302377ms
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      3.62078807s
├─ Worst:     4.758553853s
├─ Completed: 40.439738799s
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      13.740963267s
├─ Worst:     18.86923842s
├─ Completed: 2m53.292416573s
└─ Errors:    0
```
#### posts100k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      6.819819ms
├─ Worst:     47.626533ms
├─ Completed: 11.478586674s
└─ Errors:    0
```
#### posts100k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      76.877824ms
├─ Worst:     706.20074ms
├─ Completed: 710.008554ms
└─ Errors:    0
```
#### posts100k - mixed read and write (simpleA list with 300 concurrent random posts100k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      52.557033ms
├─ Worst:     891.453077ms
├─ Completed: 892.539732ms
└─ Errors:    0
```
#### posts100k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      7.899993ms
├─ Worst:     17.378488ms
├─ Completed: 110.583697ms
└─ Errors:    0
```
#### posts100k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      8.43324ms
├─ Worst:     24.071703ms
├─ Completed: 120.574557ms
└─ Errors:    0
```
#### posts100k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      9.499214ms
├─ Worst:     30.963919ms
├─ Completed: 159.270582ms
└─ Errors:    0
```
#### posts100k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      10.529029ms
├─ Worst:     55.749988ms
├─ Completed: 197.112069ms
└─ Errors:    0
```
#### posts100k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      6.582212ms
├─ Worst:     22.417116ms
├─ Completed: 97.849181ms
└─ Errors:    0
```
#### posts100k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      33.37595ms
├─ Worst:     105.074439ms
├─ Completed: 426.704607ms
└─ Errors:    0
```
#### posts100k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.44586ms
├─ Worst:     27.395385ms
├─ Completed: 41.446581ms
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      53.634231ms
├─ Worst:     124.105056ms
├─ Completed: 704.926144ms
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.815705ms
├─ Worst:     14.161797ms
├─ Completed: 48.474877ms
└─ Errors:    0
```
#### posts100k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      39.678832ms
├─ Worst:     99.834069ms
├─ Completed: 543.212918ms
└─ Errors:    0
```
#### posts100k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      37.374679ms
├─ Worst:     101.870771ms
├─ Completed: 504.760996ms
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      2.114784026s
├─ Worst:     5.285758976s
├─ Completed: 38.664103379s
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      945.097205ms
├─ Worst:     2.962834255s
├─ Completed: 22.685251263s
└─ Errors:    0
```
#### posts100k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      114.25122ms
├─ Worst:     161.551867ms
├─ Completed: 1.279824841s
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      6.807803995s
├─ Worst:     9.799595118s
├─ Completed: 1m21.886652058s
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      30.001357334s
├─ Worst:     30.011085848s
├─ Completed: 5m0.039512921s
└─ Errors:    98
```

## Go vs JS route execution
#### JS route (high concurrency) [reqs:500, conc:500]
```
┌─ Best:      552.968602ms
├─ Worst:     4.610804199s
├─ Completed: 4.613081059s
└─ Errors:    0
```
#### Go route (high concurrency) [reqs:500, conc:500]
```
┌─ Best:      524.013432ms
├─ Worst:     4.411184566s
├─ Completed: 4.412871221s
└─ Errors:    0
```
#### JS route (medium concurrency) [reqs:500, conc:50]
```
┌─ Best:      203.614369ms
├─ Worst:     632.289109ms
├─ Completed: 4.413166435s
└─ Errors:    0
```
#### Go route (medium concurrency) [reqs:500, conc:50]
```
┌─ Best:      186.347509ms
├─ Worst:     592.542709ms
├─ Completed: 4.349221953s
└─ Errors:    0
```
#### JS route (no concurrency) [reqs:500, conc:1]
```
┌─ Best:      11.645861ms
├─ Worst:     93.058285ms
├─ Completed: 8.636131438s
└─ Errors:    0
```
#### Go route (no concurrency) [reqs:500, conc:1]
```
┌─ Best:      11.360497ms
├─ Worst:     59.539586ms
├─ Completed: 8.761299841s
└─ Errors:    0
```

## Go vs JS hooks execution
#### JS OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]
```
┌─ Best:      1.696334ms
├─ Worst:     22.083347ms
├─ Completed: 46.22519ms
└─ Errors:    0
```
#### Go OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]
```
┌─ Best:      1.372771ms
├─ Worst:     13.975718ms
├─ Completed: 32.538315ms
└─ Errors:    0
```

## Deleting records
#### deleting 100 posts10k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      1.805535ms
├─ Worst:     14.150919ms
├─ Completed: 40.192299ms
└─ Errors:    0
```
#### deleting 100 posts10k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      1.743534ms
├─ Worst:     10.241687ms
├─ Completed: 33.509402ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      1.864216ms
├─ Worst:     42.119914ms
├─ Completed: 77.082727ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      1.748735ms
├─ Worst:     8.031187ms
├─ Completed: 28.176476ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      1.765095ms
├─ Worst:     13.413272ms
├─ Completed: 40.816182ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.032857ms
├─ Worst:     5.405805ms
├─ Completed: 32.065708ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      1.711254ms
├─ Worst:     6.052491ms
├─ Completed: 27.232107ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      1.861336ms
├─ Worst:     65.326384ms
├─ Completed: 100.621079ms
└─ Errors:    0
```
#### deleting 100 users - with cascade deleting all associated posts [conc:10, rule:`""`]
```
┌─ Best:      325.83979ms
├─ Worst:     13.07376384s
├─ Completed: 32.936688502s
└─ Errors:    0
```
#### deleting 100 organizations - with cascade deleting all users and associated posts [conc:10, rule:`""`]
```
┌─ Best:      974.488896ms
├─ Worst:     1m9.765943807s
├─ Completed: 2m2.218844127s
└─ Errors:    0
```
