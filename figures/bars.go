package main

import (
	"github.com/go-dataframe/dataframe"
	"github.com/go-dataframe/dataframe/dfplot"
	"github.com/go-plot/plot"
)

func init() { register("bars", bars) }

// df.groupby("city")[["rain"]].mean().plot.bar()
func bars() *plot.Figure {
	g := dataframe.Must(rainfall().GroupBy("city"))
	means := dataframe.Must(g.Mean())
	fig, err := dfplot.BarOpts(means, []string{"rain"}, []dfplot.Option{dfplot.WithTitle("Mean yearly rainfall, 2022–2024 (mm)")})
	if err != nil {
		panic(err)
	}
	return fig
}
