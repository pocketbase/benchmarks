## Creating organizations (100)
#### Creating 50 organizations [reqs:50, conc:10, rule:`""`]
```
┌─ Best:      1.429815ms
├─ Worst:     35.031931ms
├─ Completed: 56.553918ms
└─ Errors:    0
```
#### Creating 50 organizations [reqs:50, conc:10, rule:`"@request.data.name != ''"`]
```
┌─ Best:      2.465026ms
├─ Worst:     40.990034ms
├─ Completed: 45.773604ms
└─ Errors:    0
```

## Creating permissions (50)
#### Creating 25 permissions [reqs:25, conc:5, rule:`""`]
```
┌─ Best:      1.91962ms
├─ Worst:     10.834394ms
├─ Completed: 21.388226ms
└─ Errors:    0
```
#### Creating 25 permissions [reqs:25, conc:5, rule:`"@request.data.name != ''"`]
```
┌─ Best:      1.731858ms
├─ Worst:     15.217361ms
├─ Completed: 27.052646ms
└─ Errors:    0
```

## Creating users (500 - expected to be slow due to passwordHash generation)
#### Creating 250 users [reqs:250, conc:50, rule:`""`]
```
┌─ Best:      11.096082828s
├─ Worst:     15.243149454s
├─ Completed: 1m8.873893507s
└─ Errors:    0
```
#### Creating 250 users [reqs:250, conc:50, rule:`"@request.data.email != '' && @request.data.permissions:length > 0"`]
```
┌─ Best:      9.146653957s
├─ Worst:     15.551403038s
├─ Completed: 1m8.714233391s
└─ Errors:    0
```

## Creating posts (10k, 25k, 50k, 100k)
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`""`]
```
┌─ Best:      5.788378ms
├─ Worst:     5.662470531s
├─ Completed: 6.137892547s
└─ Errors:    0
```
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      11.523556ms
├─ Worst:     6.770019363s
├─ Completed: 6.976065363s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`""`]
```
┌─ Best:      6.854629ms
├─ Worst:     13.511503909s
├─ Completed: 14.780976674s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      5.226092ms
├─ Worst:     13.803579438s
├─ Completed: 17.971852544s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`""`]
```
┌─ Best:      4.777288ms
├─ Worst:     17.172627565s
├─ Completed: 29.938191332s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      4.610646ms
├─ Worst:     27.014788005s
├─ Completed: 36.373393532s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`""`]
```
┌─ Best:      4.06376ms
├─ Worst:     18.039652422s
├─ Completed: 1m0.956436764s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      4.427043ms
├─ Worst:     41.87523285s
├─ Completed: 1m13.390319277s
└─ Errors:    0
```

