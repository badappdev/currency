package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUSD(t *testing.T) {

	amount := NewUSD(1, 99)

	t.Logf("succesfully created new amount $%d.%d", amount.dollars, amount.cents)
}

func TestUSDValue(t *testing.T) {

	amount := NewUSD(1, 99)

	d, c := amount.Value()

	assert.Equal(t, int64(1), amount.dollars)
	assert.Equal(t, int64(99), amount.cents)
	assert.Equal(t, amount.dollars, d)
	assert.Equal(t, amount.cents, c)

	t.Logf("successfully found values")
}

func TestUSDAdd(t *testing.T) {

	amount := NewUSD(1, 99)
	tip := NewUSD(0, 1)

	assert.Equal(t, int64(1), amount.dollars)
	assert.Equal(t, int64(99), amount.cents)
	assert.Equal(t, int64(0), tip.dollars)
	assert.Equal(t, int64(1), tip.cents)

	amount.Add(tip)

	assert.Equal(t, int64(2), amount.dollars)
	assert.Equal(t, int64(0), amount.cents)

	t.Logf("successfully added two amounts for a total of $%d.%d", amount.dollars, amount.cents)
}
