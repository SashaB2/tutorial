package main

import (
	"tutorial/go-tutorial/pattern/structural/adapter/shop/adapter"
	outer_libs2 "tutorial/go-tutorial/pattern/structural/adapter/shop/outer-libs"
)

func main() {
	ol := outer_libs2.OldPaymentSystem{}
	ol.MakePayment(3.3)

	npl := &outer_libs2.NewPaymentLibraty{}

	nl := &adapter.PaymentAdapter{
		npl,
	}
	nl.MakePayment(3.4)
}
