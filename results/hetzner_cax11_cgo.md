## Creating organizations (100)
#### Creating 50 organizations [reqs:50, conc:10, rule:`""`]
```
┌─ Best:      2.096591ms
├─ Worst:     15.576868ms
├─ Completed: 31.372299ms
└─ Errors:    0
```
#### Creating 50 organizations [reqs:50, conc:10, rule:`"@request.data.name != ''"`]
```
┌─ Best:      2.307874ms
├─ Worst:     15.10882ms
├─ Completed: 33.906534ms
└─ Errors:    0
```

## Creating permissions (50)
#### Creating 25 permissions [reqs:25, conc:5, rule:`""`]
```
┌─ Best:      1.620424ms
├─ Worst:     8.347721ms
├─ Completed: 20.388816ms
└─ Errors:    0
```
#### Creating 25 permissions [reqs:25, conc:5, rule:`"@request.data.name != ''"`]
```
┌─ Best:      2.218712ms
├─ Worst:     9.62194ms
├─ Completed: 23.228576ms
└─ Errors:    0
```

## Creating users (500 - expected to be slow due to passwordHash generation)
#### Creating 250 users [reqs:250, conc:50, rule:`""`]
```
┌─ Best:      11.476821377s
├─ Worst:     14.903141118s
├─ Completed: 1m8.341327115s
└─ Errors:    0
```
#### Creating 250 users [reqs:250, conc:50, rule:`"@request.data.email != '' && @request.data.permissions:length > 0"`]
```
┌─ Best:      10.891861253s
├─ Worst:     15.073448354s
├─ Completed: 1m8.798511303s
└─ Errors:    0
```

## Creating posts (10k, 25k, 50k, 100k)
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`""`]
```
┌─ Best:      7.027441ms
├─ Worst:     4.238301037s
├─ Completed: 4.971361623s
└─ Errors:    0
```
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      13.312514ms
├─ Worst:     5.318260597s
├─ Completed: 5.846209381s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`""`]
```
┌─ Best:      4.214767ms
├─ Worst:     10.395502453s
├─ Completed: 11.837979829s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      12.53798ms
├─ Worst:     9.467885308s
├─ Completed: 14.399764196s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`""`]
```
┌─ Best:      5.617502ms
├─ Worst:     18.222600452s
├─ Completed: 24.067531112s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      6.854233ms
├─ Worst:     19.644992286s
├─ Completed: 28.876947876s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`""`]
```
┌─ Best:      4.010641ms
├─ Worst:     22.823858999s
├─ Completed: 48.151610917s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      6.021419ms
├─ Worst:     25.659961374s
├─ Completed: 57.389942944s
└─ Errors:    0
```

