package main

import "testing"

func TestEmpty(t *testing.T) {
	p := newProcessor()
	if v := p.readRegistry("x"); v != 0 {
		t.Errorf("Empty read, expected 0, got %v",v)
	}
}

func TestLoad(t *testing.T) {
	p := newProcessor()
	p.load(123, "x")
	v := p.readRegistry("x")
	if v != 123 {
		t.Errorf("Expected %v, got %v", 123, v)
	}
}
// x LSHIFT 2 -> f
// y RSHIFT 2 -> g
// NOT x -> h
// NOT y -> i

func TestExample(t *testing.T) {
	p := newProcessor()
	p.load(123, "x")
	p.load(456, "y")
	p.and(p.readRegistry("x"), p.readRegistry("y"), "d")
	p.or(p.readRegistry("x"), p.readRegistry("y"), "e")
	p.lshift(p.readRegistry("x"), 2, "f")
	p.rshift(p.readRegistry("y"), 2, "g")
	p.not(p.readRegistry("x"), "h")
	p.not(p.readRegistry("y"), "i")

	cmp := func (reg string, exp uint16) {
		v := p.readRegistry(reg)
		if v != exp {
			t.Errorf("Expected %v, got %v in register %q", exp, v, reg)
		}
	}

	expected := make(map[string]uint16)
	expected["d"] = 72
	expected["e"] = 507
	expected["f"] = 492
	expected["g"] = 114
	expected["h"] = 65412
	expected["i"] = 65079
	expected["x"] = 123
	expected["y"] = 456

	for key := range expected {
		cmp(key, expected[key])
	}
}