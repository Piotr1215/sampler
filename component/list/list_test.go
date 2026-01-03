package list

import (
	"runtime"
	"testing"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/sqshq/sampler/config"
	"github.com/sqshq/sampler/console"
	"github.com/sqshq/sampler/data"
)

func TestNewList(t *testing.T) {
	border := true
	label1 := "Label1"
	label2 := "Label2"
	sample1 := "echo test1"
	sample2 := "echo test2"

	cfg := config.ListConfig{
		ComponentConfig: config.ComponentConfig{
			Title: "Test List",
		},
		Items: []config.Item{
			{Label: &label1, SampleScript: &sample1},
			{Label: &label2, SampleScript: &sample2},
		},
		Border: &border,
	}

	palette := console.GetPalette(console.ThemeDark)
	list := NewList(cfg, palette)

	if list == nil {
		t.Fatal("NewList returned nil")
	}

	if len(list.order) != 2 {
		t.Errorf("expected 2 items in order, got %d", len(list.order))
	}

	if list.order[0] != "Label1" {
		t.Errorf("expected first label 'Label1', got %s", list.order[0])
	}

	// Test sample channel updates items
	list.SampleChannel <- &data.Sample{Label: "Label1", Value: "test-value"}

	// Give goroutine time to process
	runtime.Gosched()
	time.Sleep(100 * time.Millisecond)

	if list.items["Label1"] != "test-value" {
		t.Errorf("expected item value 'test-value', got %s", list.items["Label1"])
	}
}

func TestListDraw(t *testing.T) {
	border := true
	label := "Test"
	sample := "echo test"

	cfg := config.ListConfig{
		ComponentConfig: config.ComponentConfig{
			Title: "Test",
		},
		Items:  []config.Item{{Label: &label, SampleScript: &sample}},
		Border: &border,
	}

	palette := console.GetPalette(console.ThemeDark)
	list := NewList(cfg, palette)

	list.SetRect(0, 0, 20, 10)
	list.items["Test"] = "value"

	buf := ui.NewBuffer(list.GetRect())
	list.Draw(buf)

	// Just verify it doesn't panic
}
