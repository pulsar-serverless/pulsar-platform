package billing

import (
	"strconv"
)

func GetMonthYear(usageMonth string) (int, int) {
	// double digit months
	if len(usageMonth) == 6 {
		month, _ := strconv.Atoi(usageMonth[0:2])
		year, _ := strconv.Atoi(usageMonth[2:])
		return month, year
	}

	month, _ := strconv.Atoi(usageMonth[0:1])
	year, _ := strconv.Atoi(usageMonth[1:])
	return month, year
}
