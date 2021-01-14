package main

import (
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
	reg := regexp.MustCompile(`(\d+) -> (\w+)`)
	loadPattern := reg.FindAllString(input, -1)
	if loadPattern != nil && len(loadPattern) == 0 {
		return operation{operation: load, arg1: 0, destinationReg: "a"}, nil
	}
	return operation{}, nil
}

func asd()  {
	fmt.Println("asd")
}