## Search records
#### posts10k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.431104ms
├─ Worst:     16.386398ms
├─ Completed: 3.540751615s
└─ Errors:    0
```
#### posts10k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      207.208189ms
├─ Worst:     1.537682961s
├─ Completed: 1.553005029s
└─ Errors:    0
```
#### posts10k - mixed read and write (simpleA list with 300 concurrent random posts10k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      68.043333ms
├─ Worst:     1.895746548s
├─ Completed: 1.901645445s
└─ Errors:    0
```
#### posts10k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      5.053848ms
├─ Worst:     59.41337ms
├─ Completed: 222.765816ms
└─ Errors:    0
```
#### posts10k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      7.729874ms
├─ Worst:     66.172794ms
├─ Completed: 298.035737ms
└─ Errors:    0
```
#### posts10k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      10.721303ms
├─ Worst:     104.944726ms
├─ Completed: 376.843972ms
└─ Errors:    0
```
#### posts10k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      7.669474ms
├─ Worst:     123.18138ms
├─ Completed: 463.388239ms
└─ Errors:    0
```
#### posts10k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      2.759506ms
├─ Worst:     50.751486ms
├─ Completed: 164.94558ms
└─ Errors:    0
```
#### posts10k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      20.188793ms
├─ Worst:     198.621301ms
├─ Completed: 893.866276ms
└─ Errors:    0
```
#### posts10k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.823538ms
├─ Worst:     44.182903ms
├─ Completed: 138.276204ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      132.86655ms
├─ Worst:     339.449407ms
├─ Completed: 2.705187632s
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      3.150351ms
├─ Worst:     55.44613ms
├─ Completed: 141.611633ms
└─ Errors:    0
```
#### posts10k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      7.535752ms
├─ Worst:     200.145793ms
├─ Completed: 730.776902ms
└─ Errors:    0
```
#### posts10k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      15.531348ms
├─ Worst:     116.998917ms
├─ Completed: 536.948968ms
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      317.277027ms
├─ Worst:     698.354067ms
├─ Completed: 6.152421838s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      125.401315ms
├─ Worst:     294.520449ms
├─ Completed: 2.389287182s
└─ Errors:    0
```
#### posts10k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      151.787126ms
├─ Worst:     507.881ms
├─ Completed: 3.765461637s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      1.3763764s
├─ Worst:     1.633595248s
├─ Completed: 15.411898205s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      2.511652503s
├─ Worst:     4.279476411s
├─ Completed: 38.232966542s
└─ Errors:    0
```
#### posts25k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      3.603315ms
├─ Worst:     19.287742ms
├─ Completed: 5.43320472s
└─ Errors:    0
```
#### posts25k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      278.339148ms
├─ Worst:     2.349299894s
├─ Completed: 2.361269687s
└─ Errors:    0
```
#### posts25k - mixed read and write (simpleA list with 300 concurrent random posts25k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      509.305285ms
├─ Worst:     2.591215283s
├─ Completed: 2.612551684s
└─ Errors:    0
```
#### posts25k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      7.073467ms
├─ Worst:     67.087952ms
├─ Completed: 328.356376ms
└─ Errors:    0
```
#### posts25k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      7.879394ms
├─ Worst:     73.21085ms
├─ Completed: 336.572373ms
└─ Errors:    0
```
#### posts25k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      10.024495ms
├─ Worst:     119.834209ms
├─ Completed: 424.844163ms
└─ Errors:    0
```
#### posts25k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      15.697267ms
├─ Worst:     121.359943ms
├─ Completed: 521.666555ms
└─ Errors:    0
```
#### posts25k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      6.743464ms
├─ Worst:     46.087394ms
├─ Completed: 203.458236ms
└─ Errors:    0
```
#### posts25k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      71.960277ms
├─ Worst:     387.505648ms
├─ Completed: 1.463950862s
└─ Errors:    0
```
#### posts25k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      3.312951ms
├─ Worst:     38.435002ms
├─ Completed: 129.724061ms
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      158.51489ms
├─ Worst:     504.959071ms
├─ Completed: 4.066066728s
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.717986ms
├─ Worst:     59.939723ms
├─ Completed: 168.231662ms
└─ Errors:    0
```
#### posts25k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      80.997361ms
├─ Worst:     387.831165ms
├─ Completed: 1.575647369s
└─ Errors:    0
```
#### posts25k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      46.648879ms
├─ Worst:     335.880716ms
├─ Completed: 1.416192505s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      1.520673695s
├─ Worst:     1.86362022s
├─ Completed: 17.218756672s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      402.180487ms
├─ Worst:     670.409402ms
├─ Completed: 5.987623971s
└─ Errors:    0
```
#### posts25k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      727.770727ms
├─ Worst:     1.023205459s
├─ Completed: 9.415546068s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      3.585851387s
├─ Worst:     4.050710765s
├─ Completed: 39.145436317s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      6.550968866s
├─ Worst:     11.354184703s
├─ Completed: 1m36.82638175s
└─ Errors:    0
```
#### posts50k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      6.050376ms
├─ Worst:     28.490623ms
├─ Completed: 9.608139928s
└─ Errors:    0
```
#### posts50k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      285.972356ms
├─ Worst:     3.608386973s
├─ Completed: 3.621762817s
└─ Errors:    0
```
#### posts50k - mixed read and write (simpleA list with 300 concurrent random posts50k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      309.42953ms
├─ Worst:     4.393991153s
├─ Completed: 4.408541488s
└─ Errors:    0
```
#### posts50k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      8.837641ms
├─ Worst:     96.071285ms
├─ Completed: 449.872302ms
└─ Errors:    0
```
#### posts50k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      15.678424ms
├─ Worst:     151.078991ms
├─ Completed: 469.942566ms
└─ Errors:    0
```
#### posts50k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      10.85234ms
├─ Worst:     143.254399ms
├─ Completed: 553.395974ms
└─ Errors:    0
```
#### posts50k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      11.92895ms
├─ Worst:     128.26262ms
├─ Completed: 606.31778ms
└─ Errors:    0
```
#### posts50k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      8.381117ms
├─ Worst:     79.481731ms
├─ Completed: 352.655925ms
└─ Errors:    0
```
#### posts50k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      121.088794ms
├─ Worst:     766.949455ms
├─ Completed: 2.423905374s
└─ Errors:    0
```
#### posts50k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.792697ms
├─ Worst:     55.597751ms
├─ Completed: 151.957238ms
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      244.052923ms
├─ Worst:     652.233154ms
├─ Completed: 4.530899809s
└─ Errors:    0
```
#### posts50k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.17238ms
├─ Worst:     59.684029ms
├─ Completed: 141.949584ms
└─ Errors:    0
```
#### posts50k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      136.725176ms
├─ Worst:     876.733374ms
├─ Completed: 2.889441425s
└─ Errors:    0
```
#### posts50k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      121.609556ms
├─ Worst:     442.759706ms
├─ Completed: 2.413731085s
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      3.076619578s
├─ Worst:     3.663612756s
├─ Completed: 34.897517703s
└─ Errors:    0
```
#### posts50k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      825.288591ms
├─ Worst:     1.367637436s
├─ Completed: 12.075720085s
└─ Errors:    0
```
#### posts50k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      1.543699743s
├─ Worst:     2.093247529s
├─ Completed: 19.735412492s
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      7.244703008s
├─ Worst:     8.534567276s
├─ Completed: 1m18.940946906s
└─ Errors:    0
```
#### posts50k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      8.245855579s
├─ Worst:     24.551537903s
├─ Completed: 3m23.588831952s
└─ Errors:    0
```
#### posts100k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      10.727377ms
├─ Worst:     52.849076ms
├─ Completed: 15.88636411s
└─ Errors:    0
```
#### posts100k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      207.297025ms
├─ Worst:     6.434044923s
├─ Completed: 6.441184427s
└─ Errors:    0
```
#### posts100k - mixed read and write (simpleA list with 300 concurrent random posts100k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      555.87212ms
├─ Worst:     7.10619607s
├─ Completed: 7.123642547s
└─ Errors:    0
```
#### posts100k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      31.887447ms
├─ Worst:     165.930252ms
├─ Completed: 691.72306ms
└─ Errors:    0
```
#### posts100k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      15.768742ms
├─ Worst:     148.569536ms
├─ Completed: 704.345052ms
└─ Errors:    0
```
#### posts100k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      23.172809ms
├─ Worst:     170.411572ms
├─ Completed: 830.429785ms
└─ Errors:    0
```
#### posts100k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      19.252853ms
├─ Worst:     203.045025ms
├─ Completed: 836.693601ms
└─ Errors:    0
```
#### posts100k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      15.969463ms
├─ Worst:     137.733278ms
├─ Completed: 617.027946ms
└─ Errors:    0
```
#### posts100k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      257.548514ms
├─ Worst:     926.010843ms
├─ Completed: 4.575856042s
└─ Errors:    0
```
#### posts100k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.773856ms
├─ Worst:     75.359957ms
├─ Completed: 159.503193ms
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      401.753969ms
├─ Worst:     1.32999343s
├─ Completed: 8.529259262s
└─ Errors:    0
```
#### posts100k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.894346ms
├─ Worst:     39.826197ms
├─ Completed: 150.982636ms
└─ Errors:    0
```
#### posts100k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      274.852188ms
├─ Worst:     1.767514552s
├─ Completed: 5.919744479s
└─ Errors:    0
```
#### posts100k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      314.402023ms
├─ Worst:     721.797879ms
├─ Completed: 4.799980372s
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      6.199829735s
├─ Worst:     7.896918943s
├─ Completed: 1m15.714308554s
└─ Errors:    0
```
#### posts100k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      1.754085162s
├─ Worst:     2.575109669s
├─ Completed: 21.145381865s
└─ Errors:    0
```
#### posts100k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      3.113214022s
├─ Worst:     4.093552651s
├─ Completed: 37.755182091s
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      13.234800285s
├─ Worst:     16.685119609s
├─ Completed: 2m40.102874464s
└─ Errors:    0
```
#### posts100k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      30.002788501s
├─ Worst:     30.021598201s
├─ Completed: 5m0.090332579s
└─ Errors:    100
```

