package main

import (
	"fmt"
	"strconv"
)

// <start;stop>
type rng struct {
	start int
	stop int
	step int
}

func eval(toks []token) ([]int, error) {
	ranges, err := parse(toks)
	if err != nil {
		return nil, err
	}

	out := []int{}
	for _, r := range ranges {
		for i := r.start; i <= r.stop; i+=r.step {
			out = append(out, i)
		}
	}
	return out,nil
}

func parse(toks []token) ([]rng, error) {
	// let's assume minutes for now
	p := &parser{
		it: toIter(toks),
		ranges: []rng{},
		min: 0,
		max: 59,
	}
	return p.parse()
}

func (p *parser) parse() ([]rng, error) {
	for {
		current, ok := p.it.current()
		if !ok {
			break
		}
		
		if current.tokType == number {
			if err := p.parseNumber(); err != nil {
				return nil, err
			}
		} else if current.tokType == wildcard {
			if err := p.parseWildcard(); err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("unexpected token: %v", current)
		}
	}
	return p.ranges, nil
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
	v, ok := p.it.current()
	if !ok {
		return nil
	}

	num1, _ := strconv.Atoi(v.lexeme)

	p.it.consume() // number

	current, ok := p.it.current()
	if !ok {
		// just number
		p.ranges = append(p.ranges, rng{start: num1, stop: num1, step: 1})
		return nil
	}

	if current.tokType == comma {
		p.it.consume() // ,
		err := p.parseNumber()
		if err != nil {
			return err
		}
	} else if current.tokType == dash {
		p.it.consume() // -
		current, ok = p.it.current()
		if !ok {
			return fmt.Errorf("unexpected end of tokens when parsing dash")
		} else if current.tokType != number {
			return fmt.Errorf("unexpected token when parsing dash, expected number, got %v", current)
		}
		num2, _ := strconv.Atoi(v.lexeme)
		p.ranges = append(p.ranges, rng{start: num1, stop: num2, step: 1})
	} else if current.tokType == div {
		if err := p.parseDiv(); err != nil {
			return err
		}
	}
	return nil
}

func (p *parser) parseWildcard() error {
	p.it.consume() // *
	current, ok := p.it.current()
	if !ok {
		p.ranges = append(p.ranges, rng{start: p.min, stop: p.max, step: 1})
		return nil
	}

	if current.tokType == div {
		if err := p.parseDiv(); err != nil {
			return err
		}
	}

	return fmt.Errorf("error during parsing wildcard, unexpected token %v", current)
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
	return nil
}