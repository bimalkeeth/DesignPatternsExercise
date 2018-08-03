package Factory

import "errors"

const (
	Cash      = 1
	DebitCard = 2
)

type PaymentMethod interface {
	Pay(amount float64) string
}

func GetPaymentMethod(m int) (PaymentMethod, error) {

	return nil, errors.New("not implemented yet")
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float64) string {
	return ""
}

func (c *DebitCardPM) Pay(amount float64) string {
	return ""
}
