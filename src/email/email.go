package email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"text/template"
	"transaction-stori/src/models"
)

func GetTextTransactionByMonth(transactionByMonth map[string]models.MonthTransaction) string {
	transactionByMonthText := ""
	for _, transaction := range transactionByMonth {
		quantity := strconv.Itoa(transaction.Quantity)
		transactionByMonthText = transactionByMonthText + "<p><b>Number of transactions in " + transaction.Month + ": " + quantity + "</b></p>"
	}
	return transactionByMonthText
}

// SendEmail to user with balances
func SendEmail(totalBalance, averageDebitAmount, averageCreditAmount float64, transactionByMonth map[string]models.MonthTransaction) error {
	var body bytes.Buffer
	from := os.Getenv("EMAIL_SENDER_USER")
	password := os.Getenv("EMAIL_SENDER_PASSWORD")
	listEmails := os.Getenv("LIST_EMAIL_TO_SEND")
	port := os.Getenv("PORT_EMAIL")
	host := os.Getenv("HOST_EMAIL")
	fmt.Println("emais", listEmails)
	toList := []string{listEmails}
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	transactionByMonthText := GetTextTransactionByMonth(transactionByMonth)

	body.Write([]byte(fmt.Sprintf("Subject: Stori | Account Summary \n%s\n\n", mimeHeaders)))
	templateTransactions, _ := template.ParseFiles("./balance_email.html")
	templateTransactions.Execute(&body, struct {
		TotalBalance           string
		AverageDebitAmount     string
		AverageCreditAmount    string
		TransactionByMonthText string
	}{
		TotalBalance:           fmt.Sprintf("%v", totalBalance),
		AverageDebitAmount:     fmt.Sprintf("%v", averageDebitAmount),
		AverageCreditAmount:    fmt.Sprintf("%v", averageCreditAmount),
		TransactionByMonthText: transactionByMonthText,
	})

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, body.Bytes())

	if err != nil {
		fmt.Println(err)
	}

	return err
}
