type jsonTime time.Time

var layout = "2006-01-02T15:04:05"

func (t jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(t).Format(layout))), nil
}

func (t *jsonTime) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	*(*time.Time)(t), err = time.Parse(layout, r)

	if err != nil {
		return err
	}
	return
}

func (t jsonTime) String() string { return time.Time(t).UTC().Format(timeFormat) }