package stats

import (
	"github.com/Aziz-Gafuroff/bank/pkg/bank/types"
)

func Avg(payments []types.Payment) types.Money {
	avgPayment := types.Money(0)
	
	for _, payment := range payments {
		avgPayment += payment.Amount
		
	}

	return avgPayment /types.Money(len(payments))
}