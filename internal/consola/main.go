package consola

import (
	"fmt"
	"strings"
)

type TableOptions struct {
	MinWidth int
	Padding  int
}

type Table struct {
	data    [][]string
	options *TableOptions
}

func NewTable(options *TableOptions) *Table {
	return &Table{
		options: options,
	}
}

func NewTableDefault() *Table {
	return &Table{
		options: &TableOptions{
			MinWidth: 10,
			Padding:  1,
		},
	}
}

func (t *Table) InsertRow(row []string) {
	t.data = append(t.data, row)
}

func (t *Table) String() string {
	columnsSize := make([]int, len(t.data[0]))

	for _, row := range t.data {
		for i, col := range row {
			colSize := len(col)

			if size := columnsSize[i]; colSize > size {
				columnsSize[i] = colSize
			}
		}
	}

	var separator string
	for _, size := range columnsSize {
		separator += "+" + strings.Repeat("-", size+t.options.Padding*2)
	}
	separator += "+\n"

	var table string
	for _, row := range t.data {
		table += separator

		for i, col := range row {
			size := columnsSize[i]

			cell := fmt.Sprintf("%-*s", size, col)
			padding := strings.Repeat(" ", t.options.Padding)

			table += fmt.Sprintf("|%s%s%s", padding, cell, padding)
		}

		table += "|\n"
	}
	table += separator

	return table
}
