package nekosrun

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type JSONObject struct {
	data map[string]interface{}
}

type JSONArray struct {
	data   []interface{}
	length int
}

func newJsonResponse(code int, msg string, data interface{}) string {
	response := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	return json_serialize(response, false, nil)
}

func loadJsonString(data []uint8) *JSONObject {
	var obj JSONObject
	mapper := make(map[string]interface{})
	if json.Unmarshal(data, &mapper) != nil {
		return nil
	}
	obj.data = mapper
	return &obj
}

func (obj JSONObject) getInt(key string) int {
	key_data := obj.data[key]
	data, _ := key_data.(float64)
	return int(data)
}

func (obj JSONObject) getFloat(key string) float64 {
	data, _ := obj.data[key].(float64)
	return data
}

func (obj JSONObject) getBool(key string) bool {
	data := obj.data[key].(bool)
	return data
}

func (obj JSONObject) getString(key string) string {
	data := obj.data[key].(string)
	return data
}

func (obj JSONObject) getJSONArray(key string) JSONArray {
	var newobj JSONArray
	newobj.data, _ = obj.data[key].([]interface{})
	newobj.length = len(newobj.data)
	return newobj
}

func (obj JSONObject) getJSONObject(key string) JSONObject {
	var newobj JSONObject
	newobj.data, _ = obj.data[key].(map[string]interface{})
	return newobj
}

func (obj JSONArray) index(key int) JSONObject {
	var newobj JSONObject
	newobj.data, _ = obj.data[key].(map[string]interface{})
	return newobj
}

func (obj JSONObject) exist(key string) bool {
	this_obj := obj.data[key]
	return this_obj != nil
}

func json_serialize(datain interface{}, needsort bool, sortlable *[]string) string {
	var sortList []string
	data, ok := datain.(map[string]interface{})
	if !ok {
		return "{}"
	}

	returnSortedStr := "{"
	if sortlable != nil {
		sortList = *sortlable
	} else {
		sortList := make([]string, 0)
		for tag, _ := range data {
			sortList = append(sortList, tag)
		}
		if needsort {
			sort.Slice(sortList, func(i, j int) bool {
				if strings.Compare(sortList[i], sortList[j]) > 0 {
					return false
				}
				return true
			})
		}
	}

	for idx, tag := range sortList {
		if idx > 0 {
			returnSortedStr += ", "
		}
		if op, ok := data[tag].(string); ok {
			returnSortedStr += "\"" + tag + "\": \"" + op + "\""
		}
		if op, ok := data[tag].(uint); ok {
			returnSortedStr += "\"" + tag + "\": " + fmt.Sprintf("%d", op)
		}
		if op, ok := data[tag].(uint8); ok {
			returnSortedStr += "\"" + tag + "\": " + fmt.Sprintf("%d", op)
		}
		if op, ok := data[tag].(uint16); ok {
			returnSortedStr += "\"" + tag + "\": " + fmt.Sprintf("%d", op)
		}
		if op, ok := data[tag].(uint32); ok {
			returnSortedStr += "\"" + tag + "\": " + fmt.Sprintf("%d", op)
		}
		if op, ok := data[tag].(uint64); ok {
			returnSortedStr += "\"" + tag + "\": " + fmt.Sprintf("%d", op)
		}
		if op, ok := data[tag].(int); ok {
			returnSortedStr += "\"" + tag + "\": " + fmt.Sprintf("%d", op)
		}
		if op, ok := data[tag].(int8); ok {
			returnSortedStr += "\"" + tag + "\": " + fmt.Sprintf("%d", op)
		}
		if op, ok := data[tag].(int16); ok {
			returnSortedStr += "\"" + tag + "\": " + fmt.Sprintf("%d", op)
		}
		if op, ok := data[tag].(int32); ok {
			returnSortedStr += "\"" + tag + "\": " + fmt.Sprintf("%d", op)
		}
		if op, ok := data[tag].(int64); ok {
			returnSortedStr += "\"" + tag + "\": " + fmt.Sprintf("%d", op)
		}
		if op, ok := data[tag].(float32); ok {
			returnSortedStr += "\"" + tag + "\": " + fmt.Sprintf("%f", op)
		}
		if op, ok := data[tag].(float64); ok {
			returnSortedStr += "\"" + tag + "\": " + fmt.Sprintf("%f", op)
		}
		if op, ok := data[tag].(bool); ok {
			if op {
				returnSortedStr += "\"" + tag + "\":\"true\""
			} else {
				returnSortedStr += "\"" + tag + "\":\"false\""
			}
		}
		if op, ok := data[tag].(map[string]interface{}); ok {
			returnSortedStr += "\"" + tag + "\": " + json_serialize(op, true, nil)
		}
		if op, ok := data[tag].([]string); ok {
			returnSortedStr += `"` + tag + `":[`
			if len(op) > 0 {
				idx := 0
				for idx = 0; idx < len(op)-1; idx++ {
					returnSortedStr += `"` + op[idx] + `",`
				}
				returnSortedStr += `"` + op[idx] + `"`
			}
			returnSortedStr += "]"
		}
		if op, ok := data[tag].([]int); ok {
			returnSortedStr += `"` + tag + `": [`
			if len(op) > 0 {
				idx := 0
				for idx = 0; idx < len(op)-1; idx++ {
					returnSortedStr += strconv.Itoa(op[idx]) + ","
				}
				returnSortedStr += strconv.Itoa(op[idx])
			}
			returnSortedStr += "]"
		}
		if op, ok := data[tag].([]interface{}); ok {
			returnSortedStr += `"` + tag + `": [`
			if len(op) > 0 {
				idx := 0
				for idx = 0; idx < len(op)-1; idx++ {
					returnSortedStr += json_serialize(op[idx], true, nil) + ","
				}
				returnSortedStr += json_serialize(op[idx], true, nil)
			}
			returnSortedStr += "]"
		}
	}
	returnSortedStr += "}"
	return returnSortedStr
}
