package controller

import (
	"net/http"
	"os"
	"transaction-stori/src/balances"
	csvLogic "transaction-stori/src/csv"
	"transaction-stori/src/email"

	"github.com/gin-gonic/gin"
)

var contentType = "application/json; charset=utf-8"
var succesfullMessage = "Successfully sent mail to all user in toList"

// Health for check service status
func Health(c *gin.Context) {
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(http.StatusText(http.StatusOK)))
}

// SendTranstactionsToEmail send transactions to users email controller
func SendTranstactionsToEmail(c *gin.Context) {
	csvName := os.Getenv("CSV_NAME")

	debitTransactions, err := csvLogic.GetDebitTransactionsFromCSV("./" + csvName + ".csv")

	if err != nil {
		c.Data(http.StatusInternalServerError, contentType, []byte(err.Error()))
		return
	}

	totalBalance, averageDebitAmount, averageCreditAmount, transactionByMonth, err := balances.GetUserBalances(debitTransactions)

	if err != nil {
		c.Data(http.StatusInternalServerError, contentType, []byte(err.Error()))
		return
	}

	err = email.SendEmail(totalBalance, averageDebitAmount, averageCreditAmount, transactionByMonth)

	if err != nil {
		c.Data(http.StatusInternalServerError, contentType, []byte(err.Error()))
		return
	}

	c.Data(http.StatusOK, contentType, []byte(succesfullMessage))
}
