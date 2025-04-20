package model

type KanbanView struct {
	View
	GroupField string                 `json:"group_field"`
	Options    map[string]interface{} `json:"options"`
}

func NewKanbanView(name string, tableID string, groupField string) *KanbanView {
	return &KanbanView{
		View: View{
			ID:      GenerateID(),
			Name:    name,
			Type:    ViewTypeKanban,
			TableID: tableID,
		},
		GroupField: groupField,
		Options:    make(map[string]interface{}),
	}
}

func (v *KanbanView) GetViewType() ViewType {
	return ViewTypeKanban
}

func (v *KanbanView) SetOption(key string, value interface{}) {
	v.Options[key] = value
}
