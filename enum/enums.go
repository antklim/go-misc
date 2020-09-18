package enum

type Weekday int

const (
	Monday Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var weekdays = [...]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

type Month int

const (
	January Month = iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var months = [...]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func (w Weekday) String() string {
	if w < 0 || w > Weekday(len(weekdays)-1) {
		return "Unknown"
	}
	return weekdays[w]
}

func (m Month) String() string {
	if m < 0 || m > Month(len(months)-1) {
		return "Unknown"
	}
	return months[m]
}
