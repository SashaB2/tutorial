package outer_libs

import "fmt"

type NewPaymentLibraty struct {
}

func (l NewPaymentLibraty) ProcessPayment(amount float64) {
	fmt.Print(amount)
}
