/// <reference path="../pb_data/types.d.ts" />

function customMiddleware(next) {
    return (c) => {
        c.set("total", 20)

        return next(c)
    }
}

routerAdd("GET", "/js", (c) => {
    const total = c.get("total")

    const records = $app.dao().findRecordsByFilter("posts10k", "title != ''", "-created", total)

    return c.json(200, records)
}, $apis.activityLogger($app), customMiddleware)

onRecordBeforeUpdateRequest((e) => {
    if (e.record.get("title") != "") {
        e.record.set("title", "js_update")
    }
}, "js")
