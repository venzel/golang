package main

import (
	"fmt"
	model "model/payment"
)

func main() {
    paypal := model.PayPal{
        Value: 100,
    }

    picpay := model.PicPay{
        Value: 100,
    }

    amountPayPal := model.Amount(paypal);
    amountPicPay := model.Amount(picpay);

    fmt.Println(amountPayPal, amountPicPay)
    
}