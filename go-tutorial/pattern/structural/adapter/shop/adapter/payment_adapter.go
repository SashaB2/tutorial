package adapter

import (
	"tutorial/go-tutorial/pattern/structural/adapter/shop/outer-libs"
)

type PaymentAdapter struct {
	NewPayment *outer_libs.NewPaymentLibraty
}

func (r *PaymentAdapter) MakePayment(amount float64) {
	r.NewPayment.ProcessPayment(amount)
}
