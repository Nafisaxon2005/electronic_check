package Commission

import "math"

func Calculate(amount int, isAlif bool) (int, int) {
	if isAlif {
		return 0, amount
	}

	commissionFloat := float64(amount) * 29 / 10000
	commission := int(math.Round(commissionFloat))

	return commission, amount + commission
}