## Go vs JS route execution
#### JS route (high concurrency) [reqs:500, conc:500]
```
┌─ Best:      798.443997ms
├─ Worst:     12.365219542s
├─ Completed: 12.369240001s
└─ Errors:    0
```
#### Go route (high concurrency) [reqs:500, conc:500]
```
┌─ Best:      301.62247ms
├─ Worst:     11.692610464s
├─ Completed: 11.707046632s
└─ Errors:    0
```
#### JS route (medium concurrency) [reqs:500, conc:50]
```
┌─ Best:      286.263895ms
├─ Worst:     1.427747388s
├─ Completed: 11.098393995s
└─ Errors:    0
```
#### Go route (medium concurrency) [reqs:500, conc:50]
```
┌─ Best:      550.161677ms
├─ Worst:     1.403578623s
├─ Completed: 10.883302019s
└─ Errors:    0
```
#### JS route (no concurrency) [reqs:500, conc:1]
```
┌─ Best:      20.314854ms
├─ Worst:     67.876545ms
├─ Completed: 13.52138176s
└─ Errors:    0
```
#### Go route (no concurrency) [reqs:500, conc:1]
```
┌─ Best:      20.46813ms
├─ Worst:     78.894562ms
├─ Completed: 13.370067651s
└─ Errors:    0
```

