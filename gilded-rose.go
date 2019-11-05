package main

type Item struct {
	Name	string
	SellIn	int
	Quality int
}

// refer to https://github.com/emilybache/GildedRose-Refactoring-Kata/blob/master/go/gilded-rose.go

const (
	agedBrie = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras = "Sulfuras, Hand of Ragnaros"
)

func updateQuality(items []*Item) {
	for i := range items {
		if items[i].Name != agedBrie && items[i].Name != backstagePasses {
			if items[i].Name != sulfuras {
				items[i].SellIn--
				if items[i].Quality > 0 {
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

		if items[i].SellIn < 0 {
			if items[i].Name != agedBrie {
				if items[i].Name != backstagePasses {
					if items[i].Name != sulfuras && items[i].Quality > 0 {
						items[i].Quality--
					}
				} else {
					items[i].Quality = 0
				}
			} else if items[i].Quality < 50 {
				items[i].Quality++
			}
		}
	}
}