package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	table := tview.NewTable()

	// 曜日を設定
	weekdays := []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday}
	for i, weekday := range weekdays {
		table.SetCell(0, i, tview.NewTableCell(weekday.String()[:2]))
	}

	// 日付を設定
	now := time.Now()
	days := getDays(now)
	row := 1
	for _, day := range days {
		tableCell := tview.NewTableCell(fmt.Sprintf("%2s", strconv.Itoa(day.Day())))

		if now.Day() == day.Day() {
			tableCell.SetAttributes(tcell.AttrReverse)
		}

		table.SetCell(row, int(day.Weekday()), tableCell)

		if day.Weekday() == time.Saturday {
			row++
		}
	}

	// 年月を設定
	header := tview.NewTextView().SetText(now.Format("January 2006")).SetSize(1, 20).SetTextAlign(tview.AlignCenter)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 1, 0, false).
		AddItem(table, 0, 1, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

func getDays(now time.Time) []time.Time {
	currentYear := now.Year()
	currentMonth := now.Month()

	firstDay := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.Local)
	nextMonth := firstDay.AddDate(0, 1, 0)

	var days []time.Time
	for i := firstDay; i.Before(nextMonth); i = i.AddDate(0, 0, 1) {
		days = append(days, i)
	}

	return days
}
