package main

import "strconv"

type UnitedEventDetailPage struct {
	*UnitedPage

	min_price, max_price uint16
}

func (p *UnitedEventDetailPage) FindPrices() (uint16, uint16) {
	price_input := p.MustElement(".areas-filter-panel__price-section input[type=number]")

	p.min_price = 0
	p.max_price = 0

	min_attribute := price_input.MustAttribute("aria-valuemin")
	max_attribute := price_input.MustAttribute("aria-valuemax")

	min_int, min_err := strconv.Atoi(*min_attribute)
	max_int, max_err := strconv.Atoi(*max_attribute)

	if min_err == nil {
		p.min_price = uint16(min_int)
	}

	if max_err == nil {
		p.max_price = uint16(max_int)
	}

	return p.min_price, p.max_price
}
