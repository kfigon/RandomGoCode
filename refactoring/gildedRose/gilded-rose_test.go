package main

import (
	"testing"
	"fmt"

)

func TestOneDay(t *testing.T) {
	testCases := []struct {
		in Item
		exp Item
	}{
		{ Item{"+5 Dexterity Vest", 10, 20}, Item{sellIn: 9, quality: 19,}},
		{ Item{"Aged Brie", 2, 0}, Item{sellIn: 1, quality: 1,}},
		{ Item{"Elixir of the Mongoose", 5, 7}, Item{sellIn: 4, quality: 6,}},
		{ Item{"Sulfuras, Hand of Ragnaros", 0, 80}, Item{sellIn: 0, quality: 80,}},
		{ Item{"Sulfuras, Hand of Ragnaros", -1, 80}, Item{sellIn: -1, quality: 80,}},
		{ Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20}, Item{sellIn: 14, quality: 21,}},
		{ Item{"Backstage passes to a TAFKAL80ETC concert", 10, 49}, Item{sellIn: 9, quality: 50,}},
		{ Item{"Backstage passes to a TAFKAL80ETC concert", 5, 49}, Item{sellIn: 4, quality: 50,}},
		{ Item{"Conjured Mana Cake", 3, 6},  Item{sellIn: 2, quality: 5,}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
			processSingleItem(&tc.in)
			if tc.exp.sellIn != tc.in.sellIn {
				t.Errorf("%v sellIn, exp: %v, got: %v",tc.in.name, tc.exp.sellIn,  tc.in.sellIn)
			}
			if tc.exp.quality != tc.in.quality {
				t.Errorf("%v quality, exp: %v, got: %v",tc.in.name, tc.exp.quality,  tc.in.quality)
			}
		})
	}
}