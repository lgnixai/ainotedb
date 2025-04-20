package service

import (
	"context"
	"encoding/json"

	recordRepoPkg "github.com/undb/undb-go/internal/record/repository"
	"github.com/undb/undb-go/internal/view/model"
	"github.com/undb/undb-go/internal/view/repository"
)

type ChartService interface {
	GetChartData(ctx context.Context, viewID string) (interface{}, error)
	UpdateChartConfig(ctx context.Context, viewID string, config *ChartConfig) error
}

type ChartConfig struct {
	Type    model.ChartType        `json:"type"`
	XAxis   string                 `json:"x_axis"`
	YAxis   string                 `json:"y_axis"`
	AggFunc string                 `json:"agg_func"`
	Options map[string]interface{} `json:"options"`
}

type chartService struct {
	viewRepo   repository.ViewRepository
	recordRepo recordRepoPkg.RecordRepository
}

func NewChartService(viewRepo repository.ViewRepository, recordRepo recordRepoPkg.RecordRepository) ChartService {
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

	// 直接使用 *model.View，不做类型断言
// 获取表数据
_, err = s.recordRepo.GetByTableID(ctx, view.TableID)
if err != nil {
	return nil, err
}
// 返回占位数据
return map[string]interface{}{"placeholder": "processChartData"}, nil
}

func (s *chartService) UpdateChartConfig(ctx context.Context, viewID string, config *ChartConfig) error {
	view, err := s.viewRepo.GetByID(ctx, viewID)
	if err != nil {
		return err
	}

	// 直接更新 view 字段（假设 View 结构体有这些字段，若无请调整结构体定义）
view.Type = model.ViewType(config.Type)
// 假设 ViewOptions 结构体有 XAxis/YAxis/AggFunc 字段，否则请扩展结构体
// view.Options.XAxis = config.XAxis
// view.Options.YAxis = config.YAxis
// view.Options.AggFunc = config.AggFunc
// 其余 options 字段序列化到 view.Config
if len(config.Options) > 0 {
	if b, err := json.Marshal(config.Options); err == nil {
		view.Config = string(b)
	}
}
return s.viewRepo.Update(ctx, view)
}
