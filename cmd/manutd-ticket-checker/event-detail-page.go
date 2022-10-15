package mutc

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type UnitedEventDetailPage struct {
	*UnitedPage

	config *Config

	min_price, max_price uint16
	seats_available      bool
}

func (p *UnitedEventDetailPage) FindMinAndMaxPrices() (uint16, uint16) {
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

func (p *UnitedEventDetailPage) HasAvailableSeats() bool {
	p.MustEval(fmt.Sprintf(`() => document.querySelector("input.areas-filter-panel__min-sum-input").value = "%d.00"`, p.config.MinPrice))
	p.MustEval(fmt.Sprintf(`() => document.querySelector("input.areas-filter-panel__max-sum-input").value = "%d.00"`, p.config.MaxPrice))

	spinner_up := p.MustElement("a.ui-spinner-up")

	for i := 0; i < p.config.NumberOfSeats; i++ {
		spinner_up.MustClick()
	}

	find_button := p.MustElement("button.areas-filter-panel__find-button")

	is_disabled, _ := find_button.Attribute("disabled")

	if is_disabled != nil {
		return false
	}

	find_button.MustClick()

	count := 0

	for {
		url := p.MustInfo().URL

		if strings.Contains(url, "login.manutd.com") {
			p.seats_available = true
			return true
		}

		if count > 10 {
			return false
		}

		time.Sleep(200 * time.Millisecond)

		count++
	}
}
