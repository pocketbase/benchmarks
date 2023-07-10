## Creating organizations (100)
#### Creating 50 organizations [reqs:50, conc:10, rule:`""`]
```
┌─ Best:      2.260574ms
├─ Worst:     28.796142ms
├─ Completed: 56.617598ms
└─ Errors:    0
```
#### Creating 50 organizations [reqs:50, conc:10, rule:`"@request.data.name != ''"`]
```
┌─ Best:      2.355615ms
├─ Worst:     22.401102ms
├─ Completed: 44.996244ms
└─ Errors:    0
```

## Creating permissions (50)
#### Creating 25 permissions [reqs:25, conc:5, rule:`""`]
```
┌─ Best:      3.429062ms
├─ Worst:     13.619966ms
├─ Completed: 32.909608ms
└─ Errors:    0
```
#### Creating 25 permissions [reqs:25, conc:5, rule:`"@request.data.name != ''"`]
```
┌─ Best:      2.031733ms
├─ Worst:     7.340527ms
├─ Completed: 22.894985ms
└─ Errors:    0
```

## Creating users (500 - expected to be slow due to passwordHash generation)
#### Creating 250 users [reqs:250, conc:50, rule:`""`]
```
┌─ Best:      756.384235ms
├─ Worst:     3.048793693s
├─ Completed: 8.991198564s
└─ Errors:    0
```
#### Creating 250 users [reqs:250, conc:50, rule:`"@request.data.email != '' && @request.data.permissions:length > 0"`]
```
┌─ Best:      585.473269ms
├─ Worst:     4.766495338s
├─ Completed: 8.839347528s
└─ Errors:    0
```

## Creating posts (10k, 25k, 50k, 100k)
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`""`]
```
┌─ Best:      5.556915ms
├─ Worst:     6.371313761s
├─ Completed: 6.477676466s
└─ Errors:    0
```
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      5.852796ms
├─ Worst:     6.939705635s
├─ Completed: 7.116583137s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`""`]
```
┌─ Best:      4.434348ms
├─ Worst:     11.747406423s
├─ Completed: 15.752041492s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      5.311712ms
├─ Worst:     17.314973231s
├─ Completed: 17.467771572s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`""`]
```
┌─ Best:      4.285187ms
├─ Worst:     17.823307373s
├─ Completed: 31.914927795s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      6.333559ms
├─ Worst:     24.842828227s
├─ Completed: 34.110502491s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`""`]
```
┌─ Best:      2.923898ms
├─ Worst:     25.160749426s
├─ Completed: 1m6.782782606s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      4.646587ms
├─ Worst:     27.699471288s
├─ Completed: 1m12.143339183s
└─ Errors:    0
```

