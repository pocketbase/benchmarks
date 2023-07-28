## Creating organizations (100)
#### Creating 50 organizations [reqs:50, conc:10, rule:`""`]
```
┌─ Best:      1.59951ms
├─ Worst:     60.864074ms
├─ Completed: 61.247238ms
└─ Errors:    0
```
#### Creating 50 organizations [reqs:50, conc:10, rule:`"@request.data.name != ''"`]
```
┌─ Best:      2.760424ms
├─ Worst:     32.501483ms
├─ Completed: 49.523073ms
└─ Errors:    0
```

## Creating permissions (50)
#### Creating 25 permissions [reqs:25, conc:5, rule:`""`]
```
┌─ Best:      2.574432ms
├─ Worst:     10.514741ms
├─ Completed: 26.970824ms
└─ Errors:    0
```
#### Creating 25 permissions [reqs:25, conc:5, rule:`"@request.data.name != ''"`]
```
┌─ Best:      2.723684ms
├─ Worst:     11.358711ms
├─ Completed: 28.964989ms
└─ Errors:    0
```

## Creating users (500 - expected to be slow due to passwordHash generation)
#### Creating 250 users [reqs:250, conc:50, rule:`""`]
```
┌─ Best:      16.810309068s
├─ Worst:     41.501394932s
├─ Completed: 2m43.332026581s
└─ Errors:    0
```
#### Creating 250 users [reqs:250, conc:50, rule:`"@request.data.email != '' && @request.data.permissions:length > 0"`]
```
┌─ Best:      17.637644382s
├─ Worst:     44.093181486s
├─ Completed: 2m42.122929331s
└─ Errors:    0
```

## Creating posts (10k, 25k, 50k, 100k)
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`""`]
```
┌─ Best:      5.917113ms
├─ Worst:     29.694340883s
├─ Completed: 30.375365877s
└─ Errors:    0
```
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      9.772401ms
├─ Worst:     38.524412934s
├─ Completed: 38.783325879s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`""`]
```
┌─ Best:      3.605875ms
├─ Worst:     54.898799014s
├─ Completed: 1m13.609889909s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      3.514113ms
├─ Worst:     1m3.20867893s
├─ Completed: 1m21.692281559s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`""`]
```
┌─ Best:      3.407613ms
├─ Worst:     1m26.482620108s
├─ Completed: 2m27.790592321s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      2.869995ms
├─ Worst:     57.92293756s
├─ Completed: 2m31.800528058s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`""`]
```
┌─ Best:      2.888287ms
├─ Worst:     1m10.677452484s
├─ Completed: 4m27.904580891s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      3.795777ms
├─ Worst:     1m21.991465976s
├─ Completed: 5m49.0756739s
└─ Errors:    0
```

## User auth with password (expected to be slow due to passwordHash verification)
#### users auth with email/pass - high concurrency [reqs:250, conc:250]
```
┌─ Best:      37.486730072s
├─ Worst:     1m29.317767484s
├─ Completed: 1m29.319401204s
└─ Errors:    0
```
#### users auth with email/pass - small concurrency [reqs:250, conc:10]
```
┌─ Best:      1.707598355s
├─ Worst:     4.487876526s
├─ Completed: 1m20.621237142s
└─ Errors:    0
```

## User auth refresh
#### users - auth refresh (high concurrency) [reqs:1000, conc:1000]
```
┌─ Best:      1.443815361s
├─ Worst:     7.952343674s
├─ Completed: 7.953799112s
└─ Errors:    0
```
#### users - auth refresh (medium concurrency) [reqs:1000, conc:100]
```
┌─ Best:      1.017452ms
├─ Worst:     4.314331905s
├─ Completed: 6.688744848s
└─ Errors:    0
```

