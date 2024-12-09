package sort

import (
	"strconv"
	"strings"
)

func Intar(s string) []int {
	var ints []int
	arr := strings.Split(s, " ")
	for _, s := range arr {
		num, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, num)
		}
	}
	return ints
}
