package main

import "github.com/go-dataframe/dataframe"

// rainfall is the table the examples share: yearly rainfall in millimetres and
// rainy days at three stations. The numbers are made up for the example.
func rainfall() *dataframe.DataFrame {
	return dataframe.Must(dataframe.New(
		dataframe.Col("city", []string{"Lyon", "Paris", "Nice", "Lyon", "Paris", "Nice", "Lyon", "Paris", "Nice"}),
		dataframe.Col("year", []int{2022, 2022, 2022, 2023, 2023, 2023, 2024, 2024, 2024}),
		dataframe.Col("rain", []float64{741.2, 512.9, 698.4, 830.5, 640.2, 702.3, 810.0, 598.7, 760.1}),
		dataframe.Col("days", []int{98, 103, 61, 112, 118, 64, 107, 111, 70}),
	))
}
