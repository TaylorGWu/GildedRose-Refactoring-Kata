package main

import (
	"errors"
	"testing"
)

const (
	DAY1  = 1
	DAY10 = 9
	DAY20 = 10
)

func Test_updateQuality(t *testing.T) {

	var items = []*Item{
		&Item{"n1", 10, 20},
		&Item{"Aged Brie", 10, 20},
	}

	if err := testDay1(items); err != nil {
		t.Fatal(err)
	}

	if err := testDay20(items); err != nil {
		t.Fatal(err)
	}

	if err := testDay30(items); err != nil {
		t.Fatal(err)
	}

}

func testDay1(items []*Item) error {
	for i := 0; i < DAY1; i++ {
		updateQuality(items)
	}

	var err = errors.New("no pass day1")
	for _, it := range items {
		switch it.Name {
		case "n1":
			if it.SellIn != 9 || it.Quality != 19 {
				return err
			}
		case "Aged Brie":
			if it.SellIn != 9 || it.Quality != 21 {
				return err
			}
		case "Backstage passes to a TAFKAL80ETC concert":

		case "Sulfuras, Hand of Ragnaros":
		default:
			return nil
		}
	}
	return nil
}

func testDay20(items []*Item) error {
	for i := 0; i < DAY10; i++ {
		updateQuality(items)
	}
	return nil
}

func testDay30(items []*Item) error {
	for i := 0; i < DAY20; i++ {
		updateQuality(items)
	}
	return nil
}
