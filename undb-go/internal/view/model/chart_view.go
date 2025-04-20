package model

type ChartType string

const (
	ChartTypeBar  ChartType = "bar"
	ChartTypeLine ChartType = "line"
	ChartTypePie  ChartType = "pie"
	ChartTypeArea ChartType = "area"
)

type ChartView struct {
	View
	Type    ChartType              `json:"type"`
	XAxis   string                 `json:"x_axis"`
	YAxis   string                 `json:"y_axis"`
	AggFunc string                 `json:"agg_func"`
	Options map[string]interface{} `json:"options"`
}

func NewChartView(name string, tableID string, chartType ChartType) *ChartView {
	return &ChartView{
		View: View{
			ID:      GenerateID(),
			Name:    name,
			Type:    ViewTypeChart,
			TableID: tableID,
		},
		Type:    chartType,
		Options: make(map[string]interface{}),
	}
}