## Search records
#### posts10k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      1.771777ms
├─ Worst:     14.520378ms
├─ Completed: 2.586367742s
└─ Errors:    0
```
#### posts10k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      197.091839ms
├─ Worst:     1.217788802s
├─ Completed: 1.23755363s
└─ Errors:    0
```
#### posts10k - mixed read and write (simpleA list with 300 concurrent random posts10k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      289.328951ms
├─ Worst:     1.420996347s
├─ Completed: 1.435805808s
└─ Errors:    0
```
#### posts10k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      3.064189ms
├─ Worst:     52.796661ms
├─ Completed: 205.594791ms
└─ Errors:    0
```
#### posts10k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      7.117867ms
├─ Worst:     68.66073ms
├─ Completed: 317.65877ms
└─ Errors:    0
```
#### posts10k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      4.692005ms
├─ Worst:     96.323551ms
├─ Completed: 353.831947ms
└─ Errors:    0
```
#### posts10k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      11.127825ms
├─ Worst:     76.337361ms
├─ Completed: 389.161236ms
└─ Errors:    0
```
#### posts10k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      5.398931ms
├─ Worst:     52.960459ms
├─ Completed: 175.682417ms
└─ Errors:    0
```
#### posts10k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      7.269028ms
├─ Worst:     114.307277ms
├─ Completed: 461.153664ms
└─ Errors:    0
```
#### posts10k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.483174ms
├─ Worst:     31.017492ms
├─ Completed: 109.578431ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      28.076064ms
├─ Worst:     239.400729ms
├─ Completed: 1.148322023s
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.986468ms
├─ Worst:     26.829132ms
├─ Completed: 117.746464ms
└─ Errors:    0
```
#### posts10k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      10.831142ms
├─ Worst:     72.654241ms
├─ Completed: 297.993831ms
└─ Errors:    0
```
#### posts10k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      5.541212ms
├─ Worst:     71.57895ms
├─ Completed: 282.630522ms
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      181.419571ms
├─ Worst:     361.809974ms
├─ Completed: 2.864551686s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      80.933633ms
├─ Worst:     228.77317ms
├─ Completed: 1.56764111s
└─ Errors:    0
```
#### posts10k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      15.237181ms
├─ Worst:     143.514734ms
├─ Completed: 741.793331ms
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      443.034777ms
├─ Worst:     657.349284ms
├─ Completed: 5.827650774s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      1.762225117s
├─ Worst:     2.01812621s
├─ Completed: 19.237183671s
└─ Errors:    0
```
#### posts25k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.357461ms
├─ Worst:     46.518824ms
├─ Completed: 3.833192808s
└─ Errors:    0
```
#### posts25k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      176.661048ms
├─ Worst:     1.652731961s
├─ Completed: 1.673146226s
└─ Errors:    0
```
#### posts25k - mixed read and write (simpleA list with 300 concurrent random posts25k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      352.461722ms
├─ Worst:     1.748661443s
├─ Completed: 1.758047448s
└─ Errors:    0
```
#### posts25k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      5.238767ms
├─ Worst:     72.504337ms
├─ Completed: 235.748538ms
└─ Errors:    0
```
#### posts25k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      7.350306ms
├─ Worst:     62.091763ms
├─ Completed: 294.454667ms
└─ Errors:    0
```
#### posts25k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      5.211727ms
├─ Worst:     98.113448ms
├─ Completed: 369.829027ms
└─ Errors:    0
```
#### posts25k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      9.185243ms
├─ Worst:     94.161171ms
├─ Completed: 369.581181ms
└─ Errors:    0
```
#### posts25k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      2.810945ms
├─ Worst:     76.108087ms
├─ Completed: 211.361908ms
└─ Errors:    0
```
#### posts25k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      19.748698ms
├─ Worst:     150.468677ms
├─ Completed: 688.766852ms
└─ Errors:    0
```
#### posts25k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.891906ms
├─ Worst:     19.784058ms
├─ Completed: 88.095473ms
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      28.263215ms
├─ Worst:     330.11909ms
├─ Completed: 1.632482646s
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.742496ms
├─ Worst:     53.597282ms
├─ Completed: 140.644584ms
└─ Errors:    0
```
#### posts25k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      18.677368ms
├─ Worst:     180.348219ms
├─ Completed: 680.183624ms
└─ Errors:    0
```
#### posts25k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      22.962406ms
├─ Worst:     132.308146ms
├─ Completed: 623.162305ms
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      564.033755ms
├─ Worst:     773.881757ms
├─ Completed: 6.837132801s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      168.853984ms
├─ Worst:     451.228143ms
├─ Completed: 3.667153119s
└─ Errors:    0
```
#### posts25k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      70.571147ms
├─ Worst:     215.125994ms
├─ Completed: 1.673521883s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      1.168272867s
├─ Worst:     1.392250643s
├─ Completed: 12.805745325s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      3.71391171s
├─ Worst:     4.453933342s
├─ Completed: 41.073295065s
└─ Errors:    0
```
#### posts50k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      3.43511ms
├─ Worst:     25.894265ms
├─ Completed: 5.46082276s
└─ Errors:    0
```
#### posts50k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      166.700565ms
├─ Worst:     1.97042375s
├─ Completed: 1.988291785s
└─ Errors:    0
```
#### posts50k - mixed read and write (simpleA list with 300 concurrent random posts50k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      59.969959ms
├─ Worst:     2.404728843s
├─ Completed: 2.434082498s
└─ Errors:    0
```
#### posts50k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      4.737321ms
├─ Worst:     55.235877ms
├─ Completed: 263.468997ms
└─ Errors:    0
```
#### posts50k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      8.311432ms
├─ Worst:     69.704162ms
├─ Completed: 340.876424ms
└─ Errors:    0
```
#### posts50k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      6.257134ms
├─ Worst:     107.144524ms
├─ Completed: 395.23997ms
└─ Errors:    0
```
#### posts50k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      10.155368ms
├─ Worst:     114.199745ms
├─ Completed: 489.839544ms
└─ Errors:    0
```
#### posts50k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      7.570025ms
├─ Worst:     58.558385ms
├─ Completed: 219.024528ms
└─ Errors:    0
```
#### posts50k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      50.604595ms
├─ Worst:     203.825675ms
├─ Completed: 1.196829863s
└─ Errors:    0
```
#### posts50k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.712415ms
├─ Worst:     43.985818ms
├─ Completed: 110.48555ms
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      71.549495ms
├─ Worst:     298.788247ms
├─ Completed: 1.856238071s
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      3.196707ms
├─ Worst:     39.834422ms
├─ Completed: 120.474915ms
└─ Errors:    0
```
#### posts50k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      22.990597ms
├─ Worst:     247.18676ms
├─ Completed: 1.181198372s
└─ Errors:    0
```
#### posts50k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      45.002386ms
├─ Worst:     154.013999ms
├─ Completed: 1.042858498s
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      1.013773928s
├─ Worst:     1.466755807s
├─ Completed: 13.836767696s
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      446.78624ms
├─ Worst:     972.185967ms
├─ Completed: 8.405181933s
└─ Errors:    0
```
#### posts50k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      168.172028ms
├─ Worst:     390.852642ms
├─ Completed: 3.214754757s
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      2.669779726s
├─ Worst:     3.087972198s
├─ Completed: 28.862837141s
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      7.924571912s
├─ Worst:     9.623899553s
├─ Completed: 1m24.68755414s
└─ Errors:    0
```
#### posts100k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      6.320932ms
├─ Worst:     32.734151ms
├─ Completed: 10.203919745s
└─ Errors:    0
```
#### posts100k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      570.506194ms
├─ Worst:     3.773377574s
├─ Completed: 3.775157989s
└─ Errors:    0
```
#### posts100k - mixed read and write (simpleA list with 300 concurrent random posts100k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      666.812865ms
├─ Worst:     3.673978766s
├─ Completed: 3.700777067s
└─ Errors:    0
```
#### posts100k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      12.734665ms
├─ Worst:     101.637678ms
├─ Completed: 464.173149ms
└─ Errors:    0
```
#### posts100k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      8.809712ms
├─ Worst:     123.391497ms
├─ Completed: 535.923978ms
└─ Errors:    0
```
#### posts100k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      9.847641ms
├─ Worst:     141.582647ms
├─ Completed: 566.886753ms
└─ Errors:    0
```
#### posts100k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      10.89701ms
├─ Worst:     163.473907ms
├─ Completed: 616.232677ms
└─ Errors:    0
```
#### posts100k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      7.20338ms
├─ Worst:     104.078217ms
├─ Completed: 389.386207ms
└─ Errors:    0
```
#### posts100k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      81.768834ms
├─ Worst:     259.054533ms
├─ Completed: 1.838264374s
└─ Errors:    0
```
#### posts100k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      3.069145ms
├─ Worst:     32.426347ms
├─ Completed: 112.98113ms
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      113.585734ms
├─ Worst:     397.671831ms
├─ Completed: 2.81585932s
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      3.59227ms
├─ Worst:     15.694969ms
├─ Completed: 93.859932ms
└─ Errors:    0
```
#### posts100k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      123.524535ms
├─ Worst:     282.968405ms
├─ Completed: 2.129259056s
└─ Errors:    0
```
#### posts100k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      91.995955ms
├─ Worst:     259.686372ms
├─ Completed: 1.874692595s
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      2.410096435s
├─ Worst:     2.992915819s
├─ Completed: 27.874946476s
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      1.291006139s
├─ Worst:     1.651875409s
├─ Completed: 15.011449281s
└─ Errors:    0
```
#### posts100k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      479.539871ms
├─ Worst:     673.881619ms
├─ Completed: 6.133046313s
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      5.457187376s
├─ Worst:     6.279035876s
├─ Completed: 58.81747828s
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      15.679239899s
├─ Worst:     18.563390653s
├─ Completed: 2m49.003147243s
└─ Errors:    0
```

