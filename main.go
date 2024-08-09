package main

import (
	"atlasnetworks-cdn-schedule-manage/slack"
	"atlasnetworks-cdn-schedule-manage/static"

	"github.com/robfig/cron"
)

const (
	// Second Minute Hour Day Month Week
	DAILY_START_ALARM    = "0 0 9 * * MON-FRI"
	DAILY_DEADLINE_ALARM = "0 0 17 * * MON-FRI"
	WEEKLY_START_ALARM   = "0 0 9 * * MON"
	TEST_CRON_SPEC       = "0-59/10 * * * * MON-FRI"
)

func main() {
	c := cron.New()
	// c.AddFunc(TEST_CRON_SPEC, test)
	c.AddFunc(DAILY_START_ALARM, SendDailyScheduleStart)
	c.AddFunc(DAILY_DEADLINE_ALARM, SendDailyScheduleEnd)
	c.AddFunc(WEEKLY_START_ALARM, SendWeeklySchedule)
	c.Start()

	select {} // main 함수 종료 방지
}

func SendDailyScheduleStart() {
	for name := range static.UserID {
		slack.SendDailySchedule(name)
	}
}

func SendDailyScheduleEnd() {
	for name := range static.UserID {
		slack.SendDailySchedule(name)
	}
}

func SendWeeklySchedule() {
	for name := range static.UserID {
		slack.SendWeeklySchedule(name)
	}
}

// func test() {
// 	for name := range static.UserID {
// 		if name == "김현진" {
// 			slack.SendWeeklySchedule(name)
// 			slack.SendDailySchedule(name)
// 		}
// 	}
// }
