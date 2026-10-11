---
title: "The API, call by call"
weight: 10
description: "Every operation of go-dataframe v0.1.0 beside the pandas call it answers to."
---

<!-- From go-dataframe/dataframe v0.1.0, README.md ("From pandas"): the text is the library's. Change it there first. -->

The behavioural reference is pandas: each row below gives pandas' answer
(values, dtypes, index, row order), checked against pandas 3.0 on 173 recorded
cases. What differs on purpose is listed in
[Design]({{< relref "design#deliberate-differences-from-pandas" >}}).

| pandas | dataframe |
|---|---|
| `pd.DataFrame({"a": [1, 2], "b": ["x", "y"]})` | `dataframe.New(dataframe.Col("a", []int{1, 2}), dataframe.Col("b", []string{"x", "y"}))` |
| `pd.Series([1.5, None])` | `dataframe.NewSeries("", []any{1.5, nil})` |
| `pd.read_csv(path, usecols=..., nrows=..., parse_dates=[...])` | `dataframe.ReadCSVFile(path, dataframe.CSVUseCols(...), dataframe.CSVNRows(n), dataframe.CSVParseDates(...))` |
| `df.to_csv(path)` | `df.WriteCSVFile(path)` |
| `pd.read_parquet(p)`, `df.to_parquet(p)` | `parquetio.ReadParquetFile(p)`, `parquetio.WriteParquetFile(p, df)` |
| `pd.read_feather(p)`, `df.to_feather(p)` | `arrowio.ReadFeather(p)`, `arrowio.WriteFeather(p, df)` |
| `df.head()`, `df.tail(3)` | `df.Head(5)`, `df.Tail(3)` |
| `df.shape`, `df.columns`, `df.dtypes`, `df.index` | `df.Shape()`, `df.Columns()`, `df.DTypes()`, `df.Index()` |
| `df["a"]`, `df[["a", "b"]]` | `df.Column("a")`, `df.Select("a", "b")` |
| `df.values`, `s.values` | `df.Values()`, `s.Values()` (`*ndarray.Array`) |
| `df.iloc[0:5]`, `df.iloc[[0, 2], 1:3]` | `df.Iloc(dataframe.Slice(0, 5))`, `df.Iloc(dataframe.Pos(0, 2), dataframe.Slice(1, 3))` |
| `df.loc["a":"c", ["x"]]` | `df.Loc(dataframe.LabelRange("a", "c"), "x")` |
| `df[df.x > 0]` | `df.Filter(dataframe.Must(x.Gt(0)))` |
| `df.query("x > 0 and s in ['a', 'b']")` | `df.Query("x > 0 and s in ['a', 'b']")` |
| `df.set_index("k")`, `df.reset_index()` | `df.SetIndex("k")`, `df.ResetIndex(false)` |
| `df.sort_values(["a", "b"], ascending=[True, False])` | `df.SortValues([]string{"a", "b"}, true, false)` |
| `df.sort_index()`, `df.reindex(labels)` | `df.SortIndex(true)`, `df.Reindex(labels)` |
| `df.dropna()`, `df.fillna(0)`, `df.isna()` | `df.DropNA()`, `df.FillNA(0)`, `df.IsNA()` |
| `df.sum()`, `df.mean()`, `df.describe()` | `df.Sum()`, `df.Mean()`, `df.Describe()` |
| `df.round(2)`, `df.T`, `df.astype({"a": "float64"})` | `df.Round(2)`, `df.T()`, `df.AsType(map[string]dataframe.DType{"a": dataframe.Float64})` |
| `df.rename(columns={"a": "b"})`, `df.copy()` | `df.Rename(map[string]string{"a": "b"})`, `df.Copy()` |
| `df.a + df.b` (aligned on the index) | `a.Add(b)` |
| `s.apply(f)`, `df.apply(f, axis=1)` | `s.Map(f)`, `df.Apply("name", f)` |
| `df.groupby(["k1", "k2"]).sum()` | `g, _ := df.GroupBy("k1", "k2"); g.Sum()` |
| `df.groupby("k").agg(x=("x", "mean"))` | `g.Agg(dataframe.Agg("x", "mean"))` |
| `pd.concat([a, b])`, `pd.concat([a, b], axis=1)` | `dataframe.Concat(a, b)`, `dataframe.ConcatWith(dataframe.ConcatOptions{Axis: 1}, a, b)` |
| `pd.merge(l, r, on="k", how="left")` | `dataframe.Merge(l, r, dataframe.MergeOptions{On: []string{"k"}, How: "left"})` |
| `df.pivot(...)`, `df.pivot_table(...)`, `df.melt(...)` | `df.Pivot(...)`, `df.PivotTable(...)`, `df.Melt(...)` |
| `df.stack()`, `s.unstack()` | `df.Stack()`, `s.Unstack()` |
| `pd.to_datetime(s)`, `pd.Timestamp("2024-01-01")` | `dataframe.ToDatetime(s)`, `dataframe.Timestamp("2024-01-01")` |
| `df.resample("D").sum()` | `r, _ := df.Resample("D"); r.Sum()` |
| `df.to_period("M")`, `s.dt.year` | `df.ToPeriod("M")`, `d, _ := s.Dt(); d.Year()` |
| `s.shift(1)`, `s.diff()`, `s.cumsum()`, `s.rolling(7).mean()` | `s.Shift(1)`, `s.Diff(1)`, `s.CumSum()`, `r, _ := s.Rolling(7); r.Mean()` |
| `df.plot()`, `df.plot.bar()`, `df.hist()`, `df.plot.scatter(x, y)` | `dfplot.Line(df)`, `dfplot.Bar(df)`, `dfplot.Hist(df, 10)`, `dfplot.Scatter(df, x, y)` |

Fallible calls return `(value, error)`; `dataframe.Must` unwraps one in a
notebook cell. Every value is immutable: operations return new frames that
share the columns they did not change.

The complete API is on
[pkg.go.dev](https://pkg.go.dev/github.com/go-dataframe/dataframe@v0.1.0).
