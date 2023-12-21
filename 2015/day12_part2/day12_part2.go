package day12_part2

import (
	"encoding/json"
)

var Data any

func getAnswer(input string) int {
	err := json.Unmarshal([]byte(input), &Data)
	if err != nil {
		panic(err)
	}
	sum := count(Data)
	//fmt.Printf("Output: %v\n", sum)
	return int(sum)
}

func count(i any) float64 {

	sum := 0.0
	switch i.(type) {
	case float64:
		sum += i.(float64)
	case []interface{}:
		arr := i.([]interface{})
		for i := 0; i < len(arr); i++ {
			sum += count(arr[i])
		}
	case map[string]interface{}:
		obj := i.(map[string]interface{})
		skip := false
		for _, v := range obj {
			if v == "red" {
				skip = true
				break
			}
		}
		if !skip {
			for _, v := range obj {
				sum += count(v)
			}
		}
	}

	return sum
}
