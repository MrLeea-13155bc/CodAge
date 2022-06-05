package Utils

import (
	"strconv"
	"strings"
)

func AddOptionsId(ids string, firstId int64) string {
	result := strings.Split(ids, ",")
	for key, item := range result {
		id, _ := strconv.ParseInt(item, 10, 64)
		result[key] = strconv.FormatInt(id+firstId, 10)
	}
	return strings.Join(result, ",")
}
