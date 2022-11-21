package currency

import "sync"

// USD united states dollar implementation of currency
type USD struct {
	cents   int64
	dollars int64
	mtx     sync.RWMutex
}

// interface verification
var _ Currency = (*USD)(nil)

// NewUSD create new usd dollar amount
func NewUSD(dollars, cents int64) *USD {
	return &USD{
		dollars: dollars,
		cents:   cents,
	}
}

// Add add existing amount with another type of currency
func (usd *USD) Add(currency Currency) {

	d, c := currency.Value()

	usd.mtx.Lock()
	defer usd.mtx.Unlock()

	usd.dollars += d
	usd.cents += c

	if usd.cents >= 100 {
		e := usd.cents / 100

		usd.cents -= e * 100
		usd.dollars += e
	}
}

// Value retrieve values of us dollar
func (usd *USD) Value() (dollars, cents int64) {
	dollars = usd.dollars
	cents = usd.cents
	return
}
