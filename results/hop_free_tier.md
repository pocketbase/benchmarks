## Creating organizations (100)
#### Creating 50 organizations [reqs:50, conc:10, rule:`""`]
```
┌─ Best:      2.095415ms
├─ Worst:     235.085728ms
├─ Completed: 253.938404ms
└─ Errors:    0
```
#### Creating 50 organizations [reqs:50, conc:10, rule:`"@request.data.name != ''"`]
```
┌─ Best:      2.341709ms
├─ Worst:     19.926168ms
├─ Completed: 34.682825ms
└─ Errors:    0
```

## Creating permissions (50)
#### Creating 25 permissions [reqs:25, conc:5, rule:`""`]
```
┌─ Best:      1.797031ms
├─ Worst:     5.665597ms
├─ Completed: 17.278527ms
└─ Errors:    0
```
#### Creating 25 permissions [reqs:25, conc:5, rule:`"@request.data.name != ''"`]
```
┌─ Best:      1.128573ms
├─ Worst:     31.400895ms
├─ Completed: 44.938177ms
└─ Errors:    0
```

## Creating users (500 - expected to be slow due to passwordHash generation)
#### Creating 250 users [reqs:250, conc:50, rule:`""`]
```
┌─ Best:      15.946086178s
├─ Worst:     42.397237938s
├─ Completed: 2m39.31391604s
└─ Errors:    0
```
#### Creating 250 users [reqs:250, conc:50, rule:`"@request.data.email != '' && @request.data.permissions:length > 0"`]
```
┌─ Best:      17.004532222s
├─ Worst:     41.851551424s
├─ Completed: 2m37.978179828s
└─ Errors:    0
```

## Creating posts (10k, 25k, 50k, 100k)
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`""`]
```
┌─ Best:      4.951459ms
├─ Worst:     23.724590102s
├─ Completed: 23.879576555s
└─ Errors:    0
```
#### Creating 5000 posts10k [reqs:5000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      2.954856ms
├─ Worst:     30.209208466s
├─ Completed: 31.719029015s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`""`]
```
┌─ Best:      2.978456ms
├─ Worst:     38.398593977s
├─ Completed: 58.602453674s
└─ Errors:    0
```
#### Creating 12500 posts25k [reqs:12500, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      3.882257ms
├─ Worst:     55.727203611s
├─ Completed: 1m8.077058856s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`""`]
```
┌─ Best:      2.851175ms
├─ Worst:     48.115968119s
├─ Completed: 1m51.689870921s
└─ Errors:    0
```
#### Creating 25000 posts50k [reqs:25000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      2.942986ms
├─ Worst:     58.786013826s
├─ Completed: 2m14.407858059s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`""`]
```
┌─ Best:      2.333998ms
├─ Worst:     59.298439568s
├─ Completed: 3m37.80780249s
└─ Errors:    0
```
#### Creating 50000 posts100k [reqs:50000, conc:500, rule:`"@request.auth.id != '' && @request.data.public:isset = true"`]
```
┌─ Best:      2.914496ms
├─ Worst:     1m34.410309048s
├─ Completed: 4m22.459095763s
└─ Errors:    0
```

## Search records
#### posts10k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.194487ms
├─ Worst:     375.649472ms
├─ Completed: 5.535335329s
└─ Errors:    0
```
#### posts10k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.408115134s
├─ Worst:     7.863342559s
├─ Completed: 7.865238822s
└─ Errors:    0
```
#### posts10k - mixed read and write (simpleA list with 300 concurrent random posts10k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      5.606676784s
├─ Worst:     12.300038595s
├─ Completed: 12.30124638s
└─ Errors:    0
```
#### posts10k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      4.314692ms
├─ Worst:     398.141037ms
├─ Completed: 1.005874111s
└─ Errors:    0
```
#### posts10k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      4.451495ms
├─ Worst:     589.854315ms
├─ Completed: 1.10335422s
└─ Errors:    0
```
#### posts10k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      8.244271ms
├─ Worst:     500.277893ms
├─ Completed: 1.892884582s
└─ Errors:    0
```
#### posts10k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      8.9401ms
├─ Worst:     695.679476ms
├─ Completed: 2.302228276s
└─ Errors:    0
```
#### posts10k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      2.405679ms
├─ Worst:     96.483057ms
├─ Completed: 395.11918ms
└─ Errors:    0
```
#### posts10k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      106.306327ms
├─ Worst:     889.234409ms
├─ Completed: 3.604544315s
└─ Errors:    0
```
#### posts10k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.815882ms
├─ Worst:     300.031181ms
├─ Completed: 456.747252ms
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      800.070772ms
├─ Worst:     2.198979349s
├─ Completed: 14.892032121s
└─ Errors:    0
```
#### posts10k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.128576ms
├─ Worst:     503.837058ms
├─ Completed: 786.634437ms
└─ Errors:    0
```
#### posts10k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      101.013983ms
├─ Worst:     804.405125ms
├─ Completed: 2.801966547s
└─ Errors:    0
```
#### posts10k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      14.418706ms
├─ Worst:     787.890744ms
├─ Completed: 2.369764505s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      1.600486021s
├─ Worst:     5.393847937s
├─ Completed: 40.112123132s
└─ Errors:    0
```
#### posts10k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      697.884628ms
├─ Worst:     2.496645407s
├─ Completed: 18.678343287s
└─ Errors:    0
```
#### posts10k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      819.243547ms
├─ Worst:     2.89658518s
├─ Completed: 20.03120488s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      6.817500591s
├─ Worst:     11.192525508s
├─ Completed: 1m35.403263162s
└─ Errors:    0
```
#### posts10k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      17.901980383s
├─ Worst:     28.907738407s
├─ Completed: 3m50.796935059s
└─ Errors:    0
```
#### posts25k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      3.659244ms
├─ Worst:     287.279639ms
├─ Completed: 6.744461787s
└─ Errors:    0
```
#### posts25k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.451486549s
├─ Worst:     13.150522853s
├─ Completed: 13.153521269s
└─ Errors:    0
```
#### posts25k - mixed read and write (simpleA list with 300 concurrent random posts25k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.890601882s
├─ Worst:     14.096278647s
├─ Completed: 14.099445375s
└─ Errors:    0
```
#### posts25k - expand author [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author`]
```
┌─ Best:      6.694141ms
├─ Worst:     696.253273ms
├─ Completed: 1.299050275s
└─ Errors:    0
```
#### posts25k - expand author (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author&fields=id,collectionId,expand.author.id`]
```
┌─ Best:      7.155927ms
├─ Worst:     600.722757ms
├─ Completed: 1.397062782s
└─ Errors:    0
```
#### posts25k - expand author.permissions [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions`]
```
┌─ Best:      21.496143ms
├─ Worst:     694.35686ms
├─ Completed: 2.192253815s
└─ Errors:    0
```
#### posts25k - expand author.permissions (limited fields) [reqs:100, conc:10, rule:`""`, query:`?perPage=20&expand=author.permissions&fields=id,collectionId,expand.author.id,expand.author.expand.permissions.id`]
```
┌─ Best:      11.086695ms
├─ Worst:     592.732628ms
├─ Completed: 2.008650692s
└─ Errors:    0
```
#### posts25k - simple auth rule [reqs:100, conc:10, rule:`"@request.auth.id != ''"`, query:`?perPage=20`]
```
┌─ Best:      4.491495ms
├─ Worst:     398.880302ms
├─ Completed: 1.087266059s
└─ Errors:    0
```
#### posts25k - author check (no index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      307.577227ms
├─ Worst:     1.183877949s
├─ Completed: 5.824358216s
└─ Errors:    0
```
#### posts25k - author check (with index) [reqs:100, conc:10, rule:`"author = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      1.969134ms
├─ Worst:     393.244182ms
├─ Completed: 565.015711ms
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (no index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      890.150152ms
├─ Worst:     1.907241474s
├─ Completed: 14.497642509s
└─ Errors:    0
```
#### posts25k - author.id (extra join) check (with index) [reqs:100, conc:10, rule:`"author.id = @request.auth.id"`, query:`?perPage=20`]
```
┌─ Best:      2.115225ms
├─ Worst:     408.310067ms
├─ Completed: 808.722608ms
└─ Errors:    0
```
#### posts25k - loose large text search (no index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      306.98872ms
├─ Worst:     1.192610884s
├─ Completed: 6.20807532s
└─ Errors:    0
```
#### posts25k - loose large text search (with index) [reqs:100, conc:10, rule:`"description ~ 'ipsum dolor'"`, query:`?perPage=20`]
```
┌─ Best:      215.481449ms
├─ Worst:     1.109590809s
├─ Completed: 6.184260739s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, match-all) [reqs:100, conc:10, rule:`"type:each != 'c'"`, query:`?perPage=20`]
```
┌─ Best:      6.901945113s
├─ Worst:     11.587983813s
├─ Completed: 1m29.118224019s
└─ Errors:    0
```
#### posts25k - multiple select :each (no index, at-least-one) [reqs:100, conc:10, rule:`"type:each ?!= 'c'"`, query:`?perPage=20`]
```
┌─ Best:      1.40797104s
├─ Worst:     3.896096903s
├─ Completed: 31.80156645s
└─ Errors:    0
```
#### posts25k - nested single relations lookup (no indexes) [reqs:100, conc:10, rule:`"author.organization.name != 'test'"`, query:`?perPage=20`]
```
┌─ Best:      2.588917278s
├─ Worst:     5.895250971s
├─ Completed: 41.570677765s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, match-all) [reqs:100, conc:10, rule:`"author.permissions.active = true"`, query:`?perPage=20`]
```
┌─ Best:      16.009575485s
├─ Worst:     22.096038372s
├─ Completed: 3m7.297377178s
└─ Errors:    0
```
#### posts25k - nested multiple relations lookup (no indexes, at-least-one) [reqs:100, conc:10, rule:`"author.permissions.active ?= true"`, query:`?perPage=20`]
```
┌─ Best:      30.002107669s
├─ Worst:     30.089054364s
├─ Completed: 5m0.201073715s
└─ Errors:    99
```
#### posts50k - simpleA (many requests, no rules, no concurrent access) [reqs:1000, conc:1, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      6.468937ms
├─ Worst:     456.383305ms
├─ Completed: 9.70115031s
└─ Errors:    0
```
#### posts50k - simpleB (many requests, no rules, concurrent access) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]
```
┌─ Best:      2.48929882s
├─ Worst:     53.186108491s
├─ Completed: 53.190105718s
└─ Errors:    79
```
#### posts50k - mixed read and write (simpleA list with 300 concurrent random posts50k updates running in the background) [reqs:1000, conc:1000, rule:`""`, query:`?perPage=20`]

---

ERROR: Benchmark run interrupted...

---

## Deleting records
#### deleting 100 posts10k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      1.931393ms
├─ Worst:     192.443996ms
├─ Completed: 426.980207ms
└─ Errors:    0
```
#### deleting 100 posts10k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.042475ms
├─ Worst:     796.387549ms
├─ Completed: 989.831656ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      2.387539ms
├─ Worst:     93.574185ms
├─ Completed: 273.66954ms
└─ Errors:    0
```
#### deleting 100 posts25k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.208727ms
├─ Worst:     902.234973ms
├─ Completed: 1.091323668s
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      1.57226ms
├─ Worst:     99.129982ms
├─ Completed: 280.25131ms
└─ Errors:    0
```
#### deleting 100 posts50k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      2.389629ms
├─ Worst:     515.267857ms
├─ Completed: 690.752434ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, no rule) [conc:10, rule:`""`]
```
┌─ Best:      2.43178ms
├─ Worst:     94.193773ms
├─ Completed: 150.416752ms
└─ Errors:    0
```
#### deleting 100 posts100k - simple (no cascade, with rule) [conc:10, rule:`"@request.auth.id != ''"`]
```
┌─ Best:      1.759191ms
├─ Worst:     94.848841ms
├─ Completed: 189.708152ms
└─ Errors:    0
```
#### deleting 100 users - with cascade deleting all associated posts [conc:10, rule:`""`]
```
┌─ Best:      1.186634021s
├─ Worst:     49.854080133s
├─ Completed: 1m14.194072092s
└─ Errors:    0
```
#### deleting 100 organizations - with cascade deleting all users and associated posts [conc:10, rule:`""`]
```
┌─ Best:      1.602578686s
├─ Worst:     2m23.300103198s
├─ Completed: 5m34.517308686s
└─ Errors:    0
```