## Go vs JS hooks execution
#### JS OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]
```
┌─ Best:      1.491221ms
├─ Worst:     48.1134ms
├─ Completed: 102.981337ms
└─ Errors:    0
```
#### Go OnRecordBeforeUpdateRequest hook handler - [reqs:100, conc:10]
```
┌─ Best:      1.159525ms
├─ Worst:     24.76775ms
├─ Completed: 93.212037ms
└─ Errors:    0
```

## Deleting records
#### deleting 100 posts10k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      2.272072ms
├─ Worst:     25.981327ms
├─ Completed: 107.804242ms
└─ Errors:    0
```
#### deleting 100 posts10k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.576916ms
├─ Worst:     36.431194ms
├─ Completed: 132.579149ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      2.078389ms
├─ Worst:     56.289233ms
├─ Completed: 135.738633ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.505195ms
├─ Worst:     29.87486ms
├─ Completed: 133.758003ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      1.963707ms
├─ Worst:     22.674959ms
├─ Completed: 107.38763ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.200471ms
├─ Worst:     22.833921ms
├─ Completed: 104.956835ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      3.253686ms
├─ Worst:     61.471703ms
├─ Completed: 147.334229ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      3.239526ms
├─ Worst:     63.997937ms
├─ Completed: 165.406ms
└─ Errors:    0
```
#### deleting 100 users - with cascade deleting all associated posts [conc:10, rule:`""`]
```
┌─ Best:      379.78148ms
├─ Worst:     16.568562281s
├─ Completed: 34.596888546s
└─ Errors:    0
```
#### deleting 100 organizations - with cascade deleting all users and associated posts [conc:10, rule:`""`]
```
┌─ Best:      497.16446ms
├─ Worst:     43.52724314s
├─ Completed: 2m2.139878511s
└─ Errors:    0
```
