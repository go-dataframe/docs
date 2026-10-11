package main

import (
	"math"

	"github.com/go-dataframe/dataframe"
	"github.com/go-dataframe/dataframe/dfplot"
	"github.com/go-plot/plot"
)

func init() { register("lines", lines) }

// Sixty days of a made-up daily temperature, and its 7-day rolling mean:
// s.rolling(7).mean(), then df.plot(). The first six means are missing,
// as in pandas, and the line starts at day 7.
func lines() *plot.Figure {
	day := make([]int, 60)
	temp := make([]float64, 60)
	for i := range day {
		day[i] = i + 1
		temp[i] = 12 + 6*math.Sin(2*math.Pi*float64(i)/60) + 2.5*math.Sin(1.7*float64(i))
	}
	s := dataframe.Must(dataframe.NewSeries("temperature", temp))
	r := dataframe.Must(s.Rolling(7))
	df := dataframe.Must(dataframe.New(
		dataframe.Col("day", day),
		dataframe.Col("temperature", s),
		dataframe.Col("7-day mean", r.Mean()),
	))
	df = dataframe.Must(df.SetIndex("day"))
	fig, err := dfplot.LineOpts(df, nil, []dfplot.Option{dfplot.WithTitle("Daily temperature (°C)")})
	if err != nil {
		panic(err)
	}
	return fig
}
