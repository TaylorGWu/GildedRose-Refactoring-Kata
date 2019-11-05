package main

type Item struct {
	Name	string
	SellIn	int
	Quality int
}

const (
	agedBrie = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras = "Sulfuras, Hand of Ragnaros"
)

func updateQuality(items []*Item) {
	for i := range items {
		if items[i].Name != agedBrie && items[i].Name != backstagePasses {
			if items[i].Quality > 0 {
				if items[i].Name != sulfuras {
					items[i].Quality--
				}
			}
		} else {
			if items[i].Quality < 50 {
				items[i].Quality++
				if items[i].Name == backstagePasses {
					if items[i].SellIn < 11 {
						items[i].Quality++
					}
				}
			}
		}
	}
}