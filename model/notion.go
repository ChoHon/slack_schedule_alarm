package model

import (
	"time"
)

type NotionDatabase struct {
	Object  string      `json:"object"`
	Results []NotionRow `json:"results"`
}

type NotionRow struct {
	Object     string           `json:"object"`
	ID         string           `json:"id"`
	Archived   bool             `json:"archived"`
	InTrash    bool             `json:"in_trash"`
	Properties ScheduleProperty `json:"properties"`
}

type NotionProperty struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Select      NotionPropertySelect   `json:"select"`
	MultiSelect []NotionPropertySelect `json:"multi_select"`
	Status      NotionPropertySelect   `json:"status"`
	Date        NotionPropertyDate     `json:"date"`
	CreatedTime time.Time              `json:"created_time"`
	People      []NotionUser           `json:"people"`
	CreatedBy   NotionUser             `json:"created_by"`
	Title       []NotionPropertyTitle  `json:"title"`
}

type NotionUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NotionPropertySelect struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NotionPropertyDate struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	TimeZone string `json:"time_zone"`
}

type NotionPropertyTitle struct {
	Text string `json:"plain_text"`
}

// type NotionMultiFilter struct {
// 	And []NotionCondition `json:"and"`
// 	Or  []NotionCondition `json:"or"`
// }

// func NewNotionMultiFilter(logical string, conditions ...NotionCondition) *NotionMultiFilter {
// 	filter := &NotionMultiFilter{}

// 	switch logical {
// 	case "and":
// 		filter.And = conditions
// 	case "or":
// 		filter.Or = conditions
// 	default:
// 		filter = nil
// 	}

// 	return filter
// }

// type NotionCondition struct {
// 	Property    string `json:"property"`
// 	Date        string `json:"date"`
// 	Status      string `json:"status"`
// 	People      string `json:"people"`
// 	Select      string `json:"select"`
// 	MultiSelect string `json:"multi_select"`
// }

// func NewNotionCondition(conditionType, operator, value string) *NotionCondition {
// 	condition := &NotionCondition{}
// 	operateStr := fmt.Sprintf(`{"%s": "%s"}`, operator, value)

// 	switch conditionType {
// 	case "date":
// 		condition.Date = operateStr
// 	case "status":
// 		condition.Status = operateStr
// 	case "people":
// 		condition.People = operateStr
// 	case "select":
// 		condition.Select = operateStr
// 	case "multi_select":
// 		condition.MultiSelect = operateStr
// 	default:
// 		condition = nil
// 	}

// 	return condition
// }