## List records
#### users - getOne for auth refresh comparison (high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`/cwlpbz4qg81iob4`]
```
┌─ Best:      98.984999ms
├─ Worst:     9.597371461s
├─ Completed: 9.598976821s
└─ Errors:    0
```
#### users - getOne for auth refresh comparison (medium concurrency) [reqs:1000, conc:100, rule:`""`, query:`/cwlpbz4qg81iob4`]
```
┌─ Best:      940.642µs
├─ Worst:     8.094554836s
├─ Completed: 8.101186516s
└─ Errors:    0
```
#### posts10k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.273607ms
├─ Worst:     296.910451ms
├─ Completed: 6.823748309s
└─ Errors:    0
```
#### posts10k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.196295229s
├─ Worst:     15.001073807s
├─ Completed: 15.002515304s
└─ Errors:    0
```
#### posts10k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      6.493707848s
├─ Worst:     14.395891461s
├─ Completed: 14.399721998s
└─ Errors:    0
```
#### posts10k - mixed read and write (simpleA list with additional 300 concurrent random posts10k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.095681675s
├─ Worst:     16.094619457s
├─ Completed: 16.098049779s
└─ Errors:    0
```
#### posts10k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      4.230271ms
├─ Worst:     192.575361ms
├─ Completed: 772.45953ms
└─ Errors:    0
```
#### posts10k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      7.074566ms
├─ Worst:     695.521045ms
├─ Completed: 1.697362455s
└─ Errors:    0
```
#### posts10k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      10.242774ms
├─ Worst:     898.121231ms
├─ Completed: 2.997087742s
└─ Errors:    0
```
#### posts10k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      11.804954ms
├─ Worst:     990.266132ms
├─ Completed: 2.198731929s
└─ Errors:    0
```
#### posts10k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      2.650762ms
├─ Worst:     384.670209ms
├─ Completed: 800.161412ms
└─ Errors:    0
```
#### posts10k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      306.742302ms
├─ Worst:     1.188714381s
├─ Completed: 6.00882634s
└─ Errors:    0
```
#### posts10k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.868993ms
├─ Worst:     101.20034ms
├─ Completed: 399.142456ms
└─ Errors:    0
```
#### posts10k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.823103ms
├─ Worst:     498.527307ms
├─ Completed: 1.086256389s
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.417231147s
├─ Worst:     2.800471937s
├─ Completed: 21.51680453s
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.030654ms
├─ Worst:     393.941638ms
├─ Completed: 793.350493ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      2.110385ms
├─ Worst:     97.425337ms
├─ Completed: 395.458097ms
└─ Errors:    0
```
#### posts10k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      192.611027ms
├─ Worst:     1.29434447s
├─ Completed: 4.302972729s
└─ Errors:    0
```
#### posts10k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      65.298636ms
├─ Worst:     988.289604ms
├─ Completed: 3.18125612s
└─ Errors:    0
```
#### posts10k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      35.011347ms
├─ Worst:     904.127041ms
├─ Completed: 3.72446347s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      2.710105243s
├─ Worst:     6.802879253s
├─ Completed: 50.707642363s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      3.403009156s
├─ Worst:     6.675698563s
├─ Completed: 48.67420065s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      397.246381ms
├─ Worst:     3.384756072s
├─ Completed: 21.705828631s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.10408668s
├─ Worst:     3.395479742s
├─ Completed: 22.513319108s
└─ Errors:    0
```
#### posts10k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      2.228267069s
├─ Worst:     3.982075248s
├─ Completed: 31.278134868s
└─ Errors:    0
```
#### posts10k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      812.012929ms
├─ Worst:     3.902212291s
├─ Completed: 31.124206685s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      6.722767838s
├─ Worst:     14.205127179s
├─ Completed: 1m57.207734287s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      6.219489948s
├─ Worst:     11.686005822s
├─ Completed: 1m23.011389291s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      19.045981466s
├─ Worst:     42.80710433s
├─ Completed: 5m40.918112789s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      22.293039112s
├─ Worst:     42.79607078s
├─ Completed: 5m20.259567882s
└─ Errors:    0
```
#### posts25k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      3.725176ms
├─ Worst:     497.076024ms
├─ Completed: 8.119665715s
└─ Errors:    0
```
#### posts25k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      5.103659498s
├─ Worst:     15.274624938s
├─ Completed: 15.2788205s
└─ Errors:    0
```
#### posts25k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      2.281741332s
├─ Worst:     14.978321962s
├─ Completed: 14.979372435s
└─ Errors:    0
```
#### posts25k - mixed read and write (simpleA list with additional 300 concurrent random posts25k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.304528426s
├─ Worst:     21.81070274s
├─ Completed: 21.813500144s
└─ Errors:    0
```
#### posts25k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      7.678106ms
├─ Worst:     1.096274776s
├─ Completed: 1.900821009s
└─ Errors:    0
```
#### posts25k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      7.534674ms
├─ Worst:     400.239486ms
├─ Completed: 1.192942312s
└─ Errors:    0
```
#### posts25k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      92.584136ms
├─ Worst:     796.103049ms
├─ Completed: 2.802645618s
└─ Errors:    0
```
#### posts25k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      10.687672ms
├─ Worst:     997.245ms
├─ Completed: 2.691765336s
└─ Errors:    0
```
#### posts25k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      4.88968ms
├─ Worst:     188.047938ms
├─ Completed: 701.423227ms
└─ Errors:    0
```
#### posts25k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      315.61387ms
├─ Worst:     1.593869079s
├─ Completed: 9.117588777s
└─ Errors:    0
```
#### posts25k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.034665ms
├─ Worst:     496.844964ms
├─ Completed: 774.879526ms
└─ Errors:    0
```
#### posts25k - author check (with index and skipTotal) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      1.866973ms
├─ Worst:     101.345435ms
├─ Completed: 400.842984ms
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.430317456s
├─ Worst:     3.896782203s
├─ Completed: 27.297882019s
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      3.097968ms
├─ Worst:     703.565404ms
├─ Completed: 1.040574448s
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (with index and skipTotal) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      2.760514ms
├─ Worst:     98.199976ms
├─ Completed: 381.91045ms
└─ Errors:    0
```
#### posts25k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      498.379873ms
├─ Worst:     1.294435233s
├─ Completed: 8.203496893s
└─ Errors:    0
```
#### posts25k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      208.164309ms
├─ Worst:     1.387487857s
├─ Completed: 5.66480473s
└─ Errors:    0
```
#### posts25k - loose large text search (with index and skipTotal) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      305.11411ms
├─ Worst:     1.009236492s
├─ Completed: 6.116285264s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      6.504165401s
├─ Worst:     13.497032479s
├─ Completed: 1m40.291992292s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, match-all, skipTotal) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      8.334516838s
├─ Worst:     15.90692333s
├─ Completed: 2m6.820108269s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      2.585372214s
├─ Worst:     7.614823345s
├─ Completed: 37.685693188s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      2.277749525s
├─ Worst:     6.579583305s
├─ Completed: 40.47297139s
└─ Errors:    0
```
#### posts25k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      3.293822683s
├─ Worst:     6.601952021s
├─ Completed: 45.701488262s
└─ Errors:    0
```
#### posts25k - nested single relations lookup (no indexes, skipTotal) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      3.097786136s
├─ Worst:     7.376774082s
├─ Completed: 49.006658756s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      17.885635435s
├─ Worst:     30.912599285s
├─ Completed: 4m6.216984258s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, match-all, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      16.302032525s
├─ Worst:     29.889180502s
├─ Completed: 3m36.020998249s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      50.996664307s
├─ Worst:     1m30.086663179s
├─ Completed: 11m30.044403582s
└─ Errors:    9
```
#### posts25k - nested multiple relations lookup (no indexes, at-least-one, skipTotal) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      52.610826798s
├─ Worst:     1m30.076361221s
├─ Completed: 12m25.036520612s
└─ Errors:    5
```
#### posts50k - simpleA (many requests, no rules, no concurrency) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      6.485361ms
├─ Worst:     395.448267ms
├─ Completed: 11.090240213s
└─ Errors:    0
```
#### posts50k - simpleB (many requests, no rules, high concurrency) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      1.560432453s
├─ Worst:     2m19.29613546s
├─ Completed: 2m19.310466589s
└─ Errors:    80
```
#### posts50k - simpleC (many requests, no rules, high concurrency, skipTotal) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20&skipTotal=1`]
```
┌─ Best:      17.998971256s
├─ Worst:     3m7.531215553s
├─ Completed: 3m7.545885646s
└─ Errors:    75
```

