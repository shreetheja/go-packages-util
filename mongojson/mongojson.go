// Package mongojson prints MongoDB documents and pipelines as readable Extended JSON.
package mongojson

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

// Dump prints a MongoDB doc or pipeline in readable, pretty Extended JSON (non-canonical).
// Dump prints a pretty-printed MongoDB document or pipeline (Extended JSON, non-canonical).
func Dump(doc interface{}) {
	var extJSON []byte
	var err error

	// If top-level value is a slice, wrap in a document
	v := reflect.ValueOf(doc)
	if v.Kind() == reflect.Slice {
		doc = bson.M{"data": doc} // wrapped for valid BSON
	}

	// Use non-canonical Extended JSON (no $numberInt)
	extJSON, err = bson.MarshalExtJSON(doc, false, true)
	if err != nil {
		log.Printf("mongojson: failed to marshal: %v", err)
		return
	}

	// Pretty-print the JSON
	var pretty interface{}
	if err := json.Unmarshal(extJSON, &pretty); err != nil {
		log.Printf("mongojson: failed to unmarshal for pretty print: %v", err)
		fmt.Println(string(extJSON))
		return
	}

	out, err := json.MarshalIndent(pretty, "", "  ")
	if err != nil {
		log.Printf("mongojson: failed to indent JSON: %v", err)
		fmt.Println(string(extJSON))
		return
	}

	fmt.Println(string(out))
}
