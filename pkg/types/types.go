package types

type Money int64

type Category string

type Status string

const (
	StatusOK Status = "OK"
	StatusFail Status = "FAIL"
	StatusProgress Status = "INPROGRESS"
)
type Payment struct {
	ID int
	Amount Money
	Category Category
	Status Status
}

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

type PAN string

type Card struct {
	ID        int
	PAN       string
	Balance   Money
	Currency  Currency
	Color     string
	Name      string
	Active    bool
}