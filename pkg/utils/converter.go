package utils

import "encoding/json"

func StructToMap(obj interface{}) map[string]interface{} {
	objMap := make(map[string]interface{})
	jsonObj, _ := json.Marshal(obj)
	json.Unmarshal(jsonObj, &objMap)
	return objMap
}
