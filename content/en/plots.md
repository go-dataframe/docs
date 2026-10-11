---
title: "Plots"
weight: 30
description: "dfplot: a frame's columns drawn with go-plot, as df.plot() draws them."
---

`dataframe/dfplot` draws a frame with [go-plot](https://github.com/go-plot/plot):
one line per column against the index, as `df.plot()` does, and `Bar`,
`Barh`, `Hist`, `Area`, `Scatter` and the `Series` variants. The legend names
the columns and the x label is the name of the index; a missing value breaks
a line. Each function returns a `*plot.Figure`, which saves as SVG or PNG and
shows in a notebook through `Display`.

Every figure on this page was drawn by go-dataframe v0.1.0 and go-plot v0.1.0
from the code shown under it; the CI draws it again and fails if one byte of
the SVG differs.

## Lines, and a rolling mean

{{< plotfig "lines" "A daily temperature over sixty days and its 7-day rolling mean, drawn by dfplot.LineOpts" >}}

## Bars from a group-by

{{< plotfig "bars" "Mean yearly rainfall of three cities, as bars, drawn by dfplot.BarOpts" >}}

The data, shared with the [notebook display]({{< relref "display" >}}) page:

{{< gofile "figures/data.go" >}}
