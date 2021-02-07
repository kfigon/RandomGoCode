package main

type Item struct {
	name            string
	sellIn, quality int
}

func processSingleItem(item *Item) {
	wrapper := createWrapper(item.name)

	wrapper.updateQuality(item)
	wrapper.updateSellIn(item)
	if item.sellIn < 0 {
		wrapper.updateQualityAfterSellinPass(item)
	}
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		processSingleItem(items[i])
	}
}

func createWrapper(name string) itemWrapper {
	switch name {
	case "Aged Brie": return agedBrieItem{}
	case "Backstage passes to a TAFKAL80ETC concert": return backstagePassItem{}
	case "Sulfuras, Hand of Ragnaros": return sulfurasItem{}
	default: return regularItem{}
	}

}

func updateSellIn(item *Item) {
	item.sellIn = item.sellIn - 1
}

func increaseQuality(item *Item) {
	if item.quality < 50 {
		item.quality = item.quality + 1
	}
}

func decreaseQuality(item *Item) {
	if item.quality > 0 {
		item.quality = item.quality - 1
	}
}

type itemWrapper interface {
	updateQuality(item *Item)
	updateSellIn(item *Item)
	updateQualityAfterSellinPass(item *Item)
}

type regularItem struct{}
func (r regularItem) updateQuality(item *Item) {
	decreaseQuality(item)
}

func (r regularItem) updateSellIn(item *Item) {
	updateSellIn(item)
}

func (r regularItem) updateQualityAfterSellinPass(item *Item) {
	decreaseQuality(item)
}

type agedBrieItem struct{}
func (a agedBrieItem) updateQuality(item *Item) {
	increaseQuality(item)
}

func (a agedBrieItem) updateSellIn(item *Item) {
	updateSellIn(item)
}

func (a agedBrieItem) updateQualityAfterSellinPass(item *Item) {
	increaseQuality(item)
}

type backstagePassItem struct{}
func (b backstagePassItem) updateQuality(item *Item) {
	howManyUpdates := b.howManyQualityIncreases(item.sellIn)
	for i := 0; i < howManyUpdates; i++ {
		increaseQuality(item)
	}
}

func (b backstagePassItem) howManyQualityIncreases(sellIn int) int {
	if sellIn < 6 {
		return 3
	} else if sellIn < 11 {
		return 2
	}
	return 1
}

func (b backstagePassItem) updateSellIn(item *Item) {
	updateSellIn(item)
}

func (b backstagePassItem) updateQualityAfterSellinPass(item *Item) {
	item.quality = 0
}

type sulfurasItem struct{}
func (s sulfurasItem) updateQuality(item *Item) {
	increaseQuality(item)
}

func (s sulfurasItem) updateSellIn(item *Item) {
	return // do nothing
}

func (s sulfurasItem) updateQualityAfterSellinPass(item *Item) {
	return // do nothing
}