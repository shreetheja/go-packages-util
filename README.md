# mongojson

📦 A simple Go utility to pretty-print MongoDB BSON documents and aggregation pipelines in Extended JSON format.

---

mongojson provides a Dump function that prints BSON documents and slices (like aggregation pipelines) in a clean, readable format. It’s built to ease the debugging and inspection of MongoDB queries inside Go applications.

---

## ✨ Features

- ✅ Pretty-prints MongoDB documents and aggregation pipelines
- ✅ Handles top-level slices (e.g., []bson.M) by wrapping in a document
- ✅ Outputs MongoDB Extended JSON (non-canonical mode)
- ✅ Handles marshalling/unmarshalling edge cases gracefully
- ✅ Zero dependencies beyond the official mongo-driver

---

## 📦 Installation

```bash
go get github.com/theja/utils/mongojson
````

---

## 🔧 Usage

```go
package main

import (
    "go.mongodb.org/mongo-driver/bson"
    "github.com/theja/utils/mongojson"
)

func main() {
    pipeline := []bson.M{
        {"$match": bson.M{"status": "active"}},
        {"$group": bson.M{"_id": "$category", "count": bson.M{"$sum": 1}}},
    }

    mongojson.Dump(pipeline)
}
```

Output:

```json
{
  "data": [
    {
      "$match": {
        "status": "active"
      }
    },
    {
      "$group": {
        "_id": "$category",
        "count": {
          "$sum": 1
        }
      }
    }
  ]
}
```

---

## 📂 API

### func Dump(doc interface{})

Dump accepts any BSON-compatible Go value (bson.M, bson.D, \[]bson.M, etc) and prints a pretty-formatted JSON string to stdout.

* If doc is a slice (e.g., \[]bson.M), it will be wrapped inside a document under the key "data" for valid Extended JSON encoding.
* Uses bson.MarshalExtJSON with relaxed formatting (non-canonical).

---

## 🔎 Why?

* bson.MarshalExtJSON does not support top-level arrays.
* Useful during development to inspect aggregation pipelines or nested documents.
* Works well alongside tools like spew, but formats output for MongoDB context.

---

## 🛠 Internals

* Uses reflect to detect slices at the top level
* Converts output to readable JSON via encoding/json.MarshalIndent
* Fallbacks to raw output if pretty-printing fails

---

## ✅ License

MIT © 2025 Shreetheja

---

## 📚 See Also

* [https://pkg.go.dev/go.mongodb.org/mongo-driver/bson](https://pkg.go.dev/go.mongodb.org/mongo-driver/bson)
* [https://docs.mongodb.com/manual/reference/mongodb-extended-json/](https://docs.mongodb.com/manual/reference/mongodb-extended-json/)

