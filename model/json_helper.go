package model

import (
	"encoding/json"
	"fmt"
)

func FromJSON(jsonStr string, prefix string) ([]map[string]interface{}, error) {
	var data map[string]interface{}
	var result []map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	processJSON(data, prefix, &result)
	return result, nil
}

func processJSON(input map[string]interface{}, prefix string, result *[]map[string]interface{}) {
	for key, value := range input {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}

		switch v := value.(type) {
		case map[string]interface{}:

			processJSON(v, fullKey, result)
		case []interface{}:

			for i, item := range v {
				itemKey := fmt.Sprintf("%s[%d]", fullKey, i)
				switch itemValue := item.(type) {
				case map[string]interface{}:
					processJSON(itemValue, itemKey, result)
				default:
					*result = append(*result, map[string]interface{}{
						"key":   itemKey,
						"value": itemValue,
					})
				}
			}
		default:

			*result = append(*result, map[string]interface{}{
				"key":   fullKey,
				"value": v,
			})
		}
	}
}

func ParseKVtoString(kvArray []map[string]interface{}) string {
	var result string
	for _, kv := range kvArray {
		result += fmt.Sprintf("%s: %v\n", kv["key"], kv["value"])
	}
	return result
}
