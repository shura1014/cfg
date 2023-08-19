package goproperties

import "testing"

func Test_PropToMap(t *testing.T) {
	data := make(map[string]any)
	propToMap(data, []string{"a", "b", "c"}, "1")
	t.Log(data)
}

func Test_PropToMap2(t *testing.T) {
	a := make(map[string]any)
	b := make(map[string]any)
	c := make(map[string]any)
	a["a"] = b
	b["b"] = c
	c["c"] = 2
	t.Log(a)
	propToMap(a, []string{"a", "b", "c"}, "1")
	t.Log(a)
}

func Test_MapToProp(t *testing.T) {
	a := make(map[string]any)
	b := make(map[string]any)
	c := make(map[string]any)
	a["a"] = b
	b["b"] = c
	c["c"] = 1
	//t.Log(a)

	receive := make(map[string]any)
	mapToProp(receive, a, "", ".")
	t.Log(receive)
}
