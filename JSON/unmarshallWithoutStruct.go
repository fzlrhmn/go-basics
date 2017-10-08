package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Baz struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	RawData string `json:"rawdata"`
	Epoch   string `json:"epoch"`
}

func main() {
	str := `{ "foo" : { "id":"234", "type":"bar", "rawdata":[{"key":"dog", "values":[123,234]}, {"key":"cat", "values":[23, 45]}], "epoch":"1433120656705"}}`
	str2 := `{ "foo": { "id":"124", "type":"baz", "rawdata":[123, 345,345],"epoch": "1433120656704"}}`

	var objmap map[string]*json.RawMessage
	if err := json.Unmarshal([]byte(str), &objmap); err != nil {
		panic(err)
	}

	typ, err := getTypeByUnmarshalToInterface(objmap["foo"])
	if err != nil {
		panic(err)
	}

	fmt.Println("type in str: ", typ)

	if err := json.Unmarshal([]byte(str2), &objmap); err != nil {
		panic(err)
	}

	typ2, err := getTypeByRegexp(objmap["foo"])
	if err != nil {
		panic(err)
	}

	fmt.Println("type in str2: ", typ2)
}

func getTypeByUnmarshalToInterface(foo *json.RawMessage) (string, error) {
	var tmp map[string]interface{}
	if err := json.Unmarshal(*foo, &tmp); err != nil {
		return "", err
	}

	if typ, ok := tmp["type"].(string); ok {
		return typ, nil
	} else {
		return "", fmt.Errorf("type should be a string")
	}
}

func getTypeByRegexp(foo *json.RawMessage) (string, error) {
	re := regexp.MustCompile(`"type"\s*:\s*"([^"]*)"`)
	matches := re.FindAllSubmatch(*foo, -1)
	if len(matches) > 0 {
		return string(matches[0][1]), nil
	} else {
		return "", fmt.Errorf("cannot find type")
	}
}