## Search records
#### posts10k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.568975ms
├─ Worst:     28.314965ms
├─ Completed: 4.333019804s
└─ Errors:    0
```
#### posts10k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      48.226122ms
├─ Worst:     741.228642ms
├─ Completed: 742.986452ms
└─ Errors:    0
```
#### posts10k - mixed read and write (simpleA list with 300 concurrent random posts10k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      144.06968ms
├─ Worst:     921.437609ms
├─ Completed: 924.946269ms
└─ Errors:    0
```
#### posts10k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      7.919686ms
├─ Worst:     23.721939ms
├─ Completed: 143.324594ms
└─ Errors:    0
```
#### posts10k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      8.833731ms
├─ Worst:     25.159866ms
├─ Completed: 150.737277ms
└─ Errors:    0
```
#### posts10k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      9.452015ms
├─ Worst:     39.372509ms
├─ Completed: 200.678128ms
└─ Errors:    0
```
#### posts10k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      11.186505ms
├─ Worst:     36.585973ms
├─ Completed: 191.043951ms
└─ Errors:    0
```
#### posts10k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      3.649221ms
├─ Worst:     13.7676ms
├─ Completed: 70.150568ms
└─ Errors:    0
```
#### posts10k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      31.903705ms
├─ Worst:     152.731447ms
├─ Completed: 992.392206ms
└─ Errors:    0
```
#### posts10k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.401094ms
├─ Worst:     21.091242ms
├─ Completed: 66.892909ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      206.566598ms
├─ Worst:     621.190486ms
├─ Completed: 4.459155435s
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      3.719021ms
├─ Worst:     26.004271ms
├─ Completed: 82.92664ms
└─ Errors:    0
```
#### posts10k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      30.141495ms
├─ Worst:     81.403072ms
├─ Completed: 576.622942ms
└─ Errors:    0
```
#### posts10k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      26.804876ms
├─ Worst:     55.669403ms
├─ Completed: 408.165364ms
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      558.365988ms
├─ Worst:     796.237447ms
├─ Completed: 6.920834044s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      190.143459ms
├─ Worst:     393.160511ms
├─ Completed: 3.397883834s
└─ Errors:    0
```
#### posts10k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      276.438394ms
├─ Worst:     630.84828ms
├─ Completed: 5.299365465s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      1.365612215s
├─ Worst:     1.900206276s
├─ Completed: 16.287619663s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      4.496357993s
├─ Worst:     5.406332003s
├─ Completed: 50.280824938s
└─ Errors:    0
```
#### posts25k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      3.788181ms
├─ Worst:     36.965811ms
├─ Completed: 7.425686516s
└─ Errors:    0
```
#### posts25k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      83.298594ms
├─ Worst:     1.302173566s
├─ Completed: 1.30300933s
└─ Errors:    0
```
#### posts25k - mixed read and write (simpleA list with 300 concurrent random posts25k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      93.454331ms
├─ Worst:     1.334905828s
├─ Completed: 1.335920874s
└─ Errors:    0
```
#### posts25k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      8.591648ms
├─ Worst:     39.838267ms
├─ Completed: 187.115263ms
└─ Errors:    0
```
#### posts25k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      10.104737ms
├─ Worst:     26.38455ms
├─ Completed: 156.62689ms
└─ Errors:    0
```
#### posts25k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      13.294995ms
├─ Worst:     32.745426ms
├─ Completed: 218.745562ms
└─ Errors:    0
```
#### posts25k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      10.086337ms
├─ Worst:     41.752397ms
├─ Completed: 269.690251ms
└─ Errors:    0
```
#### posts25k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      6.891199ms
├─ Worst:     18.340504ms
├─ Completed: 103.562987ms
└─ Errors:    0
```
#### posts25k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      133.768079ms
├─ Worst:     236.516982ms
├─ Completed: 1.7772758s
└─ Errors:    0
```
#### posts25k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.635655ms
├─ Worst:     45.193056ms
├─ Completed: 95.030538ms
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      639.674664ms
├─ Worst:     895.34611ms
├─ Completed: 7.507976005s
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.876296ms
├─ Worst:     45.97154ms
├─ Completed: 112.683917ms
└─ Errors:    0
```
#### posts25k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      83.411272ms
├─ Worst:     211.521036ms
├─ Completed: 1.411765346s
└─ Errors:    0
```
#### posts25k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      52.178775ms
├─ Worst:     157.024808ms
├─ Completed: 957.985417ms
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      1.701139868s
├─ Worst:     2.12141534s
├─ Completed: 19.230615571s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      645.579556ms
├─ Worst:     923.240841ms
├─ Completed: 8.174903873s
└─ Errors:    0
```
#### posts25k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      889.661837ms
├─ Worst:     1.499110102s
├─ Completed: 12.932420975s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      3.937598361s
├─ Worst:     4.469416231s
├─ Completed: 41.722004167s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      11.728819905s
├─ Worst:     12.862691835s
├─ Completed: 2m2.749792516s
└─ Errors:    0
```
#### posts50k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      6.298155ms
├─ Worst:     40.871505ms
├─ Completed: 10.771119546s
└─ Errors:    0
```
#### posts50k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      92.460828ms
├─ Worst:     1.663200344s
├─ Completed: 1.670013461s
└─ Errors:    0
```
#### posts50k - mixed read and write (simpleA list with 300 concurrent random posts50k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      78.699993ms
├─ Worst:     1.647557775s
├─ Completed: 1.649766067s
└─ Errors:    0
```
#### posts50k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      15.999527ms
├─ Worst:     28.203835ms
├─ Completed: 217.344434ms
└─ Errors:    0
```
#### posts50k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      13.452034ms
├─ Worst:     46.050693ms
├─ Completed: 235.431933ms
└─ Errors:    0
```
#### posts50k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      20.584953ms
├─ Worst:     66.968728ms
├─ Completed: 342.712402ms
└─ Errors:    0
```
#### posts50k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      18.280141ms
├─ Worst:     52.313207ms
├─ Completed: 300.60721ms
└─ Errors:    0
```
#### posts50k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      9.893134ms
├─ Worst:     31.913935ms
├─ Completed: 176.987211ms
└─ Errors:    0
```
#### posts50k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      206.455693ms
├─ Worst:     391.986832ms
├─ Completed: 3.070762211s
└─ Errors:    0
```
#### posts50k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.833736ms
├─ Worst:     38.588012ms
├─ Completed: 89.33477ms
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      425.765494ms
├─ Worst:     646.980148ms
├─ Completed: 5.328205816s
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      3.58202ms
├─ Worst:     21.003595ms
├─ Completed: 76.046496ms
└─ Errors:    0
```
#### posts50k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      183.608326ms
├─ Worst:     391.210904ms
├─ Completed: 2.740838057s
└─ Errors:    0
```
#### posts50k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      125.677968ms
├─ Worst:     246.806232ms
├─ Completed: 1.891226878s
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      3.593998806s
├─ Worst:     9.257015823s
├─ Completed: 44.670287064s
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      1.332419942s
├─ Worst:     1.845895226s
├─ Completed: 16.362916613s
└─ Errors:    0
```
#### posts50k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      2.358460321s
├─ Worst:     3.146529241s
├─ Completed: 27.161948388s
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      8.201656542s
├─ Worst:     9.676671372s
├─ Completed: 1m28.422721023s
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      23.77031738s
├─ Worst:     25.488189917s
├─ Completed: 4m7.085401399s
└─ Errors:    0
```
#### posts100k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      10.801954ms
├─ Worst:     63.742754ms
├─ Completed: 17.18938312s
└─ Errors:    0
```
#### posts100k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      272.628703ms
├─ Worst:     3.331621507s
├─ Completed: 3.337397768s
└─ Errors:    0
```
#### posts100k - mixed read and write (simpleA list with 300 concurrent random posts100k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      563.638134ms
├─ Worst:     3.838758327s
├─ Completed: 3.839636416s
└─ Errors:    0
```
#### posts100k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      20.330492ms
├─ Worst:     68.23711ms
├─ Completed: 400.989133ms
└─ Errors:    0
```
#### posts100k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      26.451875ms
├─ Worst:     91.219468ms
├─ Completed: 447.418369ms
└─ Errors:    0
```
#### posts100k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      28.219612ms
├─ Worst:     61.528398ms
├─ Completed: 416.75332ms
└─ Errors:    0
```
#### posts100k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      26.439874ms
├─ Worst:     63.932062ms
├─ Completed: 426.813818ms
└─ Errors:    0
```
#### posts100k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      21.137378ms
├─ Worst:     44.53026ms
├─ Completed: 302.359082ms
└─ Errors:    0
```
#### posts100k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      377.000357ms
├─ Worst:     533.860855ms
├─ Completed: 4.434342684s
└─ Errors:    0
```
#### posts100k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      3.94364ms
├─ Worst:     23.056157ms
├─ Completed: 93.290637ms
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.16506846s
├─ Worst:     2.040819326s
├─ Completed: 17.153560348s
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      3.9764ms
├─ Worst:     28.998453ms
├─ Completed: 94.568394ms
└─ Errors:    0
```
#### posts100k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      683.417442ms
├─ Worst:     961.964767ms
├─ Completed: 8.381201242s
└─ Errors:    0
```
#### posts100k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      483.036483ms
├─ Worst:     736.617181ms
├─ Completed: 6.141526909s
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      16.333833393s
├─ Worst:     18.989824743s
├─ Completed: 3m1.368297679s
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      2.66983136s
├─ Worst:     3.73091359s
├─ Completed: 33.949190149s
└─ Errors:    0
```
#### posts100k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      11.59806877s
├─ Worst:     14.323507347s
├─ Completed: 2m11.992577345s
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      16.237369982s
├─ Worst:     19.29900161s
├─ Completed: 2m56.656117557s
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      30.00229888s
├─ Worst:     30.02049572s
├─ Completed: 5m0.055608387s
└─ Errors:    100
```

