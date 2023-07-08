package adapter

import outer_libs "tutorial/go-tutorial/pattern/adapter/shop/outer-libs"

type PaymentAdapter struct {
	NewPayment *outer_libs.NewPaymentLibraty
}

func (r *PaymentAdapter) MakePayment(amount float64) {
	r.NewPayment.ProcessPayment(amount)
}
