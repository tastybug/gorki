package util

import "time"

const isoLayout = `2006-01-02`

func ISODateToRSSDateTime(isoDate string) string {
	dateTime, err := time.Parse(isoLayout, isoDate)
	PanicOnError(err)

	// RSS asks for RFC822 date formats, see https://www.w3.org/Protocols/rfc822/#z28
	// nonetheless the RSS validator at https://validator.w3.org/feed/check.cgi asks for day of the week
	// which you get with RC1123 only, so using this instead of the 822 formatter
	return dateTime.Format(time.RFC1123Z)
}

func GetNowAsRSSDateTime() string {
	return time.Now().Format(time.RFC1123Z)
}