## Go vs JS route execution
#### JS route (high concurrency) [reqs:500, conc:500]
```
┌─ Best:      2.95446154s
├─ Worst:     16.578319758s
├─ Completed: 16.582675188s
└─ Errors:    0
```
#### Go route (high concurrency) [reqs:500, conc:500]
```
┌─ Best:      1.108482905s
├─ Worst:     13.069150309s
├─ Completed: 13.072766774s
└─ Errors:    0
```
#### JS route (medium concurrency) [reqs:500, conc:50]
```
┌─ Best:      498.554779ms
├─ Worst:     1.917309376s
├─ Completed: 14.768783129s
└─ Errors:    0
```
#### Go route (medium concurrency) [reqs:500, conc:50]
```
┌─ Best:      369.677474ms
├─ Worst:     1.905998821s
├─ Completed: 13.566823934s
└─ Errors:    0
```
#### JS route (no concurrency) [reqs:500, conc:1]
```
┌─ Best:      20.480021ms
├─ Worst:     52.951762ms
├─ Completed: 13.658129247s
└─ Errors:    0
```
#### Go route (no concurrency) [reqs:500, conc:1]
```
┌─ Best:      20.394378ms
├─ Worst:     70.067955ms
├─ Completed: 13.968352463s
└─ Errors:    0
```

