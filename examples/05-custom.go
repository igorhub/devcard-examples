package examples

import (
	"os"

	"github.com/igorhub/devcard"
	"github.com/wcharczuk/go-chart/v2"
)

type chartCell struct {
	devcard.CustomCell

	data    []datum
	tempDir string
}

type datum struct {
	name  string
	value int
}

// addBar adds a bar to the right of the chart cell.
func (c *chartCell) addBar(name string, value int) {
	c.data = append(c.data, datum{name, value})
}

// graph turns the cell's data into a chart.BarChart instance.
func (c *chartCell) graph() chart.BarChart {
	var values []chart.Value
	for _, d := range c.data {
		values = append(values, chart.Value{Label: d.name, Value: float64(d.value)})
	}

	return chart.BarChart{
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Width: 640,
		YAxis: chart.YAxis{
			Range:          &chart.ContinuousRange{Min: 0, Max: 60},
			ValueFormatter: chart.IntValueFormatter,
		},
		UseBaseValue: true,
		BaseValue:    0,
		Bars:         values,
	}
}

func (c *chartCell) Cast() devcard.Cell {
	f, err := os.CreateTemp(c.tempDir, "*.png")
	if err != nil {
		return devcard.NewErrorCell("chartCell error", err)
	}
	defer f.Close()
	err = c.graph().Render(chart.PNG, f)
	if err != nil {
		return devcard.NewErrorCell("chartCell error", err)
	}

	return devcard.NewImageCell(c.tempDir, f.Name())
}

func DevcardCustomCell(dc *devcard.Devcard) {
	dc.SetTitle("Custom cell")

	dc.Md("Sometimes you may want something more sophisticated than builtin cells. ",
		"For that, you can use a custom cell.")

	dc.Md("Suppose you want a cell that shows a bar chart. ",
		"Let's start by making a type for it:")

	// TODO: Change it to dc.Source call when it's ready.
	code(dc, `type chartCell struct {
	devcard.CustomCell

	data    []datum
	tempDir string
}

type datum struct {
	name  string
	value int
}`)

	dc.Md("Note the embedded `devcard.CustomCell` in `chartCell`. ",
		"It implements all three methods of the `Cell` interface: `Type`, `Append`, and `Erase`. ",
		"The implementations of `Append` and `Erase` do nothing but panic when called; ",
		"you need to provide your own implementations if you want to use `dc.Append` or `dc.Erase`. ",
		"For this chart cell, this is not required—we're going to implement our own appending function:")
	dc.Source("examples.addBar")

	dc.Md("To render a custom cell, it first needs to be casted it into one of the builtin cells: ",
		"`HTMLCell`, `MarkdownCell`, `MonospaceCell`, `ValueCell`, `AnnotatedValueCell`, `ImageCell`, or `ErrorCell`. ",
		"For that, we need to implement `Cast` method:")
	dc.Source("examples.Cast")

	dc.Md("Now the chart cell is ready. Let's add it to the devcard:")
	code(dc, `chart := &chartCell{
	tempDir: dc.TempDir,
	data: []datum{
		{"Monday", 17},
		{"Tuesday", 56},
		{"Wednesday", 37},
		{"Thursday", 39},
		{"Friday", 29},
	},
}
dc.Custom(chart)`)

	chart := &chartCell{
		tempDir: dc.TempDir,
		data: []datum{
			{"Monday", 17},
			{"Tuesday", 56},
			{"Wednesday", 37},
			{"Thursday", 39},
			{"Friday", 29},
		},
	}
	dc.Custom(chart)

	dc.Md("Since we opted out of imlementing `Append` method, we need to let the devcards app know when we're altering our cell. ",
		"It's done by calling `dc.Update`:")

	code(dc, `chart.add("Saturday", 38)
chart.add("Sunday", 30)
dc.Update(chart)`)

	chart = &chartCell{tempDir: dc.TempDir, data: chart.data}
	dc.Custom(chart)

	chart.addBar("Saturday", 38)
	chart.addBar("Sunday", 30)
	dc.Update(chart)
}
