package payment

type Payment interface {
    tax() float32
}

func Amount(p Payment) float32 {
    return p.tax()
}

type PayPal struct {
    Value float32
}

type PicPay struct {
    Value float32
}

func (p PayPal) tax() float32 {
    return (p.Value * 0.10) + p.Value
}

func (p PicPay) tax() float32 {
    return (p.Value * 0.15) + p.Value
}