---

ERROR: Benchmark run interrupted...

---

## Deleting records
#### deleting 100 posts10k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      2.502717ms
├─ Worst:     93.548447ms
├─ Completed: 242.403613ms
└─ Errors:    0
```
#### deleting 100 posts10k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.217113ms
├─ Worst:     588.161926ms
├─ Completed: 895.339886ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      2.258233ms
├─ Worst:     92.896411ms
├─ Completed: 169.348231ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.632228ms
├─ Worst:     94.300765ms
├─ Completed: 299.632971ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      1.589067ms
├─ Worst:     703.600607ms
├─ Completed: 901.570216ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.389185ms
├─ Worst:     95.403907ms
├─ Completed: 279.743315ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      2.558447ms
├─ Worst:     502.061361ms
├─ Completed: 961.754277ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.411375ms
├─ Worst:     96.292268ms
├─ Completed: 263.342735ms
└─ Errors:    0
```
#### deleting 100 users - with cascade deleting all associated posts [conc:10, rule:`""`]
```
┌─ Best:      1.221159355s
├─ Worst:     1m6.595950878s
├─ Completed: 1m52.60988842s
└─ Errors:    0
```
#### deleting 100 organizations - with cascade deleting all users and associated posts [conc:10, rule:`""`]
```
┌─ Best:      4.588953554s
├─ Worst:     2m53.81755804s
├─ Completed: 7m17.603306853s
└─ Errors:    0
```
