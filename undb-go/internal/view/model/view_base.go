package model

//
//// ViewType represents the type of the view
//type ViewType string
//
//const (
//	ViewTypeGrid    ViewType = "grid"
//	ViewTypeKanban  ViewType = "kanban"
//	ViewTypeGallery  ViewType = "gallery"
//	ViewTypeList     ViewType = "list"
//	ViewTypeCalendar  ViewType = "calendar"
//	ViewTypePivot     ViewType = "pivot"
//)
//
//// ViewBase provides common fields for all views
//type ViewBase struct {
//	ID      string   `json:"id"`
//	Name    string   `json:"name"`
//	Type    ViewType `json:"type"`
//	TableID string   `json:"table_id"`
//}
//
// GenerateID generates a new ID for the view

//
//// NewViewBase creates a new instance of ViewBase
//func NewViewBase(name string, viewType ViewType, tableID string) *ViewBase {
//	return &ViewBase{
//		ID:      GenerateID(),
//		Name:    name,
//		Type:    viewType,
//		TableID: tableID,
//	}
//}
