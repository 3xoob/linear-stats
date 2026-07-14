# linear-stats

A small command-line Go program that computes a least-squares linear regression line and the Pearson correlation coefficient from a list of numeric values.

## Overview

`linear-stats` reads a text file containing one numeric value per line. Each value is treated as a `y` value, with its 0-based line index used as the corresponding `x` value. From these `(x, y)` pairs, the program calculates the best-fit line (`y = mx + b`) and the Pearson correlation coefficient, then prints both results to standard output.

## Features

- Reads a list of `y` values from a plain text file (one value per line)
- Computes the least-squares linear regression line, printing slope (`m`) and intercept (`b`)
- Computes the Pearson correlation coefficient (`r`) for the data
- Basic error handling for a missing file argument, an unreadable file, or a non-numeric line

## Technologies

- Go (module targets Go 1.23, per `go.mod`)
- Go standard library only: `bufio`, `fmt`, `math`, `os`, `strconv`

## Project Structure

```
.
├── main.go          # Program entry point and all logic
├── go.mod           # Go module definition
├── LICENSE          # License terms
└── COPYRIGHT.md     # Copyright notice
```

## Requirements

- Go 1.23 or later

## Installation

```bash
git clone https://github.com/3xoob/linear-stats.git
cd linear-stats
go build -o linear-stats main.go
```

## Usage

Run directly with `go run`:

```bash
go run main.go <path-to-file>
```

Or build first and run the resulting binary:

```bash
go build -o linear-stats main.go
./linear-stats <path-to-file>
```

The input file must contain one numeric value per line. Each line's value is used as `y`, and its line index (starting at 0) is used as `x`.

## Example

Given a file `data.txt` containing:

```
1
2
3
4
5
```

Running:

```bash
go run main.go data.txt
```

Produces:

```
Linear Regression Line: y = 1.000000x + 1.000000
Pearson Correlation Coefficient: 1.0000000000
```

## License

See [LICENSE](LICENSE) and [COPYRIGHT.md](COPYRIGHT.md). All rights reserved; the source is published for portfolio/viewing purposes only, with no permission granted to copy, modify, distribute, or reuse it without prior written permission from the copyright holder.
