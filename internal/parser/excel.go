package parser

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Recipient struct {
	Name  string
	Email string
}

func ParseExcel(filePath, sheetName string) ([]Recipient, error) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening excel file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	rows, err := file.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("error reading sheet: %w", err)
	}

	var recipients []Recipient

	for _, row := range rows[1:] {
		if len(row) < 2 {
			continue
		}

		recipients = append(recipients, Recipient{
			Name:  row[0],
			Email: row[1],
		})
	}

	return recipients, nil
}
