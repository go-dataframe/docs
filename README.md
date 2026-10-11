# docs

[![docs](https://github.com/go-dataframe/docs/actions/workflows/docs.yml/badge.svg)](https://github.com/go-dataframe/docs/actions/workflows/docs.yml)
[![License](https://img.shields.io/badge/license-BSD--3--Clause-A93A82?style=flat-square)](LICENSE)

**The go-dataframe documentation site**, published at
<https://go-dataframe.github.io/docs/>.

A [Hugo](https://gohugo.io/) site with the [tannevaled/hextra](https://github.com/tannevaled/hextra)
fork of the [Hextra](https://github.com/imfing/hextra) theme, built the same way
as the [go-plot documentation](https://github.com/go-plot/docs), with
go-dataframe's branding. It documents
[`go-dataframe/dataframe`](https://github.com/go-dataframe/dataframe).

## Versions

A version of the documentation is **the minor version of
[`go-dataframe/dataframe`](https://github.com/go-dataframe/dataframe) it describes**, or `main`
while no release is tagged:

| Source | Published under |
| --- | --- |
| branch `main`, with `params.dataframe.version` in `hugo.yaml` | `https://go-dataframe.github.io/docs/<MAJOR.MINOR>/`, or `/docs/main/` while it says `main` |
| the newest release (else `main`) | also `https://go-dataframe.github.io/docs/latest/` |

`https://go-dataframe.github.io/docs/` redirects to `latest/`. Each push to `main`
replaces the directory of the version it describes, and nothing else. The
pages describe **v0.1.0**, published under `0.1/`.

To document a new release, set `params.dataframe.version` to its tag, require the
same version in `figures/go.mod`, redraw the figures, and update the pages.

## Figures and outputs

Every figure and every printed table on the site is produced by go-dataframe
v0.1.0 (the figures drawn by [go-plot](https://github.com/go-plot/plot)).
`figures/` is a Go module, one file per example; `go run .` there writes
`static/images/figures/<name>.svg` for a figure, and
`static/outputs/<name>.txt` and `<name>.html` for what a frame prints and the
table a notebook shows. The `plotfig` and `dfout` shortcodes show the result
and then **read the file that produced it**, so the code on the page is the
code that ran. The `figures` job of the CI produces them all again and fails
if a file differs from the committed one, or if `figures/go.mod` requires
another go-dataframe than the version the pages describe.

```sh
cd figures && go run .
```

## Layout

| Path | What |
| --- | --- |
| `content/en/` | the pages; `_index.md` is the home, `weight` orders the sidebar |
| `i18n/en.yaml` | the site's strings, beside the theme's |
| `hugo.yaml` | `baseURL`, the go-dataframe version described, the navbar (version › theme › search › GitHub), the brand mounts |
| `figures/` | the Go module that produces every figure and output, one file each |
| `layouts/_shortcodes/plotfig.html` | a figure and the code that drew it |
| `layouts/_shortcodes/dfout.html` | code, what it printed, and the HTML table it shows |
| `layouts/_shortcodes/gofile.html` | a Go file of this repository, verbatim |
| `static/images/figures/`, `static/outputs/` | the figures and outputs, as `figures/` writes them |
| `layouts/_partials/custom/version-select.html` | the version selector (reads `versions.json`) |
| `layouts/_partials/custom/content-begin.html` | the outdated-version banner (reads `versions.json`) |
| `layouts/_partials/favicons.html` | the brand's favicons |
| `assets/css/custom.css` | the brand magenta as Hextra's primary colour, the table headers, a white card under figures, the notebook tables |
| `themes/hextra` | submodule: [tannevaled/hextra](https://github.com/tannevaled/hextra), pinned to a release tag |
| `branding` | submodule: [go-dataframe/brand](https://github.com/go-dataframe/brand) (the mark, its PNGs and ICO) |
| `scripts/publish-version.sh` | puts one build into a `gh-pages` checkout and rewrites `versions.json`, `latest` and the root redirect |

## Working locally

```sh
git clone --recurse-submodules https://github.com/go-dataframe/docs.git
cd docs
mkdir -p data && sh themes/hextra/scripts/page-history.sh > data/pagehistory.json
hugo server --baseURL http://localhost:1313/docs/0.1/
```

then open <http://localhost:1313/docs/0.1/>. Hugo 0.146 or later; no Python,
no Node. Go 1.27.2 or later only to redraw the figures.

## Publication

`.github/workflows/docs.yml`:

| Job | When | What |
| --- | --- | --- |
| `build` | pull requests, `main` | builds with `--baseURL …/docs/<version>/`, checks internal links and anchors (lychee, offline), uploads the site |
| `links` | pull requests, `main` | fetches every external URL the pages name |
| `figures` | pull requests, `main` | produces every figure and output again and compares it with the committed file |
| `deploy` | `main` | `scripts/publish-version.sh` into `gh-pages` (started empty on the first deploy), then a commit and a push |

GitHub Pages serves the `gh-pages` branch.

## Licence

BSD-3-Clause.
