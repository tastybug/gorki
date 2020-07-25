package util

import "time"

const isoLayout = `2006-01-02`

func Iso8601ToRfc822Date(isoDate string) string {
	dateTime, err := time.Parse(isoLayout, isoDate)
	PanicOnError(err)

	// https://www.w3.org/Protocols/rfc822/#z28
	return dateTime.Format(`02 Jan 06 15:04 UT`)
}
