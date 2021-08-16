package foo

import (
	"bytes"
	"encoding/binary"
	"testing"
)

type data struct {
	foo int32
	bar int32
	x bool
}

func TestAsd(t *testing.T) {
	d := data{1,0x12345678,true}
	exp := []byte{0,0,0,1, 0x12, 0x34,0x56,0x78,  1 }
	
	buf := bytes.Buffer{}
	err := binary.Write(&buf, binary.BigEndian, &d)

	if err != nil {
		t.Fatal("Error occured, got",err)
	}
	
	got := buf.Bytes()

	if len(exp) != len(got) {
		t.Error("Invalid len, got", len(got))
	}
	for i := 0; i < len(exp); i++ {
		if exp[i] != got[i] {
			t.Error("Invalid data on",i)
		}
	}
}