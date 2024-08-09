package schedule

import (
	"atlasnetworks-cdn-schedule-manage/model"
	"atlasnetworks-cdn-schedule-manage/notion"
	"atlasnetworks-cdn-schedule-manage/static"
	"sort"
	"time"
)

func GetTodayDeadlineSchedule(user string) ([]model.NotionRow, error) {
	today := time.Now().Format("2006-01-02")

	bodySturct := model.JSON{
		"filter": model.JSON{
			"and": []model.JSON{
				{
					"property": "마감 날짜",
					"date": model.JSON{
						"equals": today,
					},
				},
				{
					"property": "사람",
					"people": model.JSON{
						"contains": static.UserID[user],
					},
				},
			},
		},
	}

	scheduleList, err := getOngoingSchedule(bodySturct)
	return scheduleList, err
}

func GetOverDeadlineSchedule(user string) ([]model.NotionRow, error) {
	today := time.Now().Format("2006-01-02")

	bodySturct := model.JSON{
		"filter": model.JSON{
			"and": []model.JSON{
				{
					"property": "마감 날짜",
					"date": model.JSON{
						"before": today,
					},
				},
				{
					"property": "사람",
					"people": model.JSON{
						"contains": static.UserID[user],
					},
				},
			},
		},
	}

	scheduleList, err := getOngoingSchedule(bodySturct)
	return scheduleList, err
}

func GetWeeklySchedule(user string) ([]model.NotionRow, error) {
	bodySturct := model.JSON{
		"filter": model.JSON{
			"and": []model.JSON{
				{
					"property": "마감 날짜",
					"date": model.JSON{
						"this_week": model.JSON{},
					},
				},
				{
					"property": "사람",
					"people": model.JSON{
						"contains": static.UserID[user],
					},
				},
			},
		},
	}

	scheduleList, err := getOngoingSchedule(bodySturct)
	return scheduleList, err
}

func getOngoingSchedule(filter model.JSON) ([]model.NotionRow, error) {
	scheduleTargetList := []model.NotionRow{}
	conditions := filter["filter"].(model.JSON)["and"].([]model.JSON)

	for _, status := range []string{"시작 전", "진행 중"} {
		filter["filter"].(model.JSON)["and"] = append(conditions, model.JSON{
			"property": "상태",
			"status": model.JSON{
				"equals": status,
			},
		})

		scheduleAllList, err := notion.GetNotionScheduleDatabase(filter)
		if err != nil {
			return nil, err
		}

		scheduleTargetList = append(scheduleTargetList, scheduleAllList...)
	}

	// 중요도 순으로 정렬
	sort.Slice(scheduleTargetList, func(i, j int) bool {
		importanceI := scheduleTargetList[i].Properties.Importance.Select.Name
		importanceJ := scheduleTargetList[j].Properties.Importance.Select.Name
		return len(importanceI) > len(importanceJ)
	})

	return scheduleTargetList, nil
}
