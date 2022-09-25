package main

import (
	"fmt"
	"sort"
)

func encodeStr(in string) string {
	return fmt.Sprintf("%v:%v", len(in), in)
}

func encodeInt(in int) string {
	return fmt.Sprintf("i%ve", in)
}

func encodeList(in []any) string {
	out := "l"
	for _, v := range in {
		out += encodeAnyType(v)
	}
	return out+"e"
}

func encodeDict(in map[string]any) string {
	out := "d"
	type entry struct {
		key string
		val any
	}
	listed := []entry{}
	for k, v := range in {
		listed = append(listed, entry{key: k, val: v})
	}
	sort.Slice(listed, func(i, j int) bool {
		return listed[i].key < listed[j].key
	})
	for _,v := range listed {
		out += encodeStr(v.key) + encodeAnyType(v.val)
	}
	return out + "e"
}

func encodeAnyType(v any) string {
	switch val := v.(type){
		case int: return encodeInt(val)
		case string: return encodeStr(val)
		case []any: return encodeList(val)
		case map[string]any: return encodeDict(val)
	}
	return ""
}