package main

import (
	"fmt"
)

// https://adventofcode.com/2015/day/7

type processor struct{
	regs map[string]uint16
}

func newProcessor() *processor {
	return &processor{
		regs: make(map[string]uint16),
	}
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

func asd()  {
	fmt.Println("asd")
}