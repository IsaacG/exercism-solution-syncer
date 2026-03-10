package savings

type AccNo string

const (
	FixedInterestRate = 0.05
	DaysPerYear = 365
	AccountNo = AccNo("XF348IJ")
	_ = iota
	Jan
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

// GetFixedInterestRate returns the FixedInterestRate constant.
func GetFixedInterestRate() float32 {
	return FixedInterestRate
}

// GetDaysPerYear returns the DaysPerYear constant.
func GetDaysPerYear() int {
	return DaysPerYear
}

// GetMonth returns the value for the given month.
func GetMonth(month int) int {
	return month - 3
}

// GetAccountNumber returns the AccountNo constant.
func GetAccountNumber() AccNo {
	return AccountNo
}
