package benchmarks

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

const schema = `
[
    {
        "id": "mum324voxqe7tp4",
        "name": "permissions",
        "type": "base",
        "system": false,
        "schema": [
            {
                "id": "qualee2c",
                "name": "name",
                "type": "text",
                "system": false,
                "required": false,
                "options": {
                    "min": null,
                    "max": null,
                    "pattern": ""
                }
            },
            {
                "id": "svbmsok7",
                "name": "active",
                "type": "bool",
                "system": false,
                "required": false,
                "options": {}
            }
        ],
        "indexes": [],
        "listRule": "",
        "viewRule": "",
        "createRule": "",
        "updateRule": "",
        "deleteRule": "",
        "options": {}
    },
    {
        "id": "mfslsmb563utlfe",
        "name": "organizations",
        "type": "base",
        "system": false,
        "schema": [
            {
                "id": "59ivonnc",
                "name": "name",
                "type": "text",
                "system": false,
                "required": false,
                "options": {
                    "min": null,
                    "max": null,
                    "pattern": ""
                }
            }
        ],
        "indexes": [],
        "listRule": "",
        "viewRule": "",
        "createRule": "",
        "updateRule": "",
        "deleteRule": "",
        "options": {}
    },
    {
        "id": "_pb_users_auth_",
        "name": "users",
        "type": "auth",
        "system": false,
        "schema": [
            {
                "id": "dzrx3l87",
                "name": "organization",
                "type": "relation",
                "system": false,
                "required": true,
                "options": {
                    "collectionId": "mfslsmb563utlfe",
                    "cascadeDelete": true,
                    "minSelect": null,
                    "maxSelect": 1,
                    "displayFields": []
                }
            },
            {
                "id": "8tmffoqz",
                "name": "permissions",
                "type": "relation",
                "system": false,
                "required": false,
                "options": {
                    "collectionId": "mum324voxqe7tp4",
                    "cascadeDelete": true,
                    "minSelect": null,
                    "maxSelect": null,
                    "displayFields": []
                }
            }
        ],
        "indexes": [],
        "listRule": "",
        "viewRule": "",
        "createRule": "",
        "updateRule": "",
        "deleteRule": "",
        "options": {
            "allowEmailAuth": true,
            "allowOAuth2Auth": true,
            "allowUsernameAuth": true,
            "exceptEmailDomains": null,
            "manageRule": null,
            "minPasswordLength": 8,
            "onlyEmailDomains": null,
            "requireEmail": false
        }
    },
    {
        "id": "8tdh70gfoplhw4i",
        "name": "posts25k",
        "type": "base",
        "system": false,
        "schema": [
            {
                "id": "c1vgnelo",
                "name": "title",
                "type": "text",
                "system": false,
                "required": false,
                "options": {
                    "min": null,
                    "max": null,
                    "pattern": ""
                }
            },
            {
                "id": "wjowv38k",
                "name": "description",
                "type": "editor",
                "system": false,
                "required": false,
                "options": {}
            },
            {
                "id": "zdxzgkue",
                "name": "public",
                "type": "bool",
                "system": false,
                "required": false,
                "options": {}
            },
            {
                "id": "jv4eb3km",
                "name": "type",
                "type": "select",
                "system": false,
                "required": false,
                "options": {
                    "maxSelect": 3,
                    "values": [
                        "a",
                        "b",
                        "c",
                        "d"
                    ]
                }
            },
            {
                "id": "3xqeysek",
                "name": "author",
                "type": "relation",
                "system": false,
                "required": false,
                "options": {
                    "collectionId": "_pb_users_auth_",
                    "cascadeDelete": true,
                    "minSelect": null,
                    "maxSelect": 1,
                    "displayFields": []
                }
            }
        ],
        "indexes": [],
        "listRule": "",
        "viewRule": "",
        "createRule": "",
        "updateRule": "",
        "deleteRule": "",
        "options": {}
    },
    {
        "id": "3wb14jxop88yw2j",
        "name": "posts50k",
        "type": "base",
        "system": false,
        "schema": [
            {
                "id": "55lheksn",
                "name": "title",
                "type": "text",
                "system": false,
                "required": false,
                "options": {
                    "min": null,
                    "max": null,
                    "pattern": ""
                }
            },
            {
                "id": "tld6qqg0",
                "name": "description",
                "type": "editor",
                "system": false,
                "required": false,
                "options": {}
            },
            {
                "id": "prb0aw0g",
                "name": "public",
                "type": "bool",
                "system": false,
                "required": false,
                "options": {}
            },
            {
                "id": "b2uhk9hr",
                "name": "type",
                "type": "select",
                "system": false,
                "required": false,
                "options": {
                    "maxSelect": 3,
                    "values": [
                        "a",
                        "b",
                        "c",
                        "d"
                    ]
                }
            },
            {
                "id": "undhpctz",
                "name": "author",
                "type": "relation",
                "system": false,
                "required": false,
                "options": {
                    "collectionId": "_pb_users_auth_",
                    "cascadeDelete": true,
                    "minSelect": null,
                    "maxSelect": 1,
                    "displayFields": []
                }
            }
        ],
        "indexes": [],
        "listRule": "",
        "viewRule": "",
        "createRule": "",
        "updateRule": "",
        "deleteRule": "",
        "options": {}
    },
    {
        "id": "fzbkdgh2fr7ak3z",
        "name": "posts10k",
        "type": "base",
        "system": false,
        "schema": [
            {
                "id": "vvmhpomj",
                "name": "title",
                "type": "text",
                "system": false,
                "required": false,
                "options": {
                    "min": null,
                    "max": null,
                    "pattern": ""
                }
            },
            {
                "id": "jddyngzm",
                "name": "description",
                "type": "editor",
                "system": false,
                "required": false,
                "options": {}
            },
            {
                "id": "utwvk5sk",
                "name": "public",
                "type": "bool",
                "system": false,
                "required": false,
                "options": {}
            },
            {
                "id": "liuxouuk",
                "name": "type",
                "type": "select",
                "system": false,
                "required": false,
                "options": {
                    "maxSelect": 3,
                    "values": [
                        "a",
                        "b",
                        "c",
                        "d"
                    ]
                }
            },
            {
                "id": "e6uxsuqu",
                "name": "author",
                "type": "relation",
                "system": false,
                "required": false,
                "options": {
                    "collectionId": "_pb_users_auth_",
                    "cascadeDelete": true,
                    "minSelect": null,
                    "maxSelect": 1,
                    "displayFields": []
                }
            }
        ],
        "indexes": [],
        "listRule": "",
        "viewRule": "",
        "createRule": "",
        "updateRule": "",
        "deleteRule": "",
        "options": {}
    },
    {
        "id": "n1ai0r1nhwpcup3",
        "name": "posts100k",
        "type": "base",
        "system": false,
        "schema": [
            {
                "id": "wpdot1nx",
                "name": "title",
                "type": "text",
                "system": false,
                "required": false,
                "options": {
                    "min": null,
                    "max": null,
                    "pattern": ""
                }
            },
            {
                "id": "k3cmuwgh",
                "name": "description",
                "type": "editor",
                "system": false,
                "required": false,
                "options": {}
            },
            {
                "id": "9dif65ys",
                "name": "public",
                "type": "bool",
                "system": false,
                "required": false,
                "options": {}
            },
            {
                "id": "uds0lj60",
                "name": "type",
                "type": "select",
                "system": false,
                "required": false,
                "options": {
                    "maxSelect": 3,
                    "values": [
                        "a",
                        "b",
                        "c",
                        "d"
                    ]
                }
            },
            {
                "id": "ne5qvfgj",
                "name": "author",
                "type": "relation",
                "system": false,
                "required": false,
                "options": {
                    "collectionId": "_pb_users_auth_",
                    "cascadeDelete": true,
                    "minSelect": null,
                    "maxSelect": 1,
                    "displayFields": []
                }
            }
        ],
        "indexes": [],
        "listRule": "",
        "viewRule": "",
        "createRule": "",
        "updateRule": "",
        "deleteRule": "",
        "options": {}
    },
    {
        "id": "bkl56v8bbsb2jvg",
        "name": "benchmarks",
        "type": "base",
        "system": false,
        "schema": [
            {
                "id": "dfbuw9mn",
                "name": "tests",
                "type": "text",
                "system": false,
                "required": false,
                "options": {
                    "min": null,
                    "max": null,
                    "pattern": ""
                }
            },
            {
                "id": "cbnu7mwn",
                "name": "result",
                "type": "text",
                "system": false,
                "required": false,
                "options": {
                    "min": null,
                    "max": null,
                    "pattern": ""
                }
            },
            {
                "id": "zawij6ef",
                "name": "error",
                "type": "text",
                "system": false,
                "required": false,
                "options": {
                    "min": null,
                    "max": null,
                    "pattern": ""
                }
            }
        ],
        "indexes": [],
        "listRule": null,
        "viewRule": null,
        "createRule": null,
        "updateRule": null,
        "deleteRule": null,
        "options": {}
    }
]
`

func (r *runner) resetSchema(deleteData bool) error {
	return r.app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(schema), &collections); err != nil {
			return err
		}

		if err := txDao.ImportCollections(collections, true, nil); err != nil {
			return err
		}

		if deleteData {
			for _, c := range collections {
				if c.Name == colBenchmarks {
					continue
				}
				if _, err := txDao.DB().NewQuery("DELETE FROM " + c.Name).Execute(); err != nil {
					return err
				}
			}
		}

		return nil
	})
}
