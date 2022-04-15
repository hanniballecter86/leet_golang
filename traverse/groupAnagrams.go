package traverse

import (
	"reflect"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	res := [][]string{{}}
	sets := []map[string]bool{}
	for _, v := range strs {
		tmp_sets := string2set(v)
		for i, set := range sets {
			if reflect.DeepEqual(set, tmp_sets) {
				res[i] = append(res[i], v)
			} else {
				sets = append(sets, tmp_sets)
				tmp_val := []string{}
				tmp_val = append(tmp_val, v)
				res = append(res, tmp_val)
			}
		}
	}
	return res
}

func string2set(str string) map[string]bool {
	str_val := strings.Split(str, "")
	str_set := map[string]bool{}
	for _, v := range str_val {
		str_set[v] = true
	}
	return str_set
}
