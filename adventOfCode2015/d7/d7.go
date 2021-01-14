package main

import (
	"strconv"
	"strings"
	"fmt"
	"regexp"
)

// https://adventofcode.com/2015/day/7

type processor struct{
	regs map[string]uint16
}

func newProcessor() *processor {
	return &processor{ regs: make(map[string]uint16) }
}

func (p *processor) readRegistry(reg string) uint16 {
	return p.regs[reg]
}

func (p *processor) load(val uint16, reg string)  {
	p.regs[reg] = val
}

func (p *processor) and(arg1, arg2 uint16, destinationReg string) {
	p.regs[destinationReg] = arg1 & arg2
}

func (p *processor) or(arg1, arg2 uint16, destinationReg string) {
	p.regs[destinationReg] = arg1 | arg2
}

func (p *processor) lshift(arg1, arg2 uint16, destinationReg string) {
	p.regs[destinationReg] = arg1 << arg2
}

func (p *processor) rshift(arg1, arg2 uint16, destinationReg string) {
	p.regs[destinationReg] = arg1 >> arg2
}

func (p *processor) not(arg uint16, destinationReg string) {
	p.regs[destinationReg] = ^arg
}

const (
	and  = iota
	load
	or 
	lshift 
	rshift 
	not
)

type operation struct {
	operation int
	arg1 uint16
	arg2 uint16

	sourceReg1 string
	sourceReg2 string
	destinationReg string
}

func parseLine(input string) (operation, error) {
	if len(input) == 0 {
		return operation{}, fmt.Errorf("Empty input provided")
	}
	parts := strings.Split(input, " -> ")
	if parts == nil || len(parts) != 2 {
		return operation{}, fmt.Errorf("Invalid input provided, no -> separator in %q", input)
	}

	targetRegister := parts[1]
	firstPart := parts[0]

	op := operation{}
	if strings.Contains(" AND ", firstPart) {
		op.operation = and
	} else if strings.Contains(" OR ", firstPart) {
		op.operation = or
	} else if strings.Contains(" LSHIFT ", firstPart) {
		op.operation = lshift
	} else if strings.Contains(" RSHIFT ", firstPart) {
		op.operation = rshift
	} else if strings.Contains("NOT ", firstPart) {
		op.operation = not
	} else if arg, ok := isLoadOperation(firstPart); ok {
		op.operation = load
		op.arg1 = arg
	} else {
		return operation{}, fmt.Errorf("Invalid operation in %q", input)
	}

	return operation{}, nil
}

func isLoadOperation(input string) (uint16, bool) {
	v, err := strconv.Atoi(input)
	if err != nil {
		return 0, false
	}
	return uint16(v), true
}

func asd()  {
	fmt.Println("asd")
}