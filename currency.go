package currency

// Currency required currency functions
type Currency interface {
	Value() (dollars, cents int64)
	Add(c Currency)
}
