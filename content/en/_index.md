---
title: "go-dataframe documentation"
linkTitle: "Home"
type: docs
cascade:
  type: docs
description: "Labelled columnar tables in pure Go (CGO=0): Series, DataFrame and Index, CSV, Arrow and Parquet, group-by, joins, reshaping and time series — on go-ndarray."
---

**go-dataframe is labelled, columnar tables in pure Go.** `Series`,
`DataFrame` and `Index`, with selection, missing data, group-by, joins,
reshaping, time series, a fast CSV reader, Arrow IPC and Parquet, notebook
display and plots. Numeric columns are
[go-ndarray](https://github.com/go-ndarray/ndarray) arrays, handed out
without copying.

pandas is the behavioural reference: the operations give pandas' answers
(values, dtypes, index, row order), checked against pandas 3.0 on 173 recorded
cases; the differences that remain are deliberate and listed in
[Design]({{< relref "design#deliberate-differences-from-pandas" >}}).

{{< plotfig "bars" "Mean yearly rainfall of three cities, as bars, drawn by dfplot from a group-by" >}}

These pages describe [go-dataframe/dataframe](https://github.com/go-dataframe/dataframe)
**v0.1.0**. Every table printed and every figure in them was produced by that
version, from the code shown beside it.

## Install

```sh
go get github.com/go-dataframe/dataframe@v0.1.0
```

The module needs Go 1.27.2 or later. It is pure Go (`CGO_ENABLED=0`).

```go
import "github.com/go-dataframe/dataframe"

df := dataframe.Must(dataframe.ReadCSVFile("weather.csv", dataframe.CSVParseDates("date")))
fmt.Println(df.Head(5))
```

## The pages

| Page | What it says |
| --- | --- |
| [The API, call by call]({{< relref "api" >}}) | every operation beside the pandas call it answers to |
| [Notebook display]({{< relref "display" >}}) | what a frame prints and the HTML table a notebook shows, produced by v0.1.0 |
| [Plots]({{< relref "plots" >}}) | `dfplot`: a frame's columns drawn with go-plot |
| [Design]({{< relref "design" >}}) | columns, the index, the algorithms, CSV, Arrow and Parquet, and the deliberate differences from pandas |
| [Performance]({{< relref "performance" >}}) | time and memory against pandas 3.0 and polars, and how they were measured |

## Speed, honestly

Medians of interleaved runs on 10⁷ rows (Apple M4 Max, a shared host under
load; the [performance page]({{< relref "performance" >}}) has the method, a
second run under heavier load, 10⁶ rows, and Arrow/Parquet). Milliseconds:

| operation | go-dataframe | pandas 3.0 | polars |
|---|---:|---:|---:|
| `read_csv` (5 mixed columns) | 407 | 4220 (C parser) / 525 (pyarrow) | 360 |
| `read_csv`, dates parsed | 362 | 5900 | 417 |
| `sort_values` (stable, float key) | 446 | 973 | 298 |
| `merge` (inner, int key) | 161 | 234 | 55 |
| `groupby("int key").sum()` | 88 | 103 | 69 |
| `df[df.x > 0.5]` | 74 | 75 | 20 |
| `groupby("str key").sum()` | 222 | 204 | 60 |

go-dataframe is faster than pandas on all of these but the string-key
group-by, which is at parity (CSV by about 10x, sorting by 2x); **polars
remains faster on most of them**. Memory is at par with pandas for numeric
columns but worse for strings (a Go `[]string` costs 16 bytes a value; pandas
3 and polars use Arrow-style string buffers): the benchmark table takes 400 MB
here, 349 MB in pandas, 269 MB in polars.

## Packages

| Package | What |
| --- | --- |
| `dataframe` | the core: Series, DataFrame, Index, CSV, display |
| `dataframe/arrowio` | Arrow IPC (file/Feather and stream) and Arrow record conversion, over apache/arrow-go |
| `dataframe/parquetio` | Parquet, over apache/arrow-go |
| `dataframe/dfplot` | plots with [go-plot](https://github.com/go-plot/plot) |

## Verification

- `testdata/gen_fixtures.py` records pandas' answers (values, dtypes, index,
  NA handling, sort stability, group-by order, merge row order, `repr` text)
  for 173 cases; `TestPandasFixtures` replays them.
- A corpus of 47 awkward CSV files (69 option sets) is compared with pandas'
  parse, and Arrow and Parquet files written by pandas/pyarrow are read back
  and compared.
- 100% statement coverage of the whole module under the race detector, on
  Linux, macOS and Windows; tests also run natively on amd64/arm64 and under
  qemu on riscv64, loong64, ppc64le and s390x (big-endian).

## Built on

[go-ndarray](https://github.com/go-ndarray/ndarray) (numeric columns),
[go-plot](https://github.com/go-plot/plot) (`dfplot`) and
[apache/arrow-go](https://github.com/apache/arrow-go) (Arrow and Parquet).
A [go-notebook](https://github.com/go-notebook/notebook) cell shows a frame
through `Display`.
