package view

import (
	"errors"
	"gorm.io/gorm"
)

type ViewService struct {
	db *gorm.DB
}

func NewViewService(db *gorm.DB) *ViewService {
	return &ViewService{db: db}
}

func (s *ViewService) CreateView(view *View) error {
	return s.db.Create(view).Error
}

func (s *ViewService) GetView(id string) (*View, error) {
	var view View
	if err := s.db.First(&view, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &view, nil
}

func (s *ViewService) ListViews(tableID string) ([]View, error) {
	var views []View
	if err := s.db.Where("table_id = ?", tableID).Find(&views).Error; err != nil {
		return nil, err
	}
	return views, nil
}

func (s *ViewService) UpdateView(view *View) error {
	return s.db.Save(view).Error
}

func (s *ViewService) DeleteView(id string) error {
	return s.db.Delete(&View{}, "id = ?", id).Error
}

func (s *ViewService) SetViewFields(id string, fields []string) error {
	return s.db.Model(&View{}).Where("id = ?", id).Update("fields", fields).Error
}

func (s *ViewService) SetViewFilter(id string, filter *Filter) error {
	return s.db.Model(&View{}).Where("id = ?", id).Update("filter", filter).Error
}

func (s *ViewService) SetViewSort(id string, sort []Sort) error {
	return s.db.Model(&View{}).Where("id = ?", id).Update("sort", sort).Error
}