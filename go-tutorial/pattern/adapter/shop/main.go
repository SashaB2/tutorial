package main

import (
	"tutorial/go-tutorial/pattern/adapter/shop/adapter"
	outer_libs "tutorial/go-tutorial/pattern/adapter/shop/outer-libs"
)

func main() {
	ol := outer_libs.OldPaymentSystem{}
	ol.MakePayment(3.3)

	npl := &outer_libs.NewPaymentLibraty{}

	nl := &adapter.PaymentAdapter{
		npl,
	}
	nl.MakePayment(3.4)
}
