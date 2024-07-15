package main

import (
	"fmt"
	"time"
)

// 计算当前UTC时间距离下一个周五0点的秒数，并返回这是年初以来的第几周
func timeToNextFridayUTC() (seconds int64, weekNumber int) {
	// 获取当前UTC时间
	now := time.Now().UTC()

	// 计算本周的周五（即本周几到周五的差值）
	weekday := int(now.Weekday())
	// 周一=0, 周二=1, ..., 周日=6
	// 如果今天是周五或之后，则计算下周的周五
	daysUntilFriday := 5 - weekday
	if daysUntilFriday <= 0 {
		daysUntilFriday += 7
	}

	// 计算下一个周五的UTC时间
	nextFriday := now.AddDate(0, 0, daysUntilFriday)

	// 如果当前时间已经过了今天的0点，则nextFriday是今天之后的周五；
	// 否则，如果现在是今天的0点之前，nextFriday实际上是“今天”的周五（如果今天是周五或之前）
	// 但因为我们总是想要下一个周五（或今天的周五如果已经过了午夜），所以上面的逻辑已经处理了这一点

	// 计算距离下一个周五0点的秒数
	secondsUntilNextFriday := nextFriday.Sub(now.Truncate(24 * time.Hour)).Seconds()

	// 计算这是年初以来的第几周
	// 注意：ISO 8601周数计算可能稍微复杂一些，这里我们简化处理，只考虑从年初到今天的总天数然后除以7
	// 注意这种方式可能不会完全符合ISO 8601标准（特别是年初和年末）
	yearStart := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	daysSinceYearStart := int(now.Sub(yearStart).Hours() / 24)
	weekNumber = (daysSinceYearStart / 7) + 1 // +1是因为我们想要的是周数，不是从0开始的索引

	return int64(secondsUntilNextFriday), weekNumber
}

func main() {
	seconds, weekNumber := timeToNextFridayUTC()
	fmt.Printf("距离下一个周五0点还有 %d 秒\n", seconds)
	fmt.Printf("这是年初以来的第 %d 周\n", weekNumber)
}
