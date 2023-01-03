package balances

import (
	"errors"
	"strconv"
	"strings"
	"transaction-stori/src/constants"
	"transaction-stori/src/models"
)

func getTotalBalance(debitTransactions []models.DebitTransaction) (float64, error) {
	sum := 0.0
	for _, transaction := range debitTransactions {
		amountTransaction, err := strconv.ParseFloat(transaction.Transaction[1:], 64)
		if err != nil {
			return 0.0, err
		}
		if transaction.Transaction[0:1] == "+" {
			sum = sum + amountTransaction
		} else if transaction.Transaction[0:1] == "-" {
			sum = sum - amountTransaction
		}
	}
	return sum, nil
}

func getAverageAmount(debitTransactions []models.DebitTransaction, operation string) (float64, error) {
	sum := 0.0
	count := 0
	for _, transaction := range debitTransactions {
		if transaction.Transaction[0:1] == operation {
			amountTransaction, err := strconv.ParseFloat(transaction.Transaction[1:], 64)
			if err != nil {
				return 0.0, err
			}
			sum = sum + amountTransaction
			count++
		}
	}
	return sum / float64(count), nil
}

func getTransactionByMonth(debitTransactions []models.DebitTransaction) map[string]models.MonthTransaction {
	emptyMonthTransaction := map[string]models.MonthTransaction{}
	for _, transaction := range debitTransactions {
		month := strings.Split(transaction.Date, "/")[0]
		emptyMonthTransaction[month] = models.MonthTransaction{
			Month:    constants.EMPTY_MONTH_TRANSACTION_DICT[month].Month,
			Quantity: emptyMonthTransaction[month].Quantity + 1}
	}
	return emptyMonthTransaction
}

// GetUserBalances return the balances from received data
func GetUserBalances(debitTransactions []models.DebitTransaction) (float64, float64, float64, map[string]models.MonthTransaction, error) {
	totalBalance, errBalance := getTotalBalance(debitTransactions)
	averageDebitAmount, errDebit := getAverageAmount(debitTransactions, "-")
	averageCreditAmount, errCredit := getAverageAmount(debitTransactions, "+")
	transactionByMonth := getTransactionByMonth(debitTransactions)
	if errBalance != nil || errDebit != nil || errCredit != nil {
		err := errors.New("error in get user balances")
		return 0.0, 0.0, 0.0, nil, err
	}
	return totalBalance, averageDebitAmount, averageCreditAmount, transactionByMonth, nil
}
