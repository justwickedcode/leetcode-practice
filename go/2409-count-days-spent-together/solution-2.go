// 2409. Count Days Spent Together (Easy)
// https://leetcode.com/problems/count-days-spent-together/
// Runtime: 0 ms (beats 100.0%) · Memory: 3.8 MB (beats 100.0%)

package main

import (
	"fmt"
	"strconv"
)

type Date struct {
	day   int
	month int
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	aliceArriveDate := dateStringToDate(arriveAlice)
	aliceLeaveDate := dateStringToDate(leaveAlice)
	bobArriveDate := dateStringToDate(arriveBob)
	bobLeaveDate := dateStringToDate(leaveBob)

	aliceStart := dateToDayOfYear(aliceArriveDate)
	aliceEnd := dateToDayOfYear(aliceLeaveDate)
	bobStart := dateToDayOfYear(bobArriveDate)
	bobEnd := dateToDayOfYear(bobLeaveDate)

	startOverlap := max(aliceStart, bobStart)
	endOverlap := min(aliceEnd, bobEnd)

	if startOverlap <= endOverlap {
		return endOverlap - startOverlap + 1
	}
	return 0
}

func dateStringToDate(date string) Date {
	day, _ := strconv.Atoi(date[3:])   // "MM-DD" -> DD is at index 3 onwards
	month, _ := strconv.Atoi(date[:2]) // "MM-DD" -> MM is at index 0 to 2
	return Date{day, month}
}

func dateToDayOfYear(d Date) int {
	daysInMonth := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	
	dayCount := d.day
	for m := 1; m < d.month; m++ {
		dayCount += daysInMonth[m]
	}
	return dayCount
}