## Go vs JS route execution
#### JS route (high concurrency) [reqs:500, conc:500]
```
┌─ Best:      366.045888ms
├─ Worst:     4.883773577s
├─ Completed: 4.899351761s
└─ Errors:    0
```
#### Go route (high concurrency) [reqs:500, conc:500]
```
┌─ Best:      301.389608ms
├─ Worst:     4.53152628s
├─ Completed: 4.539715635s
└─ Errors:    0
```
#### JS route (medium concurrency) [reqs:500, conc:50]
```
┌─ Best:      159.188664ms
├─ Worst:     732.77741ms
├─ Completed: 4.390272244s
└─ Errors:    0
```
#### Go route (medium concurrency) [reqs:500, conc:50]
```
┌─ Best:      61.728463ms
├─ Worst:     724.08434ms
├─ Completed: 4.613569965s
└─ Errors:    0
```
#### JS route (no concurrency) [reqs:500, conc:1]
```
┌─ Best:      11.630845ms
├─ Worst:     48.339605ms
├─ Completed: 7.977412125s
└─ Errors:    0
```
#### Go route (no concurrency) [reqs:500, conc:1]
```
┌─ Best:      11.4568ms
├─ Worst:     31.757962ms
├─ Completed: 7.974308683s
└─ Errors:    0
```

