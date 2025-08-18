package plotter

import (
	"fmt"
	"mult_table/eth_usd/csv_process"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

func GeneratePlotFor(pricePairs []csv_process.EtheriumPrice) (*plot.Plot, error) {
	points := make(plotter.XYs, len(pricePairs))

	for i := range pricePairs {
		points[i].X = float64(pricePairs[i].Date.Unix())
		points[i].Y = pricePairs[i].Pricw
	}

	pricePlot := plot.New()

	pricePlot.Title.Text = "Ethereum Price Chart"
	pricePlot.X.Label.Text = "Date"
	pricePlot.Y.Label.Text = "Price (USD)"

	if err := plotutil.AddLinePoints(pricePlot, "ETC", points); err != nil {
		return nil, fmt.Errorf("error adding line to plot: %w", err)
	}

	pricePlot.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}

	return pricePlot, nil
}
