
package service

import (
	"context"
	"github.com/undb/undb-go/internal/view/model"
)

type ChartService interface {
	GetChartData(ctx context.Context, viewID string) (interface{}, error)
	UpdateChartConfig(ctx context.Context, viewID string, config *ChartConfig) error
}

type ChartConfig struct {
	Type    model.ChartType `json:"type"`
	XAxis   string         `json:"x_axis"`
	YAxis   string         `json:"y_axis"`
	AggFunc string         `json:"agg_func"`
	Options map[string]interface{} `json:"options"`
}

type chartService struct {
	viewRepo   ViewRepository
	recordRepo record.RecordRepository
}

func NewChartService(viewRepo ViewRepository, recordRepo record.RecordRepository) ChartService {
	return &chartService{
		viewRepo:   viewRepo,
		recordRepo: recordRepo,
	}
}

func (s *chartService) GetChartData(ctx context.Context, viewID string) (interface{}, error) {
	view, err := s.viewRepo.GetByID(ctx, viewID)
	if err != nil {
		return nil, err
	}

	chartView, ok := view.(*model.ChartView)
	if !ok {
		return nil, ErrInvalidViewType
	}

	// Get and aggregate data based on chart configuration
	records, err := s.recordRepo.GetByTableID(ctx, view.GetTableID())
	if err != nil {
		return nil, err
	}

	// Process data according to chart type and configuration
	// This is a placeholder - actual implementation would depend on specific requirements
	data := processChartData(records, chartView)

	return data, nil
}

func (s *chartService) UpdateChartConfig(ctx context.Context, viewID string, config *ChartConfig) error {
	view, err := s.viewRepo.GetByID(ctx, viewID)
	if err != nil {
		return err
	}

	chartView, ok := view.(*model.ChartView)
	if !ok {
		return ErrInvalidViewType
	}

	// Update chart configuration
	chartView.Type = config.Type
	chartView.XAxis = config.XAxis
	chartView.YAxis = config.YAxis
	chartView.AggFunc = config.AggFunc
	chartView.Options = config.Options

	return s.viewRepo.Update(ctx, chartView)
}
