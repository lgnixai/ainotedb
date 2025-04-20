
package model

type ChartType string

const (
	ChartTypeBar   ChartType = "bar"
	ChartTypeLine  ChartType = "line"
	ChartTypePie   ChartType = "pie"
	ChartTypeArea  ChartType = "area"
)

type ChartView struct {
	ViewBase
	Type       ChartType               `json:"type"`
	XAxis      string                  `json:"x_axis"`
	YAxis      string                  `json:"y_axis"`
	AggFunc    string                  `json:"agg_func"`
	Options    map[string]interface{}  `json:"options"`
}

func NewChartView(name string, tableID string, chartType ChartType) *ChartView {
	return &ChartView{
		ViewBase: ViewBase{
			ID:      GenerateID(),
			Name:    name,
			Type:    ViewTypeChart,
			TableID: tableID,
		},
		Type:    chartType,
		Options: make(map[string]interface{}),
	}
}
