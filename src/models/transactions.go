package models

type DebitTransaction struct {
	Id          int
	Date        string
	Transaction string
}

type MonthTransaction struct {
	Month    string
	Quantity int
}
