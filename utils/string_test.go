package utils

import "testing"

func TestString(t *testing.T) {
	
	// t.Error("")
	var params = map[string]string {
		"b": "BB",
		"a": "AA",
		"c": "CC",
	}
	println(ParamsSortToUrl(params, []string{"b"}))

}