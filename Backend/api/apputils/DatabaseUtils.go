package apputils

import "strconv"

func IdToString(id int64) string {
	return strconv.FormatInt(id, 10)
}
