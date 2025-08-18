package main

import (
	"fmt"
	"mult_table/eth_usd/csv_process"
	"mult_table/eth_usd/plotter"
	"os"

	"gonum.org/v1/plot/vg"
)

func main() {
	pricePairs, err := csv_process.LoadDataFrom("prices.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading data from CSV file: %v\n", err)
		os.Exit(1)
	}

	pricePlot, err := plotter.GeneratePlotFor(pricePairs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating plot: %v\n", err)
		os.Exit(1)
	}

	if err := pricePlot.Save(vg.Inch*15, vg.Inch*4, "etherium_prices.png"); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving plot: %v\n", err)
		os.Exit(1)
	}
}
