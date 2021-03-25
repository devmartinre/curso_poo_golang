package main

// Invoice is the struct of a invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  Customer
	items   Items
}

//New returns a new Invoice
func NewInvoice(country, city string, client Customer, items Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}
