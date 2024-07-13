package main

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

func main() {

	schemaLoader := gojsonschema.NewReferenceLoader("file://./yaml-graph.schema.json")
	documentLoader := gojsonschema.NewReferenceLoader("file://./yaml-graph.json")

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		json, _ := documentLoader.LoadJSON()
		o, _ := yaml.Marshal(json)
		fmt.Printf("%s", o)
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
