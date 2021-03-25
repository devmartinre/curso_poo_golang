package main

// Item contains information of a invoiceitem
type Item struct {
	id      uint
	product string
	value   float64
}

// New returns a new Item
func NewItem(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// Value getter of Item.Value
func (i Item) Value() float64 {
	return i.value
}
