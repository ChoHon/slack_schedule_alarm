package notion

import (
	"atlasnetworks-cdn-schedule-manage/model"
	"atlasnetworks-cdn-schedule-manage/static"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetNotionScheduleDatabase(bodyStruct model.JSON) ([]model.NotionRow, error) {
	body, err := json.Marshal(bodyStruct)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", static.NOTION_SCHEDULE_DATABASE_URL),
		strings.NewReader(string(body)),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", static.NOTION_AUTHORIZATION))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2022-02-22")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	scheduleList := &model.NotionDatabase{}
	err = json.Unmarshal(respBody, &scheduleList)
	if err != nil {
		return nil, err
	}

	return scheduleList.Results, nil
}

// func MakeFilterJSON(filter model.NotionMultiFilter) string {
// 	filterJSON := ""
// 	if filter.And != nil {
// 		conditions := ""
// 		for _, condition := range filter.And {
// 			if conditions != "" {
// 				conditions += ", "
// 			}
// 			conditions += makeConditionJSON(condition)
// 		}

// 		filterJSON += fmt.Sprintf(`{"and" : [ %s ]}`, conditions)
// 	} else {
// 		conditions := ""
// 		for _, condition := range filter.Or {
// 			if conditions != "" {
// 				conditions += ", "
// 			}
// 			conditions += makeConditionJSON(condition)
// 		}

// 		filterJSON += fmt.Sprintf(`{"or" : [ %s ]}`, conditions)
// 	}

// 	result := fmt.Sprintf(`{"filter" : %s}`, filterJSON)
// 	return result
// }

// func makeConditionJSON(condition model.NotionCondition) string {
// 	switch condition.Property {
// 	case "date":
// 		conditionJSON += fmt.Sprintf(``)
// 	}

// 	result := fmt.Sprintf(`{"property" : %s}`)
// 	return result
// }
