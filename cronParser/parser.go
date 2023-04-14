package main

import (
	"fmt"
	"strconv"
)

// <start;stop>
type rng struct {
	start int
	stop int
}

func eval(toks []token) ([]int, error) {
	ranges,step, err := parse(toks)
	if err != nil {
		return nil, err
	}

	out := []int{}
	for _, r := range ranges {
		for i := r.start; i <= r.stop; i+=step {
			out = append(out, i)
		}
	}
	return out,nil
}

func parse(toks []token) ([]rng, int, error) {
	// let's assume minutes for now
	p := &parser{
		it: toIter(toks),
		ranges: []rng{},
		min: 0,
		max: 59,
	}
	return p.parse()
}

func (p *parser) parse() ([]rng, int, error) {
	for {
		current, ok := p.it.current()
		if !ok {
			break
		}
		
		if current.tokType == number {
			if err := p.parseNumber(); err != nil {
				return nil, 0, err
			}
		} else if current.tokType == wildcard {
			if err := p.parseWildcard(); err != nil {
				return nil, 0, err
			}
		}else if current.tokType == div {
			if err := p.parseDiv(); err != nil {
				return nil,0, err
			}
		} else {
			return nil, 0, fmt.Errorf("unexpected token: %v", current)
		}
	}
	step := 1
	if p.divisor != nil {
		step = *p.divisor
	}
	return p.ranges, step, nil
}

type parser struct {
	it *iter[token]
	ranges []rng
	divisor *int
	min int
	max int
}

func (p *parser) peekNext(tokType tokenType) bool {
	next, nextOk := p.it.peek()
	if !nextOk {
		return false
	}
	return next.tokType == tokType
}

func (p *parser) parseNumber() error {
	current, ok := p.it.current()
	num1, _ := strconv.Atoi(current.lexeme)
	if !ok {
		p.ranges = append(p.ranges, rng{start: num1, stop: num1})
		return nil
	}
	p.it.consume() // num
	current, ok = p.it.current()

	if ok && current.tokType == dash {
		p.it.consume() // -
		current, ok := p.it.current()
		if !ok {
			return fmt.Errorf("unexpected end of tokens when parsing dash")
		} else if current.tokType != number {
			return fmt.Errorf("unexpected token when parsing dash, expected number, got %v", current)
		}
		num2, _ := strconv.Atoi(current.lexeme)
		p.ranges = append(p.ranges, rng{start: num1, stop: num2})
		
		if p.peekNext(comma) {
			p.it.consume() // ,	
		}
		p.it.consume() // num2
	} else if ok && current.tokType == comma {
		p.it.consume() // ,
		p.ranges = append(p.ranges, rng{start: num1, stop: num1})
	} else {
		p.ranges = append(p.ranges, rng{start: num1, stop: num1})
	}
	return nil
}

func (p *parser) parseWildcard() error {
	p.it.consume() // *
	p.ranges = append(p.ranges, rng{start: p.min, stop: p.max})
	return nil
}

func (p *parser) parseDiv() error {
	p.it.consume() // /
	current, ok := p.it.current()

	if !ok {
		return fmt.Errorf("unexpected end when parsing div, expected number")
	} else if current.tokType != number {
		return fmt.Errorf("unexpected token when parsing div, expected number, got %v", current)
	} else if p.divisor != nil {
		return fmt.Errorf("divisor already set, previous %v, this: %v", *p.divisor, current)
	}

	num, _ := strconv.Atoi(current.lexeme)
	p.divisor = &num
	p.it.consume() // num
	return nil
}