## Go vs JS hooks execution
#### JS OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]
```
┌─ Best:      2.249895ms
├─ Worst:     60.218647ms
├─ Completed: 99.609994ms
└─ Errors:    0
```
#### Go OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]
```
┌─ Best:      2.195618ms
├─ Worst:     16.447272ms
├─ Completed: 54.946412ms
└─ Errors:    0
```

## Deleting records
#### deleting 100 posts10k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      3.757585ms
├─ Worst:     18.464525ms
├─ Completed: 67.955659ms
└─ Errors:    0
```
#### deleting 100 posts10k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      3.108261ms
├─ Worst:     16.357351ms
├─ Completed: 74.218941ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      3.795105ms
├─ Worst:     37.223571ms
├─ Completed: 92.601066ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      3.374422ms
├─ Worst:     24.812887ms
├─ Completed: 84.074607ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      3.460023ms
├─ Worst:     11.469957ms
├─ Completed: 67.930059ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      3.367143ms
├─ Worst:     16.546231ms
├─ Completed: 81.515711ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      3.464864ms
├─ Worst:     13.130768ms
├─ Completed: 64.258433ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      4.142388ms
├─ Worst:     45.459547ms
├─ Completed: 103.74774ms
└─ Errors:    0
```
#### deleting 100 users - with cascade deleting all associated posts [conc:10, rule:`""`]
```
┌─ Best:      388.139417ms
├─ Worst:     20.704943231s
├─ Completed: 36.892867299s
└─ Errors:    0
```
#### deleting 100 organizations - with cascade deleting all users and associated posts [conc:10, rule:`""`]
```
┌─ Best:      947.684778ms
├─ Worst:     1m21.154466611s
├─ Completed: 2m5.087531467s
└─ Errors:    0
```
