package main

import "github.com/go-dataframe/dataframe"

func init() { registerOutput("groupby", groupby) }

// Group by city: the mean and the maximum yearly rainfall, and the rainy days
// summed over the three years.
func groupby() (string, string) {
	g := dataframe.Must(rainfall().GroupBy("city"))
	out := dataframe.Must(g.Agg(
		dataframe.Agg("rain", "mean", "max"),
		dataframe.Agg("days", "sum"),
	))
	return out.String(), out.HTML()
}
