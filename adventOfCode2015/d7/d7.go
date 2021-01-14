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

	twoArgOperation := parsePattern(input, `(\w+)?(\d+)? (\w+) (\w+)?(\d+)? -> (\w+)`)
	if len(twoArgOperation) != 0 {
		
	}

	// targetRegister := parts[1]
	// firstPart := parts[0]
	

	return operation{}, nil
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