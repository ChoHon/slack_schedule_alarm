package model

type ScheduleProperty struct {
	Importance NotionProperty `json:"중요도"`
	Tag        NotionProperty `json:"태그"`
	Status     NotionProperty `json:"상태"`
	Deadline   NotionProperty `json:"마감 날짜"`
	CreatedAt  NotionProperty `json:"생성 일시"`
	People     NotionProperty `json:"사람"`
	CreatedBy  NotionProperty `json:"생성자"`
	Title      NotionProperty `json:"이름"`
}

type JSON map[string]interface{}
