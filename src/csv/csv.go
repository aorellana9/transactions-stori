package csv

import (
	"encoding/csv"
	"os"
	"strconv"
	"transaction-stori/src/models"
)

func readCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func mapDebitTransactions(data [][]string) ([]models.DebitTransaction, error) {
	var debitTransactions []models.DebitTransaction
	for i, line := range data {
		if i > 0 { // omit header line
			var rec models.DebitTransaction
			for j, field := range line {
				if j == 0 {
					id, err := strconv.Atoi(field)
					rec.Id = id
					if err != nil {
						return nil, err
					}
				} else if j == 1 {
					rec.Date = field
				} else if j == 2 {
					rec.Transaction = field
				}
			}
			debitTransactions = append(debitTransactions, rec)
		}
	}
	return debitTransactions, nil
}

// GetDebitTransactionsFromCSV return the response from csv
func GetDebitTransactionsFromCSV(csvName string) ([]models.DebitTransaction, error) {
	records, err := readCsvFile(csvName)

	if err != nil {
		return nil, err
	}

	return mapDebitTransactions(records)
}
