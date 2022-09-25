package main

func decode(in string) bencodeObj {
	return nil
}

type bencodeObj interface{
	dummy()
}

type stringObj string
func (_ stringObj) dummy(){}

type intObj int
func (_ intObj) dummy(){}

type listObj []any
func (_ listObj) dummy(){}

type dictObj map[string]any
func (_ dictObj) dummy(){}