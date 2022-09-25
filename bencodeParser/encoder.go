package main

import "fmt"

func encodeStr(in string) string {
	return fmt.Sprintf("%v:%v", len(in), in)
}

func encodeInt(in int) string {
	return fmt.Sprintf("i%ve", in)
}

func encodeList(in []any) string {
	out := "l"
	for _, v := range in {
		switch v.(type){
		case int:
		case string:
		case []any:
		case map[string]any:
		}
	}
	return out+"e"
}

func encodeDict(in map[string]any) string {
	return ""
}