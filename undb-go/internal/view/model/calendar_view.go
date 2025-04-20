package model

//
//import (
//	"time"
//)
//
//// Assuming these structs are defined elsewhere in your project
//type View struct {
//	// ...fields of View struct...
//}
//
//type Record struct {
//	ID     string                 `json:"id"`
//	Values map[string]interface{} `json:"values"`
//}
//
//func (r *Record) GetDisplayValue() string {
//	//Implementation to get display value
//	return ""
//}
//
//type CalendarEvent struct {
//	ID    string      `json:"id"`
//	Title string      `json:"title"`
//	Start time.Time   `json:"start"`
//	End   time.Time   `json:"end"`
//	Color interface{} `json:"color,omitempty"`
//}
//
//type CalendarView struct {
//	View
//	TimeField    string    `json:"timeField"`
//	TimeScale    string    `json:"timeScale"` // day, week, month
//	ColorField   string    `json:"colorField,omitempty"`
//	ShowWeekends bool      `json:"showWeekends"`
//	StartTime    time.Time `json:"startTime"`
//	EndTime      time.Time `json:"endTime"`
//}
//
//func (v *CalendarView) GetEvents(records []Record) []CalendarEvent {
//	events := make([]CalendarEvent, 0)
//	for _, record := range records {
//		if timeValue, ok := record.Values[v.TimeField].(time.Time); ok {
//			event := CalendarEvent{
//				ID:    record.ID,
//				Title: record.GetDisplayValue(),
//				Start: timeValue,
//				End:   timeValue.Add(24 * time.Hour),
//			}
//			if v.ColorField != "" {
//				event.Color = record.Values[v.ColorField]
//			}
//			events = append(events, event)
//		}
//	}
//	return events
//}
//
//func main1() {
//	// Example usage (replace with your actual data)
//	records := []Record{
//		{ID: "1", Values: map[string]interface{}{"timeField": time.Now(), "colorField": "blue", "otherField": "value"}},
//		{ID: "2", Values: map[string]interface{}{"timeField": time.Now().AddDate(0, 0, 1), "colorField": "red"}},
//	}
//
//	calendarView := CalendarView{
//		TimeField:  "timeField",
//		TimeScale:  "day",
//		ColorField: "colorField",
//		StartTime:  time.Now(),
//		EndTime:    time.Now().AddDate(0, 1, 0),
//	}
//
//	events := calendarView.GetEvents(records)
//	println(events)
//}
