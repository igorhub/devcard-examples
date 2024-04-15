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
func (cc *chartCell) addBar(name string, value int) {
	cc.data = append(cc.data, datum{name, value})
}

// graph turns the cell's data into a chart.BarChart instance.
func (cc *chartCell) graph() chart.BarChart {
	var values []chart.Value
	for _, d := range cc.data {
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

func (cc *chartCell) Cast() devcard.Cell {
	f, err := os.CreateTemp(cc.tempDir, "*.png")
	if err != nil {
		return devcard.NewErrorCell("chartCell error", err)
	}
	defer f.Close()
	err = cc.graph().Render(chart.PNG, f)
	if err != nil {
		return devcard.NewErrorCell("chartCell error", err)
	}

	return devcard.NewImageCell(cc.tempDir, f.Name())
}

func DevcardCustomCell(c *devcard.Devcard) {
	c.SetTitle("4. Custom cell")

	c.Md("Sometimes we want something more sophisticated than builtin cells.")
	c.Append("For that, we may use a custom cell.")

	c.Md("Let's say we want a cell that shows a bar chart.")
	c.Append("We start by making a type for it:")

	// TODO: Change it to c.Source call when it's ready.
	code(c, `type chartCell struct {
	devcard.CustomCell

	data    []datum
	tempDir string
}

type datum struct {
	name  string
	value int
}`)

	c.Md("Note that we embedded `devcard.CustomCell` in our struct.")
	c.Append("It implements all three methods of the `Cell` interface: `Type`, `Append`, and `Erase`.")
	c.Append("The implementations of `Append` and `Erase` do nothing but panic when called;")
	c.Append("we need to provide our own implementations if we want to use `c.Append` or `c.Erase`.")
	c.Append("For our chart cell, this is not required—we're going to implement our own appending function:")
	c.Source("examples.addBar")

	c.Md("To render a custom cell, it needs to be casted into one of the builtin cells:")
	c.Append("`HTMLCell`, `MarkdownCell`, `MonospaceCell`, `ValueCell`, `AnnotatedValueCell`, `ImageCell`, or `ErrorCell`.")
	c.Append("For that, we need to implement `Cast` method:")
	c.Source("examples.Cast")

	c.Md("Now our chart cell is ready. Let's add it to the devcard:")
	code(c, `chart := &chartCell{
	tempDir: c.TempDir,
	data: []datum{
		{"Monday", 17},
		{"Tuesday", 56},
		{"Wednesday", 37},
		{"Thursday", 39},
		{"Friday", 29},
	},
}
c.Custom(chart)`)

	chart := &chartCell{
		tempDir: c.TempDir,
		data: []datum{
			{"Monday", 17},
			{"Tuesday", 56},
			{"Wednesday", 37},
			{"Thursday", 39},
			{"Friday", 29},
		},
	}
	c.Custom(chart)

	c.Md("Since we opted out of imlementing `Append` method, we need to let the devcards app know when we're altering our cell.")
	c.Append("We do it by calling `c.Update`:")

	code(c, `chart.add("Saturday", 38)
chart.add("Sunday", 30)
c.Update(chart)`)

	chart = &chartCell{tempDir: c.TempDir, data: chart.data}
	c.Custom(chart)

	chart.addBar("Saturday", 38)
	chart.addBar("Sunday", 30)
	c.Update(chart)
}
