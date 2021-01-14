package main

import (
	"strconv"
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

	if op, isOk := parseTwoArgCommand(input); isOk {
		return op, nil
	} else if op, isOk := parseLoad(input); isOk {
		return op, nil
	} else if op, isOk := parseNot(input); isOk {
		return op, nil
	}

	return operation{}, fmt.Errorf("invalid input %q", input)
}

func parseTwoArgCommand(input string) (operation, bool) {
	twoArgOperation := parsePattern(input, `(\w+)?(\d+)? (\w+) (\w+)?(\d+)? -> (\w+)`)
	if twoArgOperation == nil || len(twoArgOperation) != 6 {
		return operation{}, false
	}

	parseOp := func(x string) int {
		switch x {
		case "AND": return and
		case "OR": return or
		case "LSHIFT": return lshift
		case "RSHIFT": return rshift
		}
		return -1
	}

	opCode := parseOp(twoArgOperation[2])
	if opCode == -1 {
		return operation{}, false
	}

	arg1 := parseInt(twoArgOperation[1])
	arg2 := parseInt(twoArgOperation[4])
	sourceReg1 := twoArgOperation[0]
	sourceReg2 := twoArgOperation[3]
	destReg := twoArgOperation[5]

	return operation{
		arg1: arg1,
		arg2: arg2,
		operation: opCode,
		sourceReg1: sourceReg1,
		sourceReg2: sourceReg2,
		destinationReg: destReg,
	}, true
}

func parseInt(x string) uint16 {
	v, err := strconv.Atoi(x)
	if err != nil {
		return 0
	}
	return uint16(v)
}

func parseLoad(input string) (operation, bool) {
	loadOperation := parsePattern(input, `(\d+) -> (\w+)`)
	if loadOperation == nil || len(loadOperation) != 2 {
		return operation{}, false
	}

	return operation{
		operation: not,
		arg1: parseInt(loadOperation[0]),
		destinationReg: loadOperation[1],
	}, true
}

func parseNot(input string) (operation, bool) {
	notOperation := parsePattern(input, `NOT (\w+) -> (\w+)`)
	if notOperation == nil || len(notOperation) != 2 {
		return operation{}, false
	}

	return operation{
		sourceReg1: notOperation[0],
		destinationReg: notOperation[1],
	}, true
}

func parsePattern(input string, pattern string) ([]string) {
	reg := regexp.MustCompile(pattern)
	result := reg.FindAllStringSubmatch(input, -1)
	if result == nil || len(result) == 0 {
		return []string{}
	}
	return result[0]
}

func asd()  {
	fmt.Println("asd")
}