## Go vs JS hooks execution
#### JS OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]
```
┌─ Best:      1.788945ms
├─ Worst:     18.897022ms
├─ Completed: 71.009025ms
└─ Errors:    0
```
#### Go OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]
```
┌─ Best:      1.676944ms
├─ Worst:     44.318653ms
├─ Completed: 104.546087ms
└─ Errors:    0
```

## Deleting records
#### deleting 100 posts10k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      1.959787ms
├─ Worst:     25.28643ms
├─ Completed: 104.23876ms
└─ Errors:    0
```
#### deleting 100 posts10k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.969241ms
├─ Worst:     36.153619ms
├─ Completed: 124.791762ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      2.14327ms
├─ Worst:     34.526995ms
├─ Completed: 114.424696ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.212791ms
├─ Worst:     41.933297ms
├─ Completed: 131.731573ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      2.503275ms
├─ Worst:     28.226868ms
├─ Completed: 107.291074ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.249191ms
├─ Worst:     23.226679ms
├─ Completed: 106.301018ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      1.706663ms
├─ Worst:     42.654184ms
├─ Completed: 133.291506ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.081668ms
├─ Worst:     37.133467ms
├─ Completed: 134.859364ms
└─ Errors:    0
```
#### deleting 100 users - with cascade deleting all associated posts [conc:10, rule:`""`]
```
┌─ Best:      352.738175ms
├─ Worst:     11.771048355s
├─ Completed: 31.903884585s
└─ Errors:    0
```
#### deleting 100 organizations - with cascade deleting all users and associated posts [conc:10, rule:`""`]
```
┌─ Best:      643.184661ms
├─ Worst:     37.403445539s
├─ Completed: 1m57.651508395s
└─ Errors:    0
```
