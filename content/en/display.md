---
title: "Notebook display"
weight: 20
description: "What a frame prints, and the HTML table a notebook shows, as go-dataframe v0.1.0 produced them."
---

`DataFrame`, `Series` and `Index` have `Display()`, which returns the value in
every format a notebook front end may show, keyed by MIME type: `text/plain`
(what `String` prints, laid out as pandas prints it) and `text/html` (the
table pandas shows in a notebook). It is the display contract of the
[go-notebook](https://github.com/go-notebook/notebook) kernel, so a cell whose
last expression is a frame shows it as a table.

Every output on this page was produced by go-dataframe v0.1.0 from the code
shown above it: `figures/<name>.go` in this site's repository writes it, and
the CI writes it again and fails if one byte differs.

## A frame

The table the examples share, built from Go slices:

{{< dfout "display" >}}

## A group-by

{{< dfout "groupby" >}}

## Describe

{{< dfout "describe" >}}

How the text layout and the HTML are produced, and where they follow pandas,
is in [Design]({{< relref "design#notebook-display" >}}).
