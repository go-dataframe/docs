// Command figures writes every figure and every printed output the
// documentation shows, with go-dataframe at the version go.mod names (and
// go-plot drawing the figures):
//
//   - a figure is a func() *plot.Figure, written to
//     ../static/images/figures/<name>.svg;
//   - an output is a func() (text, html string): what a frame prints, and the
//     HTML table a notebook shows, written to ../static/outputs/<name>.txt and
//     <name>.html.
//
// Each is one file here, and the pages include that file verbatim, so the
// code shown is the code that produced what is beside it.
//
//	cd figures && go run .
//
// The CI runs it and fails if any file differs from the committed one.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/go-plot/plot"
)

var (
	figures = map[string]func() *plot.Figure{}
	outputs = map[string]func() (text, html string){}
)

func register(name string, f func() *plot.Figure)              { figures[name] = f }
func registerOutput(name string, f func() (text, html string)) { outputs[name] = f }

func main() {
	root := filepath.Join("..", "static")
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	figDir, outDir := filepath.Join(root, "images", "figures"), filepath.Join(root, "outputs")
	for _, d := range []string{figDir, outDir} {
		if err := os.MkdirAll(d, 0o755); err != nil {
			panic(err)
		}
	}
	for _, n := range sorted(figures) {
		f, err := os.Create(filepath.Join(figDir, n+".svg"))
		if err != nil {
			panic(err)
		}
		if err := figures[n]().SaveSVG(f); err != nil {
			panic(err)
		}
		if err := f.Close(); err != nil {
			panic(err)
		}
		fmt.Println(n + ".svg")
	}
	for _, n := range sorted(outputs) {
		text, html := outputs[n]()
		write(filepath.Join(outDir, n+".txt"), text+"\n")
		fmt.Println(n + ".txt")
		if html != "" {
			write(filepath.Join(outDir, n+".html"), html+"\n")
			fmt.Println(n + ".html")
		}
	}
}

func write(path, s string) {
	if err := os.WriteFile(path, []byte(s), 0o644); err != nil {
		panic(err)
	}
}

func sorted[V any](m map[string]V) []string {
	names := make([]string, 0, len(m))
	for n := range m {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}
