package main

import "fmt"

// let's assume minutes for now
const min = 0
const max = 60

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
	p := &parser{
		it: toIter(toks),
		ranges: []rng{},
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
			
		} else if current.tokType == wildcard {

		} else {
			return nil, fmt.Errorf("unexpected token: %v", current)
		}
		p.it.next() // todo - remove, proceed in own sections
	}
	return p.ranges, nil
}

type parser struct {
	it *iter[token]
	ranges []rng
}

func (p *parser) peekNext(it *iter[token], tokType tokenType) bool {
	next, nextOk := it.peek()
	if !nextOk {
		return false
	}
	return next.tokType == tokType
}