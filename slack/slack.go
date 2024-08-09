package slack

import (
	"atlasnetworks-cdn-schedule-manage/model"
	"atlasnetworks-cdn-schedule-manage/schedule"
	"atlasnetworks-cdn-schedule-manage/static"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func SendDailySchedule(user string) error {
	today := time.Now().Format("2006.01.02")
	workList := []model.NotionRow{}

	overDeadlineList, err := schedule.GetOverDeadlineSchedule(user)
	if err != nil {
		return err
	}
	workList = append(workList, overDeadlineList...)

	todayDeadlineList, err := schedule.GetTodayDeadlineSchedule(user)
	if err != nil {
		return err
	}
	workList = append(workList, todayDeadlineList...)

	if len(workList) == 0 {
		return nil
	}

	workListStr := makeSlackWork(workList)
	alarm := makeSlackScheduleStr(user, today, "마감 기한이 지났거나 오늘까지인 업무 목록입니다", workListStr)

	return sendMessageToSlack(alarm)
}

func SendWeeklySchedule(user string) error {
	week := fmt.Sprintf("%s~%d", time.Now().Format("2006.01.02"), time.Now().AddDate(0, 0, 6).Day())
	workList, err := schedule.GetWeeklySchedule(user)
	if err != nil {
		return err
	}

	if len(workList) == 0 {
		return nil
	}

	workListStr := makeSlackWork(workList)
	alarm := makeSlackScheduleStr(user, week, "이번 주에 예정된 업무 목록입니다", workListStr)

	return sendMessageToSlack(alarm)
}

func makeSlackScheduleStr(user, date, context, WorkListStr string) string {

	mention := fmt.Sprintf(`"text": {"type": "mrkdwn", "text": "<@%s>"}`, static.MemberID[user])
	button := fmt.Sprintf(
		`"accessory": {"type": "button", "text": {"type": "plain_text", "text": "업무 보드 노션 페이지"}, "url": "%s"}`,
		static.NOTION_BOARD_PAGE_URL,
	)
	header := fmt.Sprintf("%s 업무 일정", date)

	message := fmt.Sprintf(`{"type": "section", %s, %s}, `, mention, button)
	message += fmt.Sprintf(`{"type": "header", "text": {"type": "plain_text", "text": "%s"}}, `, header)
	message += fmt.Sprintf(`{"type": "context", "elements": [{"type": "plain_text", "text": "%s"}]}, `, context)
	message += `{"type": "divider"}, `

	message += WorkListStr

	message += `, {"type": "divider"}`

	return fmt.Sprintf(`{"blocks": [%s]}`, message)
}

func makeSlackWork(scheduleList []model.NotionRow) string {
	WorkListJSON := ""
	for _, schedule := range scheduleList {
		tagStr := ""
		for _, tag := range schedule.Properties.Tag.MultiSelect {
			if tagStr != "" {
				tagStr += ", "
			}
			tagStr += fmt.Sprintf("`%s`", tag.Name)
		}

		date, _ := time.Parse("2006-01-02", schedule.Properties.Deadline.Date.Start)

		if WorkListJSON != "" {
			WorkListJSON += `, {"type": "divider"}, `
		}

		WorkListJSON += fmt.Sprintf(
			`{"type": "section", "text": {"type": "mrkdwn", "text": "*%s*"}}, `, schedule.Properties.Title.Title[0].Text,
		)

		contextStr := ""
		contextStr += fmt.Sprintf(
			`{"type": "mrkdwn", "text": "*마감 기한:*\n%s"}, `, date.Format("2006.01.02"),
		)
		contextStr += fmt.Sprintf(
			`{"type": "mrkdwn", "text": "*상태:*\n%s"}, `, fmt.Sprintf("`%s`", schedule.Properties.Status.Status.Name),
		)
		contextStr += fmt.Sprintf(
			`{"type": "mrkdwn", "text": "*중요도:%s*\n%s"}, `, "        ", schedule.Properties.Importance.Select.Name,
		)
		contextStr += fmt.Sprintf(
			`{"type": "mrkdwn", "text": "*태그:*\n%s"}`, tagStr,
		)

		WorkListJSON += fmt.Sprintf(
			`{"type": "context", "elements": [%s]}`, contextStr,
		)
	}

	return WorkListJSON
}

func sendMessageToSlack(body string) error {
	req, err := http.NewRequest(
		"POST",
		static.SLACK_WEBHOOK_URL,
		strings.NewReader(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
