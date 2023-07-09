package outer_libs

import "fmt"

type OldPaymentSystem struct {
}

func (s OldPaymentSystem) MakePayment(amount float64) {
	fmt.Print(amount)
}
