/// <reference path="../pb_data/types.d.ts" />

function customMiddleware(e) {
    e.set("total", 20)

    return e.next()
}

routerAdd("GET", "/js", (e) => {
    const total = e.get("total")

    const records = $app.findRecordsByFilter("posts10k", "title != ''", "-created", total)

    return e.json(200, records)
}, customMiddleware)

onRecordUpdateRequest((e) => {
    if (e.record.get("title") != "") {
        e.record.set("title", "js_update")
    }

    return e.next()
}, "js")
