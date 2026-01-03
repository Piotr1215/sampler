package list

import (
	"fmt"
	"image"

	ui "github.com/gizak/termui/v3"
	"github.com/sqshq/sampler/component"
	"github.com/sqshq/sampler/config"
	"github.com/sqshq/sampler/console"
	"github.com/sqshq/sampler/data"
)

type List struct {
	*ui.Block
	*data.Consumer
	alert *data.Alert
	items map[string]string
	order []string
	style ui.Style
}

func NewList(c config.ListConfig, palette console.Palette) *List {

	items := make(map[string]string)
	order := make([]string, 0, len(c.Items))

	for _, item := range c.Items {
		label := *item.Label
		order = append(order, label)
		items[label] = ""
	}

	l := &List{
		Block:    component.NewBlock(c.Title, *c.Border, palette),
		Consumer: data.NewConsumer(),
		items:    items,
		order:    order,
		style:    ui.NewStyle(palette.BaseColor),
	}

	go func() {
		for {
			select {
			case sample := <-l.SampleChannel:
				l.items[sample.Label] = sample.Value
			case alert := <-l.AlertChannel:
				l.alert = alert
			}
		}
	}()

	return l
}

func (l *List) Draw(buffer *ui.Buffer) {

	l.Block.Draw(buffer)

	y := l.Inner.Min.Y
	for _, label := range l.order {
		if y >= l.Inner.Max.Y-1 {
			break
		}

		value := l.items[label]
		line := fmt.Sprintf("%s: %s", label, value)

		cells := ui.ParseStyles(line, ui.Theme.Paragraph.Text)
		cells = ui.TrimCells(cells, l.Inner.Dx()-2)

		for _, cx := range ui.BuildCellWithXArray(cells) {
			x, cell := cx.X, cx.Cell
			cell.Style = l.style
			buffer.SetCell(cell, image.Pt(x+1+l.Inner.Min.X, y))
		}
		y++
	}

	component.RenderAlert(l.alert, l.Rectangle, buffer)
}
