package csv_process

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type EtheriumPrice struct {
	Pricw float64
	Date  time.Time
}

func LoadDataFrom(path string) ([]EtheriumPrice, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error oppening file: %w", err)
	}

	defer file.Close()

	r := csv.NewReader(file)

	var pricePairs []EtheriumPrice

	for {
		record, err := r.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("error reading CSV file: %w", err)
		}

		parsedDate, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			return nil, fmt.Errorf("cannot parse date: %w", err)
		}

		parsedPrice, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse price: %w", err)
		}

		pricePairs = append(pricePairs, EtheriumPrice{
			Pricw: parsedPrice,
			Date:  parsedDate,
		})
	}

	return pricePairs, nil
}
