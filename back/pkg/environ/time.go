package environ

import "time"

// LocationJST - JSTのLocationを返す
func LocationJST() *time.Location {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		jst = time.FixedZone("Asia/Tokyo", 9*60*60)
	}

	return jst
}

// LocationUTC - UTCのLocationを返す
func LocationUTC() *time.Location {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		utc = time.FixedZone("UTC", 0)
	}

	return utc
}

// InUTC - 引数をUTC変換したTimeを返す
func InUTC(t time.Time) time.Time {
	return t.In(LocationUTC())
}

// InJST - 引数をJST変換したTimeを返す
func InJST(t time.Time) time.Time {
	return t.In(LocationJST())
}

// NowUTC - UTC変換した現在日時を返す
func NowUTC() time.Time {
	return InUTC(time.Now())
}

// NowJST - JST変換した現在日時を返す
func NowJST() time.Time {
	return InJST(time.Now())
}
