package constants

import "transaction-stori/src/models"

var EMPTY_MONTH_TRANSACTION_DICT = map[string]models.MonthTransaction{
	"1":  {Month: "January", Quantity: 0},
	"2":  {Month: "February", Quantity: 0},
	"3":  {Month: "March", Quantity: 0},
	"4":  {Month: "April", Quantity: 0},
	"5":  {Month: "May", Quantity: 0},
	"6":  {Month: "June", Quantity: 0},
	"7":  {Month: "July", Quantity: 0},
	"8":  {Month: "August", Quantity: 0},
	"9":  {Month: "September", Quantity: 0},
	"10": {Month: "October", Quantity: 0},
	"11": {Month: "November", Quantity: 0},
	"12": {Month: "December", Quantity: 0},